FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

EXPOSE 8000
ADD go-password-generator /
CMD ["/go-password-generator"]