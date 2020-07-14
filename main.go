package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/gorilla/mux"
)

type UploadResponse struct {
	Status     string `json:"status"`
	UploadName string `json:"uploadName"`
	UploadSize int64  `json:"uploadSize"`
	URL        string `json:"url"`
}

func main() {
	// configure the songs directory name and port
	const mediaDir = "media"
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := mux.NewRouter()
	HomeFolder := "media"
	router.Handle("/media/{id}", streamHandler(http.StripPrefix("/media/", http.FileServer(http.Dir(HomeFolder)))))
	router.HandleFunc("/upload", uploadHandler).Methods("POST")
	http.Handle("/", router)
	fmt.Printf("Starting server on %v\n", port)
	// serve and log errors
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("error Retrieving file")
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	tempFile, err := ioutil.TempFile("media", "upload-*.mp4")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	tempFile.Write(fileBytes)

	inputFilePath := tempFile.Name()
	outputFilePath := TrimSuffix(inputFilePath, ".mp4")
	outputFilePath = fmt.Sprintf("%s.m3u8", outputFilePath)
	cmd := transcodeVideo(inputFilePath, outputFilePath)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fmt.Println(string(stdout))
	defer os.Remove(tempFile.Name())

	baseURL := "localhost:8080/"
	baseURL = fmt.Sprintf("%s%s", baseURL, outputFilePath)
	// fmt.Fprintf(w, "Successfully Uploaded File!\n")
	w.Header().Set("Content-Type", "application/json")
	jsonResponse := UploadResponse{
		Status:     "Success",
		UploadName: handler.Filename,
		UploadSize: handler.Size / 1024,
		URL:        string(baseURL),
	}

	json.NewEncoder(w).Encode(jsonResponse)

}

// addHeaders will act as middleware to give us CORS support
func streamHandler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	}
}

func transcodeVideo(inputFilePath, outFilePath string) *exec.Cmd {
	return exec.Command("ffmpeg",
		"-i", inputFilePath,
		"-profile:v", "baseline",
		"-level", "3.0",
		"-s", "640x360",
		"-start_number", "0",
		"-hls_time", "10",
		"-hls_list_size", "0",
		"-f", "hls",
		outFilePath,
	)
}

func TrimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}
