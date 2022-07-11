.PHONY: validate
validate:
	@protoc --validate_out=lang=go:./types *.proto
           

.PHONY: gen-rpc
gen-rpc:
	@goctl rpc protoc *.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.
	
.PHONY: gen-rpc-validate
gen-rpc-validate:
	@make gen-rpc
	@make validate

.PHONY: run
run:
	@go run user.go -f etc/user.yaml

.PHONY: test-create
test-create:
	@go test -v user_test.go -test.run TestCreate