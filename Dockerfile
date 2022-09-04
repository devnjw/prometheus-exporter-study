FROM golang as builder

WORKDIR /workspace

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .

FROM docker
WORKDIR /
COPY --from=builder /workspace/my_first_exporter /

CMD ["/my_first_exporter"]
