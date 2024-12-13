package ad

import "time"

type Entity struct {
	ID          string    `db:"id" bson:"_id"`
	Title       string    `db:"title" bson:"title"`
	Description string    `db:"description" bson:"description"`
	Price       float64   `db:"price" bson:"price"`
	Status      string    `db:"status" bson:"status"`
	CreatedAt   time.Time `db:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" bson:"updated_at"`
}
