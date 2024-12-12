-- +goose Up
-- +goose StatementBegin
CREATE TABLE orders (
                        id SERIAL PRIMARY KEY,
                        ad_id VARCHAR NOT NULL,
                        company_id VARCHAR NOT NULL,
                        influencer_id VARCHAR NOT NULL,
                        status VARCHAR DEFAULT 'pending', -- Статус заказа
                        price DECIMAL NOT NULL,
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
