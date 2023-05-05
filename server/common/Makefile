generate:
	kitex -module github.com/dove-one/dove ./../idl/rpc/$(service).thrift

auth:
	make service=auth generate
user:
	make service=user generate

all:
	make auth
	make user
