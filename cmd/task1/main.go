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

	rundhallen := maps.Point{18, 27}
	strossa := maps.Point{32, 40}

	ast := astar.New(rundhallen, strossa, m)

	err = ast.Visualize(os.Stdout)
	if err != nil {
		panic(err)
	}
}
