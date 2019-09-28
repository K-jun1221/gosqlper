package lib

import (
	"database/sql"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
)

// SelectRow only allowed string field
func SelectRow(db *sql.DB, sql SelectSQL, obj interface{}) error {

	mapping, err := tagCheck(sql.Select, obj)
	if err != nil {
		return err
	}

	rawSQL, err := sql.MakeSQL()
	if err != nil {
		return err
	}

	// TODO 可変長の引数をいい感じに渡したい
	columns := []string{}
	for i := 0; i < len(sql.Select); i++ {
		str := ""
		columns = append(columns, str)
	}
	err = db.QueryRow(rawSQL).Scan(&columns[0], &columns[1])
	// TODO END 可変長の引数をいい感じに渡したい

	v := reflect.Indirect(reflect.ValueOf(obj))

	for i, str := range columns {
		subv := v.Field(mapping[i])
		subv.SetString(str)
	}

	return nil
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
