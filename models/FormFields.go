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

type ProvinceField struct {
	Id       int
	Province string
}

type ProvinceModel struct {
	Id       int
	Province string
	Cantons  []CantonModel
}

type CantonField struct {
	Id       int
	Canton string
}

type CantonModel struct {
	Id       int
	Canton string
	Districts []*DistrictField
}

type DistrictField struct {
	Id       int
	District string
}
