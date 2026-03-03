package handler

import (
	"regexp"
	"sort"
	"strings"
	"toychart/errcode"
	"toychart/kit/ebay"

	"github.com/labstack/echo/v4"
)

type popmartSeriesItem struct {
	Title     string `json:"title"`
	ImageURL  string `json:"imageUrl"`
	AvgPrice  string `json:"avgPrice"`
	SoldCount int    `json:"soldCount"`
	Href      string `json:"href"`
}

type popmartSeriesItemsResponse struct {
	IP     popmartDirectoryItem `json:"ip"`
	Series string               `json:"series"`
	Items  []popmartSeriesItem  `json:"items"`
}

type popmartSeriesItemResponse struct {
	IP     popmartDirectoryItem `json:"ip"`
	Series string               `json:"series"`
	Item   popmartSeriesItem    `json:"item"`
}

var popmartSeriesSlugAliases = map[string]string{
	"labubu-exciting-macaron": "THE MONSTERS - Exciting Macaron Vinyl Face",
}

var popmartTitleNoisePattern = regexp.MustCompile(`(?i)\b(authentic|official|genuine|new|sealed|in\s*hand|instock|open\s*box|opened|us\s*seller|free\s*shipping|blind\s*box|figure|vinyl|plush|pendant|toy|collectible)\b`)
var popmartNonWordPattern = regexp.MustCompile(`[^a-z0-9]+`)
var excitingMacaronSeriesName = "THE MONSTERS - Exciting Macaron Vinyl Face"
var excitingMacaronCanonicalOrder = []string{
	"Soymilk",
	"Lychee Berry",
	"Green Grape",
	"Sea Salt Coconut",
	"Toffee",
	"Sesame Bean",
	"Chestnut Cocoa (Secret)",
}
var excitingMacaronItemSlug = map[string]string{
	"Soymilk":                 "soymilk",
	"Lychee Berry":            "lychee-berry",
	"Green Grape":             "green-grape",
	"Sea Salt Coconut":        "sea-salt-coconut",
	"Toffee":                  "toffee",
	"Sesame Bean":             "sesame-bean",
	"Chestnut Cocoa (Secret)": "chestnut-cocoa",
}

func normalizeToken(value string) string {
	s := strings.ToLower(strings.TrimSpace(value))
	return strings.Trim(popmartNonWordPattern.ReplaceAllString(s, " "), " ")
}

func matchExcitingMacaronName(rawTitle string) string {
	title := normalizeToken(rawTitle)
	if title == "" {
		return ""
	}
	if strings.Contains(title, "chestnut cocoa") {
		return "Chestnut Cocoa (Secret)"
	}
	if strings.Contains(title, "lychee berry") {
		return "Lychee Berry"
	}
	if strings.Contains(title, "green grape") {
		return "Green Grape"
	}
	if strings.Contains(title, "sea salt coconut") || strings.Contains(title, "sea salt") {
		return "Sea Salt Coconut"
	}
	if strings.Contains(title, "toffee") {
		return "Toffee"
	}
	if strings.Contains(title, "sesame bean") || strings.Contains(title, "sesame") {
		return "Sesame Bean"
	}
	if strings.Contains(title, "soymilk") || strings.Contains(title, "soy milk") {
		return "Soymilk"
	}
	return ""
}

func resolvePopmartSeriesBySlug(slug string) string {
	s := strings.TrimSpace(slug)
	if s == "" {
		return ""
	}
	if alias, ok := popmartSeriesSlugAliases[s]; ok {
		return alias
	}
	for _, rule := range popmartSeriesRules {
		if popmartSeriesToSlug(rule.series) == s {
			return rule.series
		}
	}
	return ""
}

func compactPopmartItemTitle(rawTitle, ipName, series string) string {
	if series == excitingMacaronSeriesName {
		if matched := matchExcitingMacaronName(rawTitle); matched != "" {
			return matched
		}
		return ""
	}

	cleaned := sanitizePopmartTitle(rawTitle)
	cleaned = regexp.MustCompile(`(?i)\bpop\s*mart\b`).ReplaceAllString(cleaned, "")
	cleaned = regexp.MustCompile(`(?i)\bthe\s+monsters\b`).ReplaceAllString(cleaned, "THE MONSTERS")
	cleaned = regexp.MustCompile(`(?i)\blabubu\b`).ReplaceAllString(cleaned, "")
	cleaned = popmartTitleNoisePattern.ReplaceAllString(cleaned, "")
	cleaned = strings.ReplaceAll(cleaned, "-", " ")
	cleaned = regexp.MustCompile(`\s+`).ReplaceAllString(cleaned, " ")
	cleaned = strings.TrimSpace(cleaned)

	// Remove IP label from item title, keep the concise variant name only.
	cleaned = regexp.MustCompile(`(?i)\b`+regexp.QuoteMeta(ipName)+`\b`).ReplaceAllString(cleaned, "")
	cleaned = regexp.MustCompile(`(?i)\b`+regexp.QuoteMeta(series)+`\b`).ReplaceAllString(cleaned, "")
	cleaned = regexp.MustCompile(`\s+`).ReplaceAllString(cleaned, " ")
	cleaned = strings.TrimSpace(cleaned)

	return strings.TrimSpace(cleaned)
}

func popmartItemSlug(series, title string) string {
	if series == excitingMacaronSeriesName {
		if slug, ok := excitingMacaronItemSlug[title]; ok {
			return slug
		}
	}
	s := strings.ToLower(strings.TrimSpace(title))
	s = strings.ReplaceAll(s, "&", " and ")
	s = regexp.MustCompile(`\([^)]*\)`).ReplaceAllString(s, "")
	s = regexp.MustCompile(`[^a-z0-9]+`).ReplaceAllString(s, "-")
	s = strings.Trim(s, "-")
	return s
}

func (h *Handler) PopmartSeriesItems(c echo.Context) error {
	ipSlug := strings.TrimSpace(c.Param("ip"))
	seriesSlug := strings.TrimSpace(c.Param("series"))
	if ipSlug == "" || seriesSlug == "" {
		return responseValidationError(c, "ip and series are required")
	}

	var selected *popmartDirectoryItem
	for i := range popmartDirectorySeed {
		if popmartDirectorySeed[i].Slug == ipSlug {
			selected = &popmartDirectorySeed[i]
			break
		}
	}
	if selected == nil {
		return responseValidationError(c, "invalid ip")
	}

	targetSeries := ""
	if seriesSlug != "all-series" {
		targetSeries = resolvePopmartSeriesBySlug(seriesSlug)
		if targetSeries == "" {
			return responseValidationError(c, "invalid series")
		}
	}

	result, err := ebay.ScrapeSoldToyPrices(strings.TrimSpace(selected.Keyword), "", 3)
	if err != nil {
		return responseError(c, errcode.InternalServerError)
	}

	grouped := map[string]popmartSeriesItem{}
	if targetSeries == excitingMacaronSeriesName {
		for _, name := range excitingMacaronCanonicalOrder {
			grouped[name] = popmartSeriesItem{
				Title:     name,
				ImageURL:  selected.ImageURL,
				AvgPrice:  "",
				SoldCount: 0,
				Href:      "/category/pop-mart/" + selected.Slug + "/" + seriesSlug + "/" + popmartItemSlug(targetSeries, name),
			}
		}
	}
	for _, sale := range result {
		saleSeries := classifyPopmartSeries(sale.Title)
		if targetSeries != "" && saleSeries != targetSeries {
			continue
		}

		itemTitle := compactPopmartItemTitle(sale.Title, selected.Name, saleSeries)
		if itemTitle == "" {
			continue
		}

		existing, ok := grouped[itemTitle]
		if !ok {
			image := strings.TrimSpace(sale.ImageURL)
			if image == "" {
				image = selected.ImageURL
			}
			grouped[itemTitle] = popmartSeriesItem{
				Title:     itemTitle,
				ImageURL:  image,
				AvgPrice:  strings.TrimSpace(sale.Price),
				SoldCount: 1,
				Href:      "/category/pop-mart/" + selected.Slug + "/" + seriesSlug + "/" + popmartItemSlug(saleSeries, itemTitle),
			}
			continue
		}

		existing.SoldCount++
		if strings.TrimSpace(existing.AvgPrice) == "" {
			existing.AvgPrice = strings.TrimSpace(sale.Price)
		}
		if strings.TrimSpace(existing.ImageURL) == "" && strings.TrimSpace(sale.ImageURL) != "" {
			existing.ImageURL = strings.TrimSpace(sale.ImageURL)
		}
		grouped[itemTitle] = existing
	}

	items := make([]popmartSeriesItem, 0, len(grouped))
	if targetSeries == excitingMacaronSeriesName {
		for _, name := range excitingMacaronCanonicalOrder {
			item, ok := grouped[name]
			if !ok {
				continue
			}
			items = append(items, item)
		}
	} else {
		for _, item := range grouped {
			items = append(items, item)
		}
		sort.Slice(items, func(i, j int) bool {
			if items[i].SoldCount != items[j].SoldCount {
				return items[i].SoldCount > items[j].SoldCount
			}
			return items[i].Title < items[j].Title
		})
	}

	seriesLabel := targetSeries
	if seriesSlug == "all-series" {
		seriesLabel = "All Series"
	}

	return responseJSON(c, popmartSeriesItemsResponse{
		IP:     *selected,
		Series: seriesLabel,
		Items:  items,
	})
}

func (h *Handler) PopmartSeriesItem(c echo.Context) error {
	ipSlug := strings.TrimSpace(c.Param("ip"))
	seriesSlug := strings.TrimSpace(c.Param("series"))
	itemSlug := strings.TrimSpace(c.Param("item"))
	if ipSlug == "" || seriesSlug == "" || itemSlug == "" {
		return responseValidationError(c, "ip, series, and item are required")
	}

	var selected *popmartDirectoryItem
	for i := range popmartDirectorySeed {
		if popmartDirectorySeed[i].Slug == ipSlug {
			selected = &popmartDirectorySeed[i]
			break
		}
	}
	if selected == nil {
		return responseValidationError(c, "invalid ip")
	}

	targetSeries := ""
	if seriesSlug != "all-series" {
		targetSeries = resolvePopmartSeriesBySlug(seriesSlug)
		if targetSeries == "" {
			return responseValidationError(c, "invalid series")
		}
	}

	result, err := ebay.ScrapeSoldToyPrices(strings.TrimSpace(selected.Keyword), "", 3)
	if err != nil {
		return responseError(c, errcode.InternalServerError)
	}

	var matched *popmartSeriesItem
	for _, sale := range result {
		saleSeries := classifyPopmartSeries(sale.Title)
		if targetSeries != "" && saleSeries != targetSeries {
			continue
		}

		itemTitle := compactPopmartItemTitle(sale.Title, selected.Name, saleSeries)
		if itemTitle == "" {
			continue
		}

		if popmartItemSlug(saleSeries, itemTitle) != itemSlug {
			continue
		}

		if matched == nil {
			image := strings.TrimSpace(sale.ImageURL)
			if image == "" {
				image = selected.ImageURL
			}
			matched = &popmartSeriesItem{
				Title:     itemTitle,
				ImageURL:  image,
				AvgPrice:  strings.TrimSpace(sale.Price),
				SoldCount: 1,
				Href:      "/category/pop-mart/" + selected.Slug + "/" + seriesSlug + "/" + itemSlug,
			}
			continue
		}

		matched.SoldCount++
		if strings.TrimSpace(matched.AvgPrice) == "" {
			matched.AvgPrice = strings.TrimSpace(sale.Price)
		}
		if strings.TrimSpace(matched.ImageURL) == "" && strings.TrimSpace(sale.ImageURL) != "" {
			matched.ImageURL = strings.TrimSpace(sale.ImageURL)
		}
	}

	if matched == nil && targetSeries == excitingMacaronSeriesName {
		for _, name := range excitingMacaronCanonicalOrder {
			if popmartItemSlug(targetSeries, name) != itemSlug {
				continue
			}
			matched = &popmartSeriesItem{
				Title:     name,
				ImageURL:  selected.ImageURL,
				AvgPrice:  "",
				SoldCount: 0,
				Href:      "/category/pop-mart/" + selected.Slug + "/" + seriesSlug + "/" + itemSlug,
			}
			break
		}
	}

	if matched == nil {
		return responseValidationError(c, "invalid item")
	}

	seriesLabel := targetSeries
	if seriesSlug == "all-series" {
		seriesLabel = "All Series"
	}

	return responseJSON(c, popmartSeriesItemResponse{
		IP:     *selected,
		Series: seriesLabel,
		Item:   *matched,
	})
}
