package models

//Tweet capture from body the message we recive
type Tweet struct {
	Message string `bson:"message" json:"message,omitempty"`
}
