package main

//#cgo CFLAGS: -I${SRCDIR}/cdevice
//#cgo LDFLAGS: -L${SRCDIR}/cdevice/ -lDiosProc

/*
#include <stdio.h>

const char* add_driver(char* driverName);
int get_all_device(int rpcPort, int httpPort);
int open_device();
int close_device();
int device_get();
int device_update();
int device_add();
int device_del();
int device_action();
int device_message();

*/
import "C"

import (
	"io/ioutil"
	"net"
	"os"
	"fmt"
	"encoding/json"
	//"unsafe"

	"github.com/osguydch/diosproc/gateway"
	"github.com/osguydch/diosproc/insecure"
	usersv1 "github.com/osguydch/diosproc/proto/device/v3"
	"github.com/osguydch/diosproc/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

type DiosProcConf struct {
	Application 	string `json:"Application"`
	SupportSSL 		string `json:"SupportSSL"`
	CertFilePath 	string `json:"CertFilePath"`
	PortMap 		[] struct {
		Type 		string `json:"Type"`
		RPCPort 	int `json:"RPCPort"`
		HTTPPort 	int `json:"HTTPPort"`
	}
}

func main() {
	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	if len(os.Args) < 2{
		fmt.Println("Usage: please provide driver config fileanme")
		return
	}

	var driverType string
	file, err := os.Open("./srvconf.json")
	if err != nil {
		panic(err)
	}

	for i,v := range os.Args {
		if i > 0 {
			cstr := C.CString(v)
			driverType = C.GoString(C.add_driver(cstr))	
			//C.free(unsafe.Pointer(cstr))		
		}

	}

	var config DiosProcConf
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		panic(err)
	}

	rpcPort := 9000
	httpPort := 9001
	for i,v := range config.PortMap {
		if v.Type == driverType {
			//找到对应的类型
			rpcPort = config.PortMap[i].RPCPort
			httpPort = config.PortMap[i].HTTPPort
		}
	}

	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)

	addr := ""
	var lis net.Listener
	for i:=0; i<10 ;i++{

		addr = fmt.Sprintf("0.0.0.0:%d", rpcPort)
		lis, err = net.Listen("tcp", addr)
		if err != nil {
			log.Fatalln("Failed to listen:", err)
			rpcPort += 2
			httpPort += 2
		} else {
			break
		}
	}


	C.get_all_device(C.int(rpcPort), C.int(httpPort))


	s := grpc.NewServer(
		// TODO: Replace with your own certificate!
		grpc.Creds(credentials.NewServerTLSFromCert(&insecure.Cert)),
	)
	usersv1.RegisterDeviceServer(s, server.New())

	// Serve gRPC Server
	log.Info("Serving gRPC on https://", addr)
	go func() {
		log.Fatal(s.Serve(lis))
	}()

	err = gateway.Run("dns:///" + addr, fmt.Sprintf("%d", httpPort))
	log.Fatalln(err)
}