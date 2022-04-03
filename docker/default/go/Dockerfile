FROM golang:1.18 AS build-base
#FROM golang:1.18
# airとかdlv後で入れる

ENV WORKDIR /go/src

WORKDIR ${WORKDIR}

COPY go.mod .
COPY go.sum .

ADD . ${WORKDIR}

RUN GOOS=linux CGO_ENABLED=0 go build -o /bin/todo-app ./cmd

FROM scratch AS run-base

COPY --from=build-base /bin/todo-app /bin/todo-app
COPY --from=build-base /go/src/config /etc/config
EXPOSE 8080
ENTRYPOINT ["/bin/todo-app"]