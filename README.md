# goliath
golang clean framwork (testing)

## Feature
- Validate
- CLI

## Query String
แบ่งออกเป็น 3 กลุ่ม คือ
1. Filter
2. Sort
3. Limit

### 1. Filter
ตัวอย่างการใช้งาน
```
/employee?country=thailand&salary=gt:30000&age=gte:30,lt:40
```
ส่วนของ value ของ query string จะอยู่ในรูปแบบ **filter-tag**:_**filter-value**_
