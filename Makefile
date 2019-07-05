COMPILER=protoc
BASEDIR=.
SPEC_DIR=$(BASEDIR)/spec
SPEC_PROTO_DIR=$(SPEC_DIR)/proto
SPEC_SWAGGER_DIR=$(SPEC_DIR)/swagger

# Proto3 library files
THIRD_PARTY=$(BASEDIR)/third_party

GEN_DIR=gen
GEN_API_DIR=$(GEN_DIR)/api

generate-services:
	$(MAKE) -compile-proto


-compile-proto:
	$(COMPILER) -I=$(SPEC_PROTO_DIR) \
				-I=$(THIRD_PARTY) \
				--go_out=plugins=grpc:$(GEN_API_DIR) \
				$(SPEC_PROTO_DIR)/service.proto

-compile-rest-gateway: -compile-proto
	$(COMPILER) -I=$(SPEC_PROTO_DIR) \
				-I=$(THIRD_PARTY) \
				--grpc-gateway_out=logtostderr=true:$(GEN_API_DIR) \
				$(SPEC_PROTO_DIR)/service.proto

-generate-swagger: -compile-rest-gateway
	$(COMPILER) -I=$(SPEC_PROTO_DIR) \
	      		-I=$(THIRD_PARTY) \
	      		--swagger_out=logtostderr=true:$(SPEC_SWAGGER_DIR) \
	      		$(SPEC_PROTO_DIR)/service.proto

compile: -generate-swagger