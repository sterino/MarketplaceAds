-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS influencer (
      id UUID PRIMARY KEY,
    email VARCHAR UNIQUE NOT NULL,
    password VARCHAR NOT NULL,
    phone_number VARCHAR,
    platforms TEXT[],
    followers_count INT DEFAULT 0,
    category VARCHAR,
    bio TEXT,
    address TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE influencer;
-- +goose StatementEnd
