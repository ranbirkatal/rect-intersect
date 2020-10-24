package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/geo/r2"
	"github.com/ranbirkatal/rect-intersect/pkg/shapes"
)

func main() {
	var input struct {
		Rects []shapes.RawRect `json:"rects"`
	}
	argLen := len(os.Args)

	if argLen <= 1 {
		log.Println("Expected a path for config file")
		return
	}

	path := os.Args[1]

	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("Failed to read file %s", err.Error())
		return
	}
	err = json.Unmarshal(data, &input)
	if err != nil {
		log.Printf("Failed to unmarshal data %s", err.Error())
		return
	}

	var rects []r2.Rect
	for idx, v := range input.Rects {
		rect := v.BuildR2Rect()
		if idx > 10 {
			break
		}
		rects = append(rects, rect)
		infoString := v.BuildInfoText()
		fmt.Printf("%+v: %+v\n", idx+1, infoString)
	}

	shapes.GenerateIntersections(rects)

}
