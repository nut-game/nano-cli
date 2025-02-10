build:
	@go build -o ./nano-cli .

build-linux:
	@GOOS=linux GOARCH=amd64 go build -o ./out/nano-cli-linux .
