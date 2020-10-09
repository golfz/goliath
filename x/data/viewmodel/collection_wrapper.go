package viewmodel

type CollectionWrapper struct {
	Size      int         `json:"size"`
	TotalSize int         `json:"total_size"`
	Page      int         `json:"page"`
	TotalPage int         `json:"total_page"`
	Limit     Limit       `json:"limit"`
	OrderBy   []OrderBy   `json:"order_by"`
	Link      LinkData    `json:"links"`
	Data      interface{} `json:"data"`
}

type Limit struct {
	Offset int `json:"offset"`
	Count  int `json:"count"`
}

type OrderBy struct {
	Field string `json:"field"`
	Desc  bool   `json:"desc"`
}

type LinkData struct {
	Self  string `json:"self,omitempty"`
	First string `json:"first,omitempty"`
	Last  string `json:"last,omitempty"`
	Prev  string `json:"prev,omitempty"`
	Next  string `json:"next,omitempty"`
	Full  string `json:"full,omitempty"`
}
