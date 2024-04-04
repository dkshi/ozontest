# OZONTEST: Тестовое задание для стажировки в Ozon Tech
---

<img src='https://fanibani.ru/wp-content/uploads/2021/07/milie004.jpg' width='350' height='300'>

## Описание
Учёл все условия из ТЗ!

gRPC сервис, реализующий методы:
- Сохранение оригинального URL в хранилище и возврат сокращённого
- Возвращение оригинального URL через сокращённый

Для выбора типа хранилища (In-memory/PostgreSQL) необходимо изменить параметр STORAGE в файле .env на postgres/inmemory соответственно.

Реализовал метод очистки от сокращённых URL срок действия которых истёк (по умолчанию 7 дней, настраивается в приложении ./internal/repository/repository.go).

## Компоненты
- Docker: [Installation Guide](https://docs.docker.com/get-docker/)
- Docker Compose: [Installation Guide](https://docs.docker.com/compose/install/)

## Установка
1. Клонируйте репозиторий:
    ```bash
    git clone https://github.com/dkshi/ozontest.git
    ```

2.Перейдите в директорию проекта:
    ```bash
    cd ozontest
    ```

3. По желанию измените `.env` файл, полобно данному примеру:
    ```plaintext
    PORT=50051
    TIME_LOCATION=UTC
    
    DB_HOST=db # Данный параметр лучше не трогать, ведь он ссылается к названию контейнера с БД
    DB_NAME=postgres
    DB_PORT=5432
    DB_USERNAME=postgres
    DB_PASSWORD=qwerty
    DB_SSLMODE=disable
    
    STORAGE=postgres # Отвечает за тип хранилища, можно менять перед запуском
    ```

## Usage
   Запустите приложение через Docker Compose:
    ```bash
    docker-compose up
    ```

