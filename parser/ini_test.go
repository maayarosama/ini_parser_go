package parser

import (
	"os"
	"reflect"
	"testing"
)

const file_path = "../testdata/sample.ini"

func TestReadFromString(t *testing.T) {
	t.Run("ReadFromString_with_comments", func(t *testing.T) {
		parser := NewParser()
		content := `
	[Profile]
	name = jarvis
	# credential
	password = secret
	`
		parser.ReadFromString(content)
		got := parser.data["Profile"]["password"]
		want := "secret"
		if got != want {
			t.Errorf("Got: %s. Expected: %s", got, want)
		}
	})

	t.Run("ReadFromString_multiple_section", func(t *testing.T) {
		parser := NewParser()
		content := `
		[Profile]
		name = jarvis
		password = secret
		[Deployment]
		project = peertube
		name = peertest
		public_ip = true
		cpu = 4
		memory = 8192
		[Owner]
		name = mo
		email = mo@peertube.com
		`
		parser.ReadFromString(content)

		got := parser.data["Owner"]["email"]
		want := "mo@peertube.com"
		if got != want {
			t.Errorf("Got: %s. Expected: %s", got, want)
		}
	})

}

func TestReadFromFile(t *testing.T) {
	parser := NewParser()
	parser.ReadFile(file_path)
	got := parser.data["Profile"]["password"]
	want := "secret"

	if got != want {
		t.Errorf("Got: %s. Expected: %s", got, want)
	}

}

func TestGet(t *testing.T) {
	t.Run("get", func(t *testing.T) {
		parser := NewParser()
		content := `
	[Profile]
	name = jarvis
	# credential
	password = secret
	`
		parser.ReadFromString(content)
		got := parser.Get("Profile", "password")
		want := "secret"
		if got != want {
			t.Errorf("Got: %s. Expected: %s", got, want)
		}
	})

}
func TestGetSection(t *testing.T) {
	t.Run("getsection_fromstring", func(t *testing.T) {
		parser := NewParser()
		content := `
		[Profile]
		name = jarvis
		password = secret
		[Deployment]
		project = peertube
		name = peertest
		public_ip = true
		cpu = 4
		memory = 8192
		[Owner]
		name = mo
		email = mo@peertube.com
	`
		parser.ReadFromString(content)
		got := parser.GetSection("Deployment")
		want := map[string]string{
			"cpu":       "4",
			"memory":    "8192",
			"name":      "peertest",
			"project":   "peertube",
			"public_ip": "true",
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Got: %v. Expected: %v", got, want)
		}
	})

	t.Run("getsection_fromfile", func(t *testing.T) {
		parser := NewParser()
		parser.ReadFile(file_path)
		got := parser.GetSection("Profile")
		want := map[string]string{
			"name":     "jarvis",
			"password": "secret",
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Got: %v. Expected: %v", got, want)
		}
	})

}

func TestGetSections(t *testing.T) {
	t.Run("getsections_fromstring", func(t *testing.T) {
		parser := NewParser()
		content := `
		[Profile]
		name = jarvis
		password = secret
		[Deployment]
		project = peertube
		name = peertest
		public_ip = true
		cpu = 4
		memory = 8192
		[Owner]
		name = mo
		email = mo@peertube.com
	`
		parser.ReadFromString(content)
		got := parser.GetSections()
		want := []string{"Profile", "Deployment", "Owner"}

		if !reflect.DeepEqual(want, got) {
			t.Errorf("Got: %v. Expected: %v", got, want)
		}
	})

	t.Run("getsections_fromfile", func(t *testing.T) {
		parser := NewParser()
		parser.ReadFile(file_path)
		got := parser.GetSections()
		want := []string{"Profile", "Deployment", "Owner"}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Got: %v. Expected: %v", got, want)
		}
	})

}
func TestGetSectionKeys(t *testing.T) {
	t.Run("getsectionkeys_fromstring", func(t *testing.T) {
		parser := NewParser()
		content := `
		[Profile]
		name = jarvis
		password = secret
		[Deployment]
		project = peertube
		name = peertest
		public_ip = true
		cpu = 4
		memory = 8192
		[Owner]
		name = mo
		email = mo@peertube.com
	`
		parser.ReadFromString(content)
		got := parser.GetSectionKeys("Profile")
		want := []string{"name", "password"}

		if !reflect.DeepEqual(want, got) {
			t.Errorf("Got: %v. Expected: %v", got, want)
		}
	})

	t.Run("getsection_fromfile", func(t *testing.T) {
		parser := NewParser()
		parser.ReadFile(file_path)
		got := parser.GetSectionKeys("Owner")
		want := []string{"name", "email"}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Got: %v. Expected: %v", got, want)
		}
	})

}

func TestWriteToFile(t *testing.T) {
	t.Run("writetofile_fromstring", func(t *testing.T) {
		parser := NewParser()
		content := `
		[Profile]
		name = jarvis
		password = secret
		[Deployment]
		project = peertube
		name = peertest
		public_ip = true
		cpu = 4
		memory = 8192
		[Owner]
		name = mo
		email = mo@peertube.com
	`
		parser.ReadFromString(content)
		got, _ := parser.WriteToFile("../testdata/writetofile_string_sample.ini")
		want := true
		_, isFile := os.Stat("../testdata/writetofile_string_sample.ini")

		if got != want {
			t.Errorf("Got: %v. Expected: %v", got, want)

		}
		if isFile != nil {
			t.Errorf("File wasn't created")

		}

	})

	t.Run("writetofile_fromfile", func(t *testing.T) {
		parser := NewParser()
		parser.ReadFile(file_path)
		got, _ := parser.WriteToFile("../testdata/writetofile_sample.ini")
		want := true
		_, isFile := os.Stat("../testdata/writetofile_sample.ini")

		if got != want {
			t.Errorf("Got: %v. Expected: %v", got, want)

		}
		if isFile != nil {
			t.Errorf("File wasn't created")

		}
	})

}
