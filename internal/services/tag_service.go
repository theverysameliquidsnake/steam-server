package services

import (
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/theverysameliquidsnake/steam-db/internal/models"
	"github.com/theverysameliquidsnake/steam-db/internal/repositories"
)

func RefreshTags() (int, error) {
	reader, err := os.Open("Steam Game Tags · SteamDB.html")
	if err != nil {
		return -1, err
	}
	defer reader.Close()

	document, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return -1, err
	}

	tagsMap := make(map[string]string)
	document.Find("a[href*='/tag/']").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		tagsMap[strings.Split(href, "/")[2]] = s.Children().Remove().End().Text()
	})

	var tags []models.Tag
	for key, value := range tagsMap {
		id, err := strconv.ParseUint(key, 10, 32)
		if err != nil {
			return -1, err
		}

		tags = append(tags, models.Tag{Id: uint32(id), Name: value})
	}

	err = repositories.DeleteTags()
	if err != nil {
		return -1, err
	}

	result, err := repositories.InsertTags(tags)
	if err != nil {
		return -1, err
	}

	return len(result), nil
}
