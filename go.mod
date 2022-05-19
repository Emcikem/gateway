module gateway

go 1.16

require (
	github.com/gin-gonic/gin v1.7.1
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/joho/godotenv v1.3.0
	github.com/onsi/gomega v1.12.0 // indirect
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a
	golang.org/x/time v0.0.0-20220411224347-583f2d630306
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v8 v8.18.2
	gopkg.in/yaml.v2 v2.4.0
	gorm.io/driver/mysql v1.1.0
	gorm.io/gorm v1.21.10
)
