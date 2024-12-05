-- +goose Up
-- +goose StatementBegin
CREATE TABLE email_verification_codes (
      email VARCHAR(255) PRIMARY KEY,
      code VARCHAR(6) NOT NULL,
      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE email_verification_codes;
-- +goose StatementEnd
