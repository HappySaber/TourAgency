-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS services_per_consultation (
    service_id INT,
    consultation_id UUID,
    PRIMARY KEY (service_id, consultation_id),
    FOREIGN KEY (service_id) REFERENCES services(id),
    FOREIGN KEY (consultation_id) REFERENCES consultations(id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS services_per_consultation;
-- +goose StatementEnd
