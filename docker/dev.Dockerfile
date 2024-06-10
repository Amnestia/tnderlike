FROM golang:1.22 AS builder

COPY go.mod go.sum /go/src/github.com/amnestia/tnderlike/service/

WORKDIR /go/src/github.com/amnestia/tnderlike/service/
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' -a \
    -o /go/bin/main ./cmd/tnderlike/main.go


FROM scratch

ENV VESHTIA_ENV=dev

COPY --from=builder /go/bin/main /go/bin/main
COPY --from=builder /go/src/github.com/amnestia/tnderlike/service/cmd/tnderlike/config/server /etc/tnderlike/config/server

EXPOSE 80
ENTRYPOINT ["/go/bin/main"]

