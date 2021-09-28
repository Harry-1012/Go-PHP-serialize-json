package main

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/haoleiqin/phpserialize"
)

type serObj struct {
}
type serializeStruct struct {
	Id    int64
	Price int64
	Name  string
}

func (serObj *serObj) Phpserialize2Json(serialize string) (jsonres string) {
	var serializeMap []interface{}
	err := phpserialize.Unmarshal([]byte(serialize), &serializeMap)
	if err != nil {
		fmt.Println(err)
	}
	var convertData serializeStruct
	var jsonMap []serializeStruct
	var jsonRes string
	listKey := 0
	for _, list := range serializeMap {
		listKey += 1
		keyIndexNumCount := 0
		for vkey, value := range list.(map[interface{}]interface{}) {
			keyIndexNumCount += 1
			switch vkey {
			case "id":
				convertData.Id = value.(int64)
			case "price":
				convertData.Price = value.(int64)
			case "name":
				convertData.Name = value.(string)
				paramNums := 3
				if keyIndexNumCount == paramNums {
					jsonMap = append(jsonMap, convertData)
					if listKey >= len(serializeMap) {
						data, errsjon := json.Marshal(jsonMap)
						if errsjon != nil {
							fmt.Println(errsjon)
						}
						enEscapeUrl, queryUnescapeErr := url.QueryUnescape(string(data))
						if queryUnescapeErr != nil {
							fmt.Println(queryUnescapeErr)
						}
						jsonRes = enEscapeUrl
					}
				}
			}
		}
	}
	return jsonRes
}
