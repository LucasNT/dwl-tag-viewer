package entities

type DwlTag struct {
	Monitor  string
	Tag      uint8
	NElement uint8
	Focus    bool
	Viewed   bool
}

type DwlTagsInformation struct {
	Tags []DwlTag
}
