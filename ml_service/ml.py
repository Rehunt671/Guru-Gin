import io
from concurrent import futures
import grpc
from PIL import Image
from ultralytics import YOLO 
import sys
import torch
sys.path.insert(1, 'C:/Users/Kingsglaive/Desktop/Work/Advance-CPE/GuruGin/ml_service/services')
import ml_pb2_grpc
import ml_pb2

device: str = "cuda" if torch.cuda.is_available() else "cpu"

class MLService(ml_pb2_grpc.MLServiceServicer):
    def __init__(self):
        # Initialize and train the YOLOv9 model
        self.model = self.train_model()

    def initialize_model(self):
        # Build a YOLOv9c model from pretrained weight
        model = YOLO("./runs/detect/train/weights/best.pt")
        # Display model information (optional) 
        model.info() 
        source1 = "./datasets/test/images/-1_jpg.rf.22c302fa8411d22ff14e27838c9cbc02.jpg"
        source2 = "./datasets/test/images/pm-3-1678776723_jpg.rf.2cfcaeb488399b3c55e15deec003c5d7.jpg"
        model.predict([source1, source2], save=True, imgsz=(640, 640), conf=0.4, augment=True, nms=True, iou=0.5) 
        return model
    
    def train_model(self):
        # Build a YOLOv9c model from pretrained weight
        model = YOLO("./yolov9c.pt")
        # Display model information (optional) 
        model.info() 
        # Train the model on the roboflow example dataset for 100 epochs
        model.train(
            data="./datasets/data.yaml",
            epochs=300,
            imgsz=512,
            batch=16,                 # Batch size
            patience=10,              # Early stopping patience (epochs)
        )
        return model

    def process_image(self, image):
        results = self.model(image)
        labels = results.names
        classifications = [labels[int(x)] for x in results.xyxy[0][:, -1]]
        return classifications
    
    def DetectObjects(self, request, context):
        image = Image.open(io.BytesIO(request.image))
        classifications = self.process_image(image)
        return ml_pb2.IngredientResponse(classifications=classifications)

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    ml_pb2_grpc.add_MLServiceServicer_to_server(MLService(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    serve()
