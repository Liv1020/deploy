FROM golang:1.12-alpine3.9 AS build
MAINTAINER pavle "lipengfei12@xin.com"

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /var/opt/

RUN mkdir -p /var/opt/code

ADD ./ /var/opt/code

RUN cd /var/opt/code \
    && CGO_ENABLED=0 go build -a -ldflags '-w -s' -mod=vendor

FROM alpine:latest
MAINTAINER pavle "lipengfei12@xin.com"

RUN apk add --no-cache curl ca-certificates

RUN mkdir -p /usr/local/code
WORKDIR /usr/local/code

COPY --from=build /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

COPY --from=build /var/opt/code/deploy .

CMD ["/usr/local/code/deploy"]
