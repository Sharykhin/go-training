package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file2")

	if err != nil {
		fmt.Println(w, "Error with reading form name")
		fmt.Fprintln(w, err)
		return
	}

	defer file.Close()

	buff := make([]byte, 512)

	_, err = file.Read(buff)

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	filetype := http.DetectContentType(buff)
	fmt.Println(filetype)
	switch filetype {
	case "image/jpeg", "image/jpg":
		fmt.Println(filetype)
	case "image/gif":
		fmt.Println(filetype)
	case "image/png":
		fmt.Println(filetype)
	default:
		fmt.Println("unknown filte type uploaded")
		fmt.Fprintf(w, "unknown filte type uploaded")
		return
	}

	out, err := os.Create("tmp/" + header.Filename)
	if err != nil {
		fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privileges")
		return
	}

	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	fmt.Fprintf(w, "File uploaded successfully : ")
	fmt.Fprintf(w, header.Filename)

}

func main() {
	http.HandleFunc("/upload", uploadHandler)
	http.ListenAndServe(":9000", nil)
}
