package order

import "time"

type Entity struct {
	ID           string    `db:"id" bson:"_id"`
	AdID         string    `db:"ad_id" bson:"ad_id"`
	CompanyID    string    `db:"company_id" bson:"company_id"`
	InfluencerID string    `db:"influencer_id" bson:"influencer_id"`
	Status       string    `db:"status" bson:"status"` // Например: "pending", "in_progress", "completed"
	Price        float64   `db:"price" bson:"price"`
	StartDate    time.Time `db:"start_date" bson:"start_date"`
	Deadline     time.Time `db:"deadline" bson:"deadline"`
	CreatedAt    time.Time `db:"created_at" bson:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" bson:"updated_at"`
}
