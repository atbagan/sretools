#FROM alpine:latest
#
#RUN apk add --no-cache git make musl-dev go
#
## Configure Go
#ENV GOROOT /usr/lib/go
#ENV GOPATH /go
#ENV PATH /go/bin:$PATH
#
#RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin ${GOPATH}/pkg ${GOPATH}/go/src/automation
#
#ADD . ${GOPATH}/go/src/automation
#
#WORKDIR ${GOPATH}/go/src/automation
#
#RUN ls
#RUN go version
#
#RUN go mod download
#RUN go get -d -t -v ./...
#
#
#ENTRYPOINT ["./automation"]
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

ENTRYPOINT [ "/app" ]
