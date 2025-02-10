build:
	@CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w" -o ./nano-cli .

build-darwin-amd:
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o ./nano-cli-darwin-amd64 .

build-linux:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./out/nano-cli-linux .

build-windows:
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./out/nano-cli.exe .
