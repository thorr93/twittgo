package models

//Relation - model to record the relation from a user to another
type Relation struct {
	UserID         string `bson:"userid" json:"userID"`
	UserRelationID string `bson:"userrelationid" json:"userRelationID"`
}
