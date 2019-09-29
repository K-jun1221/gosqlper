package gosqlper

import (
	"errors"
	"strings"
)

// SQLStatement @required: MakeSQL() method
type SQLStatement interface {
	MakeSQL() (string, error)
}

// SelectSQL @required: Select, From @optional: Join, Where
type SelectSQL struct {
	Select []string
	From   string
	Join   string
	Where  string
	Others string
}

// MakeSQL @required: Select, From @optional: Join, Where
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

	if s.Others != "" {
		sql += s.Others
	}

	return sql, nil
}

// InsertSQL @required: Into, Values
type InsertSQL struct {
	Into   string
	Values string
	Others string
}

// MakeSQL @required: Into, Values
func (s *InsertSQL) MakeSQL() (string, error) {
	if s.Into == "" || s.Values == "" {
		return "", errors.New("lack required args")
	}

	sql := "INSERT INTO " + s.Into + " VALUES " + s.Values

	if s.Others != "" {
		sql += s.Others
	}

	return sql, nil
}

// UpdateSQL @required: Update, Set @optional: Where
type UpdateSQL struct {
	Update string
	Set    string
	Where  string
	Others string
}

// MakeSQL @required: Update, Set @optional: Where
func (s *UpdateSQL) MakeSQL() (string, error) {
	if s.Update == "" || s.Set == "" {
		return "", errors.New("lack required args")
	}

	sql := "UPDATE " + s.Update + " SET " + s.Set

	if s.Where != "" {
		sql += " WHERE " + s.Where
	}

	if s.Others != "" {
		sql += s.Others
	}

	return sql, nil
}

// DeleteSQL @required: From @optional: Where
type DeleteSQL struct {
	From   string
	Where  string
	Others string
}

// MakeSQL @required: From @optional: Where
func (s *DeleteSQL) MakeSQL() (string, error) {
	if s.From == "" {
		return "", errors.New("lack required args")
	}

	sql := "DELETE FROM " + s.From

	if s.Where != "" {
		sql += " WHERE " + s.Where
	}

	if s.Others != "" {
		sql += s.Others
	}

	return sql, nil
}

// CustomSQL @required: SQL
type CustomSQL struct {
	SQL string
}

// MakeSQL @required: SQL
func (s *CustomSQL) MakeSQL() (string, error) {
	if s.SQL == "" {
		return "", errors.New("SQL is empty")
	}

	return s.SQL, nil
}
