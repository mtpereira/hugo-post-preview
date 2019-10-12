package screenshotter

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestDebug(t *testing.T) {
	tests := map[string]struct {
		enabled bool
	}{
		"enabled": {
			enabled: true,
		},
		"disabled": {
			enabled: false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := New(Debug(tt.enabled))
			if diff := cmp.Diff(tt.enabled, got.debug); diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func TestTimeout(t *testing.T) {
	tests := map[string]struct {
		timeout string
	}{
		"less than zero": {
			timeout: "-1s",
		},
		"greater than zero": {
			timeout: "10s",
		},
		"zero": {
			timeout: "0s",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			d, err := time.ParseDuration(tt.timeout)
			if err != nil {
				t.Fatal(err)
			}

			got := New(Timeout(d))
			if diff := cmp.Diff(d, got.timeout); diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
