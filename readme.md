# Список микросервисов
## Основные адреса сервисов
### В случае разворачивания на локальной машине
```yaml
services:
  sso_addr: "0.0.0.0:6100" # Адрес сервиса SSO (Single Sign-On)
  apps_addr: "0.0.0.0:6110" # Адрес сервиса Apps
  product_sk_addr: "0.0.0.0:6200" # Адрес сервиса Products SK
  locations_addr: "0.0.0.0:6210" # Адрес сервиса Locations
  statuses_addr: "0.0.0.0:6220" # Адрес сервиса Statuses
  movements_addr: "0.0.0.0:6230" # Адрес сервиса Movements
  products_sk_statuses_addr: "0.0.0.0:6240" # Адрес сервиса Products SK Statuses
  production_tasks_addr: "0.0.0.0:6250" # Адрес сервиса Production Tasks
  location_types_addr: "0.0.0.0:6260" # Адрес сервиса Location Types
```
### В случае запуска в Docker контейнерах:
```yaml
services:
  sso_addr: "sso:6100" # Адрес сервиса SSO
  apps_addr: "apps:6110" # Адрес сервиса Apps
  product_sk_addr: "products_sk:6200" # Адрес сервиса Product SK
  locations_addr: "locations:6210" # Адрес сервиса Locations
  statuses_addr: "statuses:6220" # Адрес сервиса Statuses
  movements_addr: "movements:6230" # Адрес сервиса Movements
  products_sk_statuses_addr: "products_sk_statuses:6240" # Адрес сервиса Products SK Statuses
  production_tasks_addr: "production_task:6250" # Адрес сервиса Production Tasks
  location_types_addr: "location_types:6260" # Адрес сервиса Location Types
```
--- 
## Gateway
### Описание
Единая точка входа для всех микросервисов. Обеспечивает RESTful API для клиентов и координирует работу с другими микросервисами через gRPC.
#### Конфигурация
```yaml
env: "production" # Режим работы (development или production)
address: "0.0.0.0" # Адрес сервера
port: 6000 # Порт сервера
timeout: "5s" # Таймаут запроса
idle_timeout: "60s" # Idle таймаут
services:
  sso_addr: "sso:6100" # Адрес сервиса SSO
  apps_addr: "apps:6110" # Адрес сервиса Apps
  product_sk_addr: "products_sk:6200" # Адрес сервиса Product SK
  locations_addr: "locations:6210" # Адрес сервиса Locations
  statuses_addr: "statuses:6220" # Адрес сервиса Statuses
  movements_addr: "movements:6230" # Адрес сервиса Movements
  products_sk_statuses_addr: "products_sk_statuses:6240" # Адрес сервиса Products SK Statuses
  production_tasks_addr: "production_task:6250" # Адрес сервиса Production Tasks
  location_types_addr: "location_types:6260" # Адрес сервиса Location Types
frontend:
  addr: "0.0.0.0:6010" # Адрес фронтенд-клиента
```
#### Флаг: 
```yaml
CONFIG_PATH=config/gateway/development.yaml
```
---
## Микросервисы
### SSO (Single Sign-On)
Сервис аутентификации и авторизации пользователей.

#### Пример запроса:
##### Вход в систему:
```bash
curl -X POST http://localhost:6000/api/v1/login \
  -H "Content-Type: application/json" \
  -d '{
    "password": 123456789,
    "app_id": 1
  }'
```
#### Флаг:
```yaml
CONFIG_PATH=config/sso/development.yaml
```
---
### Apps
Сервис управления приложениями.

#### Примеры запросов:
##### Получить список приложений:
```yaml
curl -X GET http://localhost:6000/api/v1/apps
```
##### Создать приложение:
```bash
curl -X POST http://localhost:6000/api/v1/apps \
  -H "Content-Type: application/json" \
  -d '{
    "name": "test",
    "secret": "test"
  }'
```
##### Обновить приложение:
```bash
curl -X PUT http://localhost:6000/api/v1/apps/ \
  -H "Content-Type: application/json" \
  -d '{
    "id": 5,
    "name": "test upd",
    "secret": "test upd"
  }'
```
##### Получить информацию о приложении:
```bash
curl -X GET http://localhost:6000/api/v1/apps/5
```
##### Удалить приложение:
```bash
curl -X DELETE http://localhost:6000/api/v1/apps/6
```

#### Флаг:
```yaml
CONFIG_PATH=config/apps/development.yaml
```
--- 
### Locations
Сервис управления локациями (складами, точками хранения и т.д.).

#### Примеры запроса:
##### Получить список локаций:
```bash 
curl -X GET http://localhost:6000/api/v1/locations
```
##### Получить информацию о локации:
```bash 
curl -X GET http://localhost:6000/api/v1/locations/1
```
##### Создать локацию:
```bash 
curl -X POST http://localhost:6000/api/v1/locations \
  -H "Content-Type: application/json" \
  -d '{
    "name": "test 1",
    "type_id": 1,
    "capacity": 82
  }'
```
##### Обновить локацию:
```bash 
curl -X PUT http://localhost:6000/api/v1/locations/13 \
  -H "Content-Type: application/json" \
  -d '{
    "id": 13,
    "name": "test 1 upd",
    "type_id": 2,
    "capacity": 82
  }'
```
##### Удалить локацию:
```bash 
curl -X DELETE http://localhost:6000/api/v1/locations/13
```

#### Флаг:
```yaml
CONFIG_PATH=config/locations/development.yaml
```
---
### Location Types
Сервис управления типами локаций.
#### Примеры запросов:
##### Получить список типов локаций:
```bash
curl -X GET http://localhost:6000/api/v1/location-types
```
##### Получить информацию о типе локации:
```bash
curl -X GET http://localhost:6000/api/v1/location-types/1
```
##### Создать тип локации:
```bash
curl -X POST http://localhost:6000/api/v1/location-types \
  -H "Content-Type: application/json" \
  -d '{
    "name": "test",
    "description": "test"
  }'
```
##### Обновить тип локации:
```bash
curl -X PUT http://localhost:6000/api/v1/location-types/4 \
  -H "Content-Type: application/json" \
  -d '{
    "id": 4,
    "name": "test upd",
    "description": "test upd"
  }'
```
##### Удалить тип локации:
```bash
curl -X DELETE http://localhost:6000/api/v1/location-types/4
```

##### Флаг:
```bash
CONFIG_PATH=config/location_types/development.yaml
```
---
### Statuses
Сервис управления общими статусами (например, статусы заказов, задач и т.д.).
#### Примеры запросов:
##### Получить список статусов:
```yaml
curl -X GET http://localhost:6000/api/v1/statuses
```
##### Получить информацию о статусе:
```yaml
curl -X GET http://localhost:6000/api/v1/statuses/1
```
##### Создать статус:
```bash
curl -X POST http://localhost:6000/api/v1/statuses \
-H "Content-Type: application/json" \
  -d '{
"name": "test 2",
"description": "test 2."
}'
```
##### Обновить статус:
```bash
curl -X PUT http://localhost:6000/api/v1/statuses/8 \
  -H "Content-Type: application/json" \
  -d '{
    "id": 8,
    "name": "test.",
    "description": "test."
  }'
```
##### Удалить статус:
```bash
curl -X DELETE http://localhost:6000/api/v1/statuses/7
```
#### Флаг:
```yaml
CONFIG_PATH=config/statuses/development.yaml
```
---
### Movements
Сервис управления перемещениями продукции между локациями.
#### Примеры запроса:
##### Получить перемещения:
```yaml
curl -X POST http://localhost:6000/api/v1/movements
```
##### Получить информацию о перемещение:
```yaml
curl -X POST http://localhost:6000/api/v1/movements/1
```
##### Создать перемещение:
```bash
curl -X POST http://localhost:6000/api/v1/movements \
  -H "Content-Type: application/json" \
  -d '{
    "product_id": "123",
    "from_location_id": 1,
    "to_location_id": 4,
    "user_id": 101,
    "comment": "Перемещение продукции в сушильную камеру"
  }'
 ```

#### Флаг:
```yaml
CONFIG_PATH=config/movements/development.yaml
```
---
### Production Task
Сервис управления производственными задачами.
#### Пример запроса:
```yaml
curl -X GET http://localhost:6000/api/v1/production-task
```
#### Флаг:
```yaml
CONFIG_PATH=config/production_task/development.yaml
```
### Products SK
Сервис управления продукцией (например, детали, узлы).
#### Пример запроса:
```yaml
curl -X POST http://localhost:6000/api/v1/products-sk
```
#### Флаг:
```yaml
CONFIG_PATH=config/products_sk/development.yaml
```
### Products SK Statuses
Описание: Сервис управления статусами продукции.
#### Пример запроса:
```yaml
curl -X PUT http://localhost:6000/api/v1/products-sk-statuses
```
#### Флаг:
```yaml
CONFIG_PATH=config/products_sk_statuses/development.yaml
```

---
## Схема таблиц

![Схема таблиц](sk.drawio.svg)