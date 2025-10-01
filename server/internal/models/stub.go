package models

type Stub struct {
	AppId        uint32 `bson:"appid"`
	Name         string `bson:"name"`
	Type         string `bson:"type"`
	New          bool   `bson:"new"`
	FirstUpdate  bool   `bson:"first_update"`
	SecondUpdate bool   `bson:"second_update"`
	Error        bool   `bson:"error"`
	Ignore       bool   `bson:"ignore"`
}
