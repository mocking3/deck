package file

import (
	"io"
	"os"
	"reflect"
	"testing"

	"github.com/hbagdi/go-kong/kong"
)

func Test_yamlFilesInDir(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name:    "empty directory",
			args:    args{"testdata/emptydir"},
			want:    nil,
			wantErr: false,
		},
		{
			name:    "directory does not exist",
			args:    args{"testdata/does-not-exist"},
			want:    nil,
			wantErr: true,
		},
		{
			name: "valid directory",
			args: args{"testdata/emptyfiles"},
			want: []string{
				"testdata/emptyfiles/Baz.YamL",
				"testdata/emptyfiles/bar.yaml",
				"testdata/emptyfiles/foo.yml",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := yamlFilesInDir(tt.args.dir)
			if (err != nil) != tt.wantErr {
				t.Errorf("yamlFilesInDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("yamlFilesInDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getReaders(t *testing.T) {
	type args struct {
		fileOrDir string
	}
	tests := []struct {
		name string
		args args
		want []io.Reader
		// length of returned array
		wantLen int
		wantErr bool
	}{
		{
			name:    "read from standard input",
			args:    args{"-"},
			want:    []io.Reader{os.Stdin},
			wantLen: 1,
			wantErr: false,
		},
		{
			name:    "directory does not exist",
			args:    args{"testdata/does-not-exist"},
			want:    nil,
			wantLen: 0,
			wantErr: true,
		},
		{
			name:    "valid directory",
			args:    args{"testdata/emptyfiles"},
			want:    nil,
			wantLen: 3,
			wantErr: false,
		},
		{
			name:    "valid file",
			args:    args{"testdata/file.yaml"},
			want:    nil,
			wantLen: 1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getReaders(tt.args.fileOrDir)
			if (err != nil) != tt.wantErr {
				t.Errorf("getReaders() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantLen != len(got) {
				t.Errorf("getReaders() mismatch in returned length: "+
					"want = %v, got = %v", tt.wantLen, len(got))
				return
			}
			if tt.want != nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getReaders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getContent(t *testing.T) {
	type args struct {
		fileOrDir string
	}
	tests := []struct {
		name    string
		args    args
		want    *Content
		wantErr bool
	}{
		{
			name:    "directory does not exist",
			args:    args{"testdata/does-not-exist"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "empty directory",
			args:    args{"testdata/emptydir"},
			want:    &Content{},
			wantErr: false,
		},
		{
			name:    "directory with empty files",
			args:    args{"testdata/emptyfiles"},
			want:    &Content{},
			wantErr: false,
		},
		{
			name:    "bad yaml",
			args:    args{"testdata/badyaml"},
			want:    nil,
			wantErr: true,
		},
		{
			name: "single file",
			args: args{"testdata/file.yaml"},
			want: &Content{
				Services: []Service{
					{
						Service: kong.Service{
							Name: kong.String("svc2"),
							Host: kong.String("2.example.com"),
						},
						Routes: []*Route{
							{
								Route: kong.Route{
									Name:  kong.String("r2"),
									Paths: kong.StringSlice("/r2"),
								},
							},
						},
					},
				},
				Plugins: []Plugin{
					{
						Plugin: kong.Plugin{
							Name: kong.String("prometheus"),
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "valid directory",
			args: args{"testdata/valid"},
			want: &Content{
				Info: Info{
					SelectorTags: []string{"tag1"},
				},
				Services: []Service{
					{
						Service: kong.Service{
							Name: kong.String("svc2"),
							Host: kong.String("2.example.com"),
						},
						Routes: []*Route{
							{
								Route: kong.Route{
									Name:  kong.String("r2"),
									Paths: kong.StringSlice("/r2"),
								},
							},
						},
					},
					{
						Service: kong.Service{
							Name: kong.String("svc1"),
							Host: kong.String("1.example.com"),
							Tags: kong.StringSlice("team-svc1"),
						},
						Routes: []*Route{
							{
								Route: kong.Route{
									Name:  kong.String("r1"),
									Paths: kong.StringSlice("/r1"),
								},
							},
						},
					},
				},
				Consumers: []Consumer{
					{
						Consumer: kong.Consumer{
							Username: kong.String("harry"),
						},
					},
				},
				Plugins: []Plugin{
					{
						Plugin: kong.Plugin{
							Name: kong.String("prometheus"),
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getContent(tt.args.fileOrDir)
			if (err != nil) != tt.wantErr {
				t.Errorf("getContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getContent() = %v, want %v", got, tt.want)
			}
		})
	}
}
