FROM golang:1.13.9 AS builder
ENV jws_version 0.1
RUN curl -LO https://github.com/nopp/jira-webhook-slack/archive/${jws_version}.tar.gz \
    && tar zxf ${jws_version}.tar.gz
WORKDIR jira-webhook-slack-${jws_version}
RUN go get -d . \
    && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM golang:1.13.9

LABEL maintainer="Carlos Augusto Malucelli <camalucelli@gmail.com>"
LABEL version="0.1"

ENV jws_version 0.1

COPY --from=builder /go/jira-webhook-slack-${jws_version}/main .
ENTRYPOINT ["./run.sh"]
