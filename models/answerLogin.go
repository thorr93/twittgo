package models

//AnswerLogin had the token which is returned with the login
type AnswerLogin struct {
	Token string `json:"token,omitempty"`
}
