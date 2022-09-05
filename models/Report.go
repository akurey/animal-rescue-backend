package models

type Report struct {
	ID         				int64 `db:"id, primarykey, autoincrement" json:"id"`
	IdAnimal        		int64 `db:"id, primarykey, autoincrement" json:"id"`
	AnimalName   			string
	ScientificName  		string
	ConservationStatusName 	string
	Abbreviaton  			string
	ClassificationName 		string
	Fields     				string `db:"option" json:"option"` //TODO: use []string instead of string
}