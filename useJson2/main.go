package main

import (
	"encoding/json"
	"fmt"
)

/*type sensorread struct {
	Name        string `json:"name"`
	Capacity    int    `json:"capacity"`
	Time        string `json:"time"`
	Information info   `json:"info"`
}

type info struct {
	Description string `json:"desc"`
}*/

func main() {

	fmt.Println("Helo World!")

	jsonSTring := `{"name": "battery sensor", "capacity": 40, "time": "12:09:06", "info":{"desc":"This is description"}}`

	//var reading sensorread

	var reading map[string]interface{}

	err := json.Unmarshal([]byte(jsonSTring), &reading)

	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf("%+v \n", reading)

}
