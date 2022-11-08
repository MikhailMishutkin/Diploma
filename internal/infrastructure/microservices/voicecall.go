package dcollect

import (
	"encoding/csv"
	"fmt"
	"graduatework/internal/model"
	"io"
	"os"
)

//...
func (m *MicroServiceStr) ReadVoiceCall() (outputData []model.VoiceCallData) {
	a := model.VoiceCallData{}

	content, err := os.Open("./simulator/voice.data")
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
			fmt.Println("voice data corrupted: ", err)
			continue
		}
		if CheckExist(record[0], mapFromFile(CountryFile)) && (record[3] == Provider4 || record[3] == Provider5 || record[3] == Provider6) {
			a.Country = record[0]
			a.Bandwidth = record[1]
			a.ResponseTime = record[2]
			a.Provider = record[3]
			fValue, _ := Conversion(record[4])
			a.ConnectionStability = fValue
			_, iValue := Conversion(record[5])
			a.TTFB = iValue
			_, iValue = Conversion(record[6])
			a.VoicePurity = iValue
			_, iValue = Conversion(record[7])
			a.MedianOfCallsTime = iValue
			outputData = append(outputData, a)
		}
	}

	return outputData
}
