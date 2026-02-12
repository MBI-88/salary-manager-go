package salary_manager

type data struct {
	StartRange string 
	EndRange   string 
	Payment    float64    
}

var payments = make(map[string][]data)

func rangePayments(key string) []data {
	payments["Week"] = []data{
		{
			StartRange: "00:01",
			EndRange:   "09:00",
			Payment:    25.0,
		},
		{
			StartRange: "09:01",
			EndRange:   "18:00",
			Payment:    15.0,
		},
		{
			StartRange: "18:01",
			EndRange:   "23:00",
			Payment:    20.0,
		},
	}
	payments["Weekend"] = []data{
		{
			StartRange: "00:01",
			EndRange:   "09:00",
			Payment:    30.0,
		},
		{
			StartRange: "09:01",
			EndRange:   "18:00",
			Payment:    20.0,
		},
		{
			StartRange: "18:01",
			EndRange:   "23:00",
			Payment:    25.0,
		},
	}

	return payments[key]
}
