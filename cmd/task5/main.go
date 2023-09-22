package main

import (
	"os"

	"github.com/ssouthcity/astar-samf/astar"
	"github.com/ssouthcity/astar-samf/maps"
	"golang.org/x/exp/slices"
)

func main() {
	m := maps.New()

	err := maps.LoadFromFile(m, "Samfundet_map_2.csv")
	if err != nil {
		panic(err)
	}

	m.Print(os.Stdout)

	klubben := maps.Point{36, 6}
	selskapsiden := maps.Point{5, 8}
	rytterhallen := maps.Point{18, 14}

	friends := astar.New(klubben, selskapsiden, m)

	path, err := friends.Solve()
	if err != nil {
		panic(err)
	}

	slices.Reverse(path)

	me := astar.NewMovingGoal(rytterhallen, path, 4, m)

	err = me.Visualize(os.Stdout)
	if err != nil {
		panic(err)
	}
}
