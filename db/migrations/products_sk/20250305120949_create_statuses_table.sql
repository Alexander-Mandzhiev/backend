-- +goose Up
-- +goose StatementBegin
CREATE TABLE statuses (
    id INT IDENTITY(1,1) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT NULL
);

INSERT INTO statuses (name, description) VALUES
    ('Ожидает загрузки', 'Продукция ожидает помещения в сушильную камеру для начала процесса вызревания.'),
    ('В процессе вызревания', 'Продукция находится в сушильной камере и проходит процесс вызревания при контролируемых условиях температуры и влажности.'),
    ('Готовая продукция', 'Продукция успешно завершила процесс вызревания и готова к упаковке или отправке на следующий этап производства.');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE statuses;
-- +goose StatementEnd