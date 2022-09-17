package bzip_test

import (
	"bytes"
	"compress/bzip2"
	"gopl2022.io/ch13/bzip"
	"io"
	"testing"
)

func TestBzip2(t *testing.T) {
	var compressed, uncompressed bytes.Buffer
	w := bzip.NewWriter(&compressed)

	// Write a repetitive message in a million pieces,
	// 写一百万条重复的消息，
	// compressing one copy but not the other.
	// 压缩一个副本但不压缩另一个。
	tee := io.MultiWriter(w, &uncompressed)
	for i := 0; i < 1000000; i++ {
		io.WriteString(tee, "hellggo")
	}
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}

	// Check the size of the compressed stream.
	// 检查压缩流的大小。
	if got, want := compressed.Len(), 255; got != want {
		t.Errorf("1 million hellos compressed to %d bytes, want %d", got, want)
	}

	// Decompress and compare with original.
	// 解压并与原始文件进行比较。
	var decompressed bytes.Buffer
	io.Copy(&decompressed, bzip2.NewReader(&compressed))
	if !bytes.Equal(uncompressed.Bytes(), decompressed.Bytes()) {
		t.Error("decompression yielded a different message")

	}
}
