package dcollect

import (
	"encoding/json"
	"fmt"
	"graduatework/internal/model"
	"io"
	"net/http"
)

//...
func (m *MicroServiceStr) ReadMMS() (outputData []model.MMSData, respStatusCode int) {
	outputData = make([]model.MMSData, 0)

	//TODO:catch the panic!!!
	response, err := http.Get("http://localhost:8383/mms")
	if err != nil {
		fmt.Print("can't GET MMS-data: ", err)
	}

	if response.StatusCode != 200 {
		return outputData, response.StatusCode
	}

	responseInBytes, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Print("can't read MMS-data: ", err)
	}

	err = json.Unmarshal(responseInBytes, &outputData)
	if err != nil {
		return outputData, response.StatusCode
	}

	reserv := make([]model.MMSData, 0)
	for i, v := range outputData {
		reserv = outputData
		if CheckExist(v.Country, mapFromFile(CountryFile)) && (v.Provider == Provider1 || v.Provider == Provider2 || v.Provider == Provider3) {
			continue
		} else {
			reserv = append(reserv[:i], reserv[i+1:]...)
		}
	}
	outputData = reserv

	return outputData, response.StatusCode

}
