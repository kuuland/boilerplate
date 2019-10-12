module github.com/kuuland/boilerplate

go 1.12

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.4.0
	github.com/jinzhu/gorm v1.9.11
	github.com/kuuland/kuu v0.0.0-20191011115836-527eb266368d
	github.com/satori/go.uuid v1.2.0
	gopkg.in/guregu/null.v3 v3.4.0
)

replace github.com/kuuland/kuu => ../../../github.com/kuuland/kuu
