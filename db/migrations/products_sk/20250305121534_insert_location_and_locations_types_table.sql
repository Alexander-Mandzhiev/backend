-- +goose Up
-- +goose StatementBegin
INSERT INTO location_types (name, description) VALUES -- Вставляем типы локаций
    ('сушилка', ''), -- dryer
    ('склад', ''); -- warehouse

INSERT INTO locations (name, type_id, capacity, current_load) VALUES -- Вставляем данные в таблицу locations
    ('Сушилка №1', 1, 82, 0),
    ('Сушилка №2', 1, 82, 0),
    ('Сушилка №3', 1, 82, 0),
    ('Сушилка №4', 1, 79, 0),
    ('Сушилка №5', 1, 72, 0),
    ('Сушилка №6', 1, 72, 0),
    ('323 камера', 1, 75, 0),
    ('314/1', 1, 47, 0),
    ('314/2', 1, 47, 0);
-- +goose StatementEnd
