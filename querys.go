package gosqlper

import (
	"database/sql"
	"errors"
	"reflect"
	"strings"

	// import mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// QueryRow only allowed string field
func QueryRow(db *sql.DB, sql string, obj interface{}) error {
	cns, err := columnGetter(sql)
	if err != nil {
		return err
	}

	// create reflect.Value
	v := reflect.Indirect(reflect.ValueOf(obj))

	// get tag mapping list
	tm, err := tagMappingGetter(cns, v)
	if err != nil {
		return err
	}

	// call scan
	columns := make([]interface{}, len(cns))
	for i := 0; i < len(cns); i++ {
		var str string
		columns[i] = &str
	}

	err = db.QueryRow(sql).Scan(columns...)
	if err != nil {
		return err
	}

	for i, column := range columns {
		subv := v.Field(tm[i])
		str, ok := column.(*string)
		if !ok {
			return errors.New("could not cast interface{} to *string type")
		}
		subv.SetString(*str)
	}

	return nil
}

// Query only allowed string field
func Query(db *sql.DB, sql string, objs interface{}) error {

	cns, err := columnGetter(sql)
	if err != nil {
		return err
	}

	// create reflect.Value
	v := reflect.Indirect(reflect.ValueOf(objs))
	if 1 > v.Cap() {
		newv := reflect.MakeSlice(v.Type(), v.Len(), 1)
		reflect.Copy(newv, v)
		v.Set(newv)
	}
	if 1 > v.Len() {
		v.SetLen(1)
	}
	vi := v.Index(0)

	// get tag mapping list
	tm, err := tagMappingGetter(cns, vi)
	if err != nil {
		return err
	}

	// call query
	rows, err := db.Query(sql)
	if err != nil {
		return err
	}

	idx := 0
	for rows.Next() {
		columns := make([]interface{}, len(cns))
		for i := 0; i < len(cns); i++ {
			var str string
			columns[i] = &str
		}

		// call query
		err = rows.Scan(columns...)
		if err != nil {
			return err
		}

		// expand as index
		if idx >= v.Cap() {
			newv := reflect.MakeSlice(v.Type(), v.Len(), idx+1)
			reflect.Copy(newv, v)
			v.Set(newv)

		}
		if idx >= v.Len() {
			v.SetLen(idx + 1)
		}
		vindex := v.Index(idx)

		for i, column := range columns {
			subv := vindex.Field(tm[i])
			str, ok := column.(*string)
			if !ok {
				return errors.New("could not cast interface{} to *string type")
			}
			subv.SetString(*str)
		}
		idx++
	}
	return nil
}

func Exec(db *sql.DB, sql string) (sql.Result, error) {

	// call exec
	return db.Exec(sql)
}

func columnGetter(sql string) ([]string, error) {
	lowerSQL := strings.ToLower(sql)
	s := []string{}

	if strings.Contains(lowerSQL, "*") {
		return s, errors.New("`*` is not allowed in gosqlper")
	}

	sidx := strings.Index(lowerSQL, "select ")
	fidx := strings.Index(lowerSQL, " from ")
	columns := strings.TrimSpace(lowerSQL[sidx+6 : fidx])

	c := strings.Split(columns, ",")

	for _, v := range c {
		s = append(s, strings.TrimSpace(v))
	}
	return s, nil
}

func tagMappingGetter(columns []string, v reflect.Value) ([]int, error) {
	tm := []int{}

	t := v.Type()

	for i := 0; i < len(columns); i++ {
		tm = append(tm, -1)
	}

	for i, v := range columns {
		for j := 0; j < t.NumField(); j++ {
			if dbTag, ok := t.Field(j).Tag.Lookup("db"); ok {
				if dbTag == v && tm[i] == -1 {
					tm[i] = j
					break
				}
			}
		}

		if tm[i] == -1 {
			return tm, errors.New("tag `" + v + "` is missing for mapping row columns")
		}
	}
	return tm, nil
}
