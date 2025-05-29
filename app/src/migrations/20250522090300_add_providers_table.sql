-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS providers (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(32), 
    addressto VARCHAR(64),
    address  VARCHAR(255),
    email VARCHAR(64),
    phonenumber  VARCHAR(16)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS providers;
-- +goose StatementEnd
