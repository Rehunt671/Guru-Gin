import sys
sys.path.append('services')
from concurrent import futures
import os
import grpc
import torch
from services import ml_pb2_grpc
from services.ml_server import MLService
from dotenv import load_dotenv

# Load environment variables from .env file
load_dotenv('../.env')
ML_PORT = os.getenv('ML_PORT')

if torch.cuda.is_available():
    print("CUDA is available. PyTorch can use the GPU.")
else:
    print("CUDA is not available. PyTorch will use the CPU.")

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    ml_pb2_grpc.add_MLServiceServicer_to_server(MLService(), server)
    server.add_insecure_port(f'[::]:{ML_PORT}')
    server.start()
    print(f"Server started on port {ML_PORT}.")
    server.wait_for_termination()

if __name__ == '__main__':
    serve()
