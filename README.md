# Salary Manager
Salary Manager
## Overview

This project is a Salary Manager written in Go. It processes employee data and calculates the total earnings for each employee based on their working hours.

## How to Run

1. Clone the repository
2. Navigate to the project directory
3. Run the command `go run main.go` followed by the path to your employee data file. For example: `go run main.go data/employee_data.txt`

## Employee Data File Format

The employee data file should be a text file where each line represents an employee's work schedule. The format of each line should be as follows:

`EMPLOYEE_NAME=DAYHH:MM-HH:MM,DAYHH:MM-HH:MM,...`

Where:
- `EMPLOYEE_NAME` is the name of the employee
- `DAY` is a two-letter representation of the day of the week (MO, TU, WE, TH, FR, SA, SU)
- `HH:MM-HH:MM` represents the start and end time of the employee's work hours for that day

## Determining the Data Type of a Variable in Go

In Go, you can determine the type of a variable using the `%T` format specifier in a `Printf` statement. For example:

