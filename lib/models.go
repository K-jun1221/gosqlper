package lib

import (
	"errors"
	"strings"
)

type SQLStatement interface {
	MakeSQL() (string, error)
}

// SelectSQL @required: Select, From @optional: Join, Where
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

// UpdateSQL @required: Into, Values
type InsertSQL struct {
	Into   string
	Values string
}

func (s *InsertSQL) MakeSQL() (string, error) {
	if s.Into == "" || s.Values == "" {
		return "", errors.New("lack required args")
	}

	sql := "INSERT INTO " + s.Into + " VALUES " + s.Values

	return sql, nil
}

// UpdateSQL @required: Update, Set @optional: Where
type UpdateSQL struct {
	Update string
	Set    string
	Where  string
}

func (s *UpdateSQL) MakeSQL() (string, error) {
	if s.Update == "" || s.Set == "" {
		return "", errors.New("lack required args")
	}

	sql := "UPDATE " + s.Update + " SET " + s.Set

	if s.Where != "" {
		sql += " WHERE " + s.Where
	}

	return sql, nil
}

// DeleteSQL @required: From @optional: Where
type DeleteSQL struct {
	From  string
	Where string
}

func (s *DeleteSQL) MakeSQL() (string, error) {
	if s.From == "" {
		return "", errors.New("lack required args")
	}

	sql := "DELETE FROM " + s.From

	if s.Where != "" {
		sql += " WHERE " + s.Where
	}

	return sql, nil
}
