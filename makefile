
start:
	go run ./cmd/app

test:
	ginkgo ./...

build:
	go build -o ./bin ./cmd/app/main.go

swag-gen:
	swag init --parseDependency -d ./cmd/app,./internal/controllers

db-deploy:
	cd sqitch && sqitch deploy

db-revert:
	cd sqitch && sqitch revert

db-verify:
	cd sqitch && sqitch verify