package sites

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/golang-module/carbon"
)

type Article struct {
	Date        string
	Description string
	Link        string
	Title       string
}

func toRssDateFromFormat(date string, fromFormat string) string {
	return carbon.ParseByFormat(date, fromFormat).ToRfc2822String()
}

func fetchSite(url string) (*goquery.Document, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	return goquery.NewDocumentFromReader(res.Body)
}
