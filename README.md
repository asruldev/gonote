# Catatann Belajar
V.1.17.3

go get -u github.com/gofiber/fiber/v2
go get github.com/joho/godotenv
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql

go get github.com/google/uuid


Run with nodemon: >
`nodemon --exec go run main.go --signal SIGTERM`


### Example Route

```
GET api/user/:userId - Get User with userId
GET api/user - Get All Users
PUT api/user/:userId - Update User with userId
```