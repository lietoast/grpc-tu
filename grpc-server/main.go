package main

import (
	"context"
	"errors"

	"github.com/lietoast/grpc-tu/pb"
)

const port = ":5001"

func main() {

}

type employeeService struct {
	pb.UnimplementedEmployeeServiceServer
}

func (e employeeService) GetByNo(c context.Context, r *pb.GetByNoRequest) (*pb.EmployeeResponse, error) {
	for _, e := range employees {
		if e.Number == r.Number {
			return &pb.EmployeeResponse{
				Employee: &e,
			}, nil
		}
	}

	return nil, errors.New("employee not found")
}
func (e employeeService) GetAll(*pb.GetAllRequest, pb.EmployeeService_GetAllServer) error {
	return nil
}
func (e employeeService) AddPhoto(pb.EmployeeService_AddPhotoServer) error {
	return nil
}
func (e employeeService) Save(context.Context, *pb.EmployeeRequest) (*pb.EmployeeResponse, error) {
	return nil, nil
}
func (e employeeService) SaveAll(pb.EmployeeService_SaveAllServer) error {
	return nil
}
