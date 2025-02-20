-- +goose Up
-- +goose StatementBegin
IF NOT EXISTS (SELECT * FROM sys.tables WHERE name = 'apps' AND schema_name(schema_id) = 'dbo')
BEGIN
CREATE TABLE apps (
                      id INT IDENTITY(1,1) PRIMARY KEY,
                      name VARCHAR(255) NOT NULL,
                      secret VARCHAR(255) NOT NULL
);
PRINT 'Table "apps" has been created.';
END
ELSE
BEGIN
    PRINT 'Table "apps" already exists.';
END
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS apps;
-- +goose StatementEnd
