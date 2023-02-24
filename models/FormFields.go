package models

type FormField struct {
	FormName     string
	FormSection  string
	FieldId      int64
	FieldName    string
	IsRequired   bool
	FieldType    string
	FieldOptions string `db:"field_options"` //TODO: use []string instead of string
}

type AddressField struct {
	Id       int64
	Province string
	Canton   string
	District string
}
