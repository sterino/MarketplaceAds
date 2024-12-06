-- +goose Up
-- +goose StatementBegin
CREATE TABLE orders (
                        id SERIAL PRIMARY KEY,
                        ad_id VARCHAR(255) NOT NULL,
                        company_id VARCHAR(255) NOT NULL,
                        influencer_id VARCHAR(255) NOT NULL,
                        status VARCHAR(50) DEFAULT 'pending', -- Статус заказа
                        price DECIMAL(10, 2) NOT NULL,
                        start_date TIMESTAMP NOT NULL,
                        deadline TIMESTAMP NOT NULL,
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE orders;
-- +goose StatementEnd
