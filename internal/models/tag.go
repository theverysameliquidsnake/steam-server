package models

type Tag struct {
	Id   uint32 `bson:"id"`
	Name string `bson:"name"`
}
