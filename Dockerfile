# Указываем базовый образ для стадии сборки
FROM golang:1.22-alpine AS build-stage

# Обновляем пакеты и устанавливаем необходимые зависимости
RUN apk update && apk add --no-cache

# Устанавливаем рабочую директорию
WORKDIR /src

# Копируем go.mod и go.sum и скачиваем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем все файлы проекта в рабочую директорию
COPY . .

# Собираем исполняемый файл
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/url-shortener ./cmd/url-shortener/main.go

# Указываем базовый образ для стадии выполнения
FROM alpine AS run-stage

# Обновляем пакеты и устанавливаем необходимые зависимости
RUN apk update && apk add --no-cache

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем исполняемый файл из стадии сборки
COPY --from=build-stage /src/bin/url-shortener /app/url-shortener
# Копируем конфигурационные файлы из стадии сборки
COPY --from=build-stage /src/configs /app/configs
# Копируем файл окружения из стадии сборки
COPY --from=build-stage /src/.env /app/.env

# Указываем команду для запуска приложения
CMD ["/app/url-shortener"]
