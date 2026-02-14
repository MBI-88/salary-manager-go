package tests

import (
	"testing"
	sm "salary_manager_go/internal"
)

func TestProcessData(t *testing.T) {
	acmTest := sm.NewSalaryManager()
	test := make(map[string]float64)
	sample := []string{
		"JOSE=MO08:00-19:00",
		"JUAN=TH12:00-17:00,FR09:01-16:00",
		"ANA=SA09:00-13:00,MO08:10-16:00",
	}

	test["JOSE"] = 179.99 // 25 + 135 + 20
	test["JUAN"] = 179.99 // 75 + 105 = 180
	test["ANA"] = 209.92 // 60 + 105

	restul := acmTest.ProcessData(sample)

	for key := range restul {
		if int(restul[key]) != int(test[key]) {
			t.Errorf("Name %s resutl => %.2f and test => %.2f\n", key, restul[key], test[key])
		}
	}

}
