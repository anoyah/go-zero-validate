.PHONY: validate
validate:
	@protoc --proto_path=. \
           --go_out=./types \
		   --proto_path=./third_party \
           --validate_out=lang=go:./types \
           user.proto

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