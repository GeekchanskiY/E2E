package static

import (
	"embed"
)

//go:embed files/*
var Fs embed.FS
