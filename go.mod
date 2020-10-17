module github.com/bondhan/go-webcounter

go 1.14

require (
	github.com/BurntSushi/toml v0.3.1 // indirect
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/guregu/null v4.0.0+incompatible
	github.com/jinzhu/gorm v1.9.16
	github.com/joho/godotenv v1.3.0
	github.com/lib/pq v1.8.0
	github.com/mattn/go-colorable v0.1.8
	github.com/onsi/ginkgo v1.14.2 // indirect
	github.com/onsi/gomega v1.10.3 // indirect
	github.com/sirupsen/logrus v1.7.0
	github.com/snowzach/rotatefilehook v0.0.0-20180327172521-2f64f265f58c
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
)

replace github.com/bondhan/go-webcounter => ./
