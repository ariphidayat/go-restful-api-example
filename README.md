# Go RESTful API Example

## Prerequisite
* Installed Go
* Installed MySQL Database
* Create database with name `go_restful_api`

## Run Application

```bash
$ go main.go
```
## APIs Example

Register User

```bash
$ curl -X POST 'localhost:8080/users' \
-H 'Content-Type: application/json' \
-d '{
	"username":"[username here]",
	"email":"[email address here]",
	"password":"[password here]"
}'
```

Get All Users

```bash
$ curl -X GET 'localhost:8080/users'
```

Get User By Id

```bash
$ curl -X GET 'localhost:8080/users/[id here]'
```

Login

```bash
$ curl -X POST 'localhost:8080/login' \
-H 'Content-Type: application/json' \
-d'{
	"email":"[email address here]",
	"password":"[password here]"
}'
```

Update User

```bash
$ curl -X PUT 'localhost:8080/users/1' \
-H 'Authorization: bearer [jwt token here]' \
-H 'Content-Type: application/json' \
-d '{
    "username" : "[new username here]",
    "email" : "[new email here]",
    "password" : "[new password here]"
}'
```

Delete User

```bash
$ curl -X DELETE 'localhost:8080/users/1' \
-H 'Authorization: bearer [jwt token here]'
```