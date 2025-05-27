-- +goose Up
-- +goose StatementBegin
CREATE TABLE locations (
    id INT IDENTITY(1,1) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    type_id INT NOT NULL REFERENCES location_types(id), -- Ссылка на тип локации
    capacity INT NOT NULL DEFAULT 0, -- Вместимость (в кг или единицах)
    current_load INT NOT NULL DEFAULT 0 -- Текущая загрузка
);

INSERT INTO locations (name, type_id, capacity, current_load) VALUES
    ('Сушилка №1', 1, 82, 0),
    ('Сушилка №2', 1, 82, 0),
    ('Сушилка №3', 1, 82, 0),
    ('Сушилка №4 (3 этаж)', 1, 79, 0),
    ('Сушилка №5 (3 этаж)', 1, 72, 0),
    ('Сушилка №6 (3 этаж)', 1, 72, 0),
    ('323 камера', 1, 75, 0),
    ('314/1', 1, 47, 0),
    ('314/2', 1, 47, 0),
    ('Сушилка №1 (5 этаж)', 1, 78, 0),
    ('Сушилка №2 (5 этаж)', 1, 34, 0),
    ('Сушилка №3 (5 этаж)', 1, 66, 0),
    ('Агросклад', 2, 60, 0);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE locations;
-- +goose StatementEnd