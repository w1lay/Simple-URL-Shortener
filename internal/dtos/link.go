package dtos

type Link struct {
	ShortLink string `db:"short_link" json:"short_link"`
	LongLink  string `db:"long_link"  json:"long_link"`
}
