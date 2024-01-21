package main

import (
	"bytes"
	"lambda-rss/sites"

	_ "embed"
	"text/template"
)

//go:embed templates/rss.xml
var bodyTemplate string

type Feed struct {
	Items []sites.Article
	Name  string
	Url   string
}

func createBody(site sites.Site, bodyBuffer *bytes.Buffer) error {
	articles, err := site.Fetch()
	if err != nil {
		return err
	}

	t, err := template.New("rss").Parse(bodyTemplate)
	if err != nil {
		return err
	}

	return t.Execute(bodyBuffer, Feed{
		Items: articles,
		Name:  site.Name,
		Url:   site.Url,
	})
}
