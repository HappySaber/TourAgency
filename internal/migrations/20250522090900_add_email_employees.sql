-- +goose Up
ALTER TABLE employees ADD COLUMN email VARCHAR(16);
-- +goose Down
ALTER TABLE employees DROP COLUMN email;