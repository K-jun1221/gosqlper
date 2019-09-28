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

func Query(db *sql.DB, sql SelectSQL, obj interface{}, objs interface{}) error {

	mapping, err := tagCheck(sql.Select, obj)
	if err != nil {
		return err
	}
	rawSQL, err := sql.MakeSQL()
	if err != nil {
		return err
	}

	fmt.Println("mapping: ", mapping)
	fmt.Println("rawSQL: ", rawSQL)

	rows, err := db.Query(rawSQL)
	if err != nil {
		return err
	}

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

		v := reflect.Indirect(reflect.ValueOf(obj))

		for i, column := range columns {
			subv := v.Field(mapping[i])
			str, _ := column.(*string)
			subv.SetString(*str)
		}

		// TODO いい感じに配列で返す
		fmt.Println(v)
		// objs = append(objs, obj)

		// TODO END いい感じに配列で返す
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
