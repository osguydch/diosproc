package server


/*

struct OpenReply {
		char* message;
		char* context;
	};

	struct DoResponse {
		char* reqName;
		char* RespResult;
		char* retContext;
	};

int open_device(int devIdx, char* devUri, char* devGroup, void** reply);
int close_device(int devIdx, char* devUri, char* devGroup, void** reply);

int device_get(int devIdx, char* devUri, char* devProperty, void** reply);
int device_set(int devIdx, char* devUri, char* devProperty, char* devValueSet, void** reply);
int device_update(int devIdx, char* devUri, char* devProperty, char* devValueUpdate, void** reply);

int device_add(int devIdx, char* devUri, char* devProperty, char* devValueAdd, void** reply);
int device_del(int devIdx, char* devUri, char* devProperty, char* devValueDel, void** reply);

int device_action(int devIdx, char* devUri, char* devFunction, char* devReqParams, void** reply);
int device_message(int devIdx, char* devUri, char* devTopic, char* devMessageValue, void** reply);


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
    //"strings"
    //"encoding/json"

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

    // if idx < 0 {
    //     return &gen.OpenReply{
    //         Message: "enum all devices ok",
    //         RetCode : "2",
    //         RetContext : string(json.Marshal(g.devs)),
    //     },nil
    // }

    var openReply  unsafe.Pointer
    ret := C.open_device(C.int(idx), C.CString(request.DeviceUri), C.CString(request.DeviceGroup), &openReply)

    if ret != 2 {
        return &gen.OpenReply{
            Message: fmt.Sprintf("grpc error for open %s",request.DeviceUri),
            RetContext : "Open Call Error",
        },nil
    }
    defer C.free_struct(openReply, C.int(3))

    var reply *C.struct_OpenReply
    reply = (*C.struct_OpenReply)(openReply)

    var context string
    context = C.GoString(reply.context)
    //context = strings.ReplaceAll(context,"\\n","")
    //context = strings.ReplaceAll(context,"\\\"","\"")
    /* var result map[string]interface{}
    err := json.Unmarshal([]byte(context), &result)
    if err != nil {
        log.Fatalf("Error occurred during unmarshaling. Error: %s", err.Error())
    }

    jsonBytes, err := json.Marshal(result)
    if err != nil {
        panic(err)
    }
    */
    return &gen.OpenReply{
        Message: C.GoString(reply.message),//fmt.Sprintf("hello %s",request.DeviceUri),
        RetCode : fmt.Sprintf("%d",ret),
        RetContext : context,
    },nil
}

func (g *DeviceServerImpl) Close( ctx context.Context, request *gen.OpenRequest) (*gen.OpenReply, error) {
    g.mu.Lock()
	defer g.mu.Unlock()

    idx := g.devs[request.DeviceUri]


    var openReply  unsafe.Pointer
    ret := C.close_device(C.int(idx), C.CString(request.DeviceUri), C.CString(request.DeviceGroup), &openReply)

    if ret != 2 {
        return &gen.OpenReply{
            Message: fmt.Sprintf("grpc error for close %s",request.DeviceUri),
            RetContext : "Close Call Error",
        },nil
    }
    defer C.free_struct(openReply, C.int(3))

    var reply *C.struct_OpenReply
    reply = (*C.struct_OpenReply)(openReply)

    var context string
    context = C.GoString(reply.context)
    //context = strings.ReplaceAll(context,"\\n","")
    //context = strings.ReplaceAll(context,"\\\"","\"")
    /* var result map[string]interface{}
    err := json.Unmarshal([]byte(context), &result)
    if err != nil {
        log.Fatalf("Error occurred during unmarshaling. Error: %s", err.Error())
    }

    jsonBytes, err := json.Marshal(result)
    if err != nil {
        panic(err)
    }
    */
    return &gen.OpenReply{
        Message: C.GoString(reply.message),//fmt.Sprintf("hello %s",request.DeviceUri),
        RetCode : fmt.Sprintf("%d",ret),
        RetContext : context,
    },nil
}

func (g *DeviceServerImpl) Get( ctx context.Context, request *gen.DoRequest) (*gen.DoResponse, error) {
    
    g.mu.Lock()
	defer g.mu.Unlock()

    idx := g.devs[request.DeviceUri]

    var openReply  unsafe.Pointer
    ret := C.device_get(C.int(idx), C.CString(request.DeviceUri), C.CString(request.ReqName), &openReply)

    if ret != 2 {
        return &gen.DoResponse{
            DeviceUri: request.DeviceUri,
            RetCode : fmt.Sprintf("%d",ret),
            RetContext : "Get Call Error",
        },nil
    }
    defer C.free_struct(openReply, C.int(3))

    var reply *C.struct_DoResponse
    reply = (*C.struct_DoResponse)(openReply)
    
    
    return &gen.DoResponse{
        DeviceUri: request.DeviceUri,
        RetCode : fmt.Sprintf("%d",ret),
        ReqName : request.ReqName,
        RespResult : C.GoString(reply.RespResult),
        ReqTransaction : request.ReqTransaction,
        RetContext : C.GoString(reply.retContext),
    },nil
}

func (g *DeviceServerImpl) Set( ctx context.Context, request *gen.DoRequest) (*gen.DoResponse, error) {

    g.mu.Lock()
	defer g.mu.Unlock()

    idx := g.devs[request.DeviceUri]

    var openReply  unsafe.Pointer
    ret := C.device_set(C.int(idx), C.CString(request.DeviceUri), C.CString(request.ReqName), 
    C.CString(request.ReqParam), &openReply)

    if ret != 2 {
        return &gen.DoResponse{
            DeviceUri: request.DeviceUri,
            RetCode : fmt.Sprintf("%d",ret),
            RetContext : "Set Call Error",
        },nil
    }
    defer C.free_struct(openReply, C.int(3))

    var reply *C.struct_DoResponse
    reply = (*C.struct_DoResponse)(openReply)
    
    
    return &gen.DoResponse{
        DeviceUri: request.DeviceUri,
        RetCode : fmt.Sprintf("%d",ret),
        ReqName : request.ReqName,
        RespResult : C.GoString(reply.RespResult),
        ReqTransaction : request.ReqTransaction,
        RetContext : C.GoString(reply.retContext),
    },nil
    
}

func (g *DeviceServerImpl) Update( ctx context.Context, request *gen.DoRequest) (*gen.DoResponse, error) {
    g.mu.Lock()
	defer g.mu.Unlock()

    idx := g.devs[request.DeviceUri]

    var openReply  unsafe.Pointer
    ret := C.device_update(C.int(idx), C.CString(request.DeviceUri), C.CString(request.ReqName), 
    C.CString(request.ReqParam), &openReply)

    if ret != 2 {
        return &gen.DoResponse{
            DeviceUri: request.DeviceUri,
            RetCode : fmt.Sprintf("%d",ret),
            RetContext : "Update Call Error",
        },nil
    }
    defer C.free_struct(openReply, C.int(3))

    var reply *C.struct_DoResponse
    reply = (*C.struct_DoResponse)(openReply)
    
    
    return &gen.DoResponse{
        DeviceUri: request.DeviceUri,
        RetCode : fmt.Sprintf("%d",ret),
        ReqName : request.ReqName,
        RespResult : C.GoString(reply.RespResult),
        ReqTransaction : request.ReqTransaction,
        RetContext : C.GoString(reply.retContext),
    },nil
}

func (g *DeviceServerImpl) Add( ctx context.Context, request *gen.DoRequest) (*gen.DoResponse, error) {
    g.mu.Lock()
	defer g.mu.Unlock()

    idx := g.devs[request.DeviceUri]

    var openReply  unsafe.Pointer
    ret := C.device_add(C.int(idx), C.CString(request.DeviceUri), C.CString(request.ReqName), 
    C.CString(request.ReqParam), &openReply)

    if ret != 2 {
        return &gen.DoResponse{
            DeviceUri: request.DeviceUri,
            RetCode : fmt.Sprintf("%d",ret),
            RetContext : "Add Call Error",
        },nil
    }
    defer C.free_struct(openReply, C.int(3))

    var reply *C.struct_DoResponse
    reply = (*C.struct_DoResponse)(openReply)
    
    
    return &gen.DoResponse{
        DeviceUri: request.DeviceUri,
        RetCode : fmt.Sprintf("%d",ret),
        ReqName : request.ReqName,
        RespResult : C.GoString(reply.RespResult),
        ReqTransaction : request.ReqTransaction,
        RetContext : C.GoString(reply.retContext),
    },nil
}

func (g *DeviceServerImpl) Del( ctx context.Context, request *gen.DoRequest) (*gen.DoResponse, error) {
    g.mu.Lock()
	defer g.mu.Unlock()

    idx := g.devs[request.DeviceUri]

    var openReply  unsafe.Pointer
    ret := C.device_del(C.int(idx), C.CString(request.DeviceUri), C.CString(request.ReqName), 
    C.CString(request.ReqParam), &openReply)

    if ret != 2 {
        return &gen.DoResponse{
            DeviceUri: request.DeviceUri,
            RetCode : fmt.Sprintf("%d",ret),
            RetContext : "Del Call Error",
        },nil
    }
    defer C.free_struct(openReply, C.int(3))

    var reply *C.struct_DoResponse
    reply = (*C.struct_DoResponse)(openReply)
    
    
    return &gen.DoResponse{
        DeviceUri: request.DeviceUri,
        RetCode : fmt.Sprintf("%d",ret),
        ReqName : request.ReqName,
        RespResult : C.GoString(reply.RespResult),
        ReqTransaction : request.ReqTransaction,
        RetContext : C.GoString(reply.retContext),
    },nil
}

func (g *DeviceServerImpl) Action( ctx context.Context, request *gen.DoRequest) (*gen.DoResponse, error) {
    g.mu.Lock()
	defer g.mu.Unlock()

    idx := g.devs[request.DeviceUri]

    var openReply  unsafe.Pointer
    ret := C.device_action(C.int(idx), C.CString(request.DeviceUri), C.CString(request.ReqName), 
    C.CString(request.ReqParam), &openReply)

    if ret != 2 {
        return &gen.DoResponse{
            DeviceUri: request.DeviceUri,
            RetCode : fmt.Sprintf("%d",ret),
            RetContext : "Action Call Error",
        },nil
    }
    defer C.free_struct(openReply, C.int(3))

    var reply *C.struct_DoResponse
    reply = (*C.struct_DoResponse)(openReply)
    
    
    return &gen.DoResponse{
        DeviceUri: request.DeviceUri,
        RetCode : fmt.Sprintf("%d",ret),
        ReqName : request.ReqName,
        RespResult : C.GoString(reply.RespResult),
        ReqTransaction : request.ReqTransaction,
        RetContext : C.GoString(reply.retContext),
    },nil
}

func (g *DeviceServerImpl) Message( ctx context.Context, request *gen.DoRequest) (*gen.DoResponse, error) {
    g.mu.Lock()
	defer g.mu.Unlock()

    idx := g.devs[request.DeviceUri]

    var openReply  unsafe.Pointer
    ret := C.device_message(C.int(idx), C.CString(request.DeviceUri), C.CString(request.ReqName), 
    C.CString(request.ReqParam), &openReply)

    if ret != 2 {
        return &gen.DoResponse{
            DeviceUri: request.DeviceUri,
            RetCode : fmt.Sprintf("%d",ret),
            RetContext : "Message Call Error",
        },nil
    }
    defer C.free_struct(openReply, C.int(3))

    var reply *C.struct_DoResponse
    reply = (*C.struct_DoResponse)(openReply)
    
    
    return &gen.DoResponse{
        DeviceUri: request.DeviceUri,
        RetCode : fmt.Sprintf("%d",ret),
        ReqName : request.ReqName,
        RespResult : C.GoString(reply.RespResult),
        ReqTransaction : request.ReqTransaction,
        RetContext : C.GoString(reply.retContext),
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