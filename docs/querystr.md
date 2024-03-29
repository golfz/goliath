# Query String
goliath ช่วยให้ทำงานกับ query string ได้ง่ายขึ้น และสามารถใช้ร่วมกับไลบรารี่ HTTP Router ได้ทุกตัว

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

#### การใช้งาน Filter-tag

---

#### 1. eq
💡 คือการ filter แบบ **เท่ากับ**

📢 **รูปแบบการใช้งาน :** `eq:value` เช่น `age=eq:12` หรือ `country=eq:thailand`

eq เป็น filter-tag พิเศษที่จะใส่ tag หรือไม่ก็ได้ ถ้าไม่ใส่ tag จะใช้แบบนี้
```
/employee?country=thailand
```
หรือจะใช้แบบใส่ tag
```
/employee?country=eq:thailand
```
ซึ่งการออกแบบให้ไม่ต้องใส่ tag ก็ได้ เพื่อให้เหมือนกับการเขียน query string ตามปกติที่ `/employee?country=thailand` ก็มีความว่า country ต้องเท่ากับ thailand อยู่แล้ว ซึ่งเหมือนกับการใส่ tag ในรูปแบบนี้ `/employee?country=eq:thailand`

---

#### 2. not
💡 คือการ filter แบบ **ไม่เท่ากับ**

📢 **รูปแบบการใช้งาน :** `not:value` เช่น `color=not:red` หรือ `country=not:usa,not:canada`

---

#### 3. gt
💡 คือการ filter แบบ **มากกว่า (>)**

📢 **รูปแบบการใช้งาน :** `gt:value` เช่น `age=gt:30` คือ การ filter ข้อมูลที่มี age > 30

---

#### 4. gte
💡 คือการ filter แบบ **มากกว่าหรือเท่ากับ (≥)**

📢 **รูปแบบการใช้งาน :** `gte:value` เช่น `age=gte:30` คือ การ filter ข้อมูลที่มี age ≥ 30

---

#### 5. lt
💡 คือการ filter แบบ **น้อยกว่า (<)**

📢 **รูปแบบการใช้งาน :** `lt:value` เช่น `age=lt:30` คือ การ filter ข้อมูลที่มี age < 30

---

#### 6. lte
💡 คือการ filter แบบ **น้อยกว่าหรือเท่ากับ (≤)**

📢 **รูปแบบการใช้งาน :** `lte:value` เช่น `age=lte:30` คือ การ filter ข้อมูลที่มี age ≤ 30

---

> **หมายเหตุ**
>
> เราสามารถใช้ `gt`, `gte`, `lt`, `lte` ร่วมกันได้
>
> โดยจะมีผลเหมือนการ `AND` กันของ `SQL` เช่น
>
> `/user?age=gte:20,lt:30` หมายถึงการ filter ข้อมูล user ที่มีเงื่อนไข `20 ≤ age < 30`
>
> เทียบกับการเขียน `SQL` คือ `SELECT * FROM user WHERE age>=20 AND age<30`
>

---

#### 7. like
💡 คือการ filter แบบ **ข้อมูลที่ต้องการ เป็นส่วนหนึ่งของข้อมูลทั้งหมด**

📢 **รูปแบบการใช้งาน :** `like:value`

ตัวอย่างการใช้งาน

เช่น
```
name=like:ad
``` 
คือ การ filter ข้อมูล name ที่มี "ad" อยู่ในชื่อในตำแหน่งใดก็ได้ ซึ่งอาจจะได้ผลลัพธ์ **Ad**rian, Br**ad**shaw, Br**ad** เป็นต้น

นอกจากนี้ยังสามารถเพิ่มความสามารถของ filter ด้วย widecard ได้อีกด้วย

**หมายเหตุ:** เมื่อใช้ widecard แล้ว ตำแหน่งและจำนวนของตัวอักษรจะถูกให้ความสำคัญด้วย

Widecard ที่สามารถใช้งานได้

`_` (underscore) ใช้แทน **ตัวอักษรใด ๆ จำนวน 1 ตัว** ในตำแหน่งที่ widecard ปรากฎอยู่

`~` (tilde) ใช้แทน **ตัวอักษรใด ๆ จำนวนกี่ตัวก็ได้** ในตำแหน่งที่ widecard ปรากฎอยู่

ตัวอย่างการใช้งาน เช่น

```
name=like:to_
``` 
คือ การ filter ข้อมูล name ที่ขึ้นต้นด้วย "to" อยู่ในชื่อในตำแหน่งขึ้นต้น และตามด้วยตัวอักษร 1 ตัว ซึ่งอาจจะได้ผลลัพธ์ **To**m, **To**y, **To**r เป็นต้น

.

```
name=like:to~
``` 
คือ การ filter ข้อมูล name ที่ขึ้นต้นด้วย "to" อยู่ในชื่อในตำแหน่งขึ้นต้น และตามด้วยตัวอักษรกี่ตัวก็ได้ ซึ่งอาจจะได้ผลลัพธ์ **To**m, **To**y, **To**r, **To**rrian, **To**shiro เป็นต้น

.

```
name=like:~a~d~
``` 
คือ การ filter ข้อมูล name ที่ขึ้นต้นด้วย "a" และ "d" อยู่ในชื่อในตำแหน่งใดก็ได้ตามลำดับ และตามด้วยตัวอักษรกี่ตัวก็ได้ ซึ่งอาจจะได้ผลลัพธ์ **Ad**am, **A**elf**d**ane, **A**rnol**d**, B**a**r**d**en,  B**a**ir**d**  เป็นต้น

---

#### 8. is
💡 คือการ filter แบบ **null หรือ not null**

📢 **รูปแบบการใช้งาน :** `is:null` หรือ `is:notnull`

**หมายเหตุ :** `is` จะไม่มี value แบบอื่น จะเป็นได้แค่ `null` และ `notnull` เท่านั้น

---

#### 9. in
💡 คือการ filter แบบ **ข้อมูลที่ต้องการจะต้องอยู่ในชุดข้อมูลที่กำหนด**

📢 **รูปแบบการใช้งาน :** `in:value1+value2+...+valueN` เช่น `country=in:thailand+usa+canada` คือ การ filter ข้อมูลที่อยู่ในประเทศ thailand **หรือ** usa **หรือ** canada

---

### 2. Sort
โดยใช้ field ชื่อ `sort_by`

ตัวอย่างการใช้งาน
```
/employee?sort_by=firstname:asc,age:desc
```
ส่วนของ value ของ query string จะอยู่ในรูปแบบ **field-name1** : _**asc|desc**_ , ... , **field-nameN** : _**asc|desc**_


### 3. Limit
เราสามารถจำกัดจำนวณข้อมูลที่ต้องการได้ด้วย 2 fields นี้ คือ

1. `limit`
2. `page`

---

#### limit
💡 คือการจำกัดจำนวนของข้อมูลที่ต้องการไม่เกินจำนวนเต็มบวกที่กำหนด (จำนวนข้อมูลต่อ 1 หน้า)

<!-- 💡 สามารถกำหนดค่า default limit ได้ทาง Environment Variables `GOLIATH_QRY_COLLECTION_LIMIT` -->

📢 **รูปแบบการใช้งาน :** `limit=positive_integer` เช่น `limit=50`, โดยที่ positive_integer จะต้องเป็นจำนวนเต็ม ≥ 1

---

#### page
💡 คือการระบุลำดับของหน้าที่ต้องการ (ต้องระบุเป็น integer ที่ ≥ 1)

💡 goliath จะไม่สนใจค่า `page` ถ้าไม่ระบุค่า `limit`

📢 **รูปแบบการใช้งาน :** `page=positive_integer&limit=positive_integer`, โดยที่ positive_integer จะต้องเป็นจำนวนเต็ม ≥ 1

เช่น `?page=2&limit=50` จะได้รับข้อมูล คือ record ที่ 51 - N

---
