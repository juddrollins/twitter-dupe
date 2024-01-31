.PHONY: build clean deploy deploy-function

build-all: build-register build-login build-authorizer build-user

build-register:
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/register cmd/register/main.go

build-login:
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/login cmd/login/main.go

build-authorizer:
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/authorizer cmd/authorizer/main.go

build-user:
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/user cmd/user/main.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose

deploy-register: clean build-register
	sls deploy function --function register --verbose

deploy-login: clean build-login
	sls deploy function --function login --verbose

deploy-authorizer: clean build-authorizer
	sls deploy function --function authorizerFunc --verbose

deploy-user: clean build-user
	sls deploy function --function user --verbose