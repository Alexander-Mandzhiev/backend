Список микросервисов
services:
sso_addr: "0.0.0.0:6100" # Адрес сервиса sso
apps_addr: "0.0.0.0:6110" # Адрес сервиса apps
product_sk_addr: "0.0.0.0:6200" # Адрес сервиса product_sk
locations_addr: "0.0.0.0:6210" # Адрес сервиса locations
statuses_addr: "0.0.0.0:6220" # Адрес сервиса statuses
movements_addr: "0.0.0.0:6230" # Адрес сервиса movements
products_sk_statuses_addr: "0.0.0.0:6240" # Адрес сервиса products_sk_statuses
production_tasks_addr: "0.0.0.0:6250" # Адрес сервиса production_tasks
Apps
Описание: Сервис управления приложениями.
Пример запроса:
curl -X GET http://localhost:6000/api/v1/apps/
Конфигурация:
Флаг: CONFIG_PATH=config/apps/development.yaml
Gateway
Описание: Единая точка входа для всех микросервисов. Обеспечивает RESTful API для клиентов и координирует работу с другими микросервисами через gRPC.
Пример запроса:
curl -X GET http://localhost:6000/api/v1/location/
Конфигурация:
Адрес: GATEWAY_ADDR=0.0.0.0:6000
Флаг: CONFIG_PATH=config/gateway/development.yaml
Locations
Описание: Сервис управления локациями (складами, точками хранения и т.д.).
Пример запроса:
curl -X GET http://localhost:6000/api/v1/location/
Конфигурация:
Флаг: CONFIG_PATH=config/locations/development.yaml
Movements
Описание: Сервис управления перемещениями продукции между локациями.
Пример запроса:
curl -X POST http://localhost:6000/api/v1/movements/
Конфигурация:
Флаг: CONFIG_PATH=config/movements/development.yaml
Production Task
Описание: Сервис управления производственными задачами.
Пример запроса:
curl -X GET http://localhost:6000/api/v1/production-task/
Конфигурация:
Флаг: CONFIG_PATH=config/production_task/development.yaml
Products SK
Описание: Сервис управления продукцией (например, детали, узлы).
Пример запроса:
curl -X POST http://localhost:6000/api/v1/products-sk/
Конфигурация:
Флаг: CONFIG_PATH=config/products_sk/development.yaml
Products SK Statuses
Описание: Сервис управления статусами продукции.
Пример запроса:
curl -X PUT http://localhost:6000/api/v1/products-sk-statuses/
Конфигурация:
Флаг: CONFIG_PATH=config/products_sk_statuses/development.yaml
SSO (Single Sign-On)
Описание: Сервис аутентификации и авторизации пользователей.
Пример запроса:
curl -X POST http://localhost:6000/api/v1/auth/login
Конфигурация:
Флаг: CONFIG_PATH=config/sso/development.yaml
Statuses
Описание: Сервис управления общими статусами (например, статусы заказов, задач и т.д.).
Пример запроса:
curl -X GET http://localhost:6000/api/v1/statuses/
Конфигурация:
Флаг: CONFIG_PATH=config/statuses/development.yaml
