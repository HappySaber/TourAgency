-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tours_per_consultation (
    tour_id UUID,
    consultation_id UUID,
    PRIMARY KEY (tour_id, consultation_id),
    FOREIGN KEY (tour_id) REFERENCES tours(id),
    FOREIGN KEY (consultation_id) REFERENCES consultations(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tours_per_consultation;
-- +goose StatementEnd
