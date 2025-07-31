-- +goose Up
ALTER TABLE services_per_consultation ADD COLUMN discount VARCHAR(64);
ALTER TABLE services_per_consultation ADD COLUMN quanity VARCHAR(64);
-- +goose Down
ALTER TABLE services_per_consultation DROP COLUMN discount;
ALTER TABLE services_per_consultation DROP COLUMN quanity;
