package models

type User struct {
	Name      string `json:"name" bson:"name"`
	Age       int    `json:"age" bson:"age"`
	Documents Documents
	CreateAt  int64 `json:"create_at" bson:"create_at"`
}

type Documents struct {
	PassportNumber string `json:"passport_number" bson:"passport_number"`
	INN            int    `json:"inn" bson:"inn"`
}
