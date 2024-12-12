package company

import "time"

type Entity struct {
	ID              string    `db:"id" bson:"_id"`
	Name            string    `db:"name" bson:"name"`
	Email           string    `db:"email" bson:"email"`
	EmailVerified   bool      `db:"email_verified" bson:"email_verified"`
	Password        string    `db:"password" bson:"password"`
	PhoneNumber     string    `db:"phone_number" bson:"phone_number"`
	AccountVerified bool      `db:"account_verified" bson:"account_verified"`
	AccountType     string    `db:"account_type" bson:"account_type"`
	Address         string    `db:"address" bson:"address"`
	OrdersID        *[]string `db:"orders_id" bson:"orders_id"`
	CreatedAt       time.Time `db:"created_at" bson:"created_at"`
	UpdatedAt       time.Time `db:"updated_at" bson:"updated_at"`
}
