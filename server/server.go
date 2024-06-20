package server

/*
int Add(int a, int b){
    return a+b;
}
*/
import "C"


import (
    "context"
    "fmt"
    //"log"
    //"net"
    "sync"

    //"github.com/gofrs/uuid"
    // importing generated stubs
    gen "github.com/osguydch/diosproc/proto/device/v3"
    //"google.golang.org/grpc"
)
// DeviceServerImpl will implement the service defined in protocol buffer definitions
type DeviceServerImpl struct {
    gen.UnimplementedDeviceServer
    mu  *sync.  RWMutex

}

// New initializes a new Backend struct.
func New() *DeviceServerImpl {

	a := C.int(10)
	b := C.int(20)
	c := C.Add(a, b)
	fmt.Println(c)
	
	return &DeviceServerImpl{
		mu: &sync.RWMutex{},
	}
}

// Open is the implementation of RPC call defined in protocol definitions.
// This will take OpenRequest message and return OpenReply
func (g *DeviceServerImpl) Open(ctx context.Context, request *gen.OpenRequest) (*gen.OpenReply, error) {
    return &gen.OpenReply{
        Message: fmt.Sprintf("hello %s",request.DeviceUri),
    },nil
}

func (g *DeviceServerImpl) Close( ctx context.Context, request *gen.OpenRequest) (*gen.CloseReply, error) {
    return &gen.CloseReply{
        Message: fmt.Sprintf("hello %s",request.DeviceUri),
    },nil
}

func (g *DeviceServerImpl) Get( ctx context.Context, request *gen.DoRequest) (*gen.DoResponse, error) {
    return &gen.DoResponse{
        DeviceUri: fmt.Sprintf("Get %s",request.DeviceUri),
    },nil
}

func (g *DeviceServerImpl) Update( ctx context.Context, request *gen.DoRequest) (*gen.DoResponse, error) {
    return &gen.DoResponse{
        DeviceUri: fmt.Sprintf("Update %s",request.DeviceUri),
    },nil
}

func (g *DeviceServerImpl) Add( ctx context.Context, request *gen.DoRequest) (*gen.DoResponse, error) {
    return &gen.DoResponse{
        DeviceUri: fmt.Sprintf("Add %s",request.DeviceUri),
    },nil
}

func (g *DeviceServerImpl) Del( ctx context.Context, request *gen.DoRequest) (*gen.DoResponse, error) {
    return &gen.DoResponse{
        DeviceUri: fmt.Sprintf("Del %s",request.DeviceUri),
    },nil
}

func (g *DeviceServerImpl) Action( ctx context.Context, request *gen.DoRequest) (*gen.DoResponse, error) {
    return &gen.DoResponse{
        DeviceUri: fmt.Sprintf("Action %s",request.DeviceUri),
    },nil
}

func (g *DeviceServerImpl) Message( ctx context.Context, request *gen.DoRequest) (*gen.DoResponse, error) {
    return &gen.DoResponse{
        DeviceUri: fmt.Sprintf("Message %s",request.DeviceUri),
    },nil
}


/*
func main() {

    a := C.int(10)
	b := C.int(20)
	c := C.Add(a, b)
	fmt.Println(c)
    
    // create new gRPC server
    server := grpc.NewServer()
    // register the GreeterServerImpl on the gRPC server
    gen.RegisterDeviceServer(server, &DeviceServerImpl{})
    // start listening on port :8080 for a tcp connection
    if l, err := net.Listen("tcp", ":9980"); err != nil {
        log.Fatal("error in listening on port :8080", err)
    } else {
        // the gRPC server
        if err:=server.Serve(l);err!=nil {
            log.Fatal("unable to start server",err)
        }
    }
}
*/