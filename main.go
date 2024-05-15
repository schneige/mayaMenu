package main

import (
	"fmt"
	"log"
	"encoding/json"
	"os"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()

	url := "https://www.mayapapaya.de/suppenangebot/"

	index := 0
	item := ""

	menu := make(map[string]map[string]string)

	days := []string{"MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY"}

	for _, day := range days {
		menu[day] = make(map[string]string)
	}

	c.OnHTML(".elementor-widget-text-editor p", func(e *colly.HTMLElement) {
		sText := e.Text
		switch {
		case index >= 20:

		case index >= 16:
			if index % 4 == 0 {
				item = sText
			} else if index % 4 == 1 {
				menu["FRIDAY"]["Item 1"] = item + " " + sText
			} else if index % 4 == 2 {
				item = sText
			} else if index % 4 == 3 {
				menu["FRIDAY"]["Item 2"] = item + " " + sText
			}
		case index >= 12:
			if index % 4 == 0 {
				item = sText
			} else if index % 4 == 1 {
				menu["THURSDAY"]["Item 1"] = item + " " + sText
			} else if index % 4 == 2 {
				item = sText
			} else if index % 4 == 3 {
				menu["THURSDAY"]["Item 2"] = item + " " + sText
			}
		case index >= 8:
			if index % 4 == 0 {
				item = sText
			} else if index % 4 == 1 {
				menu["WEDNESDAY"]["Item 1"] = item + " " + sText
			} else if index % 4 == 2 {
				item = sText
			} else if index % 4 == 3 {
				menu["WEDNESDAY"]["Item 2"] = item + " " + sText
			}
		case index >= 4:
			if index % 4 == 0 {
				item = sText
			} else if index % 4 == 1 {
				menu["TUESDAY"]["Item 1"] = item + " " + sText
			} else if index % 4 == 2 {
				item = sText
			} else if index % 4 == 3 {
				menu["TUESDAY"]["Item 2"] = item + " " + sText
			}
		case index >= 0:
			if index % 4 == 0 {
				item = sText
			} else if index % 4 == 1 {
				menu["MONDAY"]["Item 1"] = item + " " + sText
			} else if index % 4 == 2 {
				item = sText
			} else if index % 4 == 3 {
				menu["MONDAY"]["Item 2"] = item + " " + sText
			}
		default:
			fmt.Println("Something went wrong")
		}

		index++
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}

	  file, err := os.Create("menu.json")
	  if err != nil {
		  fmt.Println("Error creating file:", err)
		  return
	  }
	  defer file.Close()
  
	  jsonData, err := json.MarshalIndent(menu, "", "  ")
		if err != nil {
			fmt.Println("Error encoding JSON:", err)
			return
		}
  
	  _, err = file.Write(jsonData)
	  if err != nil {
		  fmt.Println("Error writing JSON to file:", err)
		  return
	  }
  
	  fmt.Println("Menu data written to menu.json successfully.")
}
