# workspace (GOPATH) configured at /go
FROM golang:1.18 as builder


#
RUN mkdir -p $GOPATH/src/vendoo/go_http_server
WORKDIR $GOPATH/src/vendoo/go_http_server

# Copy the local package files to the container's workspace.
COPY . ./

# installing depends and build
RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    go mod vendor && \
    make build && \
    mv ./bin/go_http_server /



FROM alpine
COPY --from=builder go_http_server .
RUN mkdir config
RUN touch .env
RUN apk update && apk add -U tzdata && cp /usr/share/zoneinfo/Asia/Tashkent /etc/localtime && apk del tzdata
#COPY ./config/rbac_model.conf ./config/rbac_model.conf
#COPY ./config/AuthKey_5RAX23V6QP.p8 ./config/AuthKey_5RAX23V6QP.p8
ENTRYPOINT ["/go_http_server"]
