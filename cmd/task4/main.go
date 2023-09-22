package main

import (
	"os"

	"github.com/ssouthcity/astar-samf/astar"
	"github.com/ssouthcity/astar-samf/maps"
)

func main() {
	m := maps.New()

	err := maps.LoadFromFile(m, "Samfundet_map_Edgar_full.csv")
	if err != nil {
		panic(err)
	}

	lyche := maps.Point{32, 28}
	klubben := maps.Point{32, 6}

	ast := astar.New(lyche, klubben, m)

	err = ast.Visualize(os.Stdout)
	if err != nil {
		panic(err)
	}
}
