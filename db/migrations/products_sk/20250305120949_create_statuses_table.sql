-- +goose Up
-- +goose StatementBegin
CREATE TABLE statuses (
    id INT IDENTITY(1,1) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE statuses;
-- +goose StatementEnd