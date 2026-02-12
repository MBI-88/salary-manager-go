package acm

import (
	"bufio"
	"os"
	"strings"
	"time"
)

type acmCompany struct {
	toPay map[string]float64
}

// AcmCompany builder
func NewAcmCompany() acmBuilder {
	return &acmCompany{}
}

func (ac acmCompany) operation(payrange []acmRange, name string, start, end time.Time) {
	for i := 0; i < len(payrange); i++ {
		if start.Compare(ac.formater(payrange[i].EndRange)) > 0 {
			continue
		} else {
			if (start.Compare(ac.formater(payrange[i].StartRange)) > -1) &&
				(start.Compare(ac.formater(payrange[i].EndRange)) < 1) && (end.Compare(ac.formater(payrange[i].EndRange)) < 1) {
				hours := end.Sub(start).Hours()
				ac.toPay[name] += hours * payrange[i].Payment
			} else if (start.Compare(ac.formater(payrange[i].StartRange)) > -1) &&
				(start.Compare(ac.formater(payrange[i].EndRange)) < 1) &&
				(end.Compare(ac.formater(payrange[i].EndRange)) > 0) {
				acmEnd := ac.formater(payrange[i].EndRange)
				hours := acmEnd.Sub(start).Hours()
				ac.toPay[name] += hours * payrange[i].Payment
				for x := 1; x <= len(payrange[i + 1:]); x++ {
					if (end.Compare(ac.formater(payrange[x].StartRange)) > -1) &&
						(end.Compare(ac.formater(payrange[x].EndRange)) < 1) {
						acmkStart := ac.formater(payrange[x].StartRange)
						hours := end.Sub(acmkStart).Hours()
						ac.toPay[name] += hours * payrange[x].Payment
					}else if end.Compare(ac.formater(payrange[x].EndRange)) > 0 {
						acmkStart := ac.formater(payrange[x].StartRange)
						acmkEnd := ac.formater(payrange[x].EndRange)
						hours := acmkEnd.Sub(acmkStart).Hours()
						ac.toPay[name] += hours * payrange[x].Payment
					}
				}
			
			}

		}
	}
}

// Convert from string to time (hours)
func (ac acmCompany) formater(hours string) time.Time {
	value, err := time.Parse("15:05", hours)
	if err != nil {
		panic(err)
	}
	return value
}

// Process data from an array
func (ac acmCompany) ProcessData(employees []string) map[string]float64 {
	ac.toPay = make(map[string]float64)
	for _, emp := range employees {
		name, data := strings.Split(emp, "=")[0], strings.Split(strings.Split(emp, "=")[1], ",")
		ac.toPay[name] = 0
		for _, dt := range data {
			day := dt[0:2]
			timeRange := strings.Split(dt[2:], "-")
			start := ac.formater(timeRange[0])
			end := ac.formater(timeRange[1])
			switch day {
			case "MO", "TU", "WE", "TH", "FR":
				payrange := rangePayments("Week")
				ac.operation(payrange, name, start, end)

			case "SA", "SU":
				payrange := rangePayments("Weekend")
				ac.operation(payrange, name, start, end)

			default:
				continue
			}

		}
	}
	return ac.toPay

}

// Load data from files
func (ac acmCompany) LoadFile(path string) []string {
	var array []string
	if _,err := os.Stat(path); os.IsExist(err) {
		panic(err)
	}
	file , err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
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


