package main

import (
	"time"

	"github.com/lietoast/grpc-tu/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var employees = []pb.Employee{
	{
		Id:        1,
		Number:    1994,
		FirstName: "Chandler",
		LastName:  "Bing",
		MonthSalary: &pb.MonthSalary{
			Basic: 5000,
			Bonus: 1250,
		},
		Status: pb.EmployeeStatus_NORMAL,
		LastModfied: &timestamppb.Timestamp{
			Seconds: time.Now().Unix(),
		},
	},
	{
		Id:        2,
		Number:    1999,
		FirstName: "Rachel",
		LastName:  "Green",
		MonthSalary: &pb.MonthSalary{
			Basic: 5600,
			Bonus: 1800,
		},
		Status: pb.EmployeeStatus_RETIRED,
		LastModfied: &timestamppb.Timestamp{
			Seconds: time.Now().Unix(),
		},
	},
	{
		Id:        3,
		Number:    2004,
		FirstName: "Ross",
		LastName:  "Geller",
		MonthSalary: &pb.MonthSalary{
			Basic: 8000,
			Bonus: 1250.9,
		},
		Status: pb.EmployeeStatus_ON_VACATION,
		LastModfied: &timestamppb.Timestamp{
			Seconds: time.Now().Unix(),
		},
	},
}
