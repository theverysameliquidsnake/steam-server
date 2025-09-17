package models

type AppDetailsGenreSteam struct {
	Id          string `json:"id"`
	Description string `json:"description"`
}

type AppDetailsScreenshotSteam struct {
	Id            uint32 `json:"id"`
	PathThumbnail string `json:"path_thumbnail"`
	PathFull      string `json:"path_full"`
}

type AppDetailsMovieEntrySteam struct {
	P480 string `json:"480"`
	Max  string `json:"max"`
}

type AppDetailsMovieSteam struct {
	Id        uint32                    `json:"id"`
	Name      string                    `json:"name"`
	Thumbnail string                    `json:"thumbnail"`
	Webm      AppDetailsMovieEntrySteam `json:"webm"`
	Mp4       AppDetailsMovieEntrySteam `json:"mp4"`
	Highlight bool                      `json:"highlight"`
}

type AppDetailsReleaseDateSteam struct {
	ComingSoon bool   `json:"coming_soon"`
	Date       string `json:"date"`
}

type AppDetailsSteam struct {
	SteamAppId       uint32                      `json:"steam_appid"`
	Type             string                      `json:"type"`
	Name             string                      `json:"name"`
	ShortDescription string                      `json:"short_description"`
	HeaderImage      string                      `json:"header_image"`
	Developers       []string                    `json:"developers"`
	Publishers       []string                    `json:"publishers"`
	Genres           []AppDetailsGenreSteam      `json:"genres"`
	Screenshots      []AppDetailsScreenshotSteam `json:"screenshots"`
	Movies           []AppDetailsMovieSteam      `json:"movies"`
	ReleaseDate      AppDetailsReleaseDateSteam  `json:"release_date"`
}
