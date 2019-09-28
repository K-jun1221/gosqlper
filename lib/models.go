package lib

import (
	"errors"
	"strings"
)

// SelectSQL, @required: Select, From @optional: Join, Where
type SelectSQL struct {
	Select []string
	From   string
	Join   string
	Where  string
}

func (s *SelectSQL) MakeSQL() (string, error) {
	if len(s.Select) == 0 || s.From == "" {
		return "", errors.New("lack required args")
	}

	sql := "SELECT " + strings.Join(s.Select, ", ") + " FROM " + s.From

	if s.Join != "" {
		sql += " JOIN " + s.Join
	}

	if s.Where != "" {
		sql += " WHERE " + s.Where
	}

	return sql, nil
}
