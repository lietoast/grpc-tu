package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/lietoast/grpc-tu/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const port = ":5001"

func main() {
	// 监听端口
	listen, err := net.Listen("tcp", port)
	failOnError(err, "Failed to listen to port "+port)

	server := grpc.NewServer()

	// 注册服务
	pb.RegisterEmployeeServiceServer(server, new(employeeService))
	// 开启 grpc 服务
	server.Serve(listen)
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

// 获取所有雇员的信息
// 通过服务端stream返回（服务端会一直返回数据，直到最后一条数据被发送到客户端）
func (e employeeService) GetAll(r *pb.GetAllRequest, stream pb.EmployeeService_GetAllServer) error {
	for _, e := range employees {
		stream.Send(&pb.EmployeeResponse{Employee: &e})
	}
	return nil
}

func (e employeeService) AddPhoto(stream pb.EmployeeService_AddPhotoServer) error {
	// 获取grpc请求的元数据
	md, ok := metadata.FromIncomingContext(stream.Context())
	if ok {
		fmt.Printf("Receiving avatar of employee No.%s\n", md["no"][0])
	}

	img := []byte{}
	for {
		data, err := stream.Recv()
		if err == io.EOF {
			fmt.Printf("Received the avatar of %d bytes\n", len(img))
			return stream.SendAndClose(&pb.AddPhotoResponse{IsOk: true}) // 通知客户端关闭流
		} else if err != nil {
			return err
		}
		fmt.Printf("Received %d bytes\n", len(data.Data))
		img = append(img, data.Data...)
	}
}

func (e employeeService) Save(context.Context, *pb.EmployeeRequest) (*pb.EmployeeResponse, error) {
	return nil, nil
}

func (e employeeService) SaveAll(stream pb.EmployeeService_SaveAllServer) error {
	for {
		empReq, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		employees = append(employees, *empReq.GetEmployee())
		stream.Send(&pb.EmployeeResponse{Employee: empReq.GetEmployee()})
	}

	for _, e := range employees {
		fmt.Printf("id: %d, no: %d, name: %s %s, salary: %f\n",
			e.GetId(),
			e.GetNumber(),
			e.GetFirstName(), e.GetLastName(),
			e.GetMonthSalary().GetBasic()+e.GetMonthSalary().GetBonus(),
		)
	}

	return nil
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %v", msg, err)
	}
}
