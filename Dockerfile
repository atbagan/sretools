FROM golang:1.16-alpine as build

WORKDIR /build
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -o app

FROM alpine:latest as certs
RUN apk --update add ca-certificates

FROM scratch
COPY --from=build /build/app /
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY .sretools.yml ./
ENTRYPOINT [ "/app" ]
