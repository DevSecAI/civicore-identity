# CIV-IAC-006: runs as root.
FROM golang:1.22
WORKDIR /src
COPY . .
RUN go build -o /out/civicore ./cmd/civicore
EXPOSE 8080
CMD ["/out/civicore"]
