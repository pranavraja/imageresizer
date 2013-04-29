package main

import (
	"errors"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
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

func resizedImageFromUrl(url string, resizedWidth int, resizedHeight int, algorithm string) (resizedImage image.Image, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	img, err := jpeg.Decode(resp.Body)
	if err != nil {
		return
	}
	if img == nil {
		return nil, errors.New("No image could be decoded from " + url)
	}
	resizedImage = resize.Resize(uint(resizedWidth), uint(resizedHeight), img, resizeAlgorithmFromString(algorithm))
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
	img, err := resizedImageFromUrl(source, width, height, r.FormValue("algorithm"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jpeg.Encode(w, img, nil)
}

func main() {
	http.HandleFunc("/", ResizeHandler)
	http.ListenAndServe(":8080", nil)
}
