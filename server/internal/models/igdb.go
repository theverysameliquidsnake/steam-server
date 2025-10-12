package models

type ExternalGameSource struct {
	Id   uint32 `json:"id"`
	Name string `json:"name"`
}

type ExternalGame struct {
	Game uint32 `json:"game"`
}

type Genre struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type Theme struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type Franchise struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type Collection struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type Keyword struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type GameIGDB struct {
	Genres     []Genre      `json:"genres"`
	Themes     []Theme      `json:"themes"`
	Franchises []Franchise  `json:"franchises"`
	Series     []Collection `json:"collections"`
	Keywords   []Keyword    `json:"keywords"`
}
