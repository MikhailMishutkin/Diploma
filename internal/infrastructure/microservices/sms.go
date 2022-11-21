package dcollect

import (
	"encoding/csv"
	"fmt"
	"graduatework/internal/model"
	"io"
	"os"
)

//...
func (m *MicroServiceStr) ReadSMS() (outputData []model.SMSData) {
	a := model.SMSData{}

	content, err := os.Open("./simulator/sms.data")
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
			fmt.Println("sms data corrupted: ", err)
			continue
		}
		if CheckExist(record[0], mapFromFile(CountryFile)) && (record[3] == Provider1 || record[3] == Provider2 || record[3] == Provider3) {
			a.Country = record[0]
			a.Bandwidth = record[1]
			a.ResponseTime = record[2]
			a.Provider = record[3]
			outputData = append(outputData, a)
		}
	}

	return outputData
}
