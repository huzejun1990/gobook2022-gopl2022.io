// The jpeg command reads a PNG image from the standard input
// jpeg 命令从标准输入中读取 PNG 图像
// and writes it as a JPEG image to the standard output.
// 并将其作为 JPEG 图像写入标准输出。
package main

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png" // register PNG decoder	// 注册PNG解码器
	"io"
	"os"
)

func main() {
	if err := toJPEG(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

func toJPEG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}

//!-main

/*
//!+with
$ go build gopl2022.io/ch3/mandelbrot

$ go build gopl2022.io/ch10/jpeg

$ ./mandelbrot | ./jpeg >mandelbrot.jpg
Input format = png
//!-with

//!+without
$ go build gopl2022.io/ch3/mandelbrot
$ go build gopl2022.io/ch10/jpeg
$ ./mandelbrot | ./jpeg >mandelbrot.jpg
bash: ./jpeg: Permission denied

//!-without
*/
