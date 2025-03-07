

#include <iostream>
#include <memory>
#include <string>

#include <fstream>
#include <sstream>
#include <vector>

#include "absl/flags/flag.h"
#include "absl/flags/parse.h"
#include "absl/log/check.h"

#include <grpc/support/log.h>
#include <grpcpp/grpcpp.h>

#ifdef BAZEL_BUILD
#include "examples/protos/helloworld.grpc.pb.h"
#else
#include "device-sm.grpc.pb.h"
#endif

ABSL_FLAG(std::string, target, "localhost:9050", "Server address");

//#include "device-sm.grpc.pb.h"

using namespace std;

using grpc::Channel;
using grpc::ClientAsyncResponseReader;
using grpc::ClientContext;
using grpc::CompletionQueue;
using grpc::Status;


std::string LoadStringFromFile(std::string path) {
  std::ifstream file(path);
  if (!file.is_open()) {
    std::cout << "Failed to open " << path << std::endl;
    abort();
  }
  std::stringstream sstr;
  sstr << file.rdbuf();
  return sstr.str();
}

#define GRPC_DEVICE_EXPORTS

#ifdef GRPC_DEVICE_EXPORTS
#define GRPC_DEVICE_API __declspec(dllexport)
#else
#define GRPC_DEVICE_API __declspec(dllimport)
#endif

class GrpcDeviceClient {

  
  public:
    GrpcDeviceClient() {}
   virtual ~GrpcDeviceClient() {}
    
    virtual int OpenDevice(string devUri, string group) = 0;
    virtual int CloseDevice() = 0;
  
    virtual string GetDeviceTarget() = 0;
    virtual int GetClientStatus() = 0;
    virtual string GetOpendDeviceResource() = 0;

    //同步接口
    virtual int Get(string devResource, string& devRes, string &calResMsg, string devUri = "",
                    string calSection = "", string calClientID = "") = 0;
    virtual int Set(string devResource, string reqParams, string& devRes, string& calResMsg,
                    string devUri = "", string calSection = "",
                    string calClientID = "") = 0;
    virtual int Update(string devResource, string reqParams, string& devRes, string& calResMsg,
                    string devUri = "", string calSection = "",
                    string calClientID = "") = 0;
    virtual int Add(string devResource, string reqParams, string& devRes, string& calResMsg,
                    string devUri = "", string calSection = "",
                    string calClientID = "") = 0;
    virtual int Del(string devResource, string reqParams, string& devRes, string& calResMsg,
                    string devUri = "", string calSection = "",
                    string calClientID = "") = 0;
    virtual int Action(string devResource, string reqParams, string& devRes, string& calResMsg,
                    string devUri = "", string calSection = "",
                    string calClientID = "") = 0;
    virtual int Message(string devResource, string reqParams, string& devRes, string& calResMsg,
                    string devUri = "", string calSection = "",
                    string calClientID = "") = 0;

    virtual int BeginAyncWait() = 0;
    virtual void WaitAllComplete() = 0;
    virtual int GetAsyncResult(int idx, string& devResource, string& devRes, string& calResMsg) = 0;
    virtual void EndAync() = 0;
    virtual int AsyncAction(string devResource, string reqParams,
                            string devUri = "", string calSection = "",
                            string calClientID = "") = 0;
    virtual int AsyncMessage(string devResource, string reqParams,
                            string devUri = "",  string calSection = "", 
                            string calClientID = "") = 0;

};

class GrpcDeviceClientImpl : public GrpcDeviceClient {
  string m_strDeviceTarget;
  string m_strDestDeviceUri;
  string m_strDeviceResource;

 public:
  GrpcDeviceClientImpl(std::shared_ptr<Channel> channel)
      : stub_(Device::V3::Device::NewStub(channel)) {
    channel_ = channel;
  }

  ~GrpcDeviceClientImpl() {}
  void SetTarget(const char* pszTarget) { m_strDeviceTarget = pszTarget; }

  int OpenDevice(string devUri, string group) override {
    // Construct Open Requet by uri and group
    Device::V3::OpenRequest request;
    request.set_deviceuri(devUri);
    request.set_devicegroup(group);

    // Cached Uri and empty resource
    m_strDestDeviceUri = devUri;
    m_strDeviceResource = "";

    // Container for the data we expect from the server.
    Device::V3::OpenReply reply;
    ClientContext context;
    Status status;

    status = stub_->Open(&context, request, &reply);

    // Act upon its status.
    if (status.ok()) {
      cout << "Open " << devUri << " Result:" << reply.message() << endl;
      int retcode = atoi(reply.retcode().c_str());
      if (2 == retcode) {
        m_strDeviceResource = reply.retcontext();
      }
      return retcode;

    } else {
      std::cout << status.error_code() << ": " << status.error_message()
                << std::endl;
      return -status.error_code();
    }
  }
  int CloseDevice() override { return 2; }

  string GetDeviceTarget() override { return m_strDeviceTarget; }
  int GetClientStatus() override { return channel_->GetState(true); }
  string GetOpendDeviceResource() override { return m_strDeviceResource; }

  int Get(string devResource, string& devRes, string& calResMsg,
          string devUri = "", string calSection = "",
          string calClientID = "") override {
    Device::V3::DoRequest getq;
    if (devUri.length() > 0)
      getq.set_deviceuri(devUri);
    else
      getq.set_deviceuri(m_strDestDeviceUri);
    getq.set_reqname(devResource);
    getq.set_reqtransaction(calSection);
    getq.set_reqclientid(calClientID);

    ClientContext context2;
    Status status;
    Device::V3::DoResponse getr;
    status = stub_->Get(&context2, getq, &getr);

    if (status.ok()) {
      if (getr.retcode() == "2") {
        cout << "Get " << devResource << " Result " << getr.retcontext()
             << endl;
        devRes = getr.retcontext();
        calResMsg = getr.respresult();
      } else
        cout << "Get " << devResource << " failed: " << getr.retcode() << endl;
      return atoi(getr.retcode().c_str());
    }
    cout << "get failed " << devResource << endl;
    cout << "error :" << status.error_message() << " with detail "
         << status.error_details() << endl;
    return -status.error_code();
  }
  int Set(string devResource, string reqParams, string& devRes,
          string& calResMsg, string devUri = "", string calSection = "",
          string calClientID = "") override {
    Device::V3::DoRequest setq;
    if (devUri.length() > 0)
      setq.set_deviceuri(devUri);
    else
      setq.set_deviceuri(m_strDestDeviceUri);
    setq.set_reqname(devResource);
    setq.set_reqparam(reqParams);
    setq.set_reqtransaction(calSection);
    setq.set_reqclientid(calClientID);

    ClientContext context2;
    Status status;
    Device::V3::DoResponse setr;
    status = stub_->Set(&context2, setq, &setr);

    if (status.ok()) {
      if (setr.retcode() == "2") {
        cout << "Set " << devResource << " Result" << setr.retcontext() << endl;
        devRes = setr.retcontext();
        calResMsg = setr.respresult();
      } else
        cout << "Set " << devResource << " failed" << setr.retcode() << endl;
      return atoi(setr.retcode().c_str());
    }
    return -100;
  }
  int Update(string devResource, string reqParams, string& devRes,
             string& calResMsg, string devUri = "", string calSection = "",
             string calClientID = "") override {
    Device::V3::DoRequest setq;
    if (devUri.length() > 0)
      setq.set_deviceuri(devUri);
    else
      setq.set_deviceuri(m_strDestDeviceUri);
    setq.set_reqname(devResource);
    setq.set_reqparam(reqParams);
    setq.set_reqtransaction(calSection);
    setq.set_reqclientid(calClientID);

    ClientContext context2;
    Status status;
    Device::V3::DoResponse setr;
    status = stub_->Update(&context2, setq, &setr);

    if (status.ok()) {
      if (setr.retcode() == "2") {
        cout << "Update " << devResource << " Result" << setr.retcontext()
             << endl;
        devRes = setr.retcontext();
        calResMsg = setr.respresult();
      } else
        cout << "Update " << devResource << " failed" << setr.retcode() << endl;
      return atoi(setr.retcode().c_str());
    }
    return -100;
  }
  int Add(string devResource, string reqParams, string& devRes,
          string& calResMsg, string devUri = "", string calSection = "",
          string calClientID = "") override {
    Device::V3::DoRequest setq;
    if (devUri.length() > 0)
      setq.set_deviceuri(devUri);
    else
      setq.set_deviceuri(m_strDestDeviceUri);
    setq.set_reqname(devResource);
    setq.set_reqparam(reqParams);
    setq.set_reqtransaction(calSection);
    setq.set_reqclientid(calClientID);

    ClientContext context2;
    Status status;
    Device::V3::DoResponse setr;
    status = stub_->Add(&context2, setq, &setr);

    if (status.ok()) {
      if (setr.retcode() == "2") {
        cout << "Add " << devResource << " Result" << setr.retcontext() << endl;
        devRes = setr.retcontext();
        calResMsg = setr.respresult();
      } else
        cout << "Add " << devResource << " failed" << setr.retcode() << endl;
      return atoi(setr.retcode().c_str());
    }
    return -100;
  }
  int Del(string devResource, string reqParams, string& devRes,
          string& calResMsg, string devUri = "", string calSection = "",
          string calClientID = "") override {
    Device::V3::DoRequest setq;
    if (devUri.length() > 0)
      setq.set_deviceuri(devUri);
    else
      setq.set_deviceuri(m_strDestDeviceUri);
    setq.set_reqname(devResource);
    setq.set_reqparam(reqParams);
    setq.set_reqtransaction(calSection);
    setq.set_reqclientid(calClientID);

    ClientContext context2;
    Status status;
    Device::V3::DoResponse setr;
    status = stub_->Del(&context2, setq, &setr);

    if (status.ok()) {
      if (setr.retcode() == "2") {
        cout << "Del " << devResource << " Result" << setr.retcontext() << endl;
        devRes = setr.retcontext();
        calResMsg = setr.respresult();
      } else
        cout << "Del " << devResource << " failed" << setr.retcode() << endl;
      return atoi(setr.retcode().c_str());
    }
    return -100;
  }
  int Action(string devResource, string reqParams, string& devRes,
             string& calResMsg, string devUri = "", string calSection = "",
             string calClientID = "") override {
    Device::V3::DoRequest setq;
    if (devUri.length() > 0)
      setq.set_deviceuri(devUri);
    else
      setq.set_deviceuri(m_strDestDeviceUri);
    setq.set_reqname(devResource);
    setq.set_reqparam(reqParams);
    setq.set_reqtransaction(calSection);
    setq.set_reqclientid(calClientID);

    ClientContext context2;
    Status status;
    Device::V3::DoResponse setr;
    status = stub_->Action(&context2, setq, &setr);

    if (status.ok()) {
      if (setr.retcode() == "2") {
        cout << "Action " << devResource << " Result: " << setr.retcontext()
             << endl;
        devRes = setr.retcontext();
        calResMsg = setr.respresult();
      } else
        cout << "Action " << devResource << " failed: " << setr.retcode()
             << endl;
      return atoi(setr.retcode().c_str());
    }
    cout << "action failed " << devResource << endl;
    cout << "error :" << status.error_message() << " with detail "
         << status.error_details() << endl;
    return -status.error_code();
  }
  int Message(string devResource, string reqParams, string& devRes,
              string& calResMsg, string devUri = "", string calSection = "",
              string calClientID = "") override {
    Device::V3::DoRequest setq;
    if (devUri.length() > 0)
      setq.set_deviceuri(devUri);
    else
      setq.set_deviceuri(m_strDestDeviceUri);
    setq.set_reqname(devResource);
    setq.set_reqparam(reqParams);
    setq.set_reqtransaction(calSection);
    setq.set_reqclientid(calClientID);

    ClientContext context2;
    Status status;
    Device::V3::DoResponse setr;
    status = stub_->Message(&context2, setq, &setr);

    if (status.ok()) {
      if (setr.retcode() == "2") {
        cout << "Message " << devResource << " Result" << setr.retcontext();
        devRes = setr.retcontext();
        calResMsg = setr.respresult();
      } else
        cout << "Message " << devResource << " failed" << setr.retcode();
      return atoi(setr.retcode().c_str());
    }
    return -100;
  }

  int AsyncAction(string devResource, string reqParams, string devUri = "", string calSection = "",
                  string calClientID = "") override {
    Device::V3::DoRequest setq;
    if (devUri.length() > 0)
      setq.set_deviceuri(devUri);
    else
      setq.set_deviceuri(m_strDestDeviceUri);
    setq.set_reqname(devResource);
    setq.set_reqparam(reqParams);
    setq.set_reqtransaction(calSection);
    setq.set_reqclientid(calClientID);

    // Call object to store rpc data
    AsyncClientCall* call = new AsyncClientCall;

    // stub_->PrepareAsyncSayHello() creates an RPC object, returning
    // an instance to store in "call" but does not actually start the RPC
    // Because we are using the asynchronous API, we need to hold on to
    // the "call" instance in order to get updates on the ongoing RPC.
    call->response_reader =
        stub_->PrepareAsyncAction(&call->context, setq, &cq_);

    // StartCall initiates the RPC call
    call->response_reader->StartCall();

    // Request that, upon completion of the RPC, "reply" be updated with the
    // server's response; "status" with the indication of whether the
    // operation was successful. Tag the request with the memory address of
    // the call object.
    call->response_reader->Finish(&call->reply, &call->status, (void*)call);

    m_reqList.push_back(call);
    m_nNeedResult = m_reqList.size();
    return m_nNeedResult;
  }
  int AsyncMessage(string devResource, string reqParams, string devUri = "", string calSection = "",
                  string calClientID = "") override {
    Device::V3::DoRequest setq;
    if (devUri.length() > 0)
      setq.set_deviceuri(devUri);
    else
      setq.set_deviceuri(m_strDestDeviceUri);
    setq.set_reqname(devResource);
    setq.set_reqparam(reqParams);
    setq.set_reqtransaction(calSection);
    setq.set_reqclientid(calClientID);

    // Call object to store rpc data
    AsyncClientCall* call = new AsyncClientCall;

    // stub_->PrepareAsyncSayHello() creates an RPC object, returning
    // an instance to store in "call" but does not actually start the RPC
    // Because we are using the asynchronous API, we need to hold on to
    // the "call" instance in order to get updates on the ongoing RPC.
    call->response_reader =
        stub_->PrepareAsyncMessage(&call->context, setq, &cq_);

    // StartCall initiates the RPC call
    call->response_reader->StartCall();

    // Request that, upon completion of the RPC, "reply" be updated with the
    // server's response; "status" with the indication of whether the
    // operation was successful. Tag the request with the memory address of
    // the call object.
    call->response_reader->Finish(&call->reply, &call->status, (void*)call);

    m_reqList.push_back(call);
    m_nNeedResult = m_reqList.size();
    return m_nNeedResult;
  }
  // Loop while listening for completed responses.
  // Prints out the response from the server.
  void AsyncCompleteRpc() {
    void* got_tag;
    bool ok = false;

    int gots = 0;
    // Block until the next result is available in the completion queue "cq".
    while (cq_.Next(&got_tag, &ok) ) {
      // The tag in this example is the memory location of the call object
      AsyncClientCall* call = static_cast<AsyncClientCall*>(got_tag);

      // Verify that the request was completed successfully. Note that "ok"
      // corresponds solely to the request for updates introduced by Finish().
      CHECK(ok);

      if (call->status.ok())
        std::cout << "Device Client received: " << call->reply.respresult()
                  << std::endl;
      else
        std::cout << "RPC failed" << std::endl;

      gots++;

      if (gots >= m_nNeedResult) {
        std::cout << "Wait All over.." << gots << endl;
        break;
      }
      std::cout << "Wait Again.... " << gots << endl;
      // Once we're complete, deallocate the call object.
      // delete call;
    }
  }

  int BeginAyncWait() override {
    // Spawn reader thread that loops indefinitely
    if (m_reqList.size() > 0)
    {
      for each (auto var in m_reqList) {
        delete var;
      }
      m_reqList.clear();
    }
    pthread_ = new std::thread(&GrpcDeviceClientImpl::AsyncCompleteRpc, this);

    return 2;
  }
  void WaitAllComplete() override {
    // wait for thread ok.
    pthread_->join();  // blocks forever

    delete pthread_;
  }
  int GetAsyncResult(int idx, string& devResource, string& devRes,
                     string& calResMsg) override {
    if (idx < m_reqList.size()) 
    {
      devResource = m_reqList[idx]->reply.reqname();
      if (m_reqList[idx]->status.ok()) {
        if (m_reqList[idx]->reply.retcode() == "2") 
        {
          cout << "Action " << devResource
                << " Result: " << m_reqList[idx]->reply.retcontext() << endl;
          devRes = m_reqList[idx]->reply.retcontext();
          calResMsg = m_reqList[idx]->reply.respresult();
        } 
        else 
        {
          cout << "Action " << devResource
            << " failed: " << m_reqList[idx]->reply.retcode() << endl;            
        }
        return atoi(m_reqList[idx]->reply.retcode().c_str());
      }
     cout << "action failed " << devResource << endl;
      cout << "error :" << m_reqList[idx]->status.error_message()
           << " with detail " << m_reqList[idx]->status.error_details() << endl;
      return -m_reqList[idx]->status.error_code();
    }
    else
    {
      cout << "invalid idx" << idx << endl;
    }
    return 0;
  }
  void EndAync() override { 
    for each (auto var in m_reqList) {
      delete var;
    }
    m_reqList.clear();
  }

 private:
    // struct for keeping state and data information
    struct AsyncClientCall {
      // Container for the data we expect from the server.
      Device::V3::DoResponse reply;

      // Context for the client. It could be used to convey extra information to
      // the server and/or tweak certain RPC behaviors.
      ClientContext context;

      // Storage for the status of the RPC upon completion.
      Status status;

      std::unique_ptr<ClientAsyncResponseReader<Device::V3::DoResponse>>
          response_reader;
    };

    // Out of the passed in Channel comes the stub, stored here, our view of the
    // server's exposed services.
    std::unique_ptr<Device::V3::Device::Stub> stub_;
    std::shared_ptr<Channel> channel_;

    // The producer-consumer queue we use to communicate asynchronously with the
    // gRPC runtime.
    CompletionQueue cq_;

    //wait Thread
    std::thread* pthread_;

    std::vector<AsyncClientCall*> m_reqList;
    int m_nNeedResult;
 };


 static string GetProcessDirectory() {
   string ret = "";
   char szFilename[MAX_PATH] = {0};
   DWORD res = GetModuleFileNameA(0, szFilename, MAX_PATH);
   if (res == 0) {
     return ret;
   }

   string fullpath = szFilename;
   string::size_type firstHit = fullpath.find_last_of('\\');
   if (firstHit == string::npos || firstHit == 0) {
     return ret;
   }

   ret = fullpath.substr(0, firstHit);  //kick last \

   return ret;
 }

GRPC_DEVICE_API GrpcDeviceClient* CreateClient(const char* pszDeviceServerTarget) {

    std::string target_str = pszDeviceServerTarget;
  cout << "try connect " << target_str << endl;
  auto channel_creds = grpc::SslCredentials(grpc::SslCredentialsOptions());


  string certFile = GetProcessDirectory();
  certFile += "\\root.crt";
  grpc::SslCredentialsOptions ssl_options;
  ssl_options.pem_root_certs = LoadStringFromFile(certFile);
  // Create a channel with SSL credentials
  cout << "with key" << ssl_options.pem_root_certs << endl;
  GrpcDeviceClientImpl* pClient = new GrpcDeviceClientImpl( grpc::CreateChannel(target_str, grpc::SslCredentials(ssl_options)));

  cout << "Connect " << target_str << " status: " <<  pClient->GetClientStatus() << endl;

  pClient->SetTarget(pszDeviceServerTarget);

  return pClient;
}


GRPC_DEVICE_API void FreeClient(GrpcDeviceClient* pClient) {
  if (pClient != nullptr) delete pClient;
}


class GreeterClient {
 public:
  explicit GreeterClient(std::shared_ptr<Channel> channel)
      : stub_(Device::V3::Device::NewStub(channel)) {}

  // Assembles the client's payload, sends it and presents the response back
  // from the server.
  std::string SayHello(const std::string& user) {
    // Data we are sending to the server.
    Device::V3::OpenRequest request;
    request.set_deviceuri(user);

    // Container for the data we expect from the server.
    Device::V3::OpenReply reply;

    // Context for the client. It could be used to convey extra information to
    // the server and/or tweak certain RPC behaviors.
    ClientContext context;

    // The producer-consumer queue we use to communicate asynchronously with the
    // gRPC runtime.
    CompletionQueue cq;

    // Storage for the status of the RPC upon completion.
    Status status;

   status = stub_->Open(&context, request, &reply);

    // Act upon its status.
    if (status.ok()) {
     cout << reply.retcontext() << endl;


     Device::V3::DoRequest getq;
     getq.set_deviceuri(user);
     getq.set_reqname("");

     ClientContext context2;
     Device::V3::DoResponse getr;
     status = stub_->Get(&context2, getq, &getr);

     if (status.ok())
     {
       if (getr.retcode() == "2")
          cout << "Get DIOSDriverList Result" <<  getr.retcontext();
       else
         cout << "Get DIOSDriverList failed" << getr.retcode();
     }
      return reply.message();
    } else {
      std::cout << status.error_code() << ": " << status.error_message()
                << std::endl;
      return "RPC failed";
    }
  }

 private:
  // Out of the passed in Channel comes the stub, stored here, our view of the
  // server's exposed services.
  std::unique_ptr<Device::V3::Device::Stub> stub_;

};



int main(int argc, char** argv) {
  absl::ParseCommandLine(argc, argv);
  // Instantiate the client. It requires a channel, out of which the actual RPCs
  // are created. This channel models a connection to an endpoint specified by
  // the argument "--target=" which is the only expected argument.

 auto clt = CreateClient("localhost:9010");
  clt->OpenDevice("DIOS/DEVICE/Detector", "");
 string isDemo, msg, devRes;
  clt->Get("AcqMode", isDemo, msg);
 clt->Action("EnterExam", "", devRes, msg);
 clt->Action("RESET", "", devRes, msg);
 clt->Action("PrepareAcquisition", "", devRes, msg);

 cout << endl << endl << "Start Aync Call action..." << endl;
 clt->BeginAyncWait();
 int nWait = 0;
 nWait = clt->AsyncAction("EnterExam", "");
 nWait = clt->AsyncAction("RESET", "");
 nWait = clt->AsyncAction("PrepareAcquisition", "");

 cout <<endl << "Action Called Start wait end...." << endl << endl;
 clt->WaitAllComplete();
 
 cout << endl << "Try Get Result .... " << endl;
 int res = 0;
 string act;
 for (int x = 0; x < nWait; x++)
 {
   res = clt->GetAsyncResult(x, act, devRes, msg);
   cout << "[" << x << "] call "<< act <<" result code=" << res << " res:" << devRes
        << " msg: " << msg << endl;
 }
 clt->EndAync();

 FreeClient(clt);

 /*
  std::string target_str = absl::GetFlag(FLAGS_target);
  cout << "try connect " << target_str << endl;
  auto channel_creds = grpc::SslCredentials(grpc::SslCredentialsOptions());
  
    grpc::SslCredentialsOptions ssl_options;
  ssl_options.pem_root_certs = LoadStringFromFile("root.crt");
  // Create a channel with SSL credentials
  cout << "with key" << ssl_options.pem_root_certs << endl;
  GreeterClient greeter(
      grpc::CreateChannel(target_str, grpc::SslCredentials(ssl_options)));
  // We indicate that the channel isn't authenticated (use of
  // InsecureChannelCredentials()).
 // GreeterClient greeter(grpc::CreateChannel(target_str, channel_creds));
  std::string user("DIOS/DEVICE/Subsystem");
  std::string reply = greeter.SayHello(user);  // The actual RPC call!
  std::cout << "Greeter received: " << reply << std::endl;
  */
  return 0;
}

/*
using grpc::Channel;
using grpc::ClientContext;
using grpc::Status;

class Client {
 public:
  Client(std::shared_ptr<grpc::Channel> channel)
      : stub_(Device::V3::Device::NewStub(channel)) {}

  int Open() {
    Device::V3::OpenRequest opReq;
    opReq.set_deviceuri("");
    opReq.set_devicegroup("");

    Device::V3::OpenReply opRes;
    grpc::ClientContext context;
    grpc::Status status = stub_->Open(&context, opReq, &opRes);
    if (status.ok()) {
      std::cout << opRes.retcontext() << std::endl;
      return 1;
    }
    return 0;
  }

 private:
  std::unique_ptr<Device::V3::Device::Stub> stub_;
};

int main() {
  cout << "Hello CMake." << endl;
  Client client(grpc::CreateChannel("locaohst::9051",
                                    grpc::InsecureChannelCredentials()));

  return 0;
}*/
