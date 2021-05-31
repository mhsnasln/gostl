package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type File struct {
	Name   string  `json:"name"`
	Facets []Facet `json:"facets"`
}

type Facet struct {
	Normal  string   `json:"normal"`
	Xaxis   float64  `json:"x"`
	Yaxis   float64  `json:"y"`
	Zaxis   float64  `json:"z"`
	Vertexs []Vertex `json:"vertexs"`
}

type Vertex struct {
	Xaxis float64 `json:"x"`
	Yaxis float64 `json:"y"`
	Zaxis float64 `json:"z"`
}

func main() {

	file, err := ioutil.ReadFile("models/test.stl")
	if err != nil {
		panic("File reading error")
	}
	
	parsed_stl := ParseAscii(file)
	fmt.Println(parsed_stl);

}

func ParseAscii(str []byte) *File {
	
	words := strings.Split(string(str), " ")

	file := File{}
	facet := Facet{}
	verts := []Vertex{}

	for i := 0; i < len(words); i++ {

		word := words[i]
		word = strings.TrimSpace(word)

		if word == "solid" && words[i+1] != "facet" {
			file.Name = words[i+1]
		} else if word == "facet" {

			fNormal := words[i+1]
			fX, _ := strconv.ParseFloat(words[i+2], 64)
			fY, _ := strconv.ParseFloat(words[i+3], 64)
			fZ, _ := strconv.ParseFloat(words[i+4], 64)

			facet = Facet{Normal: fNormal, Xaxis: fX, Yaxis: fY, Zaxis: fZ}

		} else if word == "vertex" {

			vX, _ := strconv.ParseFloat(words[i+1], 64)
			vY, _ := strconv.ParseFloat(words[i+2], 64)
			vZ, _ := strconv.ParseFloat(words[i+3], 64)

			vertexa := Vertex{Xaxis: vX, Yaxis: vY, Zaxis: vZ}
			verts = append(verts, vertexa)

		} else if word == "endfacet" {

			facet.Vertexs = verts
			file.Facets = append(file.Facets, facet);
			facet = Facet{}
			verts = []Vertex{}

		}

	}

	return &file

}