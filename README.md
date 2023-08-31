## Avito_app DevOps
Приложение - телефонная книга
-----
БД Redis содержит пары "key"-"val", где key - имена, val - номера телефонов
- в контейнерах базовый образ Debian
- Docker Compose version > 3.3
- Golang и Redis
- Redis поддерживает аутентификацию (super_secret_password)
- Redis работает через TLS соединение
