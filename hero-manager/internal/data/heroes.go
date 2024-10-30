package data

import (
	"encoding/json"
	"strings"
	"time"
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
