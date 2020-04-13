FROM golang:1.12 as builder

WORKDIR /go/src/github.com/knative/docs/helloworld
COPY . .

ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -v -o atereko


FROM alpine
RUN apk add --no-cache ca-certificates

COPY --from=builder /go/src/github.com/knative/docs/helloworld/helloworld /atereko

CMD ["/atereko"]

