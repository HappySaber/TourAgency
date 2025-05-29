-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS employees(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    firstname VARCHAR(32) NOT NULL,
    lastname VARCHAR(32) NOT NULL,
    middlename VARCHAR(32),
    address VARCHAR(255) NOT NULL,
    phonenumber VARCHAR(10),
    dateofbirth DATE, 
    dateofhiring DATE,
    position UUID,
    FOREIGN KEY (position) REFERENCES positions(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS employees;
-- +goose StatementEnd
