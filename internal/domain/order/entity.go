package order

import "time"

type Entity struct {
	ID           string    `db:"id" bson:"_id"`
	AdID         string    `db:"ad_id" bson:"ad_id"`
	CompanyID    string    `db:"company_id" bson:"company_id"`
	InfluencerID string    `db:"influencer_id" bson:"influencer_id"`
	Status       string    `db:"status" bson:"status"`
	Price        float64   `db:"price" bson:"price"`
	Description  string    `db:"description"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}
