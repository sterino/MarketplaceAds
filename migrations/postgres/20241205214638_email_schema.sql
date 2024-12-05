-- +goose Up
-- +goose StatementBegin
CREATE TABLE email_verification_codes (
      email VARCHAR PRIMARY KEY,
      code VARCHAR NOT NULL,
      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE email_verification_codes;
-- +goose StatementEnd
