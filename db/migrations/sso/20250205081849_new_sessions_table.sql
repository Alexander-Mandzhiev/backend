-- +goose Up
-- +goose StatementBegin
IF NOT EXISTS (SELECT * FROM sys.tables WHERE name = 'sessions' AND schema_name(schema_id) = 'dbo')
BEGIN
    CREATE TABLE sessions (
        jti NVARCHAR(255) PRIMARY KEY,
        user_id NVARCHAR(255) NOT NULL,
        app_id INT NOT NULL,
        is_active BIT NOT NULL DEFAULT 1,
        created_at DATETIME NOT NULL DEFAULT GETDATE(),
        expires_at DATETIME NOT NULL
    );
END
ELSE
BEGIN
    PRINT 'Table "sessions" already exists.';
END
CREATE INDEX idx_sessions_user_id ON sessions (user_id);
CREATE INDEX idx_sessions_app_id ON sessions (app_id);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS sessions;
-- +goose StatementEnd
