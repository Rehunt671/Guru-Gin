from ultralytics import YOLO
import os

def main():
    # Build a YOLOv9c model from pretrained weight
    model = YOLO("./yolov9c.pt")
    # Display model information 
    model.info() 
    # Specify the output path
    output_path = "./models/food_detect"
    # Train the model on the roboflow example dataset for 50 epochs
    try:
        model.train(
            data="./datasets/data.yaml",
            epochs=50,
            imgsz=512,
            patience=10,
            project=output_path
        )
    except Exception as e:
        print(f"An error occurred: {e}")
    finally:
        if os.path.exists("yolov8n.pt"):
            os.remove("yolov8n.pt")
        if os.path.exists("yolov9c.pt"):
            os.remove("yolov9c.pt")

if __name__ == "__main__":
    main()
