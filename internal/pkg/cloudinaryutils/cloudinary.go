package cloudinaryutils

import (
	"context"
	"log"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadImage(ctx context.Context, fileName string, folder string) (string, error) {
	cld, err := cloudinary.NewFromParams(os.Getenv("CLOUDINARY_NAME"), os.Getenv("CLOUDINARY_API_KEY"), os.Getenv("CLOUDINARY_API_SECRET"))
	if err != nil {
		log.Fatalf("Failed to initialize Cloudinary, %v", err)
		return "", err
	}
	uploadResp, err := cld.Upload.Upload(ctx, fileName, uploader.UploadParams{Folder: folder})
	if err != nil {
		log.Fatalf("Failed to upload image to cloudinary, %v", err)
		return "", err
	}
	return uploadResp.SecureURL, nil
}
