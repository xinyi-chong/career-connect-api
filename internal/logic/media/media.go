package media

import (
	"context"
	"fmt"
	"gf_demo/internal/consts"
	"gf_demo/internal/dao"
	"gf_demo/internal/model/do"
	"gf_demo/internal/service"
	"mime/multipart"
	"path"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	s3config "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/gogf/gf/os/genv"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/google/uuid"
)

type (
	sMedia struct{}
)

func init() {
	service.RegisterMedia(New())
}

func New() service.IMedia {
	return &sMedia{}
}

func (s *sMedia) CreateMedia(ctx context.Context, file *ghttp.UploadFile, accountID string) (*string, error) {
	url, key, err := s.UploadMedia(ctx, *file, accountID)
	if err != nil {
		return nil, err
	}

	id := uuid.New().String()
	_, err = dao.Media.Ctx(ctx).Data(do.Media{
		Id:        id,
		Url:       url,
		Key:       key,
	}).Insert()

	if err != nil {
		g.Log().Error(ctx, "Failed to Create Media: ", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Create Media: " + err.Error())
		return nil, err
	}

	return &id, nil
}

func (s *sMedia) UpdateMediaByID(ctx context.Context, id string, file *ghttp.UploadFile, accountID string) (error) {
	url, key, err := s.UploadMedia(ctx, *file, accountID)
	if err != nil {
		return err
	}

	_, err = dao.Media.Ctx(ctx).Data(do.Media{
		Url:       url,
		Key:       key,
	}).Where(do.Media{
		Id: id,
	}).Update()

	if err != nil {
		g.Log().Error(ctx, "Failed to Update Media: ", err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Update Media: " + err.Error())
		return err
	}

	g.Log().Info(ctx, consts.SUCCESS_UPDATE, " Media: ", id)
	return nil
}

func (s *sMedia) DeleteMediaByID(ctx context.Context, id string) (error) {
	_, err := dao.Media.Ctx(ctx).Where(do.Media{
		Id: id,
	}).Delete()

	if err != nil {
		g.Log().Error(ctx, "Failed to Delete Media By ID: ", id, err)
		err = gerror.NewCode(gcode.CodeDbOperationError, "Failed to Delete Media: " + err.Error())
		return err
	}

	g.Log().Info(ctx, consts.SUCCESS_DELETE, " Media: ", id)
	return nil
}

func (s *sMedia) UploadMedia(ctx context.Context, file ghttp.UploadFile, accountID string) (url *string, objectKey *string, err error) {
	err = ValidateFile(file)
	if err != nil {
		return nil, nil, err
	}

	buffer, err := file.Open()
	if err != nil {
		g.Log().Error(ctx, "Failed to open file: ", err)
		err = gerror.NewCode(gcode.CodeOperationFailed, "Failed to Open File: " + err.Error())
		return nil, nil, err
	}
	defer buffer.Close()

	objectKey, err = GenerateObjectKey(file.Filename, accountID)
	if err != nil {
		return nil, nil, err
	}

	// upload file
	bucket := genv.Get(consts.AWS_BUCKET)
	region := genv.Get(consts.AWS_REGION)

	err = UploadFileToAWS(bucket, region, *objectKey, buffer)
	if err != nil {
		return nil, nil, err
	}
	
	urlStr := "https://" + bucket + ".s3." + region + ".amazonaws.com/" + *objectKey
	
	url = &urlStr

	return
}

func IsImageFile(filename string) bool {
	allowedExtensions := []string{".jpg", ".jpeg", ".png", ".gif", ".bmp", ".svg"}
	for _, s := range allowedExtensions {
		if strings.Contains(filename, s) {
			return true
		}
	}
	return false
}

func ValidateFile(file ghttp.UploadFile) error {
	if file.Size > 10*1024*1024 {
		err := gerror.NewCode(gcode.CodeValidationFailed, "file size is too large, maximum 10MB")
		return err
	} else if file.Filename == "" {
		err := gerror.NewCode(gcode.CodeValidationFailed, "file name is required")
		return err
	}
		
	return nil
}

func GenerateObjectKey(filename string, accountID string) (*string, error){
		// get epoch timestamp in string
	epochTimestamp := fmt.Sprintf("%d", time.Now().Unix())

	// encode filename
	filename = fmt.Sprintf("%s%s", epochTimestamp, path.Ext(filename))
	if !IsImageFile(filename) {
		err := gerror.NewCode(gcode.CodeValidationFailed, "Invalid file type, only image files are allowed")
		return nil, err
	}

	// get object key by adding current datetime string , account id and file name
	objectKey := fmt.Sprintf("%s-%s/%s", time.Now().Format("200601021504"), accountID, filename)
	return &objectKey, nil
}

func UploadFileToAWS(bucket string, region string, objectKey string, buffer multipart.File) error {
	s3Client, err := ConnectS3()
	if err != nil {
		return err
	}

	_, err = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &objectKey,
		Body:   buffer,
		ACL:    types.ObjectCannedACL(*aws.String("public-read")),
	})
	if err != nil {
		g.Log().Error(context.Background(), "Failed to Upload File: ", err)
		err := gerror.NewCode(gcode.CodeOperationFailed, "Failed to upload file: " + err.Error())
		return err
	}
	return nil
}

func ConnectS3() (S3 *s3.Client, err error) {
	AWS_ACCESS_KEY_ID := genv.Get(consts.AWS_ACCESS_KEY_ID)
	AWS_SECRET_ACCESS_KEY := genv.Get(consts.AWS_SECRET_ACCESS_KEY)
	AWS_REGION := genv.Get(consts.AWS_REGION)
	
	customProvider := credentials.NewStaticCredentialsProvider(AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, "")

	cfg, err := s3config.LoadDefaultConfig(context.TODO(),
		s3config.WithCredentialsProvider(customProvider),
		s3config.WithRegion(AWS_REGION),
	)

	if err != nil {
		g.Log().Error(context.Background(), "Failed to load s3 config", err)
		err := gerror.NewCode(gcode.CodeInvalidConfiguration, "Failed to load s3 config: " + err.Error())
		return nil, err
	}

	S3 = s3.NewFromConfig(cfg)
	return S3, nil
}