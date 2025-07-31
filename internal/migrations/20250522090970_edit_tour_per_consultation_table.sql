-- +goose Up
ALTER TABLE tours_per_consultation ADD COLUMN discount VARCHAR(64);
ALTER TABLE tours_per_consultation ADD COLUMN quanity VARCHAR(64);
-- +goose Down
ALTER TABLE tours_per_consultation DROP COLUMN discount;
ALTER TABLE tours_per_consultation DROP COLUMN quanity;
