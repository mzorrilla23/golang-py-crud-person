package models

type Address struct {
	State   string `json:"state" bson:"state"`
	City    string `json:"city" bson:"city"`
	Pincode int    `json:"pincode" bson:"pincode"`
}

type Person struct {
	Id        int     `json:"id" bson:"id"`
	FirstName string  `json:"firstName" bson:"first_name"`
	LastName  string  `json:"lastName" bson:"last_name"`
	Address   Address `json:"address" bson:"person_address"`
}
