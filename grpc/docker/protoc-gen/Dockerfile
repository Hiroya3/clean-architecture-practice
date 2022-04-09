ARG GO_VERSION=1.18
ARG PROTOC_VERSION=3.15.8

FROM golang:${GO_VERSION}-buster

WORKDIR /home

RUN apt-get update && apt-get install -y --no-install-recommends \
		make \
		unzip

ARG PROTOC_VERSION
ARG PROTOC_ZIP=protoc-${PROTOC_VERSION}-linux-x86_64.zip

RUN curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/$PROTOC_ZIP && \
    unzip -o $PROTOC_ZIP -d /usr bin/protoc && \
    unzip -o $PROTOC_ZIP -d /usr 'include/*' && \
    rm -f $PROTOC_ZIP \

RUN go get github.com/golang/protobuf/protoc-gen-go

CMD ["protoc"]