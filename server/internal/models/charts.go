package models

type StubByStatus struct {
	Status string
	Count  uint32
}

type StubTypeCount struct {
	Type  string `bson:"_id"`
	Count uint32 `bson:"count"`
}

type GameReleaseYearCount struct {
	ReleaseYear int    `bson:"_id"`
	Count       uint32 `bson:"count"`
}

type GameReleaseYearDataset struct {
	TotalUnreleasedYetGames   uint32
	TotalGamesReleasedByYears []GameReleaseYearCount
}

type ChartsDatasets struct {
	TotalStubsByStatus []StubByStatus
	TotalStubsByType   []StubTypeCount
	GamesByYearDataset GameReleaseYearDataset
}
