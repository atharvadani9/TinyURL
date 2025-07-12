-- +goose Up
-- +goose StatementBegin
ALTER TABLE tinyurl 
ALTER COLUMN original_url VARCHAR(1024);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE tinyurl 
ALTER COLUMN original_url VARCHAR(255);
-- +goose StatementEnd