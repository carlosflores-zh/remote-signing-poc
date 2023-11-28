run-webhook:
	go run server/*.go

build-client:
	go build -o cli client/cli/main.go