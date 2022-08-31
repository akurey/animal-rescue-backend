package models

type FormField struct {
	FormName   string
	FieldName  string
	IsRequired bool
	FieldType  string
	Option     string `db:"option" json:"option"` //TODO: use []string instead of string
}
