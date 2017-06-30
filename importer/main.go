package importer

import (
	"fmt"

	"github.com/artonge/go-csv-tag"
)

func Do() {
	tab := []test{}                   // Create the slice where to put the file content
	err := csvtag.Load(csvtag.Config{ // Load your csv with the appropriate configuration
		Path:      "file.csv", // Path of the csv file
		Dest:      &tab,       // A pointer to the create slice
		Separator: ';',        // Optional - if your csv use something else than ',' to separate values
		// Header:    []string{"header1", "header2", "header3"}, // Optional - if your csv does not contains a header
	})
	if err != nil {
		fmt.Println(err)
	}

	for i, v := range tab {
		fmt.Printf("tab %d %#v\n", i, v)
	}
}
