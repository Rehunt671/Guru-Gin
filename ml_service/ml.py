import sys
sys.path.insert(1, 'C:/Users/Kingsglaive/Desktop/Work/Advance-CPE/GuruGin/ml_service/services')
import io
import os
import py_hot_reload
from concurrent import futures
import grpc
from PIL import Image  
from ultralytics import YOLO 
import ml_pb2_grpc
import ml_pb2
import torch
torch.cuda.is_available()

class MLService(ml_pb2_grpc.MLServiceServicer):
    def __init__(self):
        # Initialize and train the YOLOv9 model
        self.model = self.initialize_model()

    def initialize_model(self):
        # Build a YOLOv9c model from pretrained weight
        model = YOLO("./runs/detect/train/weights/best.pt")
        # Display model information (optional) 
        model.info() 
        return model
    
    def train_model(self):
        # Build a YOLOv9c model from pretrained weight
        model = YOLO("./yolov9c.pt")
        # Display model information (optional) 
        model.info() 
        # Train the model on the roboflow example dataset for 100 epochs
        model.train(
            data="./datasets/data.yaml",
            epochs=100,
            imgsz=512,
            patience=10,              # Early stopping patience (epochs)
        )
        return model

    def process_image(self, imagePath: str):
        classifications = set();
        results = self.model.predict(
            imagePath, 
            imgsz=(640, 640), 
        )

        for result in results[0]:
            boxes = result.boxes  # Boxes object for bbox outputs
            for box in boxes:
                class_of_object = results[0].names.get(box.cls.item())
                confidence = box.conf.item()  # confidence scores
                if confidence < 0.6:
                    continue
                classifications.add(class_of_object)
        return classifications 

    
    def DetectObjects(self, request, context):
        image = Image.open(io.BytesIO(request.image))
        image.save("image.jpg") 
        classifications = self.process_image("image.jpg")
        os.remove("image.jpg")
        return ml_pb2.ImageResponse(classifications=classifications)

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    ml_pb2_grpc.add_MLServiceServicer_to_server(MLService(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    print("server is running on port 50051")
    server.wait_for_termination()

if __name__ == '__main__':
    py_hot_reload.run_with_reloader(serve)
