package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewDefaultConfig does some very crude and basic testing for NewDefaultConfig.
func TestNewDefaultConfig(t *testing.T) {
	type args struct {
		appName string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "app",
			args: args{
				appName: "app",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotNil(t, NewDefaultConfig(tt.args.appName))
		})
	}
}
