package models

type Stub struct {
	AppId       uint32 `bson:"appid"`
	Name        string `bson:"name"`
	Type        string `bson:"type"`
	NeedsUpdate bool   `bson:"needs_update"`
	Skip        bool   `bson:"skip"`
}
