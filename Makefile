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
	go mod tidy -v

deps:
	go get github.com/labstack/echo/v4

dev-deps:
	go install github.com/air-verse/air@latest
