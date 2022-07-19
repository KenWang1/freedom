package project

func init() {
	content["/go.mod"] = modTemplate()
}

func modTemplate() string {
	return `
module {{.PackagePath}}

go 1.18

require (
	github.com/KenWang1/freedom {{.VersionNum}}
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/iris-contrib/swagger/v12 v12.0.1
	github.com/kataras/iris/v12 v12.2.0-beta3.0.20220623200152-8bfea48cd652
	gopkg.in/go-playground/validator.v9 v9.31.0
	gorm.io/driver/mysql v1.3.5
	gorm.io/gorm v1.23.8
)

`
}
