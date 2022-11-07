package handler

import (
	"encoding/json"
	"graduatework/internal/model"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

//...
type Handler struct {
	sm ServiceManager
}

//...
type ServiceManager interface {
	SortSMS() [][]model.SMSData
	SortMMS() ([][]model.MMSData, int)
	SortEmailBySpeed(f func() []model.EmailData) map[string][][]model.EmailData
	SortWorkLoad() ([]int, int)
	SortIncident() (sortData []model.IncidentData, respStatusCode int)
	GetResultData(wg *sync.WaitGroup) (r model.ResultSetT)
}

func NewHandler(s ServiceManager) *Handler {
	return &Handler{sm: s}
}

//...
func (h *Handler) RegisterR(router *mux.Router) {
	router.HandleFunc("/", h.HandleConnection)
}

//...
func (h *Handler) HandleConnection(w http.ResponseWriter, r *http.Request) {

	result := &model.ResultT{}
	var a, b, c int
	_, a = h.sm.SortMMS()
	_, b = h.sm.SortIncident()
	_, c = h.sm.SortWorkLoad()
	if a != 200 || b != 200 || c != 200 {
		result.Status = false
		result.Error = "Error on collect data"
	} else {
		result.Status = true
		result.Data = h.sm.GetResultData(&sync.WaitGroup{})
	}

	j, err := json.MarshalIndent(result.Data, "", "   ")
	if err != nil {
		result.Error = "Error on marshal data"
	}
	w.Write([]byte(j))
}
