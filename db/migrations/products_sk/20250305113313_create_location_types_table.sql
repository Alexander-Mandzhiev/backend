-- +goose Up
-- +goose StatementBegin
CREATE TABLE location_types (
    id INT IDENTITY(1,1) PRIMARY KEY,
    name VARCHAR(150) NOT NULL, -- Например, "сушилка", "склад"
    description TEXT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE location_types;
-- +goose StatementEnd