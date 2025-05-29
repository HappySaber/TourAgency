-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS services (
    id SERIAL PRIMARY KEY,
    name VARCHAR(32), 
    price VARCHAR(64)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS services;
-- +goose StatementEnd
