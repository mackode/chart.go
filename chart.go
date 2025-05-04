package main

import (
	"fmt"
	"github.com/wcharczuk/go-chart/v2"
	"os"
	"time"
)

func main() {
	viewings, err := readHistory()
	if err != nil {
		panic(err)
	}

	xVals := []time.Time{}
	yVals := []float64{}
	prevMonth := time.Time{}
	for i := len(viewings) - 1; i >= 0; i-- {
		v := viewings[i]
		month := time.Date(v.Date.Year(), v.Date.Month(), 1, 0, 0, 0, 0, v.Date.Location())
		if prevMonth.IsZero() || prevMonth != month {
			xVals = append(xVals, month)
			yVals = append(yVals, 0)
			prevMonth = month
		}
		yVals[len(yVals) - 1] += 1
	}

	chartData := chart.TimeSeries{
		XValues: xVals,
		YValues: yVals,
		Style: chart.Style{
			StrokeWidth: 2.0,
			StrokeColor: chart.ColorBlack,
			FillColor:   chart.ColorGreen,
		},
	}

	graph := chart.Chart{
		Series: []chart.Series{
			chartData,
		},
	}

	f, err := os.Create("netflix.png")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer f.Close()

	graph.Render(chart.PNG, f)
}