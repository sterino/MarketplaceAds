package company

import "time"

type Entity struct {
	ID          string    `db:"id" bson:"_id"`
	Name        string    `db:"name" bson:"name"`
	Email       string    `db:"email" bson:"email"`
	Password    string    `db:"password" bson:"password"`
	PhoneNumber string    `db:"phone_number" bson:"phone_number"`
	Address     string    `db:"address" bson:"address"`
	CreatedAt   time.Time `db:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" bson:"updated_at"`
}
