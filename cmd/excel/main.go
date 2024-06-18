package main

import (
	"encoding/json"
	"fmt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/tealeg/xlsx"
	"log"
	"os"
	data "stash/pkg"
)

func main() {
	excelFileName := "stash.xlsx"
	xlFile, openErr := xlsx.OpenFile(excelFileName)
	if openErr != nil {
		fmt.Printf("open failed: %s\n", openErr)
		return
	}

	rowNb := 0
	table := make([]data.Entry, 0)
	for _, sheet := range xlFile.Sheets {
		if sheet.Name != "Feuille 3" {
			continue
		}
		fmt.Println(sheet.Name)
		firstRow := sheet.Rows[0]
		for i, col := range firstRow.Cells {
			fmt.Println(i, col.String())

		}
		for _, row := range sheet.Rows {
			if len(row.Cells) == 0 {
				continue
			}
			date, _ := row.Cells[0].GetTime(false)
			sd := date.Format("2006-01-02")
			if sd == "0001-01-01" {
				continue
			}
			rowNb++

			total, err := row.Cells[15].Float()
			if err != nil {
				fmt.Println("Unexpected cell value", row.Cells[10].String())
			}
			entry := data.Entry{ObjectID: sd, Date: sd, Total: total, Epoch: date.Unix()}
			table = append(table, entry)
			fmt.Println(entry)
		}
		fmt.Println("Nb of row: ", rowNb)
	}
	// Marshal table in json
	bytes, _ := json.Marshal(table)
	err := os.WriteFile("stash.json", bytes, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Save in algolia
	client := search.NewClient("1QMZVCS1V5", "65b2efe47f21aed0a95b7ece43635e83")

	// Create a new index and add a record
	index := client.InitIndex("stash")
	resSave, err := index.SaveObjects(table)
	if err != nil {
		panic(err)
	}
	resSave.Wait()

}
