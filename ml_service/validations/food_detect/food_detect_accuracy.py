from ultralytics import YOLO

# Load a model
model = YOLO("./models/food_detect/train/weights/best.pt")
output_path = "./metrics/food_detect"
def main():
    metrics = model.val(plots=True,split="test",project=output_path)
    metrics.box.map    # map50-95
    metrics.box.map50  # map50
    metrics.box.map75  # map75
    metrics.box.maps   # a list contains map50-95 of each categor

if __name__ == "__main__":
    main()
