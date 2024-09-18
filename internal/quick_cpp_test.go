package internal

import "testing"

func Test_check(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test check function",
			args: args{
				err: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			check(tt.args.err)
		})
	}
}

func Test_readFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readFile(tt.args.path); got != tt.want {
				t.Errorf("readFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_add_files(t *testing.T) {
	type args struct {
		projectName string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			add_files(tt.args.projectName)
		})
	}
}

func Test_printProgressBar(t *testing.T) {
	type args struct {
		done chan error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printProgressBar(tt.args.done)
		})
	}
}

func Test_installMsys2(t *testing.T) {
	type args struct {
		done chan error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			installMsys2(tt.args.done)
		})
	}
}

func Test_installToolchain(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			installToolchain()
		})
	}
}

func Test_msys2(t *testing.T) {
	type args struct {
		projectName string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msys2(tt.args.projectName)
		})
	}
}

func TestCreate(t *testing.T) {
	type args struct {
		projectName string
		full        bool
		skipMsys    bool
		noFiles     bool
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Create(tt.args.projectName, tt.args.full, tt.args.skipMsys, tt.args.noFiles)
		})
	}
}

func TestRevert(t *testing.T) {
	type args struct {
		projectName string
		msys2       bool
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Revert(tt.args.projectName, tt.args.msys2)
		})
	}
}
