NAME := positron
CMD_POSITRON := cmd/positron
PACKAGE := github.com/ehsanx64/$(NAME)
TMP := ./tmp

run:
	go run $(PACKAGE)/$(CMD_POSITRON)

build:
	go build -o $(TMP)/$(NAME) $(CMD_POSITRON)/main.go

dev:
	air

tidy:
	GOPROXY=https://goproxy.io,direct
	go mod tidy -v

deps:
	GOPROXY=https://goproxy.io,direct
	go get github.com/labstack/echo/v4
	go get github.com/spf13/viper
	go get -u gorm.io/gorm
	go get -u gorm.io/driver/sqlite

dev-deps:
	GOPROXY=https://goproxy.io,direct
	go install github.com/air-verse/air@latest
