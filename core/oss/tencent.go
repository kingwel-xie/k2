package oss

type TencentCOS struct{}
//
//// UploadFile upload file to COS
//func (*Tencent) UploadFile(file *multipart.FileHeader) (string, string, error) {
//	client := NewClient()
//	f, openError := file.Open()
//	if openError != nil {
//		log.Error("function file.Open() Filed", zap.Any("err", openError.Error()))
//		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
//	}
//	defer f.Close() // 创建文件 defer 关闭
//	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename)
//
//	_, err := client.Object.Put(context.Background(), global.Config.Tencent.PathPrefix+"/"+fileKey, f, nil)
//	if err != nil {
//		panic(err)
//	}
//	return global.Config.Tencent.BaseURL + "/" + global.Config.Tencent.PathPrefix + "/" + fileKey, fileKey, nil
//}
//
//// DeleteFile delete file form COS
//func (*Tencent) DeleteFile(key string) error {
//	client := NewClient()
//	name := global.Config.Tencent.PathPrefix + "/" + key
//	_, err := client.Object.Delete(context.Background(), name)
//	if err != nil {
//		log.Error("function bucketManager.Delete() Filed", zap.Any("err", err.Error()))
//		return errors.New("function bucketManager.Delete() Filed, err:" + err.Error())
//	}
//	return nil
//}
//
//// NewClient init COS client
//func NewClient() *cos.Client {
//	urlStr, _ := url.Parse("https://" + global.Config.Tencent.Bucket + ".cos." + global.Config.Tencent.Region + ".myqcloud.com")
//	baseURL := &cos.BaseURL{BucketURL: urlStr}
//	client := cos.NewClient(baseURL, &http.Client{
//		Transport: &cos.AuthorizationTransport{
//			SecretID:  global.Config.Tencent.SecretID,
//			SecretKey: global.Config.Tencent.SecretKey,
//		},
//	})
//	return client
//}
