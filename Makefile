run-app:
	go run main.go

run-container:
	docker run -e DB_URL=$(DB_URL) -e SIGNING_KEY=$(SIGNING_KEY) accelerator-app/inventory:v1

build:
	docker build -t accelerator-app/inventory:v1 .

test:
	go test -failfast -v ./...

# test single function in a repo
# usage: make test-repo-fxn fxn=TestProductRepository_Get
test-repo-fxn:
	cd repository && go test -failfast -v -run $(fxn)
