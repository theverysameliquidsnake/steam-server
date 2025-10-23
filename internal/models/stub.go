package models

type Stub struct {
	AppId           uint32 `bson:"appid"`
	Name            string `bson:"name"`
	Type            string `bson:"type"`
	New             bool   `bson:"new"`
	SteamUpdate     bool   `bson:"steam_update"`
	SteamCMDUpdate  bool   `bson:"steamcmd_update"`
	IGDBUpdate      bool   `bson:"igdb_update"`
	GamalyticUpdate bool   `bson:"gamalytic_update"`
	Error           bool   `bson:"error"`
	Ignore          bool   `bson:"ignore"`
}
