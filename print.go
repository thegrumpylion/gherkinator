package main

import (
	"fmt"

	messages "github.com/cucumber/common/messages/go/v19"
)

func printScenario(s *messages.Scenario) {
  fmt.Println("Scenario:", s.Name)
  for _, stp := range s.Steps {
    fmt.Printf("  %s%s\n", stp.Keyword, stp.Text)
    printDataTable(stp.DataTable)
  }
}

func printDataTable(dt *messages.DataTable) {
  if dt == nil {
    return
  }
  for _, r := range dt.Rows {
    for _, c := range r.Cells {
      fmt.Print(c.Value)
      fmt.Print(" - ")
    }
  }
}

