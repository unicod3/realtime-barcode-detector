# Real-Time Barcode Detector

A Gocv and Zbar example for detecting and reading barcodes on webcam.

## Running the code

First you need to install GoCV library, which depends on opencv

GoCv installation instructions is well documented in their website:

[Linux](https://gocv.io/getting-started/linux/)

[MacOs](https://gocv.io/getting-started/macos/)

[Windows](https://gocv.io/getting-started/windows/)


After you installed the GoCv, you need to install zbar and get the Go wrapper for zbar as shown below:

Installing zbar can be tricky you migth want to check out below:

[Linux](http://zbar.sourceforge.net/download.html)

[MacOs](http://macappstore.org/zbar/)

[Windows](http://zbar.sourceforge.net/download.html)

```bash
go get github.com/bieber/barcode
```

Then you can clone this repo and run it
```bash
git clone https://github.com/unicod3/realtime-barcode-detector.git rbd
cd rbd
go run main.go

# or to specify the device id:
go run main.go --device-id 1234
```

Then you can show your code128 barcodes to your cam.

![Logo](http://sinanulker.com/github/realtime-barcode-detector.gif)

### Note:
It will open your default webcam since the device id set to zero.

