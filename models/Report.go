package models

type Report struct {
	ID            int64 `db:"id, primarykey, autoincrement" json:"id"`
	Created_at    string 
	IsApproved    bool
	AnimalName    string
	PlaceOfRescue string
}
