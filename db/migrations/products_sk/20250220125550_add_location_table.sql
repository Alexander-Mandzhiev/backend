-- +goose Up
-- +goose StatementBegin
CREATE TABLE locations (
    id INT IDENTITY(1,1) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(50) NOT NULL, -- Например, "сушилка", "склад"
    capacity INT NOT NULL DEFAULT 0, -- Вместимость (в кг или единицах)
    current_load INT NOT NULL DEFAULT 0 -- Текущая загрузка
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE locations;
-- +goose StatementEnd