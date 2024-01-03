package main

import "fmt"

func main() {
	for _, res := range GenerateCheck() {
		if res.status == "pass" {
			fmt.Println(res.ServiceID)
		}
	}
}

type HealthCheck struct {
	ServiceID int
	status    string
}

func GenerateCheck() []HealthCheck {
	const (
		PassStatus = "pass"
		FailStatus = "fail"
	)
	var testVar []HealthCheck
	for i := 0; i <= 5; i++ {
		entry := []HealthCheck{{ServiceID: i, status: func() string {
			if i%2 == 0 {
				return PassStatus
			}
			return FailStatus
		}()}}
		testVar = append(testVar, entry...)
	}
	return testVar
}
