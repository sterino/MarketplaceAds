-- +goose Up
-- +goose StatementBegin
CREATE TABLE ads (
                     id SERIAL PRIMARY KEY,
                     title VARCHAR NOT NULL,
                     description TEXT NOT NULL,
                     price DECIMAL NOT NULL,
                     status VARCHAR DEFAULT 'open', -- Статус объявления
                     orders_id TEXT[] DEFAULT NULL, -- Массив ID заказов, связанных с этим объявлением
                     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE ads;
-- +goose StatementEnd
