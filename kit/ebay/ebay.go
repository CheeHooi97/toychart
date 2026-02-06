package ebay

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

type CardData struct {
	ItemID string
	Title  string
	Price  string
	Image  string
	Link   string
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
				const rows = Array.from(
					document.querySelectorAll('ul.srp-results li.s-card')
				);

				return rows.map(row => {
					const linkEl  = row.querySelector('a.s-card__link');
					const titleEl = row.querySelector('.s-card__title');
					const priceEl = row.querySelector('.s-card__price');
					const imgEl   = row.querySelector('.s-card__image img');

					let link = linkEl ? linkEl.href : "";
					let itemId = "";

					if (link) {
						const m = link.match(/\/itm\/(\d+)/);
						if (m) itemId = m[1];
					}

					return {
						itemId: itemId,
						title: titleEl ? titleEl.innerText.trim() : "",
						price: priceEl ? priceEl.innerText.trim() : "N/A",
						image: imgEl ? imgEl.src : "",
						link: link
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
