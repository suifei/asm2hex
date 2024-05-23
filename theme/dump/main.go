package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetImageFiles(dir, pattern string) []string {
	// implementation of GetImageFiles function goes here
	var files []string
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(strings.ToUpper(path), strings.ToUpper(pattern)) {
			files = append(files, path)
		}
		return nil
	})
	return files
}

func MakeGolangBytesString(file string) string {
	// implementation of MakeGolangBytesString function goes here
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	} else {
		defer f.Close()
		fi, _ := f.Stat()
		size := fi.Size()
		data := make([]byte, size)
		f.Read(data)

		res_name := fi.Name()
		buf_str := toByteList(data)
		return fmt.Sprintf(
			`var %s_RES = &fyne.StaticResource{ 
	StaticName: "%s",
	StaticContent: []byte(%s),
}`, strings.ToUpper(strings.ReplaceAll(res_name, ".", "_")), res_name, buf_str)
	}
	return ""
}

func toByteList(data []byte) string {
	// 16 bytes per line
	var buf strings.Builder
	buf.WriteString("\t\t")
	for i, b := range data {
		if i%16 == 0 {
			buf.WriteString("\n\t\t")
		}
		buf.WriteString(fmt.Sprintf("0x%02x,", b))
	}
	return strings.TrimSuffix(buf.String(), ",")
}

func main() {

	files := GetImageFiles("..", ".png")
	for _, file := range files {
		fmt.Println(MakeGolangBytesString(file))
	}
}
