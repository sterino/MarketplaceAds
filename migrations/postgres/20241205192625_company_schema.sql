-- +goose Up
-- +goose StatementBegin
CREATE TABLE companies (
                           id SERIAL PRIMARY KEY,
                           name VARCHAR NOT NULL,
                           email VARCHAR UNIQUE NOT NULL,
                           email_verified BOOLEAN DEFAULT FALSE,
                           password VARCHAR NOT NULL,
                           phone_number VARCHAR,
                           account_verified BOOLEAN DEFAULT FALSE,
                           account_type VARCHAR DEFAULT "company",
                           address TEXT,
                           orders_id TEXT[] DEFAULT NULL, -- Массив строк для заказов
                           created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                           updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE company;
-- +goose StatementEnd
