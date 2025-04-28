# 🛠 Integration Service 🏄

Легкий и расширяемый интеграционный сервис на Go + Fiber + MongoDB. Позволяет без перекомпиляции и перезапуска добавлять новые проксируемые запросы через базу данных

>Lightweight and extensible integration service on Go + Fiber+ MongoDB. Allows you to add new proxied queries through the database without recompilation and restart

## 📖 Описание


**Integration Service** — это прокси-сервис, который позволяет управлять интеграциями с внешними системами через REST-запросы, храня информацию о запросах в MongoDB.

Особенности:
- Добавление новых запросов через базу данных без изменения кода
- Легкий старт и масштабирование
- Автоматическая инициализация тестового запроса
- Удобный логгер на Zap
- Минималистичный и понятный код

---

## 📖 Description

**Integration Service** — a proxy service that allows you to manage integrations with external systems via REST requests, storing all request data in MongoDB.

Features:
- Add new proxy requests without code redeploys
- Easy to start and scale
- Automatic database seeding with a test request
- Powerful logging with Zap
- Minimalistic and clean code

---

## 🚀 Технологии / Technologies

- [Golang 1.23+](https://golang.org/)
- [Fiber](https://gofiber.io/)
- [MongoDB](https://www.mongodb.com/)
- [Zap Logger](https://pkg.go.dev/go.uber.org/zap)

---

## 📂 Структура проекта / Project structure

```bash
integration_service/
├── cmd/                 # Точка входа в приложение (main.go)
├── config/              # Конфигурация проекта
├── internal/
│   ├── model/           # Описание моделей данных
│   ├── request/         # Логика отправки проксируемых запросов
│   ├── seed/            # Инициализация базы данных
│   └── server/          # HTTP сервер на Fiber
├── pkg/
│   └── logger/          # Обертка над Zap логгером
├── go.mod               # Go модули
└── README.md            # Документация
```

## 🛠 Как запустить
- Установить Go 1.23+

- Установить MongoDB (или запустить docker-compose) и запустить сервер

- Склонировать проект:
```bash
git clone https://github.com/yourusername/integration_service.git
cd integration_service
```

- Настроить .env

- Установить зависимости и запустить:
``` bash
go mod tidy
go run cmd/main.go
```

## 🛠 How to run
- Install Go 1.23+

- Install (or run docker-compose) and run MongoDB server

- Clone the project:
```bash
git clone https://github.com/yourusername/integration_service.git
cd integration_service
```

- Set up .env

- Install dependencies and run:
``` bash
go mod tidy
go run cmd/main.go
```

## 🎯 Пример запроса / Example request
```bash
curl -X POST http://localhost:3000/proxy/test_request_id
```

## ⭐ Обратная связь | Feedback
Если у вас есть предложения или вы нашли ошибки, создайте Issue или отправьте Pull Request. Если вам понравилось, то можете дать ⭐ этому репозиторию
> If you have suggestions or find any issues, feel free to open an Issue or submit a Pull Request! If you liked it, you can give ⭐ to this repository