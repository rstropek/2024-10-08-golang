package data

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/lib/pq"
)

type Hero struct {
	ID        int       `json:"id"`
	FirstSeen time.Time `json:"firstSeen"`
	Name      string    `json:"name"`
	CanFly    bool      `json:"canFly"`
	RealName  string    `json:"realName,omitempty"`
	Abilities []string  `json:"-"`
	Version   int32     `json:"version"`
}

func (h Hero) MarshalJSON() ([]byte, error) {
	var abilities string
	if len(h.Abilities) > 0 {
		abilities = strings.Join(h.Abilities, ", ")
	}

	type HeroAlias Hero

	helper := struct {
		HeroAlias
		Abilities string `json:"abilities"`
	}{
		HeroAlias: HeroAlias(h),
		Abilities: abilities,
	}

	return json.Marshal(helper)
}

func (h Hero) Validate() error {
	if len(h.Name) < 3 {
		return errors.New("name must be at least 3 characters long")
	}

	// ...

	return nil
}

type HeroRepository struct {
	DB *sql.DB
}

type AppRepository struct {
	Heroes HeroRepository
}

func NewAppRepository(db *sql.DB) *AppRepository {
	return &AppRepository{
		Heroes: HeroRepository{DB: db},
	}
}

func (r *HeroRepository) Insert(hero *Hero) error {
	query := `
	INSERT INTO heroes (name, can_fly, first_seen, realName, abilities)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id`

	queryArgs := []any{
		hero.Name,
		hero.CanFly,
		hero.FirstSeen,
		hero.RealName,
		pq.Array(hero.Abilities),
	}

	return r.DB.QueryRow(query, queryArgs...).Scan(&hero.ID)
}

func (r *HeroRepository) Get(id int) (*Hero, error) {
	query := `SELECT id, first_seen, name, can_fly, realName, abilities, version FROM heroes WHERE id = $1`

	var hero Hero

	err := r.DB.QueryRow(query, id).Scan(
		&hero.ID,
		&hero.FirstSeen,
		&hero.Name,
		&hero.CanFly,
		&hero.RealName,
		pq.Array(&hero.Abilities),
		&hero.Version,
	)
	if err != nil {
		return nil, err
	}

	return &hero, nil
}

func (r *HeroRepository) GetAll(page, pageSize int, sort string) ([]*Hero, error) {
	// SUPER, SUPER, SUPER IMPORTANT: Check the sort parameter to avoid SQL injection!!!!!!!!!!!!!!!!!!!
	query := fmt.Sprintf(`
		SELECT id, first_seen, name, can_fly, realName, abilities, version 
		FROM heroes 
		ORDER BY %s ASC
		LIMIT $2 OFFSET $1`, sort)

	rows, err := r.DB.Query(query, page*pageSize, pageSize)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	heroes := []*Hero{}

	for rows.Next() {
		var hero Hero

		err := rows.Scan(
			&hero.ID,
			&hero.FirstSeen,
			&hero.Name,
			&hero.CanFly,
			&hero.RealName,
			pq.Array(&hero.Abilities),
			&hero.Version,
		)
		if err != nil {
			return nil, err
		}

		heroes = append(heroes, &hero)
	}

	return heroes, nil
}
