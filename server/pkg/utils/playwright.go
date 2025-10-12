package utils

import (
	"os"

	"github.com/playwright-community/playwright-go"
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
		Headless: playwright.Bool(true),
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

/*func ParseSteamPage(appId uint32) (models.Game, error) {
	timeout, err := strconv.ParseFloat(os.Getenv("PLAYWRIGHT_TIMEOUT"), 64)
	if err != nil {
		return models.Game{}, err
	}

	userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36') Chrome/85.0.4183.121 Safari/537.36"
	page, err := browser.NewPage(playwright.BrowserNewPageOptions{
		UserAgent: &userAgent,
	})
	if err != nil {
		return models.Game{}, err
	}
	defer page.Close()

	if _, err := page.Goto(fmt.Sprintf("https://store.steampowered.com/app/%d", appId), playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateLoad,
		Timeout:   &timeout,
	}); err != nil {
		return models.Game{}, err
	}

	// Check for birthday verification
	count, err := page.Locator(".agegate_birthday_desc").Count()
	if err != nil {
		return models.Game{}, err
	}

	if count > 0 {
		if _, err = page.Locator("#ageDay").SelectOption(playwright.SelectOptionValues{
			Values: &[]string{"1"},
		}); err != nil {
			return models.Game{}, err
		}

		if _, err = page.Locator("#ageMonth").SelectOption(playwright.SelectOptionValues{
			Values: &[]string{"April"},
		}); err != nil {
			return models.Game{}, err
		}
		if _, err = page.Locator("#ageYear").SelectOption(playwright.SelectOptionValues{
			Values: &[]string{"1970"},
		}); err != nil {
			return models.Game{}, err
		}

		if err = page.Locator("#view_product_page_btn").Click(); err != nil {
			return models.Game{}, err
		}
	}

	// Reviews
	if err := page.Locator("button[aria-controls='review_type_flyout']").Hover(); err != nil {
		return models.Game{}, err
	}

	positiveText, err := page.Locator("label[for='review_type_positive'] > .user_reviews_count").TextContent()
	if err != nil {
		return models.Game{}, err
	}
	positive, err := strconv.ParseUint(Strip(positiveText), 10, 32)
	if err != nil {
		return models.Game{}, err
	}

	negativeText, err := page.Locator("label[for='review_type_negative'] > .user_reviews_count").TextContent()
	if err != nil {
		return models.Game{}, err
	}
	negative, err := strconv.ParseUint(Strip(negativeText), 10, 32)
	if err != nil {
		return models.Game{}, err
	}

	// Tags
	if err := page.Locator(".app_tag.add_button").Click(); err != nil {
		return models.Game{}, err
	}

	elems, err := page.Locator(".app_tag_control.popular").All()
	if err != nil {
		return models.Game{}, err
	}

	var tags []string
	for _, elem := range elems {
		tag, err := elem.TextContent()
		if err != nil {
			return models.Game{}, err
		}
		tags = append(tags, tag)
	}

	return models.Game{
		Tags:            tags,
		ReviewsPositive: uint32(positive),
		ReviewsNegative: uint32(negative),
	}, nil
}*/
