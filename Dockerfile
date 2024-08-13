FROM golang:1.21-alpine AS build

LABEL maintainer="Mohammad Reza Fadaei <mohrezfadaei@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:latest

ENV DEBUG=false
ENV ADDRESS=0.0.0.0
ENV PORT=8080
ENV DB_HOST=
ENV DB_PORT=
ENV DB_USER=
ENV DB_NAME=
ENV DB_PASSWORD=
ENV INFLUXDB_HOST=
ENV INFLUXDB_TOKEN=
ENV INFLUXDB_ORG=
ENV INFLUXDB_BUCKET=
ENV HEALTH_CHECK_INTERVAL=300

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=build /app/main .

COPY --from=build /app/internal/db/migrations /app/internal/db/migrations

EXPOSE 8080

CMD ["./main"]
