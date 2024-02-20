.PHONY: mock
mock:
	@mockgen  mockgen -source=../service/type.go -package=svcmocks -destination=../service/mock/type.mock.gen.go
	@go mod tidy
	#可以直接读取文件写到etcd
	@#etcdctl --endpoint=127.0.0.1:12379 put /webook/config "$(<dev.yaml>)"
	#考虑生成后要移动一下位置
	@mockgen protoc --go_out=paths=source_relative --go_out=.  --go-grpc_opt=paths=source_relative --go-grpc_out=. webook.proto