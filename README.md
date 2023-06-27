# TechnicalTest Golang Kredit Plus

Name : Aria shabry (aria.shabry@gmail.com)



**Setting Environtment**

1. Update your environtment at env.yaml file like DB_Host, DB_PORT, DB_NAME, and etc.
2. If your environtment error, please check env.yaml file and ConnectionString at ./helpers/env/env.go
3. Database Name = kreditplus
4. Import Database using kreditplus.sql file
5. Application address : http://localhost:5000/api/

**Answer**
Project ini hanya simulasi dan cuma memuat MVP


1. Insert data Konsumen
- api : POST/http://localhost:5000/api/konsumen
- request body : raw/json
- example :
```
{
    "NIK":"1922511504970002",
    "Gender": "Laki-laki",
    "Fullname": "Aria Shabry",
    "LegalName": "Aria Shabry",
    "TempatLahir": "Padang",
    "TanggalLahir": "15-04-1997",
    "Gaji":5000000,
    "FotoKTP":"http://kreditplus.com/storage/data/image/1922511504970002_300.jpg",
    "FotoSelfie":"http://kreditplus.com/storage/data/image/1922511504970002_600.jpg"
}
```
- screenshot hasil : ./result/1_Register_Konsumen.png
![Register_Konsumen](https://github.com/Ariashabry/KreditPlus/blob/main/results/1_%20Register_Konsumen.png?raw=true)



2. Ajukan Pinjaman
- api : POST/http://localhost:5000/api/pinjam
- request body : raw/json
- example
```
{
    "UserID": 1,
    "Amount":500000,
    "Tenor":5
}
```
- screenshot hasil : ./results/2_Ajukan_Pinjaman.png
![RequestPinjaman](https://github.com/Ariashabry/KreditPlus/blob/main/results/2_Ajukan_Pinjaman.png?raw=true)



3. Setujui Pinjaman
- api : PUT/http://localhost:5000/api/approve/:id
- :id merupakan id_pinjaman
- screenshot hasil : ./results/3_Setujui_Pinjaman.png
![SetujuiPinjaman](https://github.com/Ariashabry/KreditPlus/blob/main/results/3_Setujui_Pinjaman.png?raw=true)


4. See Status Konsumen
- api : GET/http://localhost:5000/api/konsumen/:id
- :id merupakan id_user
- screenshot hasil : ./results/4_See_Status_Konsumen.png
  ![SetujuiPinjaman](https://github.com/Ariashabry/KreditPlus/blob/main/results/4_See_Status_Konsumen.png?raw=true)

5. Penerapan Concurrecny
- Concurrency diterapkan pada api : http://localhost:5000/api//paymentloan/:id
- Dimana :id merupakan id pinjaman
- ./handlers/pinjamanHandler.go
- nama function : func (c *Context) PayMent(ctx echo.Context) error {}

6. Penerapan Unit Test
- Unit Test diterapkan pada function payment
- file : ./handlers/payment_test.go

7. Untuk ERD
![ERD](https://github.com/Ariashabry/KreditPlus/blob/main/results/erd.jpg?raw=true)


Jika ada pertanyaan, silahkan hubungi saya di : aria.shabry@gmail.com


