-- +goose Up
-- +goose StatementBegin
CREATE TABLE products_sk (
    id BIGINT PRIMARY KEY,
    part_name VARCHAR(255) NOT NULL,
    nomenclature VARCHAR(255) NOT NULL,
    number_frame INT NOT NULL,
    weight_sp_kg FLOAT NOT NULL,
    weight_gp_kg FLOAT NULL,
    manufacturing_date DATETIME NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE products_sk;
-- +goose StatementEnd