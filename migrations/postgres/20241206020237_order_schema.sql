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
                        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                        FOREIGN KEY (company_id) REFERENCES companies(id),
                        FOREIGN KEY (ad_id) REFERENCES ads(id),
                        FOREIGN KEY (influencer_id) REFERENCES influencers(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE orders;
-- +goose StatementEnd
