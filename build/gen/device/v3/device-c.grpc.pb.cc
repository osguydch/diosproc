// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: device/v3/device-c.proto

#include "device/v3/device-c.pb.h"
#include "device/v3/device-c.grpc.pb.h"

#include <functional>
#include <grpcpp/support/async_stream.h>
#include <grpcpp/support/async_unary_call.h>
#include <grpcpp/impl/channel_interface.h>
#include <grpcpp/impl/client_unary_call.h>
#include <grpcpp/support/client_callback.h>
#include <grpcpp/support/message_allocator.h>
#include <grpcpp/support/method_handler.h>
#include <grpcpp/impl/rpc_service_method.h>
#include <grpcpp/support/server_callback.h>
#include <grpcpp/impl/server_callback_handlers.h>
#include <grpcpp/server_context.h>
#include <grpcpp/impl/service_type.h>
#include <grpcpp/support/sync_stream.h>
namespace Device {
namespace V3 {

static const char* Device_method_names[] = {
  "/Device.V3.Device/Open",
  "/Device.V3.Device/Close",
  "/Device.V3.Device/Get",
  "/Device.V3.Device/Set",
  "/Device.V3.Device/Update",
  "/Device.V3.Device/Add",
  "/Device.V3.Device/Del",
  "/Device.V3.Device/Action",
  "/Device.V3.Device/Message",
};

std::unique_ptr< Device::Stub> Device::NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options) {
  (void)options;
  std::unique_ptr< Device::Stub> stub(new Device::Stub(channel, options));
  return stub;
}

Device::Stub::Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options)
  : channel_(channel), rpcmethod_Open_(Device_method_names[0], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_Close_(Device_method_names[1], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_Get_(Device_method_names[2], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_Set_(Device_method_names[3], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_Update_(Device_method_names[4], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_Add_(Device_method_names[5], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_Del_(Device_method_names[6], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_Action_(Device_method_names[7], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_Message_(Device_method_names[8], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  {}

::grpc::Status Device::Stub::Open(::grpc::ClientContext* context, const ::Device::V3::OpenRequest& request, ::Device::V3::OpenReply* response) {
  return ::grpc::internal::BlockingUnaryCall< ::Device::V3::OpenRequest, ::Device::V3::OpenReply, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_Open_, context, request, response);
}

void Device::Stub::async::Open(::grpc::ClientContext* context, const ::Device::V3::OpenRequest* request, ::Device::V3::OpenReply* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::Device::V3::OpenRequest, ::Device::V3::OpenReply, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_Open_, context, request, response, std::move(f));
}

void Device::Stub::async::Open(::grpc::ClientContext* context, const ::Device::V3::OpenRequest* request, ::Device::V3::OpenReply* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_Open_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::Device::V3::OpenReply>* Device::Stub::PrepareAsyncOpenRaw(::grpc::ClientContext* context, const ::Device::V3::OpenRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::Device::V3::OpenReply, ::Device::V3::OpenRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_Open_, context, request);
}

::grpc::ClientAsyncResponseReader< ::Device::V3::OpenReply>* Device::Stub::AsyncOpenRaw(::grpc::ClientContext* context, const ::Device::V3::OpenRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncOpenRaw(context, request, cq);
  result->StartCall();
  return result;
}

::grpc::Status Device::Stub::Close(::grpc::ClientContext* context, const ::Device::V3::OpenRequest& request, ::Device::V3::OpenReply* response) {
  return ::grpc::internal::BlockingUnaryCall< ::Device::V3::OpenRequest, ::Device::V3::OpenReply, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_Close_, context, request, response);
}

void Device::Stub::async::Close(::grpc::ClientContext* context, const ::Device::V3::OpenRequest* request, ::Device::V3::OpenReply* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::Device::V3::OpenRequest, ::Device::V3::OpenReply, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_Close_, context, request, response, std::move(f));
}

void Device::Stub::async::Close(::grpc::ClientContext* context, const ::Device::V3::OpenRequest* request, ::Device::V3::OpenReply* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_Close_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::Device::V3::OpenReply>* Device::Stub::PrepareAsyncCloseRaw(::grpc::ClientContext* context, const ::Device::V3::OpenRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::Device::V3::OpenReply, ::Device::V3::OpenRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_Close_, context, request);
}

::grpc::ClientAsyncResponseReader< ::Device::V3::OpenReply>* Device::Stub::AsyncCloseRaw(::grpc::ClientContext* context, const ::Device::V3::OpenRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncCloseRaw(context, request, cq);
  result->StartCall();
  return result;
}

::grpc::Status Device::Stub::Get(::grpc::ClientContext* context, const ::Device::V3::DoRequest& request, ::Device::V3::DoResponse* response) {
  return ::grpc::internal::BlockingUnaryCall< ::Device::V3::DoRequest, ::Device::V3::DoResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_Get_, context, request, response);
}

void Device::Stub::async::Get(::grpc::ClientContext* context, const ::Device::V3::DoRequest* request, ::Device::V3::DoResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::Device::V3::DoRequest, ::Device::V3::DoResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_Get_, context, request, response, std::move(f));
}

void Device::Stub::async::Get(::grpc::ClientContext* context, const ::Device::V3::DoRequest* request, ::Device::V3::DoResponse* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_Get_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::Device::V3::DoResponse>* Device::Stub::PrepareAsyncGetRaw(::grpc::ClientContext* context, const ::Device::V3::DoRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::Device::V3::DoResponse, ::Device::V3::DoRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_Get_, context, request);
}

::grpc::ClientAsyncResponseReader< ::Device::V3::DoResponse>* Device::Stub::AsyncGetRaw(::grpc::ClientContext* context, const ::Device::V3::DoRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncGetRaw(context, request, cq);
  result->StartCall();
  return result;
}

::grpc::Status Device::Stub::Set(::grpc::ClientContext* context, const ::Device::V3::DoRequest& request, ::Device::V3::DoResponse* response) {
  return ::grpc::internal::BlockingUnaryCall< ::Device::V3::DoRequest, ::Device::V3::DoResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_Set_, context, request, response);
}

void Device::Stub::async::Set(::grpc::ClientContext* context, const ::Device::V3::DoRequest* request, ::Device::V3::DoResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::Device::V3::DoRequest, ::Device::V3::DoResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_Set_, context, request, response, std::move(f));
}

void Device::Stub::async::Set(::grpc::ClientContext* context, const ::Device::V3::DoRequest* request, ::Device::V3::DoResponse* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_Set_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::Device::V3::DoResponse>* Device::Stub::PrepareAsyncSetRaw(::grpc::ClientContext* context, const ::Device::V3::DoRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::Device::V3::DoResponse, ::Device::V3::DoRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_Set_, context, request);
}

::grpc::ClientAsyncResponseReader< ::Device::V3::DoResponse>* Device::Stub::AsyncSetRaw(::grpc::ClientContext* context, const ::Device::V3::DoRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncSetRaw(context, request, cq);
  result->StartCall();
  return result;
}

::grpc::Status Device::Stub::Update(::grpc::ClientContext* context, const ::Device::V3::DoRequest& request, ::Device::V3::DoResponse* response) {
  return ::grpc::internal::BlockingUnaryCall< ::Device::V3::DoRequest, ::Device::V3::DoResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_Update_, context, request, response);
}

void Device::Stub::async::Update(::grpc::ClientContext* context, const ::Device::V3::DoRequest* request, ::Device::V3::DoResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::Device::V3::DoRequest, ::Device::V3::DoResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_Update_, context, request, response, std::move(f));
}

void Device::Stub::async::Update(::grpc::ClientContext* context, const ::Device::V3::DoRequest* request, ::Device::V3::DoResponse* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_Update_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::Device::V3::DoResponse>* Device::Stub::PrepareAsyncUpdateRaw(::grpc::ClientContext* context, const ::Device::V3::DoRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::Device::V3::DoResponse, ::Device::V3::DoRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_Update_, context, request);
}

::grpc::ClientAsyncResponseReader< ::Device::V3::DoResponse>* Device::Stub::AsyncUpdateRaw(::grpc::ClientContext* context, const ::Device::V3::DoRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncUpdateRaw(context, request, cq);
  result->StartCall();
  return result;
}

::grpc::Status Device::Stub::Add(::grpc::ClientContext* context, const ::Device::V3::DoRequest& request, ::Device::V3::DoResponse* response) {
  return ::grpc::internal::BlockingUnaryCall< ::Device::V3::DoRequest, ::Device::V3::DoResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_Add_, context, request, response);
}

void Device::Stub::async::Add(::grpc::ClientContext* context, const ::Device::V3::DoRequest* request, ::Device::V3::DoResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::Device::V3::DoRequest, ::Device::V3::DoResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_Add_, context, request, response, std::move(f));
}

void Device::Stub::async::Add(::grpc::ClientContext* context, const ::Device::V3::DoRequest* request, ::Device::V3::DoResponse* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_Add_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::Device::V3::DoResponse>* Device::Stub::PrepareAsyncAddRaw(::grpc::ClientContext* context, const ::Device::V3::DoRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::Device::V3::DoResponse, ::Device::V3::DoRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_Add_, context, request);
}

::grpc::ClientAsyncResponseReader< ::Device::V3::DoResponse>* Device::Stub::AsyncAddRaw(::grpc::ClientContext* context, const ::Device::V3::DoRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncAddRaw(context, request, cq);
  result->StartCall();
  return result;
}

::grpc::Status Device::Stub::Del(::grpc::ClientContext* context, const ::Device::V3::DoRequest& request, ::Device::V3::DoResponse* response) {
  return ::grpc::internal::BlockingUnaryCall< ::Device::V3::DoRequest, ::Device::V3::DoResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_Del_, context, request, response);
}

void Device::Stub::async::Del(::grpc::ClientContext* context, const ::Device::V3::DoRequest* request, ::Device::V3::DoResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::Device::V3::DoRequest, ::Device::V3::DoResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_Del_, context, request, response, std::move(f));
}

void Device::Stub::async::Del(::grpc::ClientContext* context, const ::Device::V3::DoRequest* request, ::Device::V3::DoResponse* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_Del_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::Device::V3::DoResponse>* Device::Stub::PrepareAsyncDelRaw(::grpc::ClientContext* context, const ::Device::V3::DoRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::Device::V3::DoResponse, ::Device::V3::DoRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_Del_, context, request);
}

::grpc::ClientAsyncResponseReader< ::Device::V3::DoResponse>* Device::Stub::AsyncDelRaw(::grpc::ClientContext* context, const ::Device::V3::DoRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncDelRaw(context, request, cq);
  result->StartCall();
  return result;
}

::grpc::Status Device::Stub::Action(::grpc::ClientContext* context, const ::Device::V3::DoRequest& request, ::Device::V3::DoResponse* response) {
  return ::grpc::internal::BlockingUnaryCall< ::Device::V3::DoRequest, ::Device::V3::DoResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_Action_, context, request, response);
}

void Device::Stub::async::Action(::grpc::ClientContext* context, const ::Device::V3::DoRequest* request, ::Device::V3::DoResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::Device::V3::DoRequest, ::Device::V3::DoResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_Action_, context, request, response, std::move(f));
}

void Device::Stub::async::Action(::grpc::ClientContext* context, const ::Device::V3::DoRequest* request, ::Device::V3::DoResponse* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_Action_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::Device::V3::DoResponse>* Device::Stub::PrepareAsyncActionRaw(::grpc::ClientContext* context, const ::Device::V3::DoRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::Device::V3::DoResponse, ::Device::V3::DoRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_Action_, context, request);
}

::grpc::ClientAsyncResponseReader< ::Device::V3::DoResponse>* Device::Stub::AsyncActionRaw(::grpc::ClientContext* context, const ::Device::V3::DoRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncActionRaw(context, request, cq);
  result->StartCall();
  return result;
}

::grpc::Status Device::Stub::Message(::grpc::ClientContext* context, const ::Device::V3::DoRequest& request, ::Device::V3::DoResponse* response) {
  return ::grpc::internal::BlockingUnaryCall< ::Device::V3::DoRequest, ::Device::V3::DoResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_Message_, context, request, response);
}

void Device::Stub::async::Message(::grpc::ClientContext* context, const ::Device::V3::DoRequest* request, ::Device::V3::DoResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::Device::V3::DoRequest, ::Device::V3::DoResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_Message_, context, request, response, std::move(f));
}

void Device::Stub::async::Message(::grpc::ClientContext* context, const ::Device::V3::DoRequest* request, ::Device::V3::DoResponse* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_Message_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::Device::V3::DoResponse>* Device::Stub::PrepareAsyncMessageRaw(::grpc::ClientContext* context, const ::Device::V3::DoRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::Device::V3::DoResponse, ::Device::V3::DoRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_Message_, context, request);
}

::grpc::ClientAsyncResponseReader< ::Device::V3::DoResponse>* Device::Stub::AsyncMessageRaw(::grpc::ClientContext* context, const ::Device::V3::DoRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncMessageRaw(context, request, cq);
  result->StartCall();
  return result;
}

Device::Service::Service() {
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      Device_method_names[0],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< Device::Service, ::Device::V3::OpenRequest, ::Device::V3::OpenReply, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](Device::Service* service,
             ::grpc::ServerContext* ctx,
             const ::Device::V3::OpenRequest* req,
             ::Device::V3::OpenReply* resp) {
               return service->Open(ctx, req, resp);
             }, this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      Device_method_names[1],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< Device::Service, ::Device::V3::OpenRequest, ::Device::V3::OpenReply, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](Device::Service* service,
             ::grpc::ServerContext* ctx,
             const ::Device::V3::OpenRequest* req,
             ::Device::V3::OpenReply* resp) {
               return service->Close(ctx, req, resp);
             }, this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      Device_method_names[2],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< Device::Service, ::Device::V3::DoRequest, ::Device::V3::DoResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](Device::Service* service,
             ::grpc::ServerContext* ctx,
             const ::Device::V3::DoRequest* req,
             ::Device::V3::DoResponse* resp) {
               return service->Get(ctx, req, resp);
             }, this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      Device_method_names[3],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< Device::Service, ::Device::V3::DoRequest, ::Device::V3::DoResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](Device::Service* service,
             ::grpc::ServerContext* ctx,
             const ::Device::V3::DoRequest* req,
             ::Device::V3::DoResponse* resp) {
               return service->Set(ctx, req, resp);
             }, this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      Device_method_names[4],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< Device::Service, ::Device::V3::DoRequest, ::Device::V3::DoResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](Device::Service* service,
             ::grpc::ServerContext* ctx,
             const ::Device::V3::DoRequest* req,
             ::Device::V3::DoResponse* resp) {
               return service->Update(ctx, req, resp);
             }, this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      Device_method_names[5],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< Device::Service, ::Device::V3::DoRequest, ::Device::V3::DoResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](Device::Service* service,
             ::grpc::ServerContext* ctx,
             const ::Device::V3::DoRequest* req,
             ::Device::V3::DoResponse* resp) {
               return service->Add(ctx, req, resp);
             }, this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      Device_method_names[6],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< Device::Service, ::Device::V3::DoRequest, ::Device::V3::DoResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](Device::Service* service,
             ::grpc::ServerContext* ctx,
             const ::Device::V3::DoRequest* req,
             ::Device::V3::DoResponse* resp) {
               return service->Del(ctx, req, resp);
             }, this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      Device_method_names[7],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< Device::Service, ::Device::V3::DoRequest, ::Device::V3::DoResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](Device::Service* service,
             ::grpc::ServerContext* ctx,
             const ::Device::V3::DoRequest* req,
             ::Device::V3::DoResponse* resp) {
               return service->Action(ctx, req, resp);
             }, this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      Device_method_names[8],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< Device::Service, ::Device::V3::DoRequest, ::Device::V3::DoResponse, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](Device::Service* service,
             ::grpc::ServerContext* ctx,
             const ::Device::V3::DoRequest* req,
             ::Device::V3::DoResponse* resp) {
               return service->Message(ctx, req, resp);
             }, this)));
}

Device::Service::~Service() {
}

::grpc::Status Device::Service::Open(::grpc::ServerContext* context, const ::Device::V3::OpenRequest* request, ::Device::V3::OpenReply* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status Device::Service::Close(::grpc::ServerContext* context, const ::Device::V3::OpenRequest* request, ::Device::V3::OpenReply* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status Device::Service::Get(::grpc::ServerContext* context, const ::Device::V3::DoRequest* request, ::Device::V3::DoResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status Device::Service::Set(::grpc::ServerContext* context, const ::Device::V3::DoRequest* request, ::Device::V3::DoResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status Device::Service::Update(::grpc::ServerContext* context, const ::Device::V3::DoRequest* request, ::Device::V3::DoResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status Device::Service::Add(::grpc::ServerContext* context, const ::Device::V3::DoRequest* request, ::Device::V3::DoResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status Device::Service::Del(::grpc::ServerContext* context, const ::Device::V3::DoRequest* request, ::Device::V3::DoResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status Device::Service::Action(::grpc::ServerContext* context, const ::Device::V3::DoRequest* request, ::Device::V3::DoResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status Device::Service::Message(::grpc::ServerContext* context, const ::Device::V3::DoRequest* request, ::Device::V3::DoResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}


}  // namespace Device
}  // namespace V3

