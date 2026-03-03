package handler

import (
	"strings"
	"sync"
	"toychart/kit/ebay"

	"github.com/labstack/echo/v4"
)

type popmartDirectoryItem struct {
	Slug        string `json:"slug"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Keyword     string `json:"keyword"`
	ImageURL    string `json:"imageUrl"`
	SoldCount   int    `json:"soldCount"`
}

var popmartDirectorySeed = []popmartDirectoryItem{
	{
		Slug:        "labubu",
		Name:        "Labubu",
		Description: "THE MONSTERS universe including Macaron, Have a Seat, and collabs.",
		Keyword:     "pop mart labubu",
		ImageURL:    "https://imagedelivery.net/j3cQhHk0Y2fW12f0X0X0Xw/placeholder-5/public",
	},
	{
		Slug:        "skullpanda",
		Name:        "Skullpanda",
		Description: "Art-driven character line with many themed blind box series.",
		Keyword:     "pop mart skullpanda",
		ImageURL:    "https://imagedelivery.net/j3cQhHk0Y2fW12f0X0X0Xw/placeholder-4/public",
	},
	{
		Slug:        "dimoo",
		Name:        "Dimoo",
		Description: "Dreamy fantasy character line with high collector demand.",
		Keyword:     "pop mart dimoo",
		ImageURL:    "https://imagedelivery.net/j3cQhHk0Y2fW12f0X0X0Xw/placeholder-6/public",
	},
	{
		Slug:        "hirono",
		Name:        "Hirono",
		Description: "Stylized narrative character series from Lang.",
		Keyword:     "pop mart hirono",
		ImageURL:    "https://imagedelivery.net/j3cQhHk0Y2fW12f0X0X0Xw/placeholder-3/public",
	},
	{
		Slug:        "molly",
		Name:        "Molly",
		Description: "Classic Pop Mart IP with broad crossovers and editions.",
		Keyword:     "pop mart molly",
		ImageURL:    "https://imagedelivery.net/j3cQhHk0Y2fW12f0X0X0Xw/placeholder-2/public",
	},
	{
		Slug:        "crybaby",
		Name:        "Crybaby",
		Description: "Emotive character line with trend-driven releases.",
		Keyword:     "pop mart crybaby",
		ImageURL:    "https://imagedelivery.net/j3cQhHk0Y2fW12f0X0X0Xw/placeholder-1/public",
	},
	{
		Slug:        "pucky",
		Name:        "Pucky",
		Description: "Fairy-tale inspired mini worlds and seasonal sets.",
		Keyword:     "pop mart pucky",
		ImageURL:    "https://imagedelivery.net/j3cQhHk0Y2fW12f0X0X0Xw/placeholder-7/public",
	},
	{
		Slug:        "hacipupu",
		Name:        "Hacipupu",
		Description: "Whimsical character line with playful themed drops.",
		Keyword:     "pop mart hacipupu",
		ImageURL:    "https://imagedelivery.net/j3cQhHk0Y2fW12f0X0X0Xw/placeholder-8/public",
	},
}

func (h *Handler) PopmartCategory(c echo.Context) error {
	items := make([]popmartDirectoryItem, len(popmartDirectorySeed))
	copy(items, popmartDirectorySeed)

	var wg sync.WaitGroup
	var mu sync.Mutex
	sem := make(chan struct{}, 4)

	for idx := range items {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			result, err := ebay.ScrapeSoldToyPrices(strings.TrimSpace(items[i].Keyword), "", 1)
			if err != nil || len(result) == 0 {
				return
			}

			mu.Lock()
			items[i].SoldCount = len(result)
			if strings.TrimSpace(result[0].ImageURL) != "" {
				items[i].ImageURL = strings.TrimSpace(result[0].ImageURL)
			}
			mu.Unlock()
		}(idx)
	}

	wg.Wait()
	return responseJSON(c, items)
}

