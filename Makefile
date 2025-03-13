NAME := positron
CMD_POSITRON := cmd/positron
PACKAGE := github.com/ehsanx64/$(NAME)
TMP := ./cmd/tmp

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
	go get github.com/spf13/viper
	go get github.com/eclipse/paho.mqtt.golang
	go get github.com/k0kubun/pp/v3
	
dev-deps:
	GOPROXY=https://goproxy.io,direct
	go install github.com/air-verse/air@latest
