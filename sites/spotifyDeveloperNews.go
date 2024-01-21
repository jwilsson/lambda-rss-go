package sites

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const BASE_URL = "https://developer.spotify.com/community"
const NAME = "Spotify Developer News"

func fetch() ([]Article, error) {
	doc, err := fetchSite(BASE_URL)
	if err != nil {
		return nil, err
	}

	articleElements := doc.Find("main [role='listitem']")
	articles := make([]Article, articleElements.Length())

	articleElements.Each(func(i int, s *goquery.Selection) {
		dateElement := s.Find("b")
		linkElement := s.Find("a")
		typeElement := s.Find("[type]")

		postDate := strings.Trim(dateElement.Text(), " ")
		postTitle := strings.Trim(linkElement.Text(), " ")
		postType := strings.Trim(typeElement.Text(), " ")
		postUrl, _ := url.JoinPath(BASE_URL, linkElement.AttrOr("href", ""))

		articles[i] = Article{
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
