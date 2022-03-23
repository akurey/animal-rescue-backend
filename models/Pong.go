package models

//Pong ...
type Pong struct {
	ID      int64  `db:"id, primarykey, autoincrement" json:"id"`
	Message string `db:"message" json:"message"`
}
