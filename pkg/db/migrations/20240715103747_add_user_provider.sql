-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
ADD provider VARCHAR(100) NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
DROP COLUMN provider;
-- +goose StatementEnd
