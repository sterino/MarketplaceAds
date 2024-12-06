-- +goose Up
-- +goose StatementBegin
CREATE TABLE companies (
                           id SERIAL PRIMARY KEY,
                           name VARCHAR(255) NOT NULL,
                           email VARCHAR(255) UNIQUE NOT NULL,
                           email_verified BOOLEAN DEFAULT FALSE,
                           password VARCHAR(255) NOT NULL,
                           phone_number VARCHAR(20),
                           account_verified BOOLEAN DEFAULT FALSE,
                           account_type VARCHAR(50),
                           address TEXT,
                           orders_id TEXT[], -- Массив строк для заказов
                           created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                           updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE company;
-- +goose StatementEnd
