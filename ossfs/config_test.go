package ossfs

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/pkg6/gfs"
	"net/url"
	"reflect"
	"testing"
)

func TestConfig_NewAdapter(t *testing.T) {
	type fields struct {
		CDN             string
		Bucket          string
		Endpoint        string
		AccessKeyID     string
		AccessKeySecret string
		OssConfig       *oss.Config
	}
	tests := []struct {
		name   string
		fields fields
		want   gfs.IAdapter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				CDN:             tt.fields.CDN,
				Bucket:          tt.fields.Bucket,
				Endpoint:        tt.fields.Endpoint,
				AccessKeyID:     tt.fields.AccessKeyID,
				AccessKeySecret: tt.fields.AccessKeySecret,
				Config:          tt.fields.OssConfig,
			}
			if got := c.NewAdapter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAdapter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_URL(t *testing.T) {
	type fields struct {
		CDN             string
		Bucket          string
		Endpoint        string
		AccessKeyID     string
		AccessKeySecret string
		OssConfig       *oss.Config
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *url.URL
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				CDN:             tt.fields.CDN,
				Bucket:          tt.fields.Bucket,
				Endpoint:        tt.fields.Endpoint,
				AccessKeyID:     tt.fields.AccessKeyID,
				AccessKeySecret: tt.fields.AccessKeySecret,
				Config:          tt.fields.OssConfig,
			}
			got, err := c.URL(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("URL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("URL() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_UseBucket(t *testing.T) {
	type fields struct {
		CDN             string
		Bucket          string
		Endpoint        string
		AccessKeyID     string
		AccessKeySecret string
		OssConfig       *oss.Config
	}
	type args struct {
		bucket string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				CDN:             tt.fields.CDN,
				Bucket:          tt.fields.Bucket,
				Endpoint:        tt.fields.Endpoint,
				AccessKeyID:     tt.fields.AccessKeyID,
				AccessKeySecret: tt.fields.AccessKeySecret,
				Config:          tt.fields.OssConfig,
			}
			if got := c.UseBucket(tt.args.bucket); got != tt.want {
				t.Errorf("UseBucket() = %v, want %v", got, tt.want)
			}
		})
	}
}
