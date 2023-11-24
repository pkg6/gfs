package ossfs

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/zzqqw/gfs"
	"net/url"
)

type Config struct {
	//https://help.aliyun.com/zh/oss/user-guide/regions-and-endpoints
	CDN             string
	Bucket          string
	Endpoint        string
	AccessKeyID     string
	AccessKeySecret string
	OssConfig       *oss.Config
}

func (c *Config) New() gfs.IAdapter {
	return NewOSS(c)
}

func (c *Config) URL(path string) (*url.URL, error) {
	bucketUrl, err := gfs.BucketURLMake(c.CDN, c.Endpoint, c.Bucket)
	if err != nil {
		return nil, err
	}
	return gfs.PublicURLMake(bucketUrl.String(), path)
}
func (c *Config) UseBucket(bucket string) string {
	if bucket != "" {
		return bucket
	}
	return c.Bucket
}
