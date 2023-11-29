package gfs_test

import (
	"fmt"
	"github.com/zzqqw/gfs"
	"github.com/zzqqw/gfs/bosfs"
	"github.com/zzqqw/gfs/cloudstoragefs"
	"github.com/zzqqw/gfs/config"
	"github.com/zzqqw/gfs/cosfs"
	"github.com/zzqqw/gfs/kodofs"
	"github.com/zzqqw/gfs/localfs"
	"github.com/zzqqw/gfs/ossfs"
	"google.golang.org/api/option"
	"testing"
)

func TestNewConfig(t *testing.T) {
	cfg := config.Config{
		LOCAL: &localfs.Config{},
		CloudStorage: &cloudstoragefs.Config{
			Bucket: "test bucket",
			Option: []option.ClientOption{
				option.WithCredentialsFile("CredentialsFile.json"),
			},
		},
		OSS: &ossfs.Config{
			Endpoint:        "oss-cn-hangzhou.aliyuncs.com",
			AccessKeyID:     "****************",
			AccessKeySecret: "****************",
		},
		KODO: &kodofs.Config{
			AccessKey: "AccessKey",
			SecretKey: "SecretKey",
			Bucket:    "test bucket",
		},
		COS: &cosfs.Config{
			BucketURL: "https://bucket-id.cos.ap-beijing.myqcloud.com",
			SecretID:  "SecretID",
			SecretKey: "SecretKey",
		},
		BOS: &bosfs.Config{
			Endpoint: bosfs.DefaultEndpoint,
			Ak:       "Ak",
			Sk:       "Sk",
			Bucket:   "test bucket",
		},
	}
	fs, err := gfs.NewConfig(&cfg)
	if err != nil {
		t.Fatal(err)
	}
	loc, err := fs.Adapter("")
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := loc.(*localfs.Adapter); !ok {
		t.Fatal("default local choose err")
	}
	oss, err := fs.Adapter("oss")
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := oss.(*ossfs.Adapter); !ok {
		t.Fatal("oss choose err")
	}
	_, err = fs.Adapter("oss2")
	if err == fmt.Errorf("unable to find %s disk", "oss2") {
		t.Fatal(err)
	}
}
