-- +goose Up
-- +goose StatementBegin
CREATE TABLE influencers (
                             id SERIAL PRIMARY KEY,
                             name VARCHAR(255) NOT NULL,
                             email VARCHAR(255) UNIQUE NOT NULL,
                             email_verified BOOLEAN DEFAULT FALSE,
                             password VARCHAR(255) NOT NULL,
                             phone_number VARCHAR(20),
                             account_verified BOOLEAN DEFAULT FALSE,
                             account_type VARCHAR(50),
                             platforms TEXT[], -- Массив строк для платформ
                             followers_count INT DEFAULT 0,
                             category VARCHAR(100),
                             bio TEXT,
                             address TEXT,
                             created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                             updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                             orders_id TEXT[] -- Массив строк для связанного списка заказов
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE influencer;
-- +goose StatementEnd
