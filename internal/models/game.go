package models

import "time"

type GameScreenshot struct {
	PathThumbnail string `bson:"path_thumbnail"`
	PathFull      string `bson:"path_full"`
}

type GameMovieEntry struct {
	P480 string `bson:"480"`
	Max  string `bson:"max"`
}

type GameMovie struct {
	Name      string         `bson:"name"`
	Thumbnail string         `bson:"thumbnail"`
	Webm      GameMovieEntry `bson:"webm"`
	Mp4       GameMovieEntry `bson:"mp4"`
}

type Game struct {
	AppId           uint32           `bson:"appid"`
	Name            string           `bson:"name"`
	Genres          []string         `bson:"genres"`
	Description     string           `bson:"description"`
	HeaderImage     string           `bson:"header_image"`
	Developers      []string         `bson:"developers"`
	Publishers      []string         `bson:"publishers"`
	Screenshots     []GameScreenshot `bson:"screenshots"`
	Movies          []GameMovie      `bson:"movies"`
	ComingSoon      bool             `bson:"coming_soon"`
	ReleaseDate     time.Time        `bson:"release_date"`
	AI              bool             `bson:"ai"`
	Tags            []string         `bson:"tags"`
	HasUnmappedTags bool             `bson:"has_unmapped_tags"`
	GenresIGDB      []string         `bson:"genres_igdb"`
	ThemesIGDB      []string         `bson:"themes_igdb"`
	SeriesIGDB      []string         `bson:"series_igdb"`
	FranchisesIGDB  []string         `bson:"franchises_igdb"`
	KeywordsIGDB    []string         `bson:"keywords_igdb"`
	CopiesSold      uint64           `bson:"copies_sold"`
	Downloads       uint64           `bson:"downloads"`
	GrossRevenue    uint64           `bson:"gross_revenue"`
	PlayersTotal    uint64           `bson:"players_total"`
	Owners          uint64           `bson:"owners"`
	ReviewScore     uint8            `bson:"review_score"`
	Reviews         uint64           `bson:"reviews"`
}
