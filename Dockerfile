FROM golang:1.23.2

WORKDIR /app

# Копируем файлы зависимостей
COPY go.mod go.sum ./

# Скачиваем зависимости
RUN go mod download

# Копируем весь проект
COPY . .

# Компилируем приложение
RUN go build -o app ./cmd/api

# Команда для запуска приложения
CMD ["./app"]