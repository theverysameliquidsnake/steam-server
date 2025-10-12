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
	AppId       uint32           `bson:"appid"`
	Name        string           `bson:"name"`
	Description string           `bson:"description"`
	HeaderImage string           `bson:"header_image"`
	Developers  []string         `bson:"developers"`
	Publishers  []string         `bson:"publishers"`
	Screenshots []GameScreenshot `bson:"screenshots"`
	Movies      []GameMovie      `bson:"movies"`
	ComingSoon  bool             `bson:"coming_soon"`
	ReleaseDate time.Time        `bson:"release_date"`
	Genres      []string         `bson:"genres"`
	Themes      []string         `bson:"themes"`
	Series      []string         `bson:"series"`
	Franchises  []string         `bson:"franchises"`
	Keywords    []string         `bson:"keywords"`
	//AI              bool             `bson:"ai"`
	//Tags            []string         `bson:"tags"`
	//Owners          uint32           `bson:"owners"`
	//ReviewScore     float32          `bson:"review_score"`
	//ReviewsPositive uint32           `bson:"reviews_positive"`
	//ReviewsNegative uint32           `bson:"reviews_negative"`
}
