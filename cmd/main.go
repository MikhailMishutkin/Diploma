package main

import "graduatework/internal/app"

func main() {
	//sms := dcollect.Output()

	// тест
	// r := dcollect.GetResultData(&sync.WaitGroup{})
	// fmt.Println(r)

	//сервер
	app.RunServer()

	// var wg sync.WaitGroup
	// wg.Add(7)
	// go func() {
	// 	defer wg.Done()

	// }()

	// go func() {
	// 	defer wg.Done()
	// 	mms, code := dcollect.ReadMMS()
	// 	fmt.Println(mms, code)
	// }()

	// go func() {
	// 	defer wg.Done()
	// 	voice := dcollect.ReadVoiceCall()
	// 	fmt.Println(voice)
	// }()
	// go func() {
	// 	defer wg.Done()
	// email := dcollect.SortBySpeed(dcollect.ReadEmail)
	// fmt.Println(email)
	// }()
	// go func() {
	// 	defer wg.Done()
	// 	billing := dcollect.ReadBilling()
	// 	fmt.Println(billing)
	// }()

	// go func() {
	// 	defer wg.Done()
	// support, code := dcollect.SortWorkLoad()
	// fmt.Println(support, code)
	// }()

	// go func() {
	// 	defer wg.Done()
	// incident, code := dcollect.SortIncident()
	// fmt.Println(incident, code)
	// }()

	// wg.Wait()
	// time.Sleep(time.Second)

}
