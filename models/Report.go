package models

type AnimalReport struct {
	ID         				int64 `db:"id, primarykey, autoincrement" json:"id"`
	IdAnimal        		int64 
	AnimalName   			string
	ScientificName  		string
	ConservationStatusName 	string
	Abbreviaton  			string
	ClassificationName 		string
	Fields     				string `db:"option" json:"option"` //TODO: use []string instead of string
}
type Report struct {
	ID            int64 `db:"id, primarykey, autoincrement" json:"id"`
	Created_at    string 
	IsApproved    bool
	AnimalName    string
	PlaceOfRescue string
}
