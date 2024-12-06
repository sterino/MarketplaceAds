-- +goose Up
-- +goose StatementBegin
CREATE TABLE ads (
                     id SERIAL PRIMARY KEY,
                     title VARCHAR(255) NOT NULL,
                     description TEXT NOT NULL,
                     price DECIMAL(10, 2) NOT NULL,
                     status VARCHAR(50) DEFAULT 'open', -- Статус объявления
                     orders_id TEXT[], -- Массив ID заказов, связанных с этим объявлением
                     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE ads;
-- +goose StatementEnd
