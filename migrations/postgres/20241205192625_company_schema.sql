-- +goose Up
-- +goose StatementBegin
CREATE TABLE companies (
                           id VARCHAR NOT NULL,
                           name VARCHAR NOT NULL,
                           email VARCHAR UNIQUE NOT NULL,
                           email_verified BOOLEAN DEFAULT FALSE,
                           password VARCHAR NOT NULL,
                           phone_number VARCHAR,
                           account_verified BOOLEAN DEFAULT FALSE,
                           account_type VARCHAR DEFAULT 'company',
                           address TEXT,
                           created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                           updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE company;
-- +goose StatementEnd
