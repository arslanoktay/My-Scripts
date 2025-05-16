import os
from PIL import Image

def resize_and_center_icon(input_path, output_path, icon_size, canvas_size):
    img = Image.open(input_path).convert("RGBA")
    img_resized = img.resize((icon_size, icon_size), Image.Resampling.LANCZOS)

    canvas = Image.new("RGBA", (canvas_size, canvas_size), (0, 0, 0, 0))
    x = (canvas_size - icon_size) // 2
    y = (canvas_size - icon_size) // 2
    canvas.paste(img_resized, (x, y), img_resized)
    canvas.save(output_path, "PNG")

def main():
    try:
        canvas_size = int(input("Final PNG boyutunu gir (örnek: 32): "))
    except ValueError:
        print("Lütfen geçerli bir sayı girin.")
        return

    icon_size = canvas_size - 6
    input_folder = os.getcwd()
    output_folder = os.path.join(input_folder, "converted")

    os.makedirs(output_folder, exist_ok=True)

    png_files = [f for f in os.listdir(input_folder) if f.lower().endswith(".png")]

    if not png_files:
        print("Bu klasörde hiç PNG dosyası yok.")
        return

    print(f"{len(png_files)} adet PNG işleniyor...")

    for filename in png_files:
        input_path = os.path.join(input_folder, filename)
        output_path = os.path.join(output_folder, filename)
        resize_and_center_icon(input_path, output_path, icon_size, canvas_size)

    print(f"Tüm ikonlar başarıyla '{output_folder}' klasörüne kaydedildi.")

if __name__ == "__main__":
    main()
