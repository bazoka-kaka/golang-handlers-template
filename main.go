package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// text/html template
var template = `
        <h1>Benzion</h1>
        <p>My Hobby is:</p>
        <ul>
          <li>Programming</li>
          <li>Gaming</li>
        </ul>
    `

// application/json data
var userData = map[string]interface{}{
	"first_name": "Benzion",
	"last_name":  "Yehezkel",
	"online":     false,
}

// text/plain
func handlerText(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(201)
	w.Write([]byte("Data created!"))
}

// text/html
func handlerHtml(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(template))
}

// application/json
func handlerJson(w http.ResponseWriter, r *http.Request) {
	jsonData, err := json.Marshal(userData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsonData))
}

// multipart/form-data
func handlerFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Upload File Endpooint Hit")

	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("image")
	if err != nil {
		fmt.Println("Error retrieving the file")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	tempFile.Write(fileBytes)

	fmt.Fprintln(w, "Success wrote file!")
}

func main() {
	http.HandleFunc("/", handlerText)
	http.HandleFunc("/index", handlerHtml)
	http.HandleFunc("/user", handlerJson)
	http.HandleFunc("/upload", handlerFile)

	fmt.Println("server running on port 3000")
	http.ListenAndServe(":3000", nil)
}
