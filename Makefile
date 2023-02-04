.PHONY: download_swag
download_swag:
	go install github.com/swaggo/swag/cmd/swag@latest

.PHONY: swagger
swagger:
	swag init -g ./cmd/main.go -o ./docs

.PHONY: fmt_swag
fmt_swag:
	swag fmt -g ./docs