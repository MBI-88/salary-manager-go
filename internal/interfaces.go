package salary_manager





type managerBuilder interface {
	ProcessData(employees []string) map[string]float64
	LoadFile(path string) []string
}