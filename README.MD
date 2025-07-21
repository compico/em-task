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
│   ├── commands                # CLI-команды (http-сервер, миграции)
│   │   ├── commands.go         # Регистрация команд
│   │   ├── http/               # HTTP-сервер как команда
│   │   │   ├── cmd.go
│   │   │   ├── wire.go         # Wire-описание зависимостей
│   │   │   └── wire_gen.go     # Сгенерировано Google Wire
│   │   └── migrate/            # Команда миграций
│   │       ├── cmd.go
│   │       ├── wire.go
│   │       └── wire_gen.go
│   ├── di/                     # DI-компоненты (логгер, сервер, конфиг)
│   │   ├── config.go
│   │   ├── http_server.go
│   │   ├── logger.go
│   │   └── wire.go
│   └── main.go                 # Точка входа при запуске бинарника
├── configs/
│   └── config.yaml.dist        # Шаблон конфига
├── docker/
│   └── go/
│       ├── Dockerfile           # Dockerfile для сборки Go-приложения
        └── Dockerfile.prod      # DockerFile для prod сборки
├── docker-compose.yml.dist      # Шаблон docker-compose
├── docker-compose.prod.yml.dist # Шаблон docker-compose для prod окружения
├── docs/
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml            # Swagger-описание API
├── go.mod
├── go.sum
├── internal/
│   └── pkg/                    # Бизнес-логика
│       ├── config/             # Работа с конфигами
│       │   ├── config.go
│       │   ├── database.go
│       │   ├── http.go
│       │   └── slog.go
│       ├── entity/             # Доменные сущности
│       │   └── subscription.go
│       ├── filter/             # Фильтры для обработки HTTP input'ов
│       │   └── subscription.go
│       ├── pgrepo/             # Реализация репозиториев на PostgreSQL
│       │   └── subscription.go
│       ├── query/              # Сущности для формирования запросов в БД
│       │   └── subscription.go
│       ├── repository/         # Интерфейсы репозиториев
│       │   └── subscription.go
│       └── service/            # Сервисная логика
│           └── subscription.go
├── Makefile                    # Набор команд для разработки
├── migrations/                 # SQL-миграции
│   ├── 000001_create_subscriptions_table.down.sql
│   └── 000001_create_subscriptions_table.up.sql
├── pkg/
│   ├── logger/                 # Утилиты логирования
│   │   └── logger.go
│   └── postgres/               # Подключение к PostgreSQL
│       ├── connection.go
│       └── database.go
├── README.MD
└── web/                    # HTTP-слой
├── handlers/               # Обработчики маршрутов
│   ├── create_subscription.go
│   ├── delete_subscription.go
│   ├── errors.go
│   ├── healthcheck.go
│   ├── list_subscription.go
│   ├── read_subscription.go
│   ├── subscriptions.go
│   ├── sum_subscription.go
│   └── update_subscription.go
├── middleware/             # Middleware
│   ├── json_response.go
│   └── with_logger.go
├── request/                # Структуры для HTTP запросов
│   ├── create_subscription_request.go
│   └── update_subscription_request.go
├── router/                 # Роутинг
│   └── routes.go
└── server.go               # HTTP-сервер
```
