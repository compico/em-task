### Ссылки
Документация:
  - Swagger: [https://em-task.compico.su/swagger](https://em-task.compico.su/swagger)

### Комманды
Понял, вот краткое и лаконичное описание `make`-команд для вставки в `README.md`:

---

### ⚙️ Команды

```sh
make download_toolchain
```
Скачивает необходимые CLI-инструменты: `air`, `wire`, `swag`, `migrate`.

```sh
make init
```
Инициализирует окружение: копирует `.dist`-файлы конфигураций и вызывает генерацию зависимостей через `wire`.

```sh
make wire
```
Генерирует зависимости с помощью [Google Wire](https://github.com/google/wire) для всех подкоманд (`http`, `migrate`).

```sh
make up
```
Запускает контейнеры через `docker-compose` и выводит логи приложения.

```sh
make recreate
```
Полностью пересоздаёт контейнеры: `down`, `build`, `up`, `logs`.

```sh
make migrate_up
```
Запускает выполнение миграций внутри контейнера `app`.

```sh
make migration name=название
```
Создаёт новую миграцию с заданным именем (в `migrations/`).

```sh
make docs
```
Генерирует Swagger-документацию по маршрутам.


### Структура проекта
```
.
├── cmd                         # Точка входа в приложение
│     ├── commands                # CLI-команды (http-сервер, миграции)
│     │     ├── commands.go         # Регистрация команд
│     │     ├── http/               # HTTP-сервер как команда
│     │     │     ├── cmd.go
│     │     │     ├── wire.go         # Wire-описание зависимостей
│     │     │     └── wire_gen.go     # Сгенерировано Google Wire
│     │     └── migrate/            # Команда миграций
│     │         ├── cmd.go
│     │         ├── wire.go
│     │         └── wire_gen.go
│     ├── di/                     # DI-компоненты (логгер, сервер, конфиг)
│     │     ├── config.go
│     │     ├── http_server.go
│     │     ├── logger.go
│     │     └── wire.go
│     └── main.go                 # Точка входа при запуске бинарника
├── configs/
│     ├── config.yaml             # Основной конфиг
│     └── config.yaml.dist        # Шаблон конфига
├── docker/
│     └── go/Dockerfile           # Dockerfile для сборки Go-приложения
├── docker-compose.yml          # Docker-compose для запуска
├── docker-compose.yml.dist     # Шаблон docker-compose
├── docs/
│     ├── docs.go
│     ├── swagger.json
│     └── swagger.yaml            # Swagger-описание API
├── go.mod
├── go.sum
├── internal/
│     └── pkg/                    # Бизнес-логика
│         ├── config/             # Работа с конфигами
│         │     ├── config.go
│         │     ├── database.go
│         │     ├── http.go
│         │     └── slog.go
│         ├── entity/             # Доменные сущности
│         │     └── subscription.go
│         ├── pgrepo/             # Реализация репозиториев на PostgreSQL
│         │     └── subscription.go
│         ├── repository/         # Интерфейсы репозиториев
│         │     └── subscription.go
│         └── service/            # Сервисная логика
│             └── subscription.go
├── Makefile                    # Набор команд для разработки
├── migrations/                 # SQL-миграции
│     ├── 000001_create_subscriptions_table.down.sql
│     └── 000001_create_subscriptions_table.up.sql
├── pkg/
│     ├── logger/                 # Утилиты логирования
│     │     └── logger.go
│     └── postgres/               # Подключение к PostgreSQL
│         ├── connection.go
│         └── database.go
├── README.MD
└── web/                        # HTTP-слой
    ├── handlers/              # Обработчики маршрутов
    │     ├── errors.go
    │     ├── healthcheck.go
    │     └── subscriptions.go
    ├── middleware/            # Middleware (логгирование, JSON)
    │     ├── content_json.go
    │     └── set_logger.go
    ├── router/                # Роутинг
    │     └── routes.go
    └── server.go              # Запуск HTTP-сервера
```

### Создание тестовых подписок
```sql
CREATE OR REPLACE FUNCTION generate_fake_subscriptions()
    RETURNS void AS $$
DECLARE
    services TEXT[] := ARRAY[
        'Yandex Plus',
        'Kinopoisk',
        'Netflix',
        'Yandex Music',
        'Spotify',
        'IVI'
        ];
    num_users INT := 800;
    total_rows INT := 1000;
    user_ids UUID[];
    i INT;
    uid UUID;
    service TEXT;
    price INT;
    month INT;
    day DATE;
BEGIN
    -- Генерируем 800 уникальных UUID и сохраняем их в массив
    FOR i IN 1..num_users LOOP
            user_ids := array_append(user_ids, gen_random_uuid());
        END LOOP;

    -- Создаём 1000 случайных подписок
    FOR i IN 1..total_rows LOOP
            uid := user_ids[1 + floor(random() * array_length(user_ids, 1))::int];
            service := services[1 + floor(random() * array_length(services, 1))::int];

            -- Цена от 199 до 899
            price := 100 * (2 + floor(random() * 8))::int - 1;

            -- Случайный месяц в 2025
            month := 1 + floor(random() * 12)::int;
            day := make_date(2025, month, 1);

            INSERT INTO subscriptions (service_name, price, user_id, start_date)
            VALUES (service, price, uid, day);
        END LOOP;
END;
$$ LANGUAGE plpgsql;

CREATE EXTENSION IF NOT EXISTS pgcrypto;
SELECT generate_fake_subscriptions();
```