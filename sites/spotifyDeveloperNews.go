package sites

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const BASE_URL = "https://developer.spotify.com"
const NAME = "Spotify Developer News"

func fetch() ([]Article, error) {
	siteUrl, _ := url.JoinPath(BASE_URL, "/community")
	doc, err := fetchSite(siteUrl)

	if err != nil {
		return nil, err
	}

	articleElements := doc.Find("main [role='listitem']")
	articles := goquery.Map(articleElements, func(_ int, s *goquery.Selection) Article {
		dateElement := s.Find("b")
		linkElement := s.Find("a")
		typeElement := s.Find("[type]")

		postDate := strings.Trim(dateElement.Text(), " ")
		postTitle := strings.Trim(linkElement.Text(), " ")
		postType := strings.Trim(typeElement.Text(), " ")
		postUrl, _ := url.JoinPath(BASE_URL, linkElement.AttrOr("href", ""))

		return Article{
			Date:        toRssDateFromFormat(postDate, "F j, Y"),
			Description: fmt.Sprintf("%s - %s", postType, postTitle),
			Link:        postUrl,
			Title:       postTitle,
		}
	})

	return articles, nil
}

func spotifyDeveloperNews() *Site {
	return &Site{
		Fetch: fetch,
		Name:  NAME,
		Url:   BASE_URL,
	}
}
