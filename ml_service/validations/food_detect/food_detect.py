from ultralytics import YOLO
import matplotlib.pyplot as plt
names = ['Akabare Khursani', 'Apple', 'Artichoke', 'Ash Gourd -Kubhindo-', 'Asparagus -Kurilo-', 'Avocado', 'Bacon', 'Bamboo Shoots -Tama-', 'Banana', 'Beans', 'Beaten Rice -Chiura-', 'Beef', 'Beetroot', 'Bethu ko Saag', 'Bitter Gourd', 'Black Lentils', 'Black beans', 'Bottle Gourd -Lauka-', 'Bread', 'Brinjal', 'Broad Beans -Bakullo-', 'Broccoli', 'Buff Meat', 'Butter', 'Cabbage', 'Capsicum', 'Carrot', 'Cassava -Ghar Tarul-', 'Cauliflower', 'Chayote-iskus-', 'Cheese', 'Chicken', 'Chicken Gizzards', 'Chickpeas', 'Chili Pepper -Khursani-', 'Chili Powder', 'Chowmein Noodles', 'Cinnamon', 'Coriander -Dhaniya-', 'Corn', 'Cornflakec', 'Crab Meat', 'Cucumber', 'Egg', 'Farsi ko Munta', 'Fiddlehead Ferns -Niguro-', 'Fish', 'Garden Peas', 'Garden cress-Chamsur ko saag-', 'Garlic', 'Ginger', 'Green Brinjal', 'Green Lentils', 'Green Mint -Pudina-', 'Green Peas', 'Green Soyabean -Hariyo Bhatmas-', 'Gundruk', 'Ham', 'Ice', 'Jack Fruit', 'Ketchup', 'Lapsi -Nepali Hog Plum-', 'Lemon -Nimbu-', 'Lime -Kagati-', 'Long Beans -Bodi-', 'Masyaura', 'Milk', 'Minced Meat', 'Moringa Leaves -Sajyun ko Munta-', 'Mushroom', 'Mutton', 'Nutrela -Soya Chunks-', 'Okra -Bhindi-', 'Olive Oil', 'Onion', 'Onion Leaves', 'Orange', 'Palak -Indian Spinach-', 'Palungo -Nepali Spinach-', 'Paneer', 'Papaya', 'Pea', 'Pear', 'Pointed Gourd -Chuche Karela-', 'Pork', 'Potato', 'Pumpkin -Farsi-', 'Radish', 'Rahar ko Daal', 'Rayo ko Saag', 'Red Beans', 'Red Lentils', 'Rice -Chamal-', 'Sajjyun -Moringa Drumsticks-', 'Salt', 'Sausage', 'Snake Gourd -Chichindo-', 'Soy Sauce', 'Soyabean -Bhatmas-', 'Sponge Gourd -Ghiraula-', 'Stinging Nettle -Sisnu-', 'Strawberry', 'Sugar', 'Sweet Potato -Suthuni-', 'Taro Leaves -Karkalo-', 'Taro Root-Pidalu-', 'Thukpa Noodles', 'Tofu', 'Tomato', 'Tori ko Saag', 'Tree Tomato -Rukh Tamatar-', 'Turnip', 'Wallnut', 'Water Melon', 'Wheat', 'Yellow Lentils', 'kimchi', 'mayonnaise', 'noodle', 'seaweed']
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
    
    plt.figure(figsize=(18, 15))
    
    # Plot mAP50-95
    plt.plot(categories, maps, label='mAP50-95', marker='o')
    plt.axhline(map50_95, color='r', linestyle='--', label='mAP50-95 Mean')
    
    # Plot mAP50
    plt.axhline(map50, color='g', linestyle='--', label='mAP50 Mean')
    
    # Plot mAP75
    plt.axhline(map75, color='b', linestyle='--', label='mAP75 Mean')
    
    plt.xlabel('Categories')
    plt.ylabel('mAP')
    plt.title('Mean Average Precision (mAP) for Each Category')
    plt.legend()
    plt.grid(True)
    plt.xticks(rotation=45)
    
    # Save and show the plot
    plt.tight_layout()
    plt.savefig(f"{output_path}/mAP_plot.png")
    plt.show()

if __name__ == "__main__":
    main()
