FROM golang:latest as builder
ENV CGO_ENABLE=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV PORT=8080
WORKDIR /app
COPY . .
RUN  go build -ldflags="-w -s" -o cep-clima main.go
CMD ["./cep-clima"]

FROM scratch
ENV PORT=8080
ENV WEATHER_API_KEY=19041265099a413dbfb183552253108
WORKDIR /app
COPY --from=builder /app/cep-clima .
ENTRYPOINT ["./cep-clima"]