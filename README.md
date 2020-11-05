# goliath
golang library สำหรับช่วยเขียน Go ให้ง่ายขึ้น ใช้เวลาในการเขียนน้อยลง 

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
ส่วนของ value ของ query string จะอยู่ในรูปแบบ **filter-tag** : _**filter-value**_ , ... , **filter-tag** : _**filter-value**_

#### Filter-tag ที่มีให้ใช้งาน
1. eq
2. not
3. gt
4. gte
5. lt
6. lte
7. like
8. is
9. in
