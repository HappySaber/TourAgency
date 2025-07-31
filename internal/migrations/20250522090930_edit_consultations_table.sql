-- +goose Up
ALTER TABLE consultations
RENAME COLUMN dateof TO timeofconsultation;

ALTER TABLE consultations
ALTER COLUMN timeofconsultation TYPE TIME USING (timeofconsultation::TIMESTAMP)::TIME;
-- +goose Down
ALTER TABLE consultations ALTER COLUMN dateof TYPE DATE;
