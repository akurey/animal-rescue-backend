package models

type Animal struct {
	ID                       int64 `db:"id, primarykey, autoincrement" json:"id"`
	Name                     string
	ScientificName           string
	Created_at               string
	ConservationStatus       string
	ConservationAbbreviation string
	ClassificationName       string
}
