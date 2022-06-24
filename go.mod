module github.com/kitabisa/golangproject

go 1.13

require (
	github.com/DataDog/datadog-go v4.2.0+incompatible
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/go-redis/redis/v8 v8.4.11
	github.com/go-sql-driver/mysql v1.4.1
	github.com/gobuffalo/packr v1.30.1 // indirect
	github.com/gorilla/mux v1.7.3
	github.com/jinzhu/gorm v1.9.11
	github.com/joho/godotenv v1.3.0
	github.com/kitabisa/perkakas/v2 v2.31.0
	github.com/lib/pq v1.2.0
	github.com/prometheus/client_golang v0.9.3
	github.com/rabbitmq/amqp091-go v1.2.0 // indirect
	github.com/rs/zerolog v1.19.0
	github.com/rubenv/sql-migrate v0.0.0-20191025130928-9355dd04f4b3
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.4.0
	github.com/ziutek/mymysql v1.5.4 // indirect
	gopkg.in/gorp.v1 v1.7.2 // indirect
	gopkg.in/gorp.v3 v3.0.2
)

replace gopkg.in/gorp.v3 => github.com/go-gorp/gorp/v3 v3.0.2
