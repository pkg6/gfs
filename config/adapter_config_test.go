package config

import (
	"github.com/zzqqw/gfs"
	"reflect"
	"testing"
)

func TestAdapterConfig_Adapters(t *testing.T) {
	type fields struct {
		Default string
		Adapter *Adapter
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
			a := &AdapterConfig{
				Default: tt.fields.Default,
				Adapter: tt.fields.Adapter,
			}
			if got := a.Adapters(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Adapters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdapterConfig_Disk(t *testing.T) {
	type fields struct {
		Default string
		Adapter *Adapter
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
			a := &AdapterConfig{
				Default: tt.fields.Default,
				Adapter: tt.fields.Adapter,
			}
			if got := a.Disk(); got != tt.want {
				t.Errorf("Disk() = %v, want %v", got, tt.want)
			}
		})
	}
}
