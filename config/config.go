package config

import (
	"github.com/pkg6/gfs/bosfs"
	"github.com/pkg6/gfs/cloudstoragefs"
	"github.com/pkg6/gfs/cosfs"
	"github.com/pkg6/gfs/kodofs"
	"github.com/pkg6/gfs/localfs"
	"github.com/pkg6/gfs/ossfs"
)

type Config struct {
	LOCAL        *localfs.Config        `gfs:"local,default"`
	OSS          *ossfs.Config          `gfs:"oss"`
	BOS          *bosfs.Config          `gfs:"bos"`
	COS          *cosfs.Config          `gfs:"cos"`
	KODO         *kodofs.Config         `gfs:"kodo"`
	CloudStorage *cloudstoragefs.Config `gfs:"cloud_storage"`
}
