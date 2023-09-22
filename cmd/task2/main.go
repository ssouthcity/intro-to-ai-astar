package main

import (
	"os"

	"github.com/ssouthcity/astar-samf/astar"
	"github.com/ssouthcity/astar-samf/maps"
)

func main() {
	m := maps.New()

	err := maps.LoadFromFile(m, "Samfundet_map_1.csv")
	if err != nil {
		panic(err)
	}

	strossa := maps.Point{32, 40}
	selskapsiden := maps.Point{5, 8}

	ast := astar.New(strossa, selskapsiden, m)

	err = ast.Visualize(os.Stdout)
	if err != nil {
		panic(err)
	}
}
