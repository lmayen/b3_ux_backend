package assets

import "embed"

//go:embed data/*
var DataFS embed.FS

//go:embed images/*
var ImageFS embed.FS
