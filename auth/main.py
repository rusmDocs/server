from concurrent import futures

import grpc

from auth.dto import auth_pb2_grpc
from auth.services.tokens import JWT


def serve():
    print("i'm alive")
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10)) 
    auth_pb2_grpc.add_AuthServiceServicer_to_server(JWT(), server) 
    server.add_insecure_port('[::]:50051') 
    server.start() 
    server.wait_for_termination() 


if __name__ == '__main__':
    serve()
