package dcollect

import (
	"encoding/csv"
	"fmt"
	"graduatework/internal/model"
	"io"
	"os"
)

//...
func (m *MicroServiceStr) ReadEmail() (outputData []model.EmailData) {
	a := model.EmailData{}

	content, err := os.Open("./simulator/email.data")
	if err != nil {
		fmt.Print(err)
	}

	reader := csv.NewReader(content)
	reader.Comma = ';'

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("email data corrupted: ", err)
			continue
		}
		if CheckExist(record[0], mapFromFile(CountryFile)) && CheckExist(record[1], mapProvFromFile(ProviderFile)) {
			a.Country = record[0]
			a.Provider = record[1]
			_, iValue := Conversion(record[2])
			a.DeliveryTime = iValue
			outputData = append(outputData, a)
		}
	}

	return outputData
}
