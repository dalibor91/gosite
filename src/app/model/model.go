package model

type DefaultModel interface {
	Save() bool
	Update() bool
	Delete() bool
	GetTableName() string
}