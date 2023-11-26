package gfs_test

import (
	"github.com/zzqqw/gfs"
	"github.com/zzqqw/gfs/bosfs"
	"github.com/zzqqw/gfs/cloudstoragefs"
	"github.com/zzqqw/gfs/config"
	"github.com/zzqqw/gfs/cosfs"
	"github.com/zzqqw/gfs/kodofs"
	"github.com/zzqqw/gfs/localfs"
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
	if fs.Disk("") != "local" {
		t.Fatal(err)
	}
	if fs.Disk("oss") != "oss" {
		t.Fatal(err)
	}
	if fs.Disk("oss2") == "" {
		t.Fatal(err)
	}
	if fs.DiskExist("oss2") == true {
		t.Fatal(err)
	}
}
