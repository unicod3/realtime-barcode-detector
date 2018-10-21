package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"

	"github.com/bieber/barcode"
	"gocv.io/x/gocv"
)

func main() {
	var deviceID int
	flag.IntVar(&deviceID, "device-id", 0, "integer value, webcam device ID")
	flag.Parse()

	webcam, err := gocv.VideoCaptureDevice(deviceID)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer webcam.Close()

	window := gocv.NewWindow("Barcode detector")
	defer window.Close()

	img := gocv.NewMat()
	defer img.Close()

	textColor := color.RGBA{255, 0, 0, 0} // red
	dotColor := color.RGBA{0, 255, 0, 0}  // green
	fmt.Printf("start reading camera device: %v\n", deviceID)
	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("cannot read device %d\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}
		webcam.Read(&img)

		// read barcode with zbar from the frame
		scanner := barcode.NewScanner().
			SetEnabledAll(true)

		imgObj, _ := img.ToImage()

		src := barcode.NewImage(imgObj)
		symbols, _ := scanner.ScanImage(src)

		for _, s := range symbols {
			data := s.Data
			points := s.Boundary // Data points that zbar returns

			x0 := points[0].X
			y0 := points[0].Y

			size := gocv.GetTextSize(data, gocv.FontHersheyPlain, 1.2, 2)
			pt := image.Pt(x0-size.X, y0-size.Y)
			gocv.PutText(&img, data, pt, gocv.FontHersheyPlain, 1.2, textColor, 2)

			for _, p := range points {
				x0 := p.X
				y0 := p.Y
				pt := image.Pt(x0, y0)
				gocv.PutText(&img, ".", pt, gocv.FontHersheyPlain, 1.2, dotColor, 2)
			}
		}

		window.IMShow(img)
		window.WaitKey(1)
	}
}
