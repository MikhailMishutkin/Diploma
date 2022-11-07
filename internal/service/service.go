package service

import (
	"fmt"
	"graduatework/internal/model"
	"os"
	"sort"
	"strings"
	"sync"
	"unicode"
)

const (
	answerTime   = 60 / 18
	CountryFile  = "/mnt/c/go_work/src/GraduateWork/alpha2eng.txt"
	ProviderFile = "/mnt/c/go_work/src/GraduateWork/providers.txt"
)

type ServiceManage struct {
	microService MicroServicer
}

type MicroServicer interface {
	ReadSMS() (outputData []model.SMSData)
	ReadMMS() (outputData []model.MMSData, respStatusCode int)
	ReadEmail() (outputData []model.EmailData)
	ReadVoiceCall() (outputData []model.VoiceCallData)
	ReadBilling() (outputData model.BillingData)
	ReadSupportData() (outputData []model.SupportData, respStatusCode int)
	ReadIncidentData() (outputData []model.IncidentData, respStatusCode int)
}

//...
func NewServiceManage(m MicroServicer) *ServiceManage {
	return &ServiceManage{microService: m}
}

//...
func (s *ServiceManage) SortSMS() [][]model.SMSData {
	//var sorted [][]model.SMSData
	sorted := make([][]model.SMSData, 2)
	m := mapFromFile(CountryFile)
	r := s.microService.ReadSMS()

	for i, v := range r {
		a := m[v.Country]
		r[i].Country = a
	}
	sort.SliceStable(r, func(i, j int) bool {
		return r[i].Provider < r[j].Provider
	})
	sorted[0] = r
	sort.SliceStable(r, func(i, j int) bool {
		return r[i].Country < r[j].Country
	})
	sorted[1] = r

	return sorted
}

//...
func (s *ServiceManage) SortMMS() ([][]model.MMSData, int) {
	//var sorted [][]model.MMSData
	sorted := make([][]model.MMSData, 2)
	m := mapFromFile(CountryFile)
	r, statusCode := s.microService.ReadMMS()
	for i, v := range r {
		a := m[v.Country]
		r[i].Country = a
	}
	sort.SliceStable(r, func(i, j int) bool {
		return r[i].Country < r[j].Country
	})

	sorted[0] = r

	sort.SliceStable(r, func(i, j int) bool {
		return r[i].Provider < r[j].Provider
	})
	sorted[1] = r

	return sorted, statusCode
}

//...
func (s *ServiceManage) SortEmailBySpeed(f func() []model.EmailData) map[string][][]model.EmailData {
	m := make(map[string][][]model.EmailData, 10)
	out := make([][]model.EmailData, 2)
	var sliceOfCountries []string
	var incr, decr []model.EmailData
	sliceOfEmailDataStr := s.microService.ReadEmail()

	var a, b string
	for _, v := range sliceOfEmailDataStr {
		a = v.Country
		if b != a {
			sliceOfCountries = append(sliceOfCountries, a)
		}
		b = a
	}

	for _, v := range sliceOfCountries {

		var sliceCountryEmailDataStruct []model.EmailData = nil
		for _, v1 := range sliceOfEmailDataStr {
			if v1.Country == v {
				sliceCountryEmailDataStruct = append(sliceCountryEmailDataStruct, v1)
			}

		}

		len := len(sliceCountryEmailDataStruct)
		for i := 0; i < len-1; i++ {
			for j := 0; j < len-i-1; j++ {
				if sliceCountryEmailDataStruct[j].DeliveryTime > sliceCountryEmailDataStruct[j+1].DeliveryTime {
					sliceCountryEmailDataStruct[j], sliceCountryEmailDataStruct[j+1] = sliceCountryEmailDataStruct[j+1], sliceCountryEmailDataStruct[j]
				}
				incr = sliceCountryEmailDataStruct[0:3]

				decr = sliceCountryEmailDataStruct[(len - 3):len]

			}

		}
		out[0] = incr
		out[1] = decr

		m[v] = append(m[v], out[0])
		m[v] = append(m[v], out[1])
	}

	return m
}

//...
func (s *ServiceManage) SortWorkLoad() ([]int, int) {
	var workLoad, waitingTime int
	var supportData []int = []int{0, 0}

	read, sCode := s.microService.ReadSupportData()

	if sCode == 500 {
		return nil, sCode
	}

	var tickets int
	for _, v := range read {
		a := v.ActiveTickets
		tickets = tickets + a
	}

	if tickets > 16 {
		workLoad = 3
	} else if tickets > 9 && tickets <= 16 {
		workLoad = 2
	} else {
		workLoad = 1
	}

	waitingTime = tickets * answerTime

	supportData[0] = workLoad
	supportData[1] = waitingTime

	return supportData, sCode
}

//...
func (s *ServiceManage) SortIncident() (sortData []model.IncidentData, respStatusCode int) {
	read, respStatusCode := s.microService.ReadIncidentData()
	if respStatusCode != 200 {
		return read, respStatusCode
	}
	for i := 0; i < len(read)-1; i++ {
		a := "active"
		for j := len(read) - i - 1; j > 0; j-- {
			if read[j].Status != a {
				continue
			} else {
				read[j].Status, read[j-1].Status = read[j-1].Status, read[j].Status
			}
		}

	}
	sortData = read

	return sortData, respStatusCode
}

//...
func (s *ServiceManage) GetResultData(wg *sync.WaitGroup) (r model.ResultSetT) {

	sortSMSChan := make(chan [][]model.SMSData)
	sortMMSChan := make(chan [][]model.MMSData)
	voiceCallChan := make(chan []model.VoiceCallData)
	emailChan := make(chan map[string][][]model.EmailData)
	billingChan := make(chan model.BillingData)
	supportChan := make(chan []int)
	incidentChan := make(chan []model.IncidentData)

	wg.Add(7)
	go func() {
		defer wg.Done()
		a := s.SortSMS()
		sortSMSChan <- a
	}()

	go func() {
		defer wg.Done()
		a, _ := s.SortMMS()
		sortMMSChan <- a
	}()

	go func() {
		defer wg.Done()
		b := s.microService.ReadVoiceCall()
		voiceCallChan <- b
	}()

	go func() {
		defer wg.Done()
		b := s.SortEmailBySpeed(s.microService.ReadEmail)
		emailChan <- b
	}()

	go func() {
		defer wg.Done()
		b := s.microService.ReadBilling()
		billingChan <- b
	}()

	go func() {
		defer wg.Done()
		b, _ := s.SortWorkLoad()
		supportChan <- b
	}()

	go func() {
		defer wg.Done()
		b, _ := s.SortIncident()
		incidentChan <- b
	}()

	r.SMS = <-sortSMSChan
	r.MMS = <-sortMMSChan
	r.VoiceCall = <-voiceCallChan
	r.Email = <-emailChan
	r.Billing = <-billingChan
	r.Support = <-supportChan
	r.Incidents = <-incidentChan

	close(sortSMSChan)
	close(sortMMSChan)
	close(voiceCallChan)
	close(emailChan)
	close(billingChan)
	close(supportChan)
	close(incidentChan)

	return r
}

//...
func mapFromFile(f string) (m map[string]string) {
	content, err := os.ReadFile(f)
	if err != nil {
		fmt.Print(err)
	}

	s := string(content)
	sl := strings.Fields(s)
	m = make(map[string]string)
	var key, val string
	for _, v := range sl {
		if len(v) == 2 && unicode.IsUpper(rune(v[1])) {
			m[key] = val
			val = ""
			key = v
		} else {
			val = val + " " + v
		}

	}
	m[key] = val
	return m
}
