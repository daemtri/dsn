package dsn

import (
	"net/url"
	"reflect"
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	type args struct {
		dataSourceName string
	}
	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.dataSourceName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_Host(t *testing.T) {
	tests := []struct {
		name string
		c    *Config
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Host(); got != tt.want {
				t.Errorf("Config.Host() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_Scheme(t *testing.T) {
	tests := []struct {
		name string
		c    *Config
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Scheme(); got != tt.want {
				t.Errorf("Config.Scheme() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_Path(t *testing.T) {
	tests := []struct {
		name string
		c    *Config
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Path(); got != tt.want {
				t.Errorf("Config.Path() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_User(t *testing.T) {
	tests := []struct {
		name string
		c    *Config
		want *url.Userinfo
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.User(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Config.User() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValues_Int(t *testing.T) {
	type args struct {
		param string
		def   int
	}
	tests := []struct {
		name string
		v    *Values
		args args
		want int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Int(tt.args.param, tt.args.def); got != tt.want {
				t.Errorf("Values.Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValues_String(t *testing.T) {
	type args struct {
		param string
		def   string
	}
	tests := []struct {
		name  string
		query string
		args  args
		want  string
	}{
		{name: "test_a=hello", query: "a=hello&a=world", args: args{param: "a", def: "no"}, want: "hello"},
		{name: "test_a=nil", query: "b=world", args: args{param: "a", def: "empty"}, want: "empty"},
		{name: "test_a=empyty", query: "a=", args: args{param: "a", def: "empty"}, want: "empty"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			values, _ := url.ParseQuery(tt.query)
			v := Values{Values: values}
			if got := v.String(tt.args.param, tt.args.def); got != tt.want {
				t.Errorf("Values.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValues_Bool(t *testing.T) {
	type args struct {
		param string
		def   bool
	}
	tests := []struct {
		name  string
		query string
		args  args
		want  bool
	}{
		{name: "test_a=true", query: "a=true", args: args{param: "a", def: false}, want: true},
		{name: "test_a=T", query: "a=T", args: args{param: "a", def: false}, want: true},
		{name: "test_a=1", query: "a=1", args: args{param: "a", def: false}, want: true},
		{name: "test_a=2", query: "a=2", args: args{param: "a", def: false}, want: false},
		{name: "test_a=nil", query: "b=false", args: args{param: "a", def: false}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			values, _ := url.ParseQuery(tt.query)
			v := Values{Values: values}
			if got := v.Bool(tt.args.param, tt.args.def); got != tt.want {
				t.Errorf("Values.Bool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValues_Duration(t *testing.T) {
	type args struct {
		param string
		def   time.Duration
	}
	tests := []struct {
		name  string
		query string
		args  args
		want  time.Duration
	}{
		{name: "test_a=1s", query: "a=1s", args: args{param: "a", def: 2 * time.Second}, want: time.Second},
		{name: "test_a=2s", query: "b=2m", args: args{param: "a", def: 2 * time.Second}, want: 2 * time.Second},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			values, _ := url.ParseQuery(tt.query)
			v := Values{Values: values}
			if got := v.Duration(tt.args.param, tt.args.def); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Values.Duration() = %v, want %v", got, tt.want)
			}
		})
	}
}
