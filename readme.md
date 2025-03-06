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
## Микросервисы
### SSO (Single Sign-On)
Сервис аутентификации и авторизации пользователей.

#### Пример запроса:
```yaml
curl -X POST http://localhost:6000/api/v1/auth/login
```
#### Флаг:
```yaml
CONFIG_PATH=config/sso/development.yaml
```
### Apps
Сервис управления приложениями.

#### Пример запроса:
```yaml
curl -X GET http://localhost:6000/api/v1/apps
```
#### Флаг:
```yaml
CONFIG_PATH=config/apps/development.yaml
```

### Locations
Сервис управления локациями (складами, точками хранения и т.д.).

#### Пример запроса:

```yaml 
curl -X GET http://localhost:6000/api/v1/locations
```
#### Флаг:
```yaml
CONFIG_PATH=config/locations/development.yaml
```
### Movements
Сервис управления перемещениями продукции между локациями.
#### Пример запроса:
```yaml
curl -X POST http://localhost:6000/api/v1/movements
```
#### Флаг:
```yaml
CONFIG_PATH=config/movements/development.yaml
```
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
### Statuses
Сервис управления общими статусами (например, статусы заказов, задач и т.д.).
#### Пример запроса:
```yaml
curl -X GET http://localhost:6000/api/v1/statuses
```
#### Флаг:
```yaml
CONFIG_PATH=config/statuses/development.yaml
```