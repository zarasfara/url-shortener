FROM golang:1.22-alpine AS build-stage

RUN apk update && apk add --no-cache

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/url-shortener ./cmd/url-shortener/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/cleanup-scheduler ./cmd/cleanup-scheduler/main.go

FROM alpine AS run-stage

RUN apk update && apk add --no-cache supervisor

WORKDIR /app

COPY --from=build-stage /src/bin/url-shortener /app/url-shortener
COPY --from=build-stage /src/bin/cleanup-scheduler /app/cleanup-scheduler

COPY --from=build-stage /src/configs /app/configs
COPY --from=build-stage /src/.env /app/.env

COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf

CMD ["supervisord", "-c", "/etc/supervisor/conf.d/supervisord.conf"]
