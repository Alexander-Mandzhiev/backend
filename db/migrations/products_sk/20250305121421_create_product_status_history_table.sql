-- +goose Up
-- +goose StatementBegin
CREATE TABLE product_status_history (
    product_id BIGINT NOT NULL,
    status_id INT NOT NULL,
    active BIT NOT NULL DEFAULT 1, -- 1 = активный статус, 0 = неактивный
    created_at DATETIME DEFAULT GETDATE(),
    PRIMARY KEY (product_id, status_id),
    FOREIGN KEY (product_id) REFERENCES products_sk(id) ON DELETE CASCADE,
    FOREIGN KEY (status_id) REFERENCES statuses(id) ON DELETE CASCADE
);

CREATE INDEX idx_product_status_history_product_id ON product_status_history(product_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE product_status_history;
-- +goose StatementEnd