package file

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/laojie0524/TSDDServerLib/config"
	"github.com/laojie0524/TSDDServerLib/pkg/log"
	"go.uber.org/zap"
	"io"
	"net/url"
)

type ServiceAWS struct {
	log.Log
	ctx *config.Context
}

// NewServiceAWS NewServiceAWS
func NewServiceAWS(ctx *config.Context) *ServiceAWS {

	return &ServiceAWS{
		Log: log.NewTLog("ServiceAWS"),
		ctx: ctx,
	}
}

// UploadFile 上传文件
func (s *ServiceAWS) UploadFile(filePath string, contentType string, copyFileWriter func(io.Writer) error) (map[string]interface{}, error) {
	awsCfg := s.ctx.GetConfig().AWS
	client, err := session.NewSession(&aws.Config{
		Region:           aws.String(awsCfg.Region),
		Endpoint:         aws.String(awsCfg.Endpoint), //minio在这里设置地址,可以兼容
		S3ForcePathStyle: aws.Bool(false),
		DisableSSL:       aws.Bool(false),
		Credentials: credentials.NewStaticCredentials(
			awsCfg.SecretID,
			awsCfg.SecretKey,
			"",
		),
	})
	if err != nil {
		return nil, err
	}

	buff := bytes.NewBuffer(make([]byte, 0))
	err = copyFileWriter(buff)
	if err != nil {
		s.Error("复制文件内容失败！", zap.Error(err))
		return nil, err
	}
	uploader := s3manager.NewUploader(client)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(awsCfg.Bucket),
        Key:    aws.String(filePath),
        Body:   buff,
        ContentType: aws.String(contentType),
	})
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{}, nil
}

func (s *ServiceAWS) DownloadURL(path string, filename string) (string, error) {
	awsCfg := s.ctx.GetConfig().AWS

	rpath, _ := url.JoinPath(awsCfg.BucketURL, path)
	return rpath, nil
}
