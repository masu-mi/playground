FROM alpine:3.11

RUN apk add --no-cache --update ca-certificates

COPY dist/linux-amd64/go-yacc /go-yacc

ENTRYPOINT ["/go-yacc"]
