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

##### การใช้งาน Filter-tag
##### 1. eq
**รูปแบบการใช้งาน :** eq:value เช่น `age=eq:12` หรือ `country=eq:thailand`

eq เป็น filter-tag พิเศษที่จะใส่ tag หรือไม่ก็ได้ ถ้าไม่ใส่ tag จะใช้แบบนี้
```
/employee?country=thailand
```
หรือจะใช้แบบใส่ tag
```
/employee?country=eq:thailand
```
ซึ่งการออกแบบให้ไม่ต้องใส่ tag ก็ได้ เพื่อให้เหมือนกับการเขียน query string ตามปกติที่ `/employee?country=thailand` ก็มีความว่า country ต้องเท่ากับ thailand อยู่แล้ว ซึ่งเหมือนกับการใส่ tag ในรูปแบบนี้ `/employee?country=eq:thailand`
