package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/lietoast/grpc-tu/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	// 加载客户端私钥和证书
	cert, err := tls.LoadX509KeyPair("../cert/client/client.crt", "../cert/client/client.key")
	failOnError(err, "Failed to load cert file")

	// 将根证书加入证书池
	certPool := x509.NewCertPool()
	rootBuf, err := ioutil.ReadFile("../cert/ca.crt") // 根证书：ca.crt
	failOnError(err, "Failed to load ca.crt")
	if !certPool.AppendCertsFromPEM(rootBuf) {
		log.Panicf("Failed to append ca")
	}

	creds := credentials.NewTLS(&tls.Config{
		ServerName: "hellosvc.com", // ServerName需要与服务器证书的CN字段保持一致
		Certificates: []tls.Certificate{cert},
		RootCAs: certPool,
	})

	conn, err := grpc.Dial("127.0.0.1:5001", grpc.WithTransportCredentials(creds))
	failOnError(err, "Failed to connect to gprc server")
	defer conn.Close()

	client := pb.NewEmployeeServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	getByNoResp, err := client.GetByNo(ctx, &pb.GetByNoRequest{Number: 1999})
	failOnError(err, "Failed to get response from grpc server")

	employee := getByNoResp.Employee
	fmt.Printf("id: %d, no: %d, name: %s %s, salary: %f\n", 
		employee.GetId(), 
		employee.GetNumber(),
		employee.GetFirstName(), employee.GetLastName(),
		employee.GetMonthSalary().GetBasic() + employee.GetMonthSalary().GetBonus(),
	)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %v", msg, err)
	}
}