package importer

import (
	"bytes"
	"fmt"
	"os"
	"regexp"

	"github.com/bernardolm/go-csv-tag"

	"github.com/bernardolm/ponto-esperto-importer/lib"
)

func Do() []Workday {
	workdays := []Workday{}

	err := csvtag.Load(csvtag.Config{
		Path:      lib.Config.FilePath,
		Dest:      &workdays,
		Separator: ';',
		Modifier:  modifyMyFile,
	})
	if err != nil {
		fmt.Println(err)
	}

	if lib.Config.Debug {
		for i, v := range workdays {
			fmt.Printf("Workday %d %+v\n", i, v)
		}
	}

	return workdays
}

func modifyMyFile(file *os.File) *string {
	if file == nil {
		return nil
	}

	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(file)
	if err != nil {
		return nil
	}
	contents := buf.String()

	reg, err := regexp.Compile(";\n")
	if err != nil {
		return nil
	}
	contents = reg.ReplaceAllString(contents, "\n")

	reg, err = regexp.Compile(`Saldo\stotal.*`)
	if err != nil {
		return nil
	}
	contents = reg.ReplaceAllString(contents, "")

	reg, err = regexp.Compile("\t")
	if err != nil {
		return nil
	}
	contents = reg.ReplaceAllString(contents, "")

	return &contents
}
