package main

import (
	"bufio"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/nfnt/resize"
)


func main() {
	// Kullanıcıdan hedef boyut alınır
	fmt.Print("Yeni boyutu girin (örn: 32): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	targetSize, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil || targetSize <= 0 {
		fmt.Println("Geçersiz boyut.")
		return
	}

	// Çıktı dizini
	outputDir := "resized"
	os.MkdirAll(outputDir, os.ModePerm)

	// Geçerli dizindeki dosyaları tarar
	files, _ := os.ReadDir(".")

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filename := file.Name()
		ext := strings.ToLower(filepath.Ext(filename))

		if ext != ".png" && ext != ".jpg" && ext != ".jpeg" && ext != ".gif" {
			continue
		}

		// Dosyayı aç
		imgFile, err := os.Open(filename)
		if err != nil {
			fmt.Printf("Dosya açılamadı: %s\n", filename)
			continue
		}

		// Dosyayı çözümle
		var img image.Image
		switch ext {
		case ".png":
			img, err = png.Decode(imgFile)
		case ".jpg", ".jpeg":
			img, err = jpeg.Decode(imgFile)
		case ".gif":
			img, err = gif.Decode(imgFile)
		}
		imgFile.Close()
		if err != nil {
			fmt.Printf("Görüntü okunamadı: %s\n", filename)
			continue
		}

		// Görüntüyü yeniden boyutlandır
		resizedImg := resize.Resize(uint(targetSize), uint(targetSize), img, resize.Lanczos3)

		// Çıktı dosyası oluştur
		outPath := filepath.Join(outputDir, filename)
		outFile, err := os.Create(outPath)
		if err != nil {
			fmt.Printf("Çıktı oluşturulamadı: %s\n", outPath)
			continue
		}
		defer outFile.Close()

		// Yeni görüntüyü yaz
		switch ext {
		case ".png":
			png.Encode(outFile, resizedImg)
		case ".jpg", ".jpeg":
			jpeg.Encode(outFile, resizedImg, nil)
		case ".gif":
			gif.Encode(outFile, resizedImg, nil)
		}

		fmt.Printf("✓ %s -> %s\n", filename, outPath)
	}

	fmt.Println("İşlem tamamlandı.")
}