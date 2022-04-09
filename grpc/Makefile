init:
	docker build -f ./docker/protoc-gen/Dockerfile -t grpc-sample-protoc-gen .

.PHONY: gen-proto
gen-proto:
	docker run --rm \
    		-v ${PWD}:/home \
    		pos-nec-protoc-gen \
    		protoc \
    			--go_out=plugins=grpc\:gen \
    			--go_opt=module=github.com/Hiroya3/clean-architecture-practice/grpc/gen \
    			-I . \
    			-I /go/src \
    			--proto_path=proto/ \
    			todo.proto
