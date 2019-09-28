package gosqlper

import "errors"

// SelectSQL, @required: Select, From @optional: Join, Where
type SelectSQL struct {
	Select string
	From   string
	Join   string
	Where  string
}

func (s *SelectSQL) MakeSQL() (string, error) {
	if s.Select == "" || s.From == "" {
		return "", errors.New("lack required args")
	}

	sql := "SELECT " + s.Select + " FROM " + s.From

	if s.Join != "" {
		sql += " JOIN " + s.Join
	}

	if s.Where != "" {
		sql += " WHERE " + s.Where
	}

	return sql, nil
}
