package file

import (
	"bytes"
	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	"github.com/laojie0524/TSDDServerLib/config"
	"github.com/laojie0524/TSDDServerLib/pkg/log"
	"go.uber.org/zap"
	"io"
	"net/url"
)

type ServiceOBS struct {
	log.Log
	ctx *config.Context
}

// NewServiceOBS NewServiceOBS
func NewServiceOBS(ctx *config.Context) *ServiceOBS {

	return &ServiceOBS{
		Log: log.NewTLog("ServiceOBS"),
		ctx: ctx,
	}
}

// UploadFile 上传文件
func (s *ServiceOBS) UploadFile(filePath string, contentType string, copyFileWriter func(io.Writer) error) (map[string]interface{}, error) {
	obsCfg := s.ctx.GetConfig().OBS
	client, err := obs.New(obsCfg.AccessKeyID, obsCfg.AccessKeySecret, obsCfg.Endpoint)
	if err != nil {
		return nil, err
	}

	buff := bytes.NewBuffer(make([]byte, 0))
	err = copyFileWriter(buff)
	if err != nil {
		s.Error("复制文件内容失败！", zap.Error(err))
		return nil, err
	}
	input := &obs.PutObjectInput{
		PutObjectBasicInput: obs.PutObjectBasicInput{
			ObjectOperationInput: obs.ObjectOperationInput{
				Bucket:  s.ctx.GetConfig().OBS.BucketName,
				Key: filePath,
			},
			HttpHeader: obs.HttpHeader{
				ContentType: contentType,
			},
		},
		Body: buff,
	}
	_, err = client.PutObject(input)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{}, nil
}

func (s *ServiceOBS) DownloadURL(path string, filename string) (string, error) {
	obsCfg := s.ctx.GetConfig().OBS

	rpath, _ := url.JoinPath(obsCfg.BucketURL, path)
	return rpath, nil
}
