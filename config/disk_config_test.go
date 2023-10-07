package config

import (
	"github.com/zzqqw/gfs"
	"reflect"
	"testing"
)

func TestDiskConfig_Adapters(t *testing.T) {
	type fields struct {
		Default string
		Disks   []Disks
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]gfs.IAdapter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &DiskConfig{
				Default: tt.fields.Default,
				Disks:   tt.fields.Disks,
			}
			if got := s.Adapters(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Adapters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiskConfig_Disk(t *testing.T) {
	type fields struct {
		Default string
		Disks   []Disks
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &DiskConfig{
				Default: tt.fields.Default,
				Disks:   tt.fields.Disks,
			}
			if got := s.Disk(); got != tt.want {
				t.Errorf("Disk() = %v, want %v", got, tt.want)
			}
		})
	}
}
