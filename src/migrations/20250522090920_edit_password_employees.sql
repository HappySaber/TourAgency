-- +goose Up
ALTER TABLE employees ALTER COLUMN password TYPE varchar(100);
-- +goose Down
ALTER TABLE employees ALTER COLUMN password TYPE varchar(16);