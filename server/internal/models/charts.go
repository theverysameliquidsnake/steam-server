package models

type StubTypeCount struct {
	Type  string `bson:"_id"`
	Count uint32 `bson:"count"`
}

type GameReleaseYearCount struct {
	ReleaseYear int    `bson:"_id"`
	Count       uint32 `bson:"count"`
}

type ChartsDatasets struct {
	TotalCountOfStubs              uint32
	TotalCountOfUntouchedStubs     uint32
	TotalCountOfUnreleasedYetGames uint32
	TotalStubsByType               []StubTypeCount
	TotalGamesReleasedByYears      []GameReleaseYearCount
}
