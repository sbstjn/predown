package main

import (
	blackfriday "gopkg.in/russross/blackfriday.v2"
)

const (
	blackfridayOptions = blackfriday.CommonExtensions | blackfriday.AutoHeadingIDs
)

func toHTML(data []byte) []byte {
	return blackfriday.Run(
		data,
		blackfriday.WithExtensions(blackfridayOptions),
	)
}
