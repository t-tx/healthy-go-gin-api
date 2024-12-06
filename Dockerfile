FROM golang:1.20 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOOS=linux go build -o /app/app
RUN chmod +x /app

FROM ubuntu

WORKDIR /

COPY --from=build-stage /app/app /app

COPY config.yml /config.yml
EXPOSE 8080


ENTRYPOINT ["/app"]
CMD ["healthy"]