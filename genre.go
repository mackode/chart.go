package main

import (
	"github.com/wcharczuk/go-chart/v2"
	"os"
)

func main() {
	series := 0
	movies := 0
	viewings, err := readHistory()
	if err != nil {
		panic(err)
	}

	for _, v := range viewings {
		_, err := omdbFetch(v.Title)
		if err == nil {
			movies += 1
		} else {
			series += 1
		}
	}

	pie := chart.PieChart{
		Width:  512,
		Height: 512,
		Values: []chart.Value{
			{Value: float64(series), Label: "Series"},
			{Value: float64(movies), Label: "Movies"},
		},
	}

	f, err := os.Create("genre.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = pie.Render(chart.PNG, f)
	if err != nil {
		panic(err)
	}
}
