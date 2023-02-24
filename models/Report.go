package models

type AnimalReport struct {
	ID                     int64 `db:"id, primarykey, autoincrement" json:"id"`
	IdAnimal               int64
	AnimalName             string
	ScientificName         string
	ConservationStatusName string
	Abbreviation           string
	ClassificationName     string
	Fields                 string //TODO: use []string instead of string
}

type Report struct {
	ID                     int64 `db:"id, primarykey, autoincrement" json:"id"`
	CreatedAt              string
	IsApproved             bool
	AnimalName             string
	AnimalId               int64  `db:"animal_id"`
	ScientificName         string `db:"scientific_name"`
	ConservationStatusName string `db:"conservation_status_name"`
	Abbreviation           string `db:"abbreviation"`
	ClassificationName     string `db:"classification_name"`
	Fields                 string `db:"fields"`
}

type ReportFieldValue struct {
	ReportID int64
	AnimalID int64
	FieldID  int64
	Value    string
}
