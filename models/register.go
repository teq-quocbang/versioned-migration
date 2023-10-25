package models

type Register struct {
}

func (Register) TableName() string {
	return "register"
}
