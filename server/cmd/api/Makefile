generate:
	hz new -idl ./../../idl/http/$(service).thrift -mod github.com/dove-one/dove/server/cmd/api -f
	rm go.mod .gitignore
	rm build.sh
	rm -rf script

auth:
	make service=auth generate

user:
	make service=user generate

all:
	make auth
	make user
