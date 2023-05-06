package tmplt

import (
	"embed"
)

var (

	//go:embed static/*
	StaticEmbed embed.FS

	//go:embed templates1/*.html
	TmplEmbed embed.FS
)
