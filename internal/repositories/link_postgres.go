package repositories

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/w1lay/Simple-URL-Shortener/internal/dtos"
)

type LinkPostgresRepository struct {
	db *sqlx.DB
}

func NewLinkPostgresRepository(db *sqlx.DB) *LinkPostgresRepository {
	return &LinkPostgresRepository{db: db}
}

func (r *LinkPostgresRepository) SaveLinkRecord(dto dtos.Link) error {
	// ToDO: сделать нормально с аргументами вида $1
	query := fmt.Sprintf(
		`INSERT INTO links (short_link, long_link) VALUES ('%s', '%s');`,
		dto.ShortLink,
		dto.LongLink,
	)
	fmt.Println(query)
	_, err := r.db.Exec(query)
	return err
}

func (r *LinkPostgresRepository) FindLinkByShortLink(shortLink string) (*dtos.Link, error) {
	query := `SELECT * FROM links WHERE short_link = '$1';`
	var link dtos.Link

	err := r.db.Get(&link, query, shortLink)
	if err != nil {
		return nil, err
	}

	return &link, nil
}
