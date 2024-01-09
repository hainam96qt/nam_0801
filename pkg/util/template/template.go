package template

import (
	"bytes"
	"html/template"
	"strconv"
	"strings"
)

func TemplateSQL(sqlString string, conditionValues map[string]interface{}) (string, error) {
	tmpl, err := template.New("baseQuery").Funcs(template.FuncMap{
		"join":      strings.Join,
		"joinInt32": joinInt32,
	}).Parse(sqlString)
	if err != nil {
		return "", err
	}

	var queryBuffer bytes.Buffer
	err = tmpl.Execute(&queryBuffer, conditionValues)
	if err != nil {
		return "", err
	}

	return queryBuffer.String(), nil
}

func joinInt32(elems []int32, sep string) string {
	var strArr []string
	for _, v := range elems {
		strArr = append(strArr, strconv.Itoa(int(v)))
	}
	return strings.Join(strArr, sep)
}
