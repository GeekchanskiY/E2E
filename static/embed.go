package static

import (
	"embed"
)

//go:embed css/* fonts/* icons/* img/* js/*
//go:embed site.webmanifest
var Fs embed.FS
