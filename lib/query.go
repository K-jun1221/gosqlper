package lib

import (
	"database/sql"
	"fmt"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
)

// QueryRow only allowed string field
func QueryRow(db *sql.DB, sql SelectSQL, obj interface{}) error {
	mapping, err := tagCheck(sql.Select, obj)
	if err != nil {
		return err
	}
	rawSQL, err := sql.MakeSQL()
	if err != nil {
		return err
	}

	columns := make([]interface{}, len(sql.Select))
	for i := 0; i < len(sql.Select); i++ {
		var str string
		columns[i] = &str
	}
	err = db.QueryRow(rawSQL).Scan(columns...)
	if err != nil {
		return err
	}

	v := reflect.Indirect(reflect.ValueOf(obj))

	for i, column := range columns {
		subv := v.Field(mapping[i])
		str, _ := column.(*string)
		subv.SetString(*str)
	}

	return nil
}

// Query only allowed string field
func Query(db *sql.DB, sql SelectSQL, obj interface{}, objs interface{}) error {

	mapping, err := tagCheck(sql.Select, obj)
	if err != nil {
		return err
	}
	rawSQL, err := sql.MakeSQL()
	if err != nil {
		return err
	}

	v := reflect.Indirect(reflect.ValueOf(objs))
	rows, err := db.Query(rawSQL)
	if err != nil {
		return err
	}

	idx := 0
	vi := reflect.Indirect(reflect.ValueOf(obj))

	for rows.Next() {
		columns := make([]interface{}, len(sql.Select))
		for i := 0; i < len(sql.Select); i++ {
			var str string
			columns[i] = &str
		}
		err = rows.Scan(columns...)
		if err != nil {
			return err
		}

		// Indexに合わせて拡張
		if idx >= v.Cap() {
			newv := reflect.MakeSlice(v.Type(), v.Len(), idx+1)
			reflect.Copy(newv, v)
			v.Set(newv)
			v.SetCap(idx + 1)
		}
		if idx >= v.Len() {
			v.SetLen(idx + 1)
		}
		vindex := v.Index(idx)

		for i, column := range columns {
			subv := vi.Field(mapping[i])
			str, _ := column.(*string)
			subv.SetString(*str)
		}

		vindex.Set(vi)
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

func tagCheck(columns []string, obj interface{}) ([]int, error) {
	idxMap := []int{}

	t := reflect.Indirect(reflect.ValueOf(obj)).Type()

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
