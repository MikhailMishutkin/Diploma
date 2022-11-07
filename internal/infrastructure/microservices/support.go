package dcollect

import (
	"encoding/json"
	"fmt"
	"graduatework/internal/model"
	"io"
	"net/http"
)

//...
func (m *MicroServiceStr) ReadSupportData() (outputData []model.SupportData, respStatusCode int) {

	outputData = make([]model.SupportData, 0)

	//TODO:catch the panic!!!
	response, err := http.Get("http://localhost:8383/support")
	if err != nil {
		fmt.Print("can't GET support-data: ", err)
	}

	if response.StatusCode != 200 {
		return outputData, response.StatusCode
	}

	responseInBytes, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Print("can't read support-data: ", err)
	}

	err = json.Unmarshal(responseInBytes, &outputData)
	if err != nil {
		return outputData, response.StatusCode
	}

	reserv := make([]model.SupportData, 0)
	for i := range outputData {
		reserv = outputData

		reserv = append(reserv[:i], reserv[i+1:]...)

	}
	outputData = reserv

	return outputData, response.StatusCode
}
