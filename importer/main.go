package importer

import (
	"bytes"
	"fmt"
	"os"
	"regexp"

	"github.com/bernardolm/go-csv-tag"
)

func Do(filePath string) {
	period := []Workday{}

	err := csvtag.Load(csvtag.Config{
		Path:      filePath,
		Dest:      &period,
		Separator: ';',
		Modifier:  modifyMyFile,
	})
	if err != nil {
		fmt.Println(err)
	}

	for i, v := range period {
		fmt.Printf("Workday %d %#v\n", i, v)
	}
}

func modifyMyFile(file *os.File) *string {
	if file == nil {
		return nil
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(file)
	contents := buf.String()

	reg, err := regexp.Compile(";\n")
	if err != nil {
		return nil
	}
	contents = reg.ReplaceAllString(contents, "\n")

	reg, err = regexp.Compile("Saldo.*")
	if err != nil {
		return nil
	}
	contents = reg.ReplaceAllString(contents, "")

	return &contents
}
