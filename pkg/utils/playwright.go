package utils

import (
	"fmt"
	"os"
	"strconv"

	"github.com/playwright-community/playwright-go"
	"github.com/theverysameliquidsnake/steam-db/internal/models"
)

var engine *playwright.Playwright
var browser playwright.Browser

func InitPlaywright() error {
	return playwright.Install(
		&playwright.RunOptions{
			Browsers: []string{"chromium"},
		},
	)
}

func StartPlaywright() error {
	pw, err := playwright.Run()
	if err != nil {
		return err
	}
	engine = pw

	br, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
		Proxy: &playwright.Proxy{
			Server: os.Getenv("PROXY_FULL"),
		},
	})
	if err != nil {
		return err
	}
	browser = br

	return nil
}

func StopPlaywright() error {
	if err := browser.Close(); err != nil {
		return err
	}

	if err := engine.Stop(); err != nil {
		return err
	}

	return nil
}

func ParseGamalyticPage(appId uint32) (*models.Game, error) {
	timeout, err := strconv.ParseFloat(os.Getenv("PLAYWRIGHT_TIMEOUT"), 64)
	if err != nil {
		return nil, err
	}

	userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36') Chrome/85.0.4183.121 Safari/537.36"
	page, err := browser.NewPage(playwright.BrowserNewPageOptions{
		UserAgent: &userAgent,
	})
	if err != nil {
		return nil, err
	}
	defer page.Close()

	if _, err := page.Goto(fmt.Sprintf("https://gamalytic.com/game/%d", appId), playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateNetworkidle,
		Timeout:   &timeout,
	}); err != nil {
		return nil, err
	}

	// Wait lazy load
	statsBlock := page.Locator("//span[text()='Stats']")
	err = statsBlock.WaitFor(playwright.LocatorWaitForOptions{
		State:   playwright.WaitForSelectorStateVisible,
		Timeout: &timeout,
	})
	if err != nil {
		return nil, err
	}

	var game models.Game

	// Copies sold
	if game.CopiesSold, err = extractStat("Copies sold: ", page); err != nil {
		return nil, err
	}

	// Downloads
	if game.Downloads, err = extractStat("Downloads: ", page); err != nil {
		return nil, err
	}

	// Gross revenue
	if game.GrossRevenue, err = extractStat("Gross revenue: ", page); err != nil {
		return nil, err
	}

	// Players total
	if game.PlayersTotal, err = extractStat("Players total: ", page); err != nil {
		return nil, err
	}

	// Owners
	if game.Owners, err = extractStat("Owners: ", page); err != nil {
		return nil, err
	}

	// Review score
	reviewScore, err := extractStat("Review score: ", page)
	if err != nil {
		return nil, err
	}

	game.ReviewScore = uint8(reviewScore)

	// Reviews
	if game.Reviews, err = extractStat("Reviews: ", page); err != nil {
		return nil, err
	}

	return &game, nil
}

func extractStat(statPrefix string, page playwright.Page) (uint64, error) {
	statCount, err := page.Locator(fmt.Sprintf("//b[text()='%s']", statPrefix)).Count()
	if err != nil {
		return 0, err
	}

	if statCount > 0 {
		statString, err := page.Locator(fmt.Sprintf("//b[text()='%s']/following-sibling::div[1]", statPrefix)).TextContent()
		if err != nil {
			return 0, err
		}

		stat, err := ParseStatEntry(statString)
		if err != nil {
			return 0, err
		}

		return stat, nil
	}

	return 0, nil
}
