from os import path, system
import random
from PIL import Image

# Number of images to download
num_images = 100

# Image size (width and height)
width = 200
height = 300

# List of desired file formats
file_formats = ['jpg', 'png', 'gif']

# Base URL for the random image service
base_url = f"https://picsum.photos/{width}/{height}"

# Download the images
for i in range(num_images):
    # Choose a random file format from the list
    file_format = random.choice(file_formats)

    # Set the filename with the selected file format
    filename = path.join("imgs",f"image_{i + 1}.{file_format}")
    
    system(f"wget --output-document {filename} {base_url}")
    print(f"Downloaded image {i + 1}/{num_images}")

print("All images downloaded and converted successfully.")
