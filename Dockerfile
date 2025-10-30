FROM golang:1.21-alpine as builder
ENV CGO_ENABLE=0
WORKDIR /app
COPY . .
RUN  go build -ldflags="-w -s" -o cep-clima main.go

FROM scratch
ENV PORT=8080
EXPOSE ${PORT}
COPY --from=builder /app/cep-clima /cep-clima
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
CMD ["/cep-clima"]