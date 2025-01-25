package config

import (
	"context"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadFile(file multipart.File, filePath string) (string, error) {
	//connect to CLoudinary
	url := os.Getenv("CLOUDINARY_URL")
	cld, err := cloudinary.NewFromURL(url)
	if err != nil {
		return "", err
	}

	//upload image
	ctx := context.Background()
	resp, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{PublicID: filePath})
	if err != nil {
		return "", err
	}

	return resp.SecureURL, nil
}
