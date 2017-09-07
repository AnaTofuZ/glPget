package glPget

import (
	"reflect"
	"testing"
)

func TestOptions_parse(t *testing.T) {
	type fields struct {
		Help    bool
		Version bool
		Trace   bool
		Procs   int
	}
	type args struct {
		argv []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := &Options{
				Help:    tt.fields.Help,
				Version: tt.fields.Version,
				Trace:   tt.fields.Trace,
				Procs:   tt.fields.Procs,
			}
			got, err := opts.parse(tt.args.argv)
			if (err != nil) != tt.wantErr {
				t.Errorf("Options.parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Options.parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
