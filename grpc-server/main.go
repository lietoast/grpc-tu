package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"
	"log"
	"net"
	"time"

	"github.com/lietoast/grpc-tu/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
)

const port = ":5001"

func main() {
	// 加载服务器私钥和证书
	cert, err := tls.LoadX509KeyPair("../cert/server/server.crt", "../cert/server/server.key")
	failOnError(err, "Failed to load cert files")

	// 生成证书池，将根证书加入证书池
	certPool := x509.NewCertPool()
	rootBuf, err := ioutil.ReadFile("../cert/ca.crt") // 根证书：ca.crt
	failOnError(err, "Failed to load ca.crt")
	if !certPool.AppendCertsFromPEM(rootBuf) {
		log.Panicf("Failed to append ca")
	}

	// 初始化TLSConfig
	tlsConf := &tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert, // 双向认证必须
		Certificates: []tls.Certificate{cert},
		ClientCAs:    certPool,
	}
	keepAliveArgs := keepalive.ServerParameters{
		Time:             10 * time.Second,
		Timeout:          20 * time.Second,
		MaxConnectionAge: 30 * time.Second,
	}

	// 监听端口
	listen, err := net.Listen("tcp", port)
	failOnError(err, "Failed to listen to port "+port)

	server := grpc.NewServer(
		grpc.KeepaliveParams(keepAliveArgs),
		grpc.MaxSendMsgSize(1024*1024*4),
		grpc.Creds(credentials.NewTLS(tlsConf)),
	)

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

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %v", msg, err)
	}
}
