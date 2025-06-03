-- +goose Up
ALTER TABLE consultations
ALTER COLUMN timeofconsultation TYPE time USING timeofconsultation::time;
-- +goose Down
ALTER TABLE consultations
RENAME COLUMN dateof TO timeofconsultation;

ALTER TABLE consultations
ALTER COLUMN timeofconsultation TYPE TIME USING (timeofconsultation::TIMESTAMP)::TIME;

