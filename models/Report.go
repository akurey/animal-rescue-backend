package models

type AnimalReport struct {
	ID         				int64 `db:"id, primarykey, autoincrement" json:"id"`
	Idanimal        		int64 
	Animalname   			string
	Scientificname  		string
	Conservationstatusname 	string
	Abbreviaton  			string
	Classificationname 		string
	Fields     				string `db:"option" json:"option"` //TODO: use []string instead of string
}
type Report struct {
	ID            int64 `db:"id, primarykey, autoincrement" json:"id"`
	Created_at    string 
	IsApproved    bool
	AnimalName    string
	PlaceOfRescue string
}
