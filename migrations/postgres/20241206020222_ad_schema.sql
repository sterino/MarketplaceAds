-- +goose Up
-- +goose StatementBegin
CREATE TABLE ads (
                     id SERIAL PRIMARY KEY,
                     company_id VARCHAR NOT NULL,
                     title VARCHAR NOT NULL,
                     description TEXT NOT NULL,
                     price DECIMAL NOT NULL,
                     status VARCHAR DEFAULT 'open',
                     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE ads;
-- +goose StatementEnd
