.PHONY: build clean deploy deploy-function

build:
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/register cmd/register/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/user cmd/user/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/login cmd/login/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/authorizer cmd/authorizer/main.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose

deploy-register: clean build
	sls deploy function --function register --verbose

deploy-login: clean build
	sls deploy function --function login --verbose

deploy-authorizer: clean build
	sls deploy function --function authorizerFunc --verbose

deploy-user: clean build
	sls deploy function --function user --verbose