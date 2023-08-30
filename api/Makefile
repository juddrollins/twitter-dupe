.PHONY: build clean deploy deploy-function

build:
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/register cmd/register/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/user cmd/user/main.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose

deploy-function: clean build
	sls deploy function --function register --verbose