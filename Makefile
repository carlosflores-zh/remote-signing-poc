run-webhook:
	go run server/*.go

build-client:
	go build -o client client/cli/main.go