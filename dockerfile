FROM golang:1.21 as build
WORKDIR /app
COPY . .
RUN  CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go mod init cep-clima && go mod tidy && go build -o cep-clima .
CMD ["./cep-clima"]

FROM scratch
WORKDIR /app
ENV WEATHER_API_KEY="19041265099a413dbfb183552253108"
COPY --from=build /app/cep-clima .
ENTRYPOINT ["./cep-clima"]