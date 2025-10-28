FROM golang:latest as builder
ENV CGO_ENABLE=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /app
COPY . .
RUN  go build -ldflags="-w -s" -o cep-clima main.go
CMD ["./cep-clima"]

FROM scratch
WORKDIR /app
COPY --from=builder /app/cep-clima .
ENTRYPOINT ["./cep-clima"]