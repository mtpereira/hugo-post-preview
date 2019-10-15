package screenshotter

import (
	"bytes"
	"net/url"
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

func TestScreenshotter_Take(t *testing.T) {
	tests := map[string]struct {
		ss      *Screenshotter
		postURL string
		element string
		output  string
		wantErr bool
	}{
		"no errors": {
			ss:      New(),
			postURL: "localhost:1337",
			element: "post",
			output:  "",
			wantErr: false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			w := &bytes.Buffer{}
			url, err := url.Parse(tt.postURL)
			if err != nil {
				t.Fatal(err)
			}
			if err := tt.ss.Take(*url, tt.element, w); (err != nil) != tt.wantErr {
				t.Errorf("Screenshotter.Take() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.output {
				t.Errorf("Screenshotter.Take() = %v, want %v", gotW, tt.output)
			}
		})
	}
}
