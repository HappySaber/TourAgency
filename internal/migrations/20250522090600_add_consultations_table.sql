-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS consultations (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    dateofconsultation DATE, 
    dateof DATE,
    client UUID,
    employee UUID,
    FOREIGN KEY (client) REFERENCES clients(id),
    FOREIGN KEY (employee) REFERENCES employees(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS consultations;
-- +goose StatementEnd
