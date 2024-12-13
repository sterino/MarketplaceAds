-- +goose Up
-- +goose StatementBegin
CREATE TABLE orders (
                        id SERIAL PRIMARY KEY,
                        ad_id VARCHAR NOT NULL,
                        company_id VARCHAR NOT NULL,
                        influencer_id VARCHAR NOT NULL,
                        status VARCHAR DEFAULT 'pending', -- Статус заказа
                        price DECIMAL NOT NULL,
                        description TEXT, -- Описание заказа
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE orders;
-- +goose StatementEnd
