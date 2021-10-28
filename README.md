
# How to Use
- Install Go & MySQL
- Create new .env file
```
$ export CONNECTION="{database_username}:{database_password}@tcp({database_host}:{database_port})/{database_name}?charset=utf8&&parseTime=True&loc=Local"
```
- Clone and Run 
```
$ git clone https://github.com/urnikrokhiyah/sagaraTest.git
$ source .env && go run main.go
```
# List of Endpoint

| Method | Endpoint | Description| Authorization
|:-----|:--------|:----------| :----------:|
| POST  | /register | Register a new user | No 
| POST | /login | Login existing user| No | No
|---|---|---|---|---|
| GET    | /products |Get list of all products | No 
| GET    | /products/{id} |Get product detail | No 
| POST | /products | create a new product| Yes
| PUT   | /products/{id} | Update existing product | Yes 
| DELETE| /products/{id} | Delete existing product | Yes 
| POST | /products/{id}/images | upload image of product| Yes

- Note: The RESTful API using Db.AutoMigrate
