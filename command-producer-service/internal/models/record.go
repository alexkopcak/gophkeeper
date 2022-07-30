package models

type Record struct {
	Opeartion   byte
	Id          int64 `gorm:"primaryKey"`
	UserId      int64
	MessageType byte
	Data        []byte
	Meta        string
}
