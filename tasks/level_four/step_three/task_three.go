// "Отчётики"

package main

import "time"

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

type Report struct {
	User
	ReportID int
	Date     string
}

func CreateReport(user User, reportDate string) Report {
	reportID := int(time.Now().UnixNano())
	return Report{
		User:     user,
		ReportID: reportID,
		Date:     reportDate,
	}
}

func GenerateUserReports(users []User, reportDate string) []Report {
	reports := make([]Report, 0, len(users))
	for _, user := range users {
		report := CreateReport(user, reportDate)
		reports = append(reports, report)
	}
	return reports
}
