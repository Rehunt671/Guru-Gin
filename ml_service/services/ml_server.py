from PIL import Image  
from ultralytics import YOLO 
import ml_pb2_grpc
import ml_pb2
import io
import os
import uuid

class MLService(ml_pb2_grpc.MLServiceServicer):
    def __init__(self):
        # Initialize model
        self.model = self.initialize_model()

    def initialize_model(self):
        model = YOLO("./runs/detect/food_detects/train/weights/best.pt")
        model.info() 
        return model
    
    def process(self, source: str):
        classifications = set()
        results = self.model.predict(
            source,      # source 
            conf=0.6,   # กรองเอาเฉพาะ confidence > 60%
            save=True,  # save รูปที่ detect ได้ไว้
        )
        for result in results[0]:
            boxes = result.boxes  # Boxes object for bbox outputs
            for box in boxes:
                class_of_object = results[0].names.get(box.cls.item())
                confidence = box.conf.item()  # confidence scores
                classifications.add(class_of_object)

        return classifications 
    
    def DetectObjects(self, request_iterator, context):
        classifications = set()
        for image_request in request_iterator:
            ext = image_request.info.image_type
            unique_filename = f"{uuid.uuid4()}{ext}"
            image_data = image_request.data
            image = Image.open(io.BytesIO(image_data))
            image.save(unique_filename)
            result = self.process(image)
            classifications.update(result)
            os.remove(unique_filename)
        
        for classification in classifications:
            yield ml_pb2.ImageResponse(classification=classification)
