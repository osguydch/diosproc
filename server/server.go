package server


/*

struct OpenReply {
		char* message;
		char* retCode;
		char* context;
	};

	struct CloseReply {
		char* message;
		char* retCode;
	};

	struct DoRequest {
		char* deviceUr;
		char* typeName;
		char* reqName;
		char* reqParam;
		char* reqTransaction;
		char* reqClientID;
		char* reqTimeStamp;
	};

	struct DoResponse {
		char* deviceUri;
		char* typeName;
		char* reqName;
		char* RespParam ;
		char* reqTransaction;
		char* reqClientID;
		char* RespTimeStamp;
		char* retContext;
	};

int open_device(int devIdx, char* devUri, char* devGroup, void** reply);
int close_device(int devIdx, char* devUri, char* devGroup, void** reply);
int device_get();
int device_update();
int device_add();
int device_del();
int device_action();
int device_message();
int free_struct(void* pPoirnt, int nItemCount);

int Add(int a, int b){
    return a+b;
}
*/
import "C"


import (
    "context"
    "fmt"
    _ "reflect"
    //"log"
    //"net"
    "sync"
    "unsafe"

    //"github.com/gofrs/uuid"
    // importing generated stubs
    gen "github.com/osguydch/diosproc/proto/device/v3"
    //"google.golang.org/grpc"
)
// DeviceServerImpl will implement the service defined in protocol buffer definitions
type DeviceServerImpl struct {
    gen.UnimplementedDeviceServer
    mu  *sync.  RWMutex
    devs map[string]int
}

// New initializes a new Backend struct.
func New(dev []string) *DeviceServerImpl {

	a := C.int(10)
	b := C.int(20)
	c := C.Add(a, b)
	fmt.Println(c)

    devs := make(map[string]int)

    for i,v:= range dev {
        devs[v] = i
    }

	return &DeviceServerImpl{
		mu: &sync.RWMutex{},
        devs :devs,
	}
}

// Open is the implementation of RPC call defined in protocol definitions.
// This will take OpenRequest message and return OpenReply
func (g *DeviceServerImpl) Open(ctx context.Context, request *gen.OpenRequest) (*gen.OpenReply, error) {
	//C.open_device()
    g.mu.Lock()
	defer g.mu.Unlock()

    idx := g.devs[request.DeviceUri]

    var openReply  unsafe.Pointer
    ret := C.open_device(C.int(idx), C.CString(request.DeviceUri), C.CString(request.DeviceGroup), &openReply)

    defer C.free_struct(openReply, C.int(3))

    var reply *C.struct_OpenReply
    reply = (*C.struct_OpenReply)(openReply)

    return &gen.OpenReply{
        Message: C.GoString(reply.message),//fmt.Sprintf("hello %s",request.DeviceUri),
        RetCode : fmt.Sprintf("%d",ret),
        RetContext : C.GoString(reply.context),
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