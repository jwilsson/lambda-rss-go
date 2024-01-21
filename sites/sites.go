package sites

type Site struct {
	Fetch func() ([]Article, error)
	Name  string
	Url   string
}

func GetSite(name string) *Site {
	sites := map[string]*Site{
		"spotify-developer-news": spotifyDeveloperNews(),
	}

	return sites[name]
}
