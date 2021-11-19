# Catatann Belajar
V.1.17.3

```
go get -u github.com/gofiber/fiber/v2
go get github.com/joho/godotenv
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql

go get github.com/google/uuid
go get golang.org/x/crypto/bcrypt
go get github.com/dgrijalva/jwt-go

go get -v github.com/swaggo/swag/cmd/swag
go get -u github.com/arsmn/fiber-swagger/v2
go get -u github.com/swaggo/swag/cmd/swag
go mod vendor -v
```

Run with nodemon: >
`nodemon --exec go run main.go --signal SIGTERM`

Run Swagger with `swag init -g main.go --output docs`

https://github.com/swaggo/swag#declarative-comments-format

### Example Route

```
GET api/user/:userId - Get User with userId
GET api/user - Get All Users
PUT api/user/:userId - Update User with userId
```
l
