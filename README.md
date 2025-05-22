<<<<<<< HEAD
# quote-service-test
=======

# 📚 Сервис "Цитатник" на Go

Простой REST API-сервис на Go для управления цитатами. Поддержка сохранения в `quotes.json`.

## 🚀 Запуск

```bash
go mod tidy
go run main.go
```

## 📋 Примеры curl-запросов

```bash
curl -X POST http://localhost:8080/quotes \
  -H "Content-Type: application/json" \
  -d '{"author":"Confucius", "quote":"Life is simple, but we insist on making it complicated."}'

curl http://localhost:8080/quotes
curl http://localhost:8080/quotes/random
curl http://localhost:8080/quotes?author=Confucius
curl -X DELETE http://localhost:8080/quotes/1
```

## ✅ Тесты

```bash
go test
```
>>>>>>> aeb90de (Цитатник на Go)
