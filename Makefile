run-webhook:
	go run server/*.go

build-client:
	go build client/cli/main.go