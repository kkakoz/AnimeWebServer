GOPATH:=$(shell go env GOPATH)


.PHONY: baseProto
baseProto:
		protoc -I. -I$(GOPATH)/include \
			--go_out=plugins=grpc:./api/base/ \
			 ./api/base/*.proto

.PHONY: animeProto
animeProto:
		protoc -I. -I$(GOPATH)/include \
			--go_out=plugins=grpc:./api/anime/ \
			--grpc-gateway_out . --grpc-gateway_opt paths=source_relative \
			 --grpc-gateway_opt logtostderr=true \
			 ./api/anime/*.proto

.PHONY: userProto
userProto:
		protoc -I. -I$(GOPATH)/include \
			--go_out=plugins=grpc:./api/user/ \
			--grpc-gateway_out . --grpc-gateway_opt paths=source_relative \
			 --grpc-gateway_opt logtostderr=true \
			 ./api/user/*.proto

.PHONY: countProto
countProto:
	protoc -I. -I$(GOPATH)/include \
		--go_out=plugins=grpc:./api/count/ \
		--grpc-gateway_out . --grpc-gateway_opt paths=source_relative \
		 --grpc-gateway_opt logtostderr=true \
		 ./api/count/*.proto
#       --go_out . --go_opt paths=source_relative \
#		-
#		--go_out=plugins=grpc:. ./api/video/v1/hello.proto
#protoc protoc --go_out=plugins=grpc:. ./api/video/v1/hello.proto
#	protoc -I$(GOPATH)/include \
#		--go_out . --go_opt paths=source_relative \
#		--go-grpc_out . --go-grpc_opt paths=source_relative \
#		--grpc-gateway_out . --grpc-gateway_opt paths=source_relative \
#		--grpc-gateway_opt logtostderr=true \
#        --grpc-gateway_opt generate_unbound_methods=true \
#        --grpc-gateway_opt register_func_suffix=GW \
#        --grpc-gateway_opt allow_delete_body=true \
#        --openapiv2_out . --openapiv2_opt logtostderr=true \
#		api/video/v1/hello.proto
