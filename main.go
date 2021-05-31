package main

import (
	"fmt"
	"strings"
)

type File struct {
	Name   string `json:"name"`
	Facets Facet  `json:"facets"`
}

type Facet struct {
	Normal string `json:"type"`
	Xaxis   float64 `json:"x"`
	Yaxis   float64 `json:"y"`
	Zaxis   float64 `json:"z"`
	Vertexs Vertexs `json:"vertexs"`
}

type Vertexs struct {
	Primary   Vertex `json:"primary"`
	Secondary Vertex `json:"secondary"`
	Tertiary  Vertex `json:"tertiary"`
}

type Vertex struct {
	Xaxis float64 `json:"x"`
	Yaxis float64 `json:"y"`
	Zaxis float64 `json:"z"`
}

func main() {

	// Temp
	// primary_vertex := Vertex{1,0,-1}
	// secondary_vertex := Vertex{0,0,-1}
	// tertiary_vertex := Vertex{1,1,1}
	// vertexs := Vertexs{primary_vertex, secondary_vertex, tertiary_vertex}
	// facet := Facet{-1, -1, 0, vertexs}
	// stl := File{"new_cube.stl", facet}
	// fmt.Println(stl)

	str := "solid new_cube facet normal 0.0 0.0 1.0 outer loop vertex 1.0 1.0 0.0 vertex -1.0 1.0 0.0 vertex 0.0 -1.0 0.0 endloop endfacet endsolid"

	words := strings.Split(string(str), " ")

	var name string
	var facets []interface{}
	var vertexs []interface{}
	var vertex []interface{}

	for i := 0; i < len(words); i++ {
		
		word := words[i];

		if word == "solid" && words[i+1] != "facet" {
			name = words[i + 1];
		}else if word == "facet" {
			facets = append(facets, words[i+1:i+5])
		}else if word == "vertex" {
			vertex = append(vertex, words[i+1:i+5])
		}else if word == "endfacet" {
			vertexs = append(vertexs, vertex)
		}

	}

	Stl := File{Name: name}

	fmt.Println(Stl);

}