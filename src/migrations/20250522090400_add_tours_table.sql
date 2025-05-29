-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tours (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(32), 
    rating VARCHAR(8),
    hotel  VARCHAR(64),
    nutrition VARCHAR(64),
    city  VARCHAR(64),
    country VARCHAR(64),
    provider UUID,
    FOREIGN KEY (provider) REFERENCES providers(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tour;
-- +goose StatementEnd
