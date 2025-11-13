package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(datastring string) (err error)
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for i, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			log.Printf("Data parsing error (line %d): %w", i+1, err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("Information formation error (line %d): %w", i+1, err)
			continue
		}

		fmt.Println(info)
	}
}
