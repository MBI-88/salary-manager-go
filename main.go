package main

import (
	"flag"
	"fmt"
	sm "salary_manager_go/internal"
)



func main() {
	flag.Parse()
	args := flag.Arg(0)
	manager := sm.NewSalaryManager()
	result := manager.ProcessData(manager.LoadFile(args))

	for key,val := range result {
		fmt.Printf("Name %s total to earn %0.2f$\n", key, val)
	}
	
}