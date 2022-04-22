package models

type Claims struct {
	Username string `json:"username,omitempty"`
	Rules    uint16 `json:"rules,omitempty"`
}
