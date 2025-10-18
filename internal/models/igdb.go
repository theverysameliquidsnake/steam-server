package models

type ExternalGameSourceIGDB struct {
	Id   uint32 `json:"id"`
	Name string `json:"name"`
}

type ExternalGameIGDB struct {
	Game uint32 `json:"game"`
}

type GenreIGDB struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type ThemeIGDB struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type FranchiseIGDB struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type CollectionIGDB struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type KeywordIGDB struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type GameIGDB struct {
	Genres     []GenreIGDB      `json:"genres"`
	Themes     []ThemeIGDB      `json:"themes"`
	Franchises []FranchiseIGDB  `json:"franchises"`
	Series     []CollectionIGDB `json:"collections"`
	Keywords   []KeywordIGDB    `json:"keywords"`
}
