-- +goose Up
-- +goose StatementBegin
CREATE TABLE influencers (
                             id SERIAL PRIMARY KEY,
                             name VARCHAR NOT NULL,
                             email VARCHAR UNIQUE NOT NULL,
                             email_verified BOOLEAN DEFAULT FALSE,
                             password VARCHAR NOT NULL,
                             phone_number VARCHAR,
                             account_verified BOOLEAN DEFAULT FALSE,
                             account_type VARCHAR DEFAULT "influencer",
                             platforms TEXT[], -- Массив строк для платформ
                             followers_count INT DEFAULT 0,
                             category VARCHAR,
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
