-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS influencer (
    id UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    name VARCHAR NOT NULL,
    email VARCHAR UNIQUE NOT NULL,
    password VARCHAR NOT NULL,
    phone_number VARCHAR,
    address VARCHAR,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE influencer;
-- +goose StatementEnd
