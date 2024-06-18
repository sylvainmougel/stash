package main

import (
	"encoding/json"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"log"
	"net/http"
	"os"
	data "stash/pkg"
)

func httpserver(w http.ResponseWriter, _ *http.Request) {
	// read data
	d, err := os.ReadFile("stash.json")
	if err != nil {
		log.Fatal(err)
	}
	var records []data.Entry
	json.Unmarshal(d, &records)

	dates := make([]string, 0)
	items := make([]opts.LineData, 0)

	for _, r := range records {
		dates = append(dates, r.Date)
		items = append(items, opts.LineData{Value: r.Total})

	}

	// create a new line instance
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Stash evolution",
			Subtitle: "foo",
		}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true, TriggerOn: "mousemove", AxisPointer: &opts.AxisPointer{Type: "shadow", Show: true}}),
	)

	// Put data into instance
	line.SetXAxis(dates).
		AddSeries("Total", items).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	line.Render(w)
}

func main() {
	http.HandleFunc("/", httpserver)
	http.ListenAndServe(":8081", nil)
}
