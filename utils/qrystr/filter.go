package qrystr

import (
	"errors"
	"fmt"
	"net/http"
	"sort"
	"strings"
)

var filterTag = []string{"eq", "not", "gt", "gte", "lt", "lte", "like", "is", "in"}

type FilterField struct {
	Field     string                 `json:"field"`
	Condition []FilterFieldCondition `json:"condition"`
}

type FilterFieldCondition struct {
	Operator string   `json:"operator"`
	Values   []string `json:"values"`
}

func getFilter(r *http.Request) []FilterField {
	out := make([]FilterField, 0)

	q := r.URL.Query()

	fieldNames := make([]string, 0)
	for field, _ := range q {
		fieldNames = append(fieldNames, field)
	}

	sort.Strings(fieldNames)

	for _, field := range fieldNames {
		value := q[field]

		if !isFilterField(field) {
			continue
		}

		var fieldConditions []FilterFieldCondition
		tags := strings.Split(value[0], ",")
		for _, v := range tags {
			condition, err := getFilterFieldCondition(v)
			if err != nil {
				continue
			}
			fieldConditions = append(fieldConditions, condition)
		}

		if len(fieldConditions) == 0 {
			continue
		}

		out = append(out, FilterField{
			Field:     field,
			Condition: fieldConditions,
		})
	}

	return out
}

func isFilterField(field string) bool {
	if _, has := hasStringInSlice(getStandardLimitAndSortFieldList(), field); has {
		return false
	}

	return true
}

func getFilterFieldCondition(s string) (FilterFieldCondition, error) {

	if s == "" {
		return FilterFieldCondition{}, errors.New("filter tag:value is empty string")
	}

	commaIndex := strings.Index(s, ",")
	if commaIndex != -1 {
		return FilterFieldCondition{}, errors.New("filter tag:value has ','(comma)")
	}

	splitTagVal := strings.Split(s, ":")

	if len(splitTagVal) > 1 {
		tag := splitTagVal[0]
		val := splitTagVal[1]
		tag = strings.TrimSpace(tag)

		if tag == "" || val == "" {
			return FilterFieldCondition{}, errors.New("filter tag or value is empty string")
		}

		if _, has := hasStringInSlice(filterTag, tag); !has {
			return FilterFieldCondition{}, errors.New("filter tag is wrong")
		}

		if tag == "in" {
			inValues := strings.Split(val, "+")
			return FilterFieldCondition{
				Operator: "in",
				Values:   inValues,
			}, nil
		}

		if tag == "is" && val != "null" && val != "notnull" {
			return FilterFieldCondition{}, errors.New("'is' tag accept only values: 'null' & 'notnull'")
		}

		return FilterFieldCondition{
			Operator: tag,
			Values:   []string{val},
		}, nil
	}

	// below this, len(splitTagVal) == 1
	val := splitTagVal[0]

	return FilterFieldCondition{
		Operator: "eq",
		Values:   []string{val},
	}, nil
}

func (q *QueryString) GetFilterSql(isNeedWhereKeyword bool) string {
	if len(q.Filter) == 0 {
		return ""
	}

	out := ""

	if isNeedWhereKeyword {
		out += " WHERE "
	}

	for i, filter := range q.Filter {
		if i != 0 {
			out += " AND "
		}
		for j, tag := range filter.Condition {
			if j != 0 {
				out += " AND "
			}
			//"eq", "not", "gt", "gte", "lt", "lte", "like", "is", "in"
			switch tag.Operator {
			case "eq":
				out += fmt.Sprintf(" `%s` = '%s' ", filter.Field, tag.Values[0])
			case "not":
				out += fmt.Sprintf(" `%s` <> '%s' ", filter.Field, tag.Values[0])
			case "gt":
				out += fmt.Sprintf(" `%s` > '%s' ", filter.Field, tag.Values[0])
			case "gte":
				out += fmt.Sprintf(" `%s` >= '%s' ", filter.Field, tag.Values[0])
			case "lt":
				out += fmt.Sprintf(" `%s` < '%s' ", filter.Field, tag.Values[0])
			case "lte":
				out += fmt.Sprintf(" `%s` <= '%s' ", filter.Field, tag.Values[0])
			case "like":
				like := strings.ReplaceAll(tag.Values[0], "~", "%")
				out += fmt.Sprintf(" `%s` LIKE '%s' ", filter.Field, like)
			case "is":
				isVal := tag.Values[0]
				if isVal == "null" {
					isVal = "IS NULL"
				} else if isVal == "notnull" {
					isVal = "IS NOT NULL"
				}
				out += fmt.Sprintf(" `%s` %s ", filter.Field, isVal)
			case "in":
				inVal := ""
				for inIndex, inItem := range tag.Values {
					if inIndex == 0 {
						inVal += "("
					}
					if inIndex > 0 {
						inVal += ", "
					}

					inVal += fmt.Sprintf("'%s'", inItem)

					if inIndex == len(tag.Values)-1 {
						inVal += ")"
					}
				}
				out += fmt.Sprintf(" `%s` IN %s ", filter.Field, inVal)
			}
		}
	}

	out = strings.ReplaceAll(out, "  AND  ", " AND ")
	out = strings.ReplaceAll(out, "WHERE  ", "WHERE ")

	return out
}
