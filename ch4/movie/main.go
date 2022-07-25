package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1976, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	// ...
}

func main() {
	{
		// !+Marshal
		data, err := json.Marshal(movies)
		if err != nil {
			log.Fatalf("JSON marshaling failed: %s", err)
		}
		fmt.Printf("%s\n", data)
		// ！-Marshal
	}

	{
		//!+MarshalIndent
		data, err := json.MarshalIndent(movies, "", "	")
		if err != nil {
			log.Fatalf("JSON marshaling failed: %s", err)
		}
		fmt.Printf("%s\n", data)
		//!-MarshalIndent

		//!+Unmarshal
		var titles []struct{ Title string }
		if err := json.Unmarshal(data, &titles); err != nil {
			log.Fatalf("JSON unmarshaling failed: %s", err)
		}
		fmt.Println(titles) // [{Casablanca} {Cool Hand Luke} {Bullitt}]
		//!-unmarshal
	}
}

/*
[{"Title":"Casablanca","released":1942,"Actors":["Humphrey Bogart","Ingrid Bergman"]},{"Title":"Cool Hand Luke","released":1976,"color":true,"Actors":["Paul Newman"]},{"T
itle":"Bullitt","released":1968,"color":true,"Actors":["Steve McQueen","Jacqueline Bisset"]}]
*/

/*
[
        {
                "Title": "Casablanca",
                "released": 1942,
                "Actors": [
                        "Humphrey Bogart",
                        "Ingrid Bergman"
                ]
        },
        {
                "Title": "Cool Hand Luke",
                "released": 1976,
                "color": true,
                "Actors": [
                        "Paul Newman"
                ]
        },
        {
                "Title": "Bullitt",
                "released": 1968,
                "color": true,
                "Actors": [
                        "Steve McQueen",
                        "Jacqueline Bisset"
                ]
        }
]

*/
