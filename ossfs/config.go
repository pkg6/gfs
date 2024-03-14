package ossfs

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/pkg6/gfs"
	"net/url"
)

type Config struct {
	CDN    string `json:"cdn" xml:"CDN" yaml:"CDN"`
	Bucket string `json:"bucket" xml:"Bucket" yaml:"Bucket"`
	//https://help.aliyun.com/zh/oss/user-guide/regions-and-endpoints
	Endpoint        string      `json:"endpoint" xml:"Endpoint" yaml:"Endpoint"`
	AccessKeyID     string      `json:"access_key_id" xml:"AccessKeyID" yaml:"AccessKeyID"`
	AccessKeySecret string      `json:"access_key_secret" xml:"AccessKeySecret" yaml:"AccessKeySecret"`
	Config          *oss.Config `json:"config" xml:"Config" yaml:"Config"`
}

func (c *Config) NewAdapter() gfs.IAdapter {
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
