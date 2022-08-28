// The thumbnail package produces thumbnail-size images from
// thumbnail 包生成缩略图大小的图像
// larger images.  Only JPEG images are currently supported.
// 更大的图像。 当前仅支持 JPEG 图像。
package thumbnail

import (
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Image returns a thumbnail-size version of src.
// 图片返回 src 的缩略图大小版本。
func Image(src image.Image) image.Image {
	// Compute thumbnail size, preserving aspect ratio.
	// 计算缩略图大小，保留纵横比。
	xs := src.Bounds().Size().X
	ys := src.Bounds().Size().Y
	width, height := 128, 128
	if aspect := float64(xs) / float64(ys); aspect < 1.0 {
		width = int(128 * aspect) //portrait	//图像
	} else {
		height = int(128 / aspect) // landscape	// 景观
	}
	xscale := float64(xs) / float64(width)
	yscale := float64(ys) / float64(height)

	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	// a very crude scaling algorithm
	// 一个非常粗糙的缩放算法
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			srcx := int(float64(x) * xscale)
			srcy := int(float64(y) * yscale)
			dst.Set(x, y, src.At(srcx, srcy))
		}
	}
	return dst
}

// ImageStream reads an image from r and
// ImageStream 从 r 中读取图像并
// writes a thumbnail-size version of it to w.
// 将其缩略图大小的版本写入 w。
func ImageStream(w io.Writer, r io.Reader) error {
	src, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	dst := Image(src)
	return jpeg.Encode(w, dst, nil)
}

// ImageFile2 reads an image from infile and writes // ImageFile2 从 infile 中读取图像并写入
// a thumbnail-size version of it to outfile.// 要输出的缩略图大小版本。
func ImageFile2(outfile, infile string) (err error) {
	in, err := os.Open(infile)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(outfile)
	if err != nil {
		return err
	}

	if err := ImageStream(out, in); err != nil {
		out.Close()
		return fmt.Errorf("scaling %s to %s: %s", infile, outfile, err)
	}
	return out.Close()
}

// ImageFile reads an image from infile and writes
//ImageFile 从 infile 中读取图像并写入
// a thumbnail-size version of it in the same directory.
//在同一目录中的缩略图大小版本。
// It returns the generated file name, e.g. "foo.thumb.jpeg".
//它返回生成的文件名，例如 “foo.thumb.jpeg”。
func ImageFile(infile string) (string, error) {
	ext := filepath.Ext(infile) // e.g., ".jpg", ".JPEG" //例如：jpg、JPEG
	outfile := strings.TrimSuffix(infile, ext) + ".thumb" + ext
	return outfile, ImageFile2(outfile, infile)
}
