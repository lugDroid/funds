package models

type Portfolio struct {
	Categories []Category
	Assets     []Asset
}

type Asset struct {
	Name   string
	Id     int
	Ticker string
	Shares float32
}

type Category struct {
	Id       int
	Name     string
	PercGoal float32
}
