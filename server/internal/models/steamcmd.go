package models

type SteamCMD struct {
	AIContentType string         `json:"aicontenttype"`
	StoreTags     map[int]string `json:"store_tags"`
}
