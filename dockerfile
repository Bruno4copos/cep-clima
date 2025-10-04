FROM golang:1.21 as build
ENV CGO_ENABLE=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /app
COPY . .
RUN  go build -o cep-clima .
CMD ["./cep-clima"]

FROM scratch
WORKDIR /app
ENV WEATHER_API_KEY=19041265099a413dbfb183552253108
ENV PORT=8080
COPY --from=build /app/cep-clima .
ENTRYPOINT ["./cep-clima"]