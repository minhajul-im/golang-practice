package img_processing

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"

	"github.com/nfnt/resize"
)

func ProcessingImages(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(20 << 32)
	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, []string{"To large file!"})
		return
	}

	files, ok := r.MultipartForm.File["files"]
	if !ok || len(files) == 0 {
		sendErrorResponse(w, http.StatusBadRequest, []string{"No files uploaded"})
		return
	}
	os.MkdirAll("./uploads/thumb", 0755)
	os.MkdirAll("./uploads/favicon", 0755)
	os.MkdirAll("./uploads/download", 0755)
	os.MkdirAll("./uploads/original", 0755)

	var failedImgList []FailedImg
	var processedFiles []string

	for _, fileHeader := range files {
		contentType := fileHeader.Header.Get("Content-Type")
		if contentType == "image/png" {
			file, err := fileHeader.Open()
			if err != nil {
				img := FailedImg{
					Name:   fileHeader.Filename,
					Reason: err.Error(),
				}
				failedImgList = append(failedImgList, img)
				continue
			}
			defer file.Close()
			data, err := io.ReadAll(file)
			if err != nil {
				img := FailedImg{
					Name:   fileHeader.Filename,
					Reason: err.Error(),
				}
				failedImgList = append(failedImgList, img)
				continue
			}
			img, _, err := image.Decode(bytes.NewReader(data))
			if err != nil {
				img := FailedImg{
					Name:   fileHeader.Filename,
					Reason: err.Error(),
				}
				failedImgList = append(failedImgList, img)
				continue
			}
			// Resize for thumbnail, favicon, download
			thumb := resize.Resize(600, 300, img, resize.Lanczos2)
			favIcon := resize.Resize(32, 32, img, resize.Lanczos2)
			download := resize.Resize(800, 800, img, resize.Lanczos2)
			original := img

			// Save thumbnail with compression

			thumbFile, err := os.Create("./uploads/thumb/" + fileHeader.Filename)
			if err != nil {
				img := FailedImg{
					Name:   fileHeader.Filename,
					Reason: err.Error(),
				}
				failedImgList = append(failedImgList, img)
				continue
			}
			defer thumbFile.Close()
			encoder := png.Encoder{CompressionLevel: png.BestCompression}
			err = encoder.Encode(thumbFile, thumb)
			if err != nil {
				img := FailedImg{
					Name:   fileHeader.Filename,
					Reason: err.Error(),
				}
				failedImgList = append(failedImgList, img)
				continue
			}
			// Save favicon with compression
			favIconFile, err := os.Create("./uploads/favicon/" + fileHeader.Filename)
			if err != nil {
				img := FailedImg{
					Name:   fileHeader.Filename,
					Reason: "Failed to create favicon file: " + err.Error(),
				}
				failedImgList = append(failedImgList, img)
				continue
			}
			defer favIconFile.Close()
			encoder = png.Encoder{CompressionLevel: png.BestCompression}
			err = encoder.Encode(favIconFile, favIcon)
			if err != nil {
				img := FailedImg{
					Name:   fileHeader.Filename,
					Reason: "Failed to encode favicon: " + err.Error(),
				}
				failedImgList = append(failedImgList, img)
				continue
			}
			// Save download with compression
			downloadFile, err := os.Create("./uploads/download/" + fileHeader.Filename)
			if err != nil {
				img := FailedImg{
					Name:   fileHeader.Filename,
					Reason: "Failed to create download file: " + err.Error(),
				}
				failedImgList = append(failedImgList, img)
				continue
			}
			defer downloadFile.Close()
			encoder = png.Encoder{CompressionLevel: png.DefaultCompression}
			err = encoder.Encode(downloadFile, download)
			if err != nil {
				img := FailedImg{
					Name:   fileHeader.Filename,
					Reason: "Failed to encode download: " + err.Error(),
				}
				failedImgList = append(failedImgList, img)
				continue
			}
			// Save original with compression (no resizing)
			originalFile, err := os.Create("./uploads/original/" + fileHeader.Filename)
			if err != nil {
				img := FailedImg{
					Name:   fileHeader.Filename,
					Reason: "Failed to create original file: " + err.Error(),
				}
				failedImgList = append(failedImgList, img)
				continue
			}
			defer originalFile.Close()
			encoder = png.Encoder{CompressionLevel: png.DefaultCompression}
			err = encoder.Encode(originalFile, original)
			if err != nil {
				img := FailedImg{
					Name:   fileHeader.Filename,
					Reason: err.Error(),
				}
				failedImgList = append(failedImgList, img)
				continue
			}

			processedFiles = append(processedFiles,
				"thumb/"+fileHeader.Filename,
				"favicon/"+fileHeader.Filename,
				"download/"+fileHeader.Filename,
				"original/"+fileHeader.Filename)

		} else if contentType == "image/jpeg" {
			file, err := fileHeader.Open()
			if err != nil {
				img := FailedImg{
					Name:   fileHeader.Filename,
					Reason: err.Error(),
				}
				failedImgList = append(failedImgList, img)
				continue
			}
			defer file.Close()
			data, err := io.ReadAll(file)
			if err != nil {
				img := FailedImg{
					Name:   fileHeader.Filename,
					Reason: err.Error(),
				}
				failedImgList = append(failedImgList, img)
				continue
			}
			img, _, err := image.Decode(bytes.NewReader(data))
			if err != nil {
				img := FailedImg{
					Name:   fileHeader.Filename,
					Reason: err.Error(),
				}
				failedImgList = append(failedImgList, img)
				continue
			}
			// Resize for thumbnail, favicon, download
			thumb := resize.Resize(600, 300, img, resize.Lanczos2)
			favIcon := resize.Resize(32, 32, img, resize.Lanczos2)
			download := resize.Resize(800, 800, img, resize.Lanczos2)
			original := img
			// Save thumbnail with compression
			thumbFile, err := os.Create("./uploads/thumb/" + fileHeader.Filename)
			if err != nil {
				img := FailedImg{
					Name:   fileHeader.Filename,
					Reason: "Failed to create thumb file: " + err.Error(),
				}
				failedImgList = append(failedImgList, img)
				continue
			}
			defer thumbFile.Close()
			err = jpeg.Encode(thumbFile, thumb, &jpeg.Options{Quality: 75})
			if err != nil {
				img := FailedImg{
					Name:   fileHeader.Filename,
					Reason: "Failed to encode thumb: " + err.Error(),
				}
				failedImgList = append(failedImgList, img)
				continue
			}
			// Save favicon with compression
			favIconFile, err := os.Create("./uploads/favicon/" + fileHeader.Filename)
			if err != nil {
				img := FailedImg{
					Name:   fileHeader.Filename,
					Reason: "Failed to create favicon file: " + err.Error(),
				}
				failedImgList = append(failedImgList, img)
				continue
			}
			defer favIconFile.Close()
			err = jpeg.Encode(favIconFile, favIcon, &jpeg.Options{Quality: 70})
			if err != nil {
				img := FailedImg{
					Name:   fileHeader.Filename,
					Reason: "Failed to encode favicon: " + err.Error(),
				}
				failedImgList = append(failedImgList, img)
				continue
			}
			// Save download with compression
			downloadFile, err := os.Create("./uploads/download/" + fileHeader.Filename)
			if err != nil {
				img := FailedImg{
					Name:   fileHeader.Filename,
					Reason: "Failed to create download file: " + err.Error(),
				}
				failedImgList = append(failedImgList, img)
				continue
			}
			defer downloadFile.Close()
			err = jpeg.Encode(downloadFile, download, &jpeg.Options{Quality: 85})
			if err != nil {
				img := FailedImg{
					Name:   fileHeader.Filename,
					Reason: "Failed to encode download: " + err.Error(),
				}
				failedImgList = append(failedImgList, img)
				continue
			}
			// Save original with compression (no resizing)
			originalFile, err := os.Create("./uploads/original/" + fileHeader.Filename)
			if err != nil {
				img := FailedImg{
					Name:   fileHeader.Filename,
					Reason: "Failed to create original file: " + err.Error(),
				}
				failedImgList = append(failedImgList, img)
				continue
			}
			defer originalFile.Close()
			err = jpeg.Encode(originalFile, original, &jpeg.Options{Quality: 95})
			if err != nil {
				img := FailedImg{
					Name:   fileHeader.Filename,
					Reason: "Failed to encode original: " + err.Error(),
				}
				failedImgList = append(failedImgList, img)
				continue
			}
			// Track processed files
			processedFiles = append(processedFiles,
				"thumb/"+fileHeader.Filename,
				"favicon/"+fileHeader.Filename,
				"download/"+fileHeader.Filename,
				"original/"+fileHeader.Filename)

		} else {
			img := FailedImg{
				Name:   fileHeader.Filename,
				Reason: "Unsupported content type: " + contentType,
			}
			failedImgList = append(failedImgList, img)
		}
	}

	data := ResData{
		SavedFiles:  processedFiles,
		FailedFiles: failedImgList,
	}

	sendSuccessResponse(w, http.StatusOK, data)

}
