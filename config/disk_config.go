package config

import (
	"github.com/zzqqw/gfs"
	"github.com/zzqqw/gfs/bosfs"
	"github.com/zzqqw/gfs/cloudstoragefs"
	"github.com/zzqqw/gfs/cosfs"
	"github.com/zzqqw/gfs/kodofs"
	"github.com/zzqqw/gfs/localfs"
	"github.com/zzqqw/gfs/ossfs"
	"reflect"
)

type DiskConfig struct {
	Default string
	Disks   []Disks
}
type Disks struct {
	DiskName     string
	LOCAL        *localfs.Config
	OSS          *ossfs.Config
	BOS          *bosfs.Config
	COS          *cosfs.Config
	KODO         *kodofs.Config
	CloudStorage *cloudstoragefs.Config
}

func (s *DiskConfig) Disk() string {
	return s.Default
}

func (s *DiskConfig) Adapters() map[string]gfs.IAdapter {
	disks := s.Disks
	sortIAdapters := map[string]gfs.IAdapter{}
	adapterMaps := map[string][]gfs.IAdapter{}
	var sortNames []string
	if disks == nil {
		return sortIAdapters
	}
	for _, disk := range disks {
		sortNames = append(sortNames, disk.DiskName)
		v := reflect.ValueOf(disk)
		for i := 0; i < v.NumField(); i++ {
			e := v.Field(i)
			if !e.IsZero() {
				if a, ok := e.Interface().(gfs.IAdapterConfig); ok {
					if a != nil {
						adapterMaps[disk.DiskName] = append(adapterMaps[disk.DiskName], a.New())
					}
				}
			}
		}
	}
	for _, name := range sortNames {
		if as, ok := adapterMaps[name]; ok {
			if len(as) > 0 {
				sortIAdapters[name] = as[0]
			}
		}
	}
	return sortIAdapters
}
