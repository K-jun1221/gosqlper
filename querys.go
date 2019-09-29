package gosqlper

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
)

// QueryRow only allowed string field
func QueryRow(db *sql.DB, sql SelectSQL, obj interface{}) error {

	// create reflect.Value
	v := reflect.Indirect(reflect.ValueOf(obj))

	// get tag mapping list
	tm, err := tagCheck(sql.Select, v)
	if err != nil {
		return err
	}

	// get raw sql statement
	rawSQL, err := sql.MakeSQL()
	if err != nil {
		return err
	}

	// call scan
	columns := make([]interface{}, len(sql.Select))
	for i := 0; i < len(sql.Select); i++ {
		var str string
		columns[i] = &str
	}
	err = db.QueryRow(rawSQL).Scan(columns...)
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
func Query(db *sql.DB, sql SelectSQL, objs interface{}) error {

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
	tm, err := tagCheck(sql.Select, vi)
	if err != nil {
		return err
	}

	// get raw sql statement
	rawSQL, err := sql.MakeSQL()
	if err != nil {
		return err
	}

	// call query
	rows, err := db.Query(rawSQL)
	if err != nil {
		return err
	}

	idx := 0
	for rows.Next() {
		columns := make([]interface{}, len(sql.Select))
		for i := 0; i < len(sql.Select); i++ {
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

func Exec(db *sql.DB, sql SQLStatement) (sql.Result, error) {
	fmt.Println("Exec was called")
	rawSQL, err := sql.MakeSQL()
	if err != nil {
		return nil, err
	}
	return db.Exec(rawSQL)
}

func tagCheck(columns []string, v reflect.Value) ([]int, error) {
	idxMap := []int{}

	t := v.Type()

	for i := 0; i < len(columns); i++ {
		idxMap = append(idxMap, -1)
	}

	for i, v := range columns {
		for j := 0; i < t.NumField(); j++ {
			if dbTag, ok := t.Field(j).Tag.Lookup("db"); ok {
				if dbTag == v && idxMap[i] == -1 {
					idxMap[i] = j
					break
				}
			}
		}
	}
	return idxMap, nil
}
