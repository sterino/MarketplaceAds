-- +goose Up
-- +goose StatementBegin
CREATE TABLE applications (
                              id SERIAL PRIMARY KEY,
                              ad_id VARCHAR NOT NULL,
                              company_id VARCHAR NOT NULL,
                              influencer_id VARCHAR NOT NULL,
                              status VARCHAR DEFAULT 'pending', -- Статус заявки
                              created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                              updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                              FOREIGN KEY (ad_id) REFERENCES ads(id),
                              FOREIGN KEY (influencer_id) REFERENCES influencers(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE applications;
-- +goose StatementEnd
