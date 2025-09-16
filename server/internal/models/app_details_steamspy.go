package models

type AppDetailsSteamSpy struct {
	AppId    uint32            `json:"appid"`
	Name     string            `json:"name"`
	Positive uint32            `json:"positive"`
	Negative uint32            `json:"negative"`
	Owners   string            `json:"owners"`
	Tags     map[string]uint32 `json:"tags"`
}
