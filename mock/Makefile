.PHONY: mock

install-tools:
	go install github.com/golang/mock/mockgen@v1.5.0

mock:
	@mockgen --build_flags=--mod=mod -package service_test github.com/yoskeoka/go-example/mock/domain User,UserGroup > service/mockgen_user_test.go
