package router

import (
	img_processing "image_processing/internal/image_process"
	"net/http"
)

func Router(mux *http.ServeMux) {
	mux.HandleFunc("POST /files/store", img_processing.ProcessingImages)

}
