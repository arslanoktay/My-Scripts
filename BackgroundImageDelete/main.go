package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
)

func isNearWhite(r, g, b uint32, tolerance uint8) bool {
	return r>>8 >= 255-uint32(tolerance) &&
		g>>8 >= 255-uint32(tolerance) &&
		b>>8 >= 255-uint32(tolerance)
}

func main() {
	var filename string
	fmt.Print("Görsel dosya adını girin (örnek: PB.png): ")
	fmt.Scanln(&filename)

	inputPath := filepath.Join("public", filename)

	file, err := os.Open(inputPath)
	if err != nil {
		fmt.Println("Dosya açılamadı:", err)
		return
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Görsel çözümlenemedi:", err)
		return
	}

	bounds := img.Bounds()
	output := image.NewNRGBA(bounds)

	tolerance := uint8(30) // 0–255 arası: ne kadar açık renklere şeffaflık uygulanacak

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			if isNearWhite(r, g, b, tolerance) {
				output.Set(x, y, color.NRGBA{0, 0, 0, 0})
			} else {
				output.Set(x, y, color.NRGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)})
			}
		}
	}

	outFile, err := os.Create(filepath.Join("public", "output.png"))
	if err != nil {
		fmt.Println("Çıktı dosyası oluşturulamadı:", err)
		return
	}
	defer outFile.Close()

	err = png.Encode(outFile, output)
	if err != nil {
		fmt.Println("PNG kaydedilemedi:", err)
		return
	}

	fmt.Println("Arka plan kaldırıldı: public/output.png")
}
