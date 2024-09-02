from ultralytics import YOLO
import matplotlib.pyplot as plt
# Load a model
model = YOLO("./models/food_detect/train/weights/best.pt")
output_path = "./metrics/food_detect"
def main():
    metrics = model.val(plots=True,split="test",project=output_path)
    # Extract mAP values
    map50_95 = metrics.box.map
    map50 = metrics.box.map50
    map75 = metrics.box.map75
    maps = metrics.box.maps
    # Plot the mAP values
    categories = [f"{i}" for i in range(len(maps))]
    
    plt.figure(figsize=(30, 15))
    # Plot mAP50-95
    plt.plot(categories, maps, label='mAP50-95', marker='o')
    plt.axhline(map50_95, color='r', linestyle='--', label='mAP50-95 Mean')
    
    # Plot mAP50
    plt.axhline(map50, color='g', linestyle='--', label='mAP50 Mean')
    
    # Plot mP75
    plt.axhline(map75, color='b', linestyle='--', label='mAP75 Mean')
    
    plt.xlabel('Categories')
    plt.ylabel('mAP')
    plt.title('Mean Average Precision (mAP) for Each Category')
    plt.legend()
    plt.grid(True)  
    plt.xticks(rotation=45)
    plt.gca().margins(x=0.01)
    # Save and show the plot
    plt.tight_layout()
    plt.savefig(f"{output_path}/val/mAP_plot.png")
    plt.show()

if __name__ == "__main__":
    main()
