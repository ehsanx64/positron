NAME := positron
CMD_POSITRON := cmd/positron
PACKAGE := github.com/ehsanx64/$(NAME)

run:
	go run $(PACKAGE)/$(CMD_POSITRON)
deps:

dev-deps:
	go install github.com/air-verse/air@latest
