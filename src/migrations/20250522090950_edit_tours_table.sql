-- +goose Up
ALTER TABLE tours ADD COLUMN price VARCHAR(64);
-- +goose Down
ALTER TABLE tours DROP COLUMN price;