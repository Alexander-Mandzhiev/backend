-- +goose Up
-- +goose StatementBegin
CREATE TABLE products_sk_statuses (
    product_id BIGINT NOT NULL,
    status_id INT NOT NULL,
    active BIT NOT NULL DEFAULT 1, -- 1 = активный статус, 0 = неактивный
    created_at DATETIME DEFAULT GETDATE(),
    PRIMARY KEY (product_id, status_id),
    FOREIGN KEY (product_id) REFERENCES products_sk(id) ON DELETE CASCADE, -- Исправлено: products -> products_sk
    FOREIGN KEY (status_id) REFERENCES statuses(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE products_sk_statuses;
-- +goose StatementEnd