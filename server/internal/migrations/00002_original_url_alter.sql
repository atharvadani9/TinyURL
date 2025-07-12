-- +goose Up
-- +goose StatementBegin
ALTER TABLE tinyurl
ALTER COLUMN original_url TYPE VARCHAR(1024);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE tinyurl
ALTER COLUMN original_url TYPE VARCHAR(255);
-- +goose StatementEnd