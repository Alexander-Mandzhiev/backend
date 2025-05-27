-- +goose Up
-- +goose StatementBegin
CREATE TABLE movements (
    id BIGINT PRIMARY KEY IDENTITY(1,1),
    product_id BIGINT NOT NULL,
    from_location_id INT NULL,
    to_location_id INT NULL,
    user_id INT NOT NULL,
    comment TEXT,
    created_at DATETIME DEFAULT GETDATE(),
    updated_at DATETIME DEFAULT GETDATE(),
    removed_at DATETIME NULL,
    FOREIGN KEY (from_location_id) REFERENCES locations(id),
    FOREIGN KEY (to_location_id) REFERENCES locations(id)
);

CREATE INDEX idx_movements_product_id ON movements(product_id);
CREATE INDEX idx_movements_from_location_id ON movements(from_location_id);
CREATE INDEX idx_movements_to_location_id ON movements(to_location_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE movements;
-- +goose StatementEnd