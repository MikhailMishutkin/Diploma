package dcollect

import (
	"fmt"
	"graduatework/internal/model"
	"graduatework/internal/service"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const (
	CountryFile  = "alpha2eng.txt"
	ProviderFile = "providers.txt"
	Provider1    = "Topolo"
	Provider2    = "Rond"
	Provider3    = "Kildy"
	Provider4    = "TransparentCalls"
	Provider5    = "E-Voice"
	Provider6    = "JustPhone"
)

type MicroServiceStr struct {
	//m *Result
}

type MicroServicer interface {
	MicroService() service.MicroServicer
}

func NewMicroService(m MicroServiceStr) *MicroServiceStr {
	return &MicroServiceStr{}
}

func (m *MicroServiceStr) MicroService() service.MicroServicer {
	//
	return m
}

// 	return s.personRepository
// }

//...
func CheckExist(forCheck string, m map[string]string) (checked bool) {
	for i := range m {
		if forCheck == i {
			return true
		} else {
			continue
		}
	}
	return false

}

//...
func mapProvFromFile(f string) (m map[string]string) {
	content, err := os.ReadFile(f)
	if err != nil {
		fmt.Print(err)
	}

	s := string(content)
	sl := strings.Fields(s)
	m = make(map[string]string)

	for i, v := range sl {
		m[v] = strconv.Itoa(i)
	}
	return m
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

// common func for services for convert strings to int or to float32
func Conversion(s string) (f float32, i int) {
	i, err := strconv.Atoi(s)
	if err != nil {
		conv, err := strconv.ParseFloat(s, 32)
		if err != nil {
			fmt.Print(err)
		}
		f = float32(conv)
		return f, 0
	}
	return 0, i

}

type ByCountryS []model.SMSData

func (a ByCountryS) Len() int           { return len(a) }
func (a ByCountryS) Less(i, j int) bool { return a[i].Country < a[j].Country }
func (a ByCountryS) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type ByProviderS []model.SMSData

func (a ByProviderS) Len() int           { return len(a) }
func (a ByProviderS) Less(i, j int) bool { return a[i].Provider < a[j].Provider }
func (a ByProviderS) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type ByCountryM []model.MMSData

func (a ByCountryM) Len() int           { return len(a) }
func (a ByCountryM) Less(i, j int) bool { return a[i].Country < a[j].Country }
func (a ByCountryM) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type ByProviderM []model.MMSData

func (a ByProviderM) Len() int           { return len(a) }
func (a ByProviderM) Less(i, j int) bool { return a[i].Provider < a[j].Provider }
func (a ByProviderM) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
