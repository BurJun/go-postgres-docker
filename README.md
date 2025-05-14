# Go PostgreSQL Docker Project

Этот проект демонстрирует использование Go с PostgreSQL внутри Docker-контейнеров.

## Запуск с Docker

1. Клонировать репозиторий:
   ```bash
   git clone https://github.com/yourusername/go-postgres-docker.git
   cd go-postgres-docker
   ```
2. Соберать и запустить контейнеры с помощью Docker Compose:
   ```bash
   docker-compose up --build
   ```
4. Приложение будет доступно на http://localhost:8080, а PostgreSQL на порту 5432.
