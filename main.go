package main

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	_ "net/http/pprof"
	"strconv"
	"time"

	"github.com/nfnt/resize"
)

var timeoutClient *http.Client

func init() {
	timeoutClient = &http.Client{
		Transport: &http.Transport{
			ResponseHeaderTimeout: 5 * time.Second,
		},
	}
}

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
	resp, err := timeoutClient.Get(source)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()
	remoteContentType := resp.Header.Get("Content-Type")
	img, err := decodeImage(resp.Body, remoteContentType)
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
	w.Header().Set("Content-Type", remoteContentType)
	w.Header().Set("ETag", resp.Header.Get("ETag"))
	w.Header().Set("Cache-Control", resp.Header.Get("Cache-Control"))
	w.Header().Set("Expires", resp.Header.Get("Expires"))
	err = encodeImage(w, resizedImage, remoteContentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", ResizeHandler)
	http.ListenAndServe(":8080", nil)
}
