TEST_NAME?=TestVpcAllExample
TIMEOUT?=3m

show:
	@echo 'make lint - run go init'
	@echo 'make tidy - run go tidy'
	@echo 'make lint - run linter'
	@echo 'make test - run basic test '

init:
	@go init $(TEST_NAME)

tidy: 
	@go mod tidy

lint:
	@golint *.go

test:
	@go test -v -timeout $(TIMEOUT) *.go
	