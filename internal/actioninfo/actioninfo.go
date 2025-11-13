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

	for i, ch := range dataset {
		err := dp.Parse(ch)
		if err != nil {
			log.Printf("Information formation error: %v", err)
		}
		i++
		continue
	}
	info, err := dp.ActionInfo()
	if err != nil {
		log.Printf("Information formation error%v", err)
	}
	fmt.Println(info)
}
