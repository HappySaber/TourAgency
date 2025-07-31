-- +goose Up
ALTER TABLE employees ADD COLUMN password VARCHAR(16);
-- +goose Down
ALTER TABLE employees DROP COLUMN password;