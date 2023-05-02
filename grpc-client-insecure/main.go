package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/lietoast/grpc-tu/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	// 不安全的连接实现
	conn, err := grpc.Dial("127.0.0.1:5001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	failOnError(err, "Failed to connect to gprc server")
	defer conn.Close()

	client := pb.NewEmployeeServiceClient(conn)

	getByNo(client)
	fmt.Println("#######################################################")
	getAll(client)
	fmt.Println("#######################################################")
	AddPhoto(client)
	fmt.Println("#######################################################")
	saveAll(client)
}

func saveAll(client pb.EmployeeServiceClient) {
	employees := []pb.Employee{
		{
			Id:        4,
			Number:    2009,
			FirstName: "John",
			LastName:  "Wick",
			MonthSalary: &pb.MonthSalary{
				Basic: 10000,
				Bonus: 1,
			},
			Status:      pb.EmployeeStatus_NORMAL,
			LastModfied: timestamppb.Now(),
		},
		{
			Id:        5,
			Number:    2013,
			FirstName: "Joey",
			LastName:  "Tribbiani",
			MonthSalary: &pb.MonthSalary{
				Basic: 3500,
				Bonus: 2500,
			},
			Status:      pb.EmployeeStatus_NORMAL,
			LastModfied: timestamppb.Now(),
		},
	}

	stream, err := client.SaveAll(context.Background())
	failOnError(err, "Failed to get stream from grpc service")

	sig := make(chan struct{})
	// 创建等待响应的goroutine
	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				sig <- struct{}{}
				break
			}
			failOnError(err, "Error occurred when receiving response")
			printEmployee(resp.Employee)
		}
	}()

	for _, e := range employees {
		err := stream.Send(&pb.EmployeeRequest{Employee: &e})
		failOnError(err, "Error occurred when sending request")
	}

	stream.CloseSend()
	<-sig
}

func AddPhoto(client pb.EmployeeServiceClient) {
	// 这里使用了陌芋的图片（https://www.pixiv.net/en/artworks/101848882）
	// 请关注pixiv陌芋捏
	imgFile, err := os.Open("委托中？.jpeg")
	failOnError(err, "Failed to open the img")
	defer imgFile.Close()

	md := metadata.New(map[string]string{
		"no": "2004",
	})
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, md)

	stream, err := client.AddPhoto(ctx)
	failOnError(err, "Failed to get stream from grpc server")

	for {
		chunk := make([]byte, 128*1024) // 一次传输128k
		chunkSize, err := imgFile.Read(chunk)
		if err == io.EOF {
			break
		}
		failOnError(err, "Error occured when transporting the img")
		if chunkSize < len(chunk) {
			chunk = chunk[:chunkSize]
		}

		stream.Send(&pb.AddPhotoRequest{Data: chunk})
	}

	resp, err := stream.CloseAndRecv() // 服务器返回响应前，阻塞在这里
	failOnError(err, "Server error")
	fmt.Println(resp.IsOk)
}

func getAll(client pb.EmployeeServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := client.GetAll(ctx, &pb.GetAllRequest{})
	failOnError(err, "Failed to get stream from grpc server")

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		failOnError(err, "Failed to receive response from stream")
		printEmployee(resp.Employee)
	}
}

func getByNo(client pb.EmployeeServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	getByNoResp, err := client.GetByNo(ctx, &pb.GetByNoRequest{Number: 1999})
	failOnError(err, "Failed to get response from grpc server")

	employee := getByNoResp.Employee
	printEmployee(employee)
}

func printEmployee(employee *pb.Employee) {
	fmt.Printf("id: %d, no: %d, name: %s %s, salary: %f\n",
		employee.GetId(),
		employee.GetNumber(),
		employee.GetFirstName(), employee.GetLastName(),
		employee.GetMonthSalary().GetBasic()+employee.GetMonthSalary().GetBonus(),
	)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %v", msg, err)
	}
}
