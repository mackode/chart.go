package main

import (
	"fmt"
	"github.com/wcharczuk/go-chart/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	viewings, err := readHistory()
	if err != nil {
		panic(err)
	}

	data := []float64{}
	for _, v := range viewings {
		md, err := omdbFetch(v.Title)
		if err != nil {
			if strings.Constains(err.Error(), "not found") {
				continue
			} else {
				panic(err)
			}
		}
		parts := strings.Split(md.Ratings, "/")
		value, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			continue
		}
		data = append(data, value)
	}

	drawHisto(data)
}

func drawHisto(data []float64) {
	binWidth := 0.5
	minValue := 0.0
	maxValue := 10.0

	numberOfBins := int((maxValue - minValue) / binWidth)
	bins := make([]int, numberOfBins)
	for _, value := range data {
		if value >= minValue && value < maxValue {
			i := int((value - minValue) / binWidth)
			bins[i]++
		}
	}

	bars := []chart.Value{}
	for i, count := range bins {
		binStart := minValue + float64(i) * binWidth
		binEnd := binStart + binWidth
		bars = append(bars, chart.Value{
			Value: float64(count),
			Label: fmt.Sprintf("%.1f - %.1f", binStart, binEnd),
		})
	}

	barChart := chart.BarChart{
		Height: 512,
		Width:  1024,
		Bars:   bars,
	}

	f, _ := os.Create("ratings.png")
	defer f.Close()
	barChart.Render(chart.PNG, f)
}
