package main

import (
	"errors"
	"github.com/nfnt/resize"
	"image/jpeg"
	"io"
	"net/http"
	_ "net/http/pprof"
	"strconv"
)

func resizeAlgorithmFromString(algorithm string) resize.InterpolationFunction {
	switch algorithm {
	case "nearestNeighbour":
		return resize.NearestNeighbor
	case "bilinear":
		return resize.Bilinear
	case "bicubic":
		return resize.Bicubic
	case "mitchellNetravali":
		return resize.MitchellNetravali
	case "lanczos2":
		return resize.Lanczos2
	case "lanczos3":
		return resize.Lanczos3
	default:
		return resize.NearestNeighbor
	}
	panic("Control should never reach here")
}

func resizeAndPipe(destination io.Writer, source io.ReadCloser, resizedWidth int, resizedHeight int, algorithm string) error {
	defer source.Close()
	img, err := jpeg.Decode(source)
	if err != nil {
		return err
	}
	if img == nil {
		return errors.New("Source image couldn't be decoded")
	}
	resizedImage := resize.Resize(uint(resizedWidth), uint(resizedHeight), img, resizeAlgorithmFromString(algorithm))
	return jpeg.Encode(destination, resizedImage, nil)
}

func ResizeHandler(w http.ResponseWriter, r *http.Request) {
	source := r.FormValue("source")
	if source == "" {
		http.Error(w, "No source URL provided", http.StatusBadRequest)
		return
	}
	width, _ := strconv.Atoi(r.FormValue("width"))
	height, _ := strconv.Atoi(r.FormValue("height"))
	resp, err := http.Get(source)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = resizeAndPipe(w, resp.Body, width, height, r.FormValue("algorithm"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", ResizeHandler)
	http.ListenAndServe(":8080", nil)
}
