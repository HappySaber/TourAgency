-- +goose Up
ALTER TABLE clients ADD COLUMN passport VARCHAR(64);
ALTER TABLE consultations ADD COLUMN notes VARCHAR(512);
-- +goose Down
ALTER TABLE consultations DROP COLUMN notes;
ALTER TABLE clients DROP COLUMN passport;
