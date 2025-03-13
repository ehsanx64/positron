package ui

import "embed"

//go:embed index.html
var Main embed.FS

//go:embed assets/*
var Assets embed.FS
