package handler

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
	"toychart/errcode"
	"toychart/kit/ebay"

	"github.com/labstack/echo/v4"
)

type popmartSeriesCard struct {
	Series      string `json:"series"`
	SeriesSlug  string `json:"seriesSlug"`
	ImageURL    string `json:"imageUrl"`
	Count       int    `json:"count"`
	SamplePrice string `json:"samplePrice"`
	Href        string `json:"href"`
}

type popmartSeriesResponse struct {
	IP    popmartDirectoryItem `json:"ip"`
	Cards []popmartSeriesCard  `json:"cards"`
}

type popmartSeriesRule struct {
	series   string
	patterns []*regexp.Regexp
}

var popmartSeriesRules = []popmartSeriesRule{
	{
		series: "THE MONSTERS - Exciting Macaron Vinyl Face",
		patterns: []*regexp.Regexp{
			regexp.MustCompile(`(?i)\bLABUBU\b.*\bMACARON\b`),
			regexp.MustCompile(`(?i)\bEXCITING\s+MACARON\b`),
			regexp.MustCompile(`(?i)\bMACARON\s+V\d+\b`),
		},
	},
	{
		series: "THE MONSTERS x One Piece Series",
		patterns: []*regexp.Regexp{
			regexp.MustCompile(`(?i)\bTHE\s+MONSTERS\b.*\bONE\s*PIECE\b`),
			regexp.MustCompile(`(?i)\bLABUBU\b.*\bONE\s*PIECE\b`),
		},
	},
	{
		series: "POP BEAN Pajama Cross Dressing Series",
		patterns: []*regexp.Regexp{
			regexp.MustCompile(`(?i)\bPAJAMA\s+CROSS\s+DRESSING\b`),
		},
	},
	{
		series: "THE MONSTERS x Coca-Cola Series",
		patterns: []*regexp.Regexp{
			regexp.MustCompile(`(?i)\bTHE\s+MONSTERS\b.*\bCOCA[\s-]?COLA\b`),
			regexp.MustCompile(`(?i)\bLABUBU\b.*\bCOCA[\s-]?COLA\b`),
		},
	},
	{
		series: "THE MONSTERS x How to Train Your Dragon Series",
		patterns: []*regexp.Regexp{
			regexp.MustCompile(`(?i)\bHOW\s+TO\s+TRAIN\s+YOUR\s+DRAGON\b`),
		},
	},
	{
		series: "THE MONSTERS x SpongeBob Series",
		patterns: []*regexp.Regexp{
			regexp.MustCompile(`(?i)\bSPONGE\s*BOB\b`),
		},
	},
	{
		series: "THE MONSTERS x Kow Yokoyama Series",
		patterns: []*regexp.Regexp{
			regexp.MustCompile(`(?i)\bKOW\s+YOKOYAMA\b`),
		},
	},
	{
		series: "THE MONSTERS - Have a Seat Series",
		patterns: []*regexp.Regexp{
			regexp.MustCompile(`(?i)\bHAVE\s+A\s+SEAT\b`),
		},
	},
	{
		series: "THE MONSTERS - Big into Energy Series",
		patterns: []*regexp.Regexp{
			regexp.MustCompile(`(?i)\bBIG\s+INTO\s+ENERGY\b`),
		},
	},
}

func sanitizePopmartTitle(value string) string {
	result := strings.TrimSpace(value)
	result = regexp.MustCompile(`(?i)\bAUTHENTIC\b`).ReplaceAllString(result, "")
	result = regexp.MustCompile(`(?i)Opens in a new window or tab`).ReplaceAllString(result, "")
	result = regexp.MustCompile(`\s+`).ReplaceAllString(result, " ")
	return strings.TrimSpace(result)
}

func classifyPopmartSeries(title string) string {
	cleaned := sanitizePopmartTitle(title)
	if cleaned == "" {
		return ""
	}

	for _, rule := range popmartSeriesRules {
		for _, pattern := range rule.patterns {
			if pattern.MatchString(cleaned) {
				return rule.series
			}
		}
	}

	withoutPrefix := regexp.MustCompile(`(?i)^(authentic\s+)?(pop\s*mart\s+)?`).ReplaceAllString(cleaned, "")
	withoutCharacter := regexp.MustCompile(`(?i)^(labubu|the monsters|skullpanda|dimoo|hirono|molly|crybaby|pucky|hacipupu)\s+`).ReplaceAllString(withoutPrefix, "")
	if strings.TrimSpace(withoutCharacter) != "" {
		return strings.TrimSpace(withoutCharacter)
	}
	return cleaned
}

func popmartSeriesToSlug(value string) string {
	if strings.Contains(strings.ToLower(value), "exciting macaron") {
		return "labubu-exciting-macaron"
	}
	s := strings.ToLower(sanitizePopmartTitle(value))
	s = strings.ReplaceAll(s, "&", " and ")
	s = regexp.MustCompile(`[^a-z0-9]+`).ReplaceAllString(s, "-")
	s = strings.Trim(s, "-")
	return s
}

func (h *Handler) PopmartIPSeries(c echo.Context) error {
	ipSlug := strings.TrimSpace(c.Param("ip"))
	if ipSlug == "" {
		return responseValidationError(c, "ip is required")
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

	maxPages := 3
	if raw := strings.TrimSpace(c.QueryParam("maxPages")); raw != "" {
		if parsed, err := strconv.Atoi(raw); err == nil {
			maxPages = parsed
		}
	}

	result, err := ebay.ScrapeSoldToyPrices(strings.TrimSpace(selected.Keyword), "", maxPages)
	if err != nil {
		return responseError(c, errcode.InternalServerError)
	}

	seriesMap := map[string]popmartSeriesCard{}
	for _, item := range result {
		series := classifyPopmartSeries(item.Title)
		if strings.TrimSpace(series) == "" {
			series = "Other Series"
		}

		existing, ok := seriesMap[series]
		if !ok {
			image := strings.TrimSpace(item.ImageURL)
			if image == "" {
				image = selected.ImageURL
			}
			seriesMap[series] = popmartSeriesCard{
				Series:      series,
				SeriesSlug:  popmartSeriesToSlug(series),
				ImageURL:    image,
				Count:       1,
				SamplePrice: strings.TrimSpace(item.Price),
				Href:        "/category/pop-mart/" + selected.Slug + "/" + popmartSeriesToSlug(series),
			}
			continue
		}

		existing.Count++
		if strings.TrimSpace(existing.ImageURL) == "" && strings.TrimSpace(item.ImageURL) != "" {
			existing.ImageURL = strings.TrimSpace(item.ImageURL)
		}
		seriesMap[series] = existing
	}

	cards := make([]popmartSeriesCard, 0, len(seriesMap))
	for _, card := range seriesMap {
		cards = append(cards, card)
	}
	sort.Slice(cards, func(i, j int) bool {
		if cards[i].Count != cards[j].Count {
			return cards[i].Count > cards[j].Count
		}
		return cards[i].Series < cards[j].Series
	})

	resp := popmartSeriesResponse{
		IP:    *selected,
		Cards: cards,
	}
	return responseJSON(c, resp)
}
