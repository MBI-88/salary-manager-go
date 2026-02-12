package salary_manager

import (
	"bufio"
	"os"
	"strings"
	"time"
)

type salaryManager struct {
	toPay map[string]float64
}

// salaryManager builder
func NewSalaryManager() managerBuilder {
	return &salaryManager{}
}

func (sm *salaryManager) operation(payrange []data, name string, start, end time.Time) {
	for i := 0; i < len(payrange); i++ {
		if start.Compare(sm.formater(payrange[i].EndRange)) > 0 {
			continue
		} else {
			if (start.Compare(sm.formater(payrange[i].StartRange)) > -1) &&
				(start.Compare(sm.formater(payrange[i].EndRange)) < 1) &&
				(end.Compare(sm.formater(payrange[i].EndRange)) < 1) {

				hours := end.Sub(start).Hours()
				sm.toPay[name] += hours * payrange[i].Payment

			} else if (start.Compare(sm.formater(payrange[i].StartRange)) > -1) &&
				(start.Compare(sm.formater(payrange[i].EndRange)) < 1) &&
				(end.Compare(sm.formater(payrange[i].EndRange)) > 0) {

				EndRange := sm.formater(payrange[i].EndRange)
				hours := EndRange.Sub(start).Hours()
				sm.toPay[name] += hours * payrange[i].Payment

				for x := i + 1; x < len(payrange); x++ {
					if (end.Compare(sm.formater(payrange[x].StartRange)) > -1) &&
						(end.Compare(sm.formater(payrange[x].EndRange)) < 1) {

						StartRange := sm.formater(payrange[x].StartRange)
						hours := end.Sub(StartRange).Hours()
						sm.toPay[name] += hours * payrange[x].Payment
						
					} else if end.Compare(sm.formater(payrange[x].EndRange)) > 0 {
						StartRange := sm.formater(payrange[x].StartRange)
						acmkEnd := sm.formater(payrange[x].EndRange)
						hours := acmkEnd.Sub(StartRange).Hours()
						sm.toPay[name] += hours * payrange[x].Payment
					}
				}

			}

		}
	}
}

// Convert from string to time (hours)
func (*salaryManager) formater(hours string) time.Time {
	value, err := time.Parse(TimeFormat, hours)
	if err != nil {
		panic(err)
	}
	return value
}

// Process data from an array
func (sm *salaryManager) ProcessData(employees []string) map[string]float64 {
	sm.toPay = make(map[string]float64)
	for _, emp := range employees {
		name, data := strings.Split(emp, "=")[0], strings.Split(strings.Split(emp, "=")[1], ",")
		sm.toPay[name] = 0
		for _, dt := range data {
			day := dt[0:2]
			timeRange := strings.Split(dt[2:], "-")
			start := sm.formater(timeRange[0])
			end := sm.formater(timeRange[1])
			switch day {
			case MO, TU, WE, TH, FR:
				payrange := rangePayments(Week)
				sm.operation(payrange, name, start, end)

			case SA, SU:
				payrange := rangePayments(Weekend)
				sm.operation(payrange, name, start, end)

			default:
				continue
			}

		}
	}
	return sm.toPay

}

// Load data from files
func (*salaryManager) LoadFile(path string) []string {
	var array []string
	if _, err := os.Stat(path); os.IsExist(err) {
		panic(err)
	}
	file, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	buffer := bufio.NewScanner(file)
	for buffer.Scan() {
		line := buffer.Text()
		array = append(array, line)
	}

	return array
}
