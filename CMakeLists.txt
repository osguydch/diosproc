# Copyright 2018 gRPC authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
# cmake build file for C++ helloworld example.
# Assumes protobuf and gRPC have been installed using cmake.
# See cmake_externalproject/CMakeLists.txt for all-in-one cmake build
# that automatically builds all the dependencies before building helloworld.

cmake_minimum_required(VERSION 3.8)

project(HelloWorld C CXX)

include(../cmake/common.cmake)

# Proto file
get_filename_component(hw_proto "../../protos/helloworld.proto" ABSOLUTE)
get_filename_component(dev_proto "../../protos/device-sm.proto" ABSOLUTE)
get_filename_component(hw_proto_path "${hw_proto}" PATH)
get_filename_component(dev_proto_path "${dev_proto}" PATH)


 
set(CMAKE_DEBUG_POSTFIX "D")
#set(CMAKE_CXX_FLAGS_RELEASE "/Zi -DNDEBUG")
#set(CMAKE_EXE_LINKER_FLAGS_RELEASE "${CMAKE_EXE_LINKER_FLAGS_RELEASE} /DEBUG")
#set(CMAKE_SHARED_LINKER_FLAGS_RELEASE "${CMAKE_SHARED_LINKER_FLAGS_RELEASE} /DEBUG")
 


# Generated sources
set(hw_proto_srcs "${CMAKE_CURRENT_BINARY_DIR}/helloworld.pb.cc")
set(hw_proto_hdrs "${CMAKE_CURRENT_BINARY_DIR}/helloworld.pb.h")
set(hw_grpc_srcs "${CMAKE_CURRENT_BINARY_DIR}/helloworld.grpc.pb.cc")
set(hw_grpc_hdrs "${CMAKE_CURRENT_BINARY_DIR}/helloworld.grpc.pb.h")

set(dev_proto_srcs "${CMAKE_CURRENT_BINARY_DIR}/device-sm.pb.cc")
set(dev_proto_hdrs "${CMAKE_CURRENT_BINARY_DIR}/device-sm.pb.h")
set(dev_grpc_srcs "${CMAKE_CURRENT_BINARY_DIR}/device-sm.grpc.pb.cc")
set(dev_grpc_hdrs "${CMAKE_CURRENT_BINARY_DIR}/device-sm.grpc.pb.h")


add_custom_command(
      OUTPUT "${hw_proto_srcs}" "${hw_proto_hdrs}" "${hw_grpc_srcs}" "${hw_grpc_hdrs}"
      COMMAND ${_PROTOBUF_PROTOC}
      ARGS --grpc_out "${CMAKE_CURRENT_BINARY_DIR}"
        --cpp_out "${CMAKE_CURRENT_BINARY_DIR}"
        -I "${hw_proto_path}"
        --plugin=protoc-gen-grpc="${_GRPC_CPP_PLUGIN_EXECUTABLE}"
        "${hw_proto}"
      DEPENDS "${hw_proto}")

add_custom_command(
      OUTPUT "${dev_proto_srcs}" "${dev_proto_hdrs}" "${dev_grpc_srcs}" "${dev_grpc_hdrs}"
      COMMAND ${_PROTOBUF_PROTOC}
      ARGS --grpc_out "${CMAKE_CURRENT_BINARY_DIR}"
        --cpp_out "${CMAKE_CURRENT_BINARY_DIR}"
        -I "${dev_proto_path}"
        --plugin=protoc-gen-grpc="${_GRPC_CPP_PLUGIN_EXECUTABLE}"
        "${dev_proto}"
      DEPENDS "${dev_proto}")

# Include generated *.pb.h files
include_directories("${CMAKE_CURRENT_BINARY_DIR}")

# hw_grpc_proto
add_library(hw_grpc_proto
  ${hw_grpc_srcs}
  ${hw_grpc_hdrs}
  ${hw_proto_srcs}
  ${hw_proto_hdrs})
target_link_libraries(hw_grpc_proto
  absl::check
  ${_REFLECTION}
  ${_GRPC_GRPCPP}
  ${_PROTOBUF_LIBPROTOBUF})

  add_library(dev_grpc_proto
  ${dev_grpc_srcs}
  ${dev_grpc_hdrs}
  ${dev_proto_srcs}
  ${dev_proto_hdrs})
target_link_libraries(dev_grpc_proto
  absl::check
  ${_REFLECTION}
  ${_GRPC_GRPCPP}
  ${_PROTOBUF_LIBPROTOBUF})


# Targets greeter_[async_](client|server)
foreach(_target
  greeter_client greeter_server
  greeter_callback_client greeter_callback_server
  greeter_async_client greeter_async_client2 greeter_async_server )
  add_executable(${_target} "${_target}.cc")
  target_link_libraries(${_target}
    hw_grpc_proto
    absl::check
    absl::flags
    absl::flags_parse
    absl::log
    ${_REFLECTION}
    ${_GRPC_GRPCPP}
    ${_PROTOBUF_LIBPROTOBUF})
endforeach()

add_executable(diosprocx64 "device_client.cc")
  target_link_libraries(diosprocx64
    dev_grpc_proto
    absl::check
    absl::flags
    absl::flags_parse
    absl::log
    ${_REFLECTION}
    ${_GRPC_GRPCPP}
    ${_PROTOBUF_LIBPROTOBUF})

add_library( devClientlib SHARED "device_client.cc")
target_link_libraries(devClientlib
    dev_grpc_proto
    absl::check
    absl::flags
    absl::flags_parse
    absl::log
    ${_REFLECTION}
    ${_GRPC_GRPCPP}
    ${_PROTOBUF_LIBPROTOBUF})
set_target_properties(devClientlib PROPERTIES VERSION 1.0 SOVERSION 1)



  message(STATUS "Dest dll ${CMAKE_CURRENT_BINARY_DIR}")


  IF(EXISTS ${CMAKE_CURRENT_BINARY_DIR}/devClientlib.dll)
    file(COPY ${CMAKE_CURRENT_BINARY_DIR}/devClientlib.dll DESTINATION "E:/1.V1.0.0.0/DeviceV3/PlatformModule/")
  ENDIF()

  IF(EXISTS ${CMAKE_CURRENT_BINARY_DIR}/devClientlib.pdb)
     file(COPY ${CMAKE_CURRENT_BINARY_DIR}/devClientlib.pdb DESTINATION "E:/1.V1.0.0.0/DeviceV3/PlatformModule/")
  ENDIF()
