package services

import (
	"github/dev-hack95/Textflow/constants"
	"github/dev-hack95/Textflow/helper"
	"github/dev-hack95/Textflow/structs"
	"github/dev-hack95/Textflow/utilities/logs"
	"log"
	"strings"

	"github.com/minio/minio-go"
)

func UploadVideo(data structs.VideoPayload) (*structs.VideoPayload, error) {
	minioClient, err := minio.New(constants.Endpoint, constants.AccessKeyID, constants.SecretAccessKey, constants.UseSSL)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	parts := strings.Split(data.Video, "/")

	dataVideo := parts[len(parts)-1]

	video, err := helper.StringGenerator(dataVideo)
	if err != nil {
		logs.Error("Error: ", err.Error())
	}

	_, err = minioClient.FPutObject(constants.BucketName, video, data.Video, minio.PutObjectOptions{})

	if err != nil {
		logs.Error("Error: ", err.Error())
	}

	logs.Info("Video uploded Succesfully")

	var ret structs.VideoPayload

	ret.Email = data.Email
	ret.Video = video

	return &ret, nil
}
