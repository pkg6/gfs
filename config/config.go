package config

import (
	"github.com/zzqqw/gfs/bosfs"
	"github.com/zzqqw/gfs/cloudstoragefs"
	"github.com/zzqqw/gfs/cosfs"
	"github.com/zzqqw/gfs/kodofs"
	"github.com/zzqqw/gfs/localfs"
	"github.com/zzqqw/gfs/ossfs"
)

type Config struct {
	LOCAL        *localfs.Config        `gfs:"local,default"`
	OSS          *ossfs.Config          `gfs:"oss"`
	BOS          *bosfs.Config          `gfs:"bos"`
	COS          *cosfs.Config          `gfs:"cos"`
	KODO         *kodofs.Config         `gfs:"kodo"`
	CloudStorage *cloudstoragefs.Config `gfs:"cloud_storage"`
}
