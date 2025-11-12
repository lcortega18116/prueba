package models

type Item struct {
	ID         int64  `db:"id" json:"id"`
	Ticker     string `db:"ticker" json:"ticker"`
	TargetFrom string `db:"target_from" json:"target_from"`
	TargetTo   string `db:"target_to" json:"target_to"`
	Company    string `db:"company" json:"company"`
	Action     string `db:"action" json:"action"`
	Brokerage  string `db:"brokerage" json:"brokerage"`
	RatingFrom string `db:"rating_from" json:"rating_from"`
	RatingTo   string `db:"rating_to" json:"rating_to"`
	Time       string `db:"time" json:"time"`
}
