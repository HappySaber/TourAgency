-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS clients(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    firstname VARCHAR(32) NOT NULL,
    lastname VARCHAR(32) NOT NULL,
    middlename VARCHAR(32),
    address VARCHAR(255) NOT NULL,
    phonenumber VARCHAR(10),
    dateofbirth DATE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS clients;
-- +goose StatementEnd
