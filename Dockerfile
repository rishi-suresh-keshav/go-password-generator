FROM golang:1.15.0 as build

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod download
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o go-password-generator cmd/go-password-generator/main.go


FROM alpine:latest AS runtime
WORKDIR /app
COPY --from=build /app/go-password-generator ./
EXPOSE 8000
CMD ["/app/go-password-generator"]
