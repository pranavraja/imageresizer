package main

import (
	"errors"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
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

func decodeImage(source io.Reader, contentType string) (decoded image.Image, err error) {
	switch contentType {
	case "image/jpeg":
		decoded, err = jpeg.Decode(source)
	case "image/png":
		decoded, err = png.Decode(source)
	default:
		err = errors.New("Unsupported content type: " + contentType)
	}
	return
}

func encodeImage(destination io.Writer, img image.Image, contentType string) (err error) {
	switch contentType {
	case "image/jpeg":
		return jpeg.Encode(destination, img, nil)
	case "image/png":
		return png.Encode(destination, img)
	default:
		err = errors.New("Unsupported content type: " + contentType)
	}
	return
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
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()
	img, err := decodeImage(resp.Body, resp.Header.Get("Content-Type"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if img == nil {
		http.Error(w, "Could not decode image from URL "+source, http.StatusInternalServerError)
		return
	}
	algorithm := r.FormValue("algorithm")
	resizedImage := resize.Resize(uint(width), uint(height), img, resizeAlgorithmFromString(algorithm))
	err = encodeImage(w, resizedImage, resp.Header.Get("Content-Type"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", ResizeHandler)
	http.ListenAndServe(":8080", nil)
}
