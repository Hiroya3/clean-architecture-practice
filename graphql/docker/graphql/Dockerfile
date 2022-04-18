FROM golang:1.18 AS build-base

RUN go install github.com/go-delve/delve/cmd/dlv@v1.8.0
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

ENV WORKDIR /go/src

WORKDIR ${WORKDIR}

ADD . ${WORKDIR}

RUN GOOS=linux CGO_ENABLED=0 go build -buildvcs=false -o /bin/graphql-app ./graph/server

FROM scratch AS run-base

COPY --from=build-base /bin/graphql-app /bin/graphql-app
COPY --from=build-base /go/src/config /etc/config

EXPOSE 8080
ENTRYPOINT ["/bin/graphql-app"]