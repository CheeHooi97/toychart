package ebay

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

type CardData struct {
	ItemID     string `json:"itemId"`
	Title      string `json:"title"`
	Price      string `json:"price"`
	ImageURL   string `json:"imageUrl"`
	ItemWebURL string `json:"itemWebUrl"`
	Subtitle   string `json:"subtitle"`
	Caption    string `json:"caption"`
}

type SoldItem struct {
	ItemID      string `json:"itemId"`
	Title       string `json:"title"`
	Price       string `json:"price"`
	Currency    string `json:"currency"`
	ImageURL    string `json:"imageUrl"`
	ItemWebURL  string `json:"itemWebUrl"`
	ItemEndDate string `json:"itemEndDate"`
}

func ScrapEbaySoldCards(baseUrl string, maxPages int) ([]CardData, error) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-blink-features", "AutomationControlled"),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("no-sandbox", true),
		chromedp.UserAgent(
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) "+
				"AppleWebKit/537.36 (KHTML, like Gecko) "+
				"Chrome/122.0.0.0 Safari/537.36",
		),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 15*time.Minute)
	defer cancel()

	results := make([]CardData, 0)
	seen := make(map[string]bool)

	for page := 1; page <= maxPages; page++ {
		pageUrl := fmt.Sprintf("%s&_pgn=%d", baseUrl, page)
		fmt.Printf("\n--- Page %d ---\n%s\n", page, pageUrl)

		var pageItems []CardData

		err := chromedp.Run(ctx,
			chromedp.Navigate(pageUrl),

			// Wait until results load
			chromedp.WaitVisible(`ul.srp-results li.s-card`, chromedp.ByQuery),

			chromedp.Evaluate(`
			(() => {
				const cleanText = (value) => {
					if (!value) return "";
					return value
						.replace(/\bAUTHENTIC\b/gi, "")
						.replace(/Opens in a new window or tab/gi, "")
						.replace(/\s+/g, " ")
						.trim();
				};

				const rows = Array.from(
					document.querySelectorAll('ul.srp-results li.s-card')
				);

				return rows.map(row => {
					const linkEl  = row.querySelector('a.s-card__link');
					const titleEl = row.querySelector('.s-card__title');
					const priceEl = row.querySelector('.s-card__price');
					const imgEl   = row.querySelector('.s-card__image img');
					const subtitleEl = row.querySelector('.s-card__subtitle');
					const captionEl = row.querySelector('.s-card__caption');

					let link = linkEl ? linkEl.href : "";
					let itemId = "";

					if (link) {
						const m = link.match(/\/itm\/(\d+)/);
						if (m) itemId = m[1];
					}

					return {
						itemId: itemId,
						title: cleanText(titleEl ? titleEl.innerText : ""),
						price: priceEl ? priceEl.innerText.trim() : "N/A",
						imageUrl: imgEl ? imgEl.src : "",
						itemWebUrl: link,
						subtitle: cleanText(subtitleEl ? subtitleEl.innerText : ""),
						caption: cleanText(captionEl ? captionEl.innerText : "")
					};
				});
			})()
			`, &pageItems),
		)

		if err != nil {
			return results, err
		}

		// No items = end
		if len(pageItems) == 0 {
			fmt.Println("No more items found. Stopping.")
			break
		}

		newCount := 0
		for _, item := range pageItems {
			if item.ItemID == "" {
				continue
			}

			// Filter noise cards
			if strings.Contains(item.Title, "Shop on eBay") {
				continue
			}

			if strings.TrimSpace(item.ItemWebURL) == "" {
				item.ItemWebURL = "https://www.ebay.com/itm/" + strings.TrimSpace(item.ItemID)
			}

			if !seen[item.ItemID] {
				seen[item.ItemID] = true
				results = append(results, item)
				newCount++
			}
		}

		fmt.Printf("Page %d | Fetched: %d | New: %d | Total: %d\n",
			page, len(pageItems), newCount, len(results),
		)

		// Small delay to avoid rate-limit
		time.Sleep(2 * time.Second)
	}

	fmt.Printf("\n=== DONE ===\nTotal unique items: %d\n", len(results))
	return results, nil
}

// https://www.ebay.com/sch/i.html?_nkw=Labubu&LH_Complete=1&LH_Sold=1

func BuildSoldSearchURL(keyword string) string {
	params := url.Values{}
	params.Set("_nkw", keyword)
	params.Set("LH_Complete", "1")
	params.Set("LH_Sold", "1")
	return "https://www.ebay.com/sch/i.html?" + params.Encode()
}

func ScrapeSoldToyPrices(keyword, pageURL string, maxPages int) ([]CardData, error) {
	targetURL := pageURL
	if targetURL == "" {
		targetURL = BuildSoldSearchURL(keyword)
	}

	if maxPages <= 0 {
		maxPages = 1
	}
	if maxPages > 10 {
		maxPages = 10
	}

	return ScrapEbaySoldCards(targetURL, maxPages)
}

func SearchSoldToyPrices(keyword string) ([]SoldItem, error) {
	result, err := SearchToyPrices(keyword)
	if err != nil {
		return nil, err
	}

	soldItems := make([]SoldItem, 0, len(result.ItemSummaries))
	now := time.Now().UTC()
	for _, item := range result.ItemSummaries {
		if item.ItemEndDate == "" {
			continue
		}

		// Ensure the listing has actually ended.
		endTime, err := time.Parse(time.RFC3339, item.ItemEndDate)
		if err != nil {
			continue
		}
		if endTime.After(now) {
			continue
		}

		// Keep auction sold entries only.
		isAuction := false
		for _, opt := range item.BuyingOptions {
			if strings.EqualFold(opt, "AUCTION") {
				isAuction = true
				break
			}
		}
		if !isAuction {
			continue
		}

		priceValue := strings.TrimSpace(item.Price.Value)
		priceCurrency := strings.TrimSpace(item.Price.Currency)
		if priceValue == "" {
			priceValue = strings.TrimSpace(item.CurrentBidPrice.Value)
			priceCurrency = strings.TrimSpace(item.CurrentBidPrice.Currency)
		}
		if priceValue == "" {
			continue
		}
		itemWebURL := strings.TrimSpace(item.ItemWebUrl)
		if itemWebURL == "" && strings.TrimSpace(item.ItemID) != "" {
			itemWebURL = "https://www.ebay.com/itm/" + strings.TrimSpace(item.ItemID)
		}

		soldItems = append(soldItems, SoldItem{
			ItemID:      item.ItemID,
			Title:       item.Title,
			Price:       priceValue,
			Currency:    priceCurrency,
			ImageURL:    item.Image.ImageUrl,
			ItemWebURL:  itemWebURL,
			ItemEndDate: item.ItemEndDate,
		})
	}

	return soldItems, nil
}
