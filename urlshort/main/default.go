package main

var (
	pathsToUrls = map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	yaml = `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`

	json = `
[
	{
		"path": "/serediuk",
		"url": "https://github.com/serediukit"
	},
	{	
		"path": "/serediuk-go",
		"url": "https://github.com/serediukit/gophercises"
	}
]
`
)
