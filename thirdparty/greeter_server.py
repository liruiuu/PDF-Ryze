# Copyright 2015 gRPC authors.
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
"""The Python implementation of the GRPC helloworld.Greeter server."""
from pdf import main_pdf
from greeter_client import run
from concurrent import futures
import logging

import grpc
import helloworld_pb2
import helloworld_pb2_grpc

connect_times = 1


class Greeter(helloworld_pb2_grpc.GreeterServicer):

    def __init__(self):
        self.server = None  # 添加server属性
        self.connect_times = 1

    def set_server(self, server):
        self.server = server  # 在服务端启动时设置server引用

    def set_connect_times(self, connect_times):
        connect_times = connect_times  # 在服务端启动时设置server引用

    def SayHello(self, request, context):
        print("接收到信息：", request.myStrings)
        print("self.connect_times11=", self.connect_times)
        if request.myStrings[0] == "closeServer":

            self.connect_times -= 1
            print("self.connect_times22=", self.connect_times)
            if self.connect_times <= 0:
                self.server.stop(grace=200)
                print("已经关闭")
            return helloworld_pb2.HelloReply(message="server closed")
        elif request.myStrings[0] == "openServer":
            self.connect_times += 1
            print("self.connect_times22=", self.connect_times)
            return helloworld_pb2.HelloReply(message="server opened")
        else:
            main_pdf(request.myStrings)
            # return helloworld_pb2.HelloReply(message="Hello, %s!" % request.myStrings)
            return helloworld_pb2.HelloReply(message="Done")


def serve():
    port = "50051"
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    greeter = Greeter()
    greeter.set_server(server)  # 设置Greeter实例的server引用
    helloworld_pb2_grpc.add_GreeterServicer_to_server(greeter, server)
    try:
        server.add_insecure_port("[::]:" + port)
        server.start()
        print("Server started, listening on " + port)
        server.wait_for_termination()
    # except OSError as e:
    except:
        run()
        # print(f"Port is already in use. Error: {e}")


if __name__ == "__main__":
    logging.basicConfig()
    serve()
