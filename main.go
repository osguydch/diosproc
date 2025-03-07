package main

//#cgo CFLAGS: -I${SRCDIR}/cdevice
//#cgo LDFLAGS: -L${SRCDIR}/cdevice/ -lDiosProc

/*
#include <stdio.h>



const char* add_driver(char* driverName);
const char* get_all_device(int rpcPort, int httpPort);




*/
import "C"

import (
	"io/ioutil"
	"net"
	"os"
	"fmt"
	"strings"
	"encoding/json"
	"strconv"
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
		RPCPort 	string `json:"RPCPort"`
		HTTPPort 	string `json:"HTTPPort"`
		ProcMap [] struct {
			Name string `json:"Name"`
			Index string `json:Index`
		}
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

	//driverType : Detector_CareRay_1800RF_1234
	drivers := strings.Split(driverType, "_")

	var config DiosProcConf
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		panic(err)
	}
	file.Close()

	var vendor string
	vendor = drivers[1] + "_" + drivers[2]+ "_" + drivers[3];

	fmt.Println("Load driver :" + vendor)

	rpcPort := 9000
	httpPort := 9001
	var wi int
	for _,v := range config.PortMap {
		if v.Type == drivers[0] {
			//找到对应的类型
			fmt.Println("Got Type PortMap config %s",v.Type )
			rpcPort,_ = strconv.Atoi(v.RPCPort)
			httpPort,_ = strconv.Atoi(v.HTTPPort)
			for _,w := range v.ProcMap {
				if strings.HasPrefix(vendor, w.Name) {
					//按厂商名匹配
					fmt.Println("Got Vendor PortMap Index config %s" ,w.Index)
					wi,_ = strconv.Atoi(w.Index)
					rpcPort += wi*2
					httpPort += wi*2
					break
				}				
			}
			break
		}
	}

	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)

	addr := ""
	var lis net.Listener
	for i:=0; i<10 ;i++{

		addr = fmt.Sprintf("localhost:%d", rpcPort)
		fmt.Println("try listen " + addr)

		lis, err = net.Listen("tcp", addr)
		if err != nil {
			fmt.Println("Failed to listen: %s err %v",addr, err)
			rpcPort += 2
			httpPort += 2
			fmt.Println("try another port %d" , rpcPort)
		} else {
			break
		}
	}

	var devices string
	devices = C.GoString( C.get_all_device(C.int(rpcPort), C.int(httpPort)))
	device := strings.Split(devices, ";")


	fmt.Println("new Server with " + addr)
	s := grpc.NewServer(
		// TODO: Replace with your own certificate!
		grpc.Creds(credentials.NewServerTLSFromCert(&insecure.Cert)),
	)
	usersv1.RegisterDeviceServer(s, server.New(device))

	// Serve gRPC Server
	log.Info("Serving gRPC on https://", addr)
	go func() {
		log.Fatal(s.Serve(lis))
	}()

	err = gateway.Run("dns:///" + addr, fmt.Sprintf("%d", httpPort), config.SupportSSL)
	log.Fatalln(err)
}