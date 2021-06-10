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
			 --validate_out="lang=go:./api/anime/" \
			 ./api/anime/*.proto

.PHONY: userProto
userProto:
		protoc -I. -I$(GOPATH)/include \
			--go_out=plugins=grpc:./api/user/ \
			--grpc-gateway_out . --grpc-gateway_opt paths=source_relative \
			--grpc-gateway_opt logtostderr=true \
		    --validate_out="lang=go:./api/user/" \
			 ./api/user/*.proto

.PHONY: countProto
countProto:
	protoc -I. -I$(GOPATH)/include \
		--go_out=plugins=grpc:./api/count/ \
		--grpc-gateway_out . --grpc-gateway_opt paths=source_relative \
		 --grpc-gateway_opt logtostderr=true \
		 --validate_out="lang=go:./api/count/" \
		 ./api/count/*.proto
