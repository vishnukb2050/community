
import sys
try:
    from PIL import Image
    img = Image.open('assets/icons/app_icon.png')
    img.save('assets/icons/app_icon_fixed.png', 'PNG')
    print("Converted successfully")
except Exception as e:
    print(f"Error: {e}")
