package parser

import (
	"os"
	"reflect"
	"testing"
)

const inisample = "./testdata/sample.ini"
const malformedinisample = "./testdata/wrong_sample.ini"

func TestReadFromString(t *testing.T) {
	t.Run("ReadFromString_with_comments", func(t *testing.T) {
		parser := NewParser()
		content := `
	[Profile]
	name = jarvis
	# credential
	password= secret
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
	parser.ReadFromFile(inisample)
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
		got, _ := parser.Get("Profile", "password")
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
		got, _ := parser.GetSection("Deployment")
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
		parser.ReadFromFile(inisample)
		got, _ := parser.GetSection("Profile")
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
		parser.ReadFromFile(inisample)
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
		parser.ReadFromFile(inisample)
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
		got := parser.WriteToFile("./testdata/writetofile_string_sample.ini")
		// want := true
		_, isFile := os.Stat("./testdata/writetofile_string_sample.ini")

		if got != nil {
			t.Errorf("Got: %v.", got)

		}
		if isFile != nil {
			t.Errorf("File wasn't created")

		}

	})

	t.Run("writetofile_fromfile", func(t *testing.T) {
		parser := NewParser()
		parser.ReadFromFile(inisample)
		got := parser.WriteToFile("./testdata/writetofile_sample.ini")
		// want := true
		_, isFile := os.Stat("./testdata/writetofile_sample.ini")

		if got != nil {
			t.Errorf("Got: %v. ", got)

		}
		if isFile != nil {
			t.Errorf("File wasn't created")

		}
	})

}

func TestWrongValues(t *testing.T) {

	t.Run("test_wrong_content_from_string", func(t *testing.T) {
		parser := NewParser()
		content := `
		[Profile
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
		err := parser.ReadFromString(content)

		if err == nil {
			t.Errorf("It should've returned an error: %v", err)

		}

	})

	t.Run("test_wrong_content_from_file", func(t *testing.T) {
		parser := NewParser()
		err := parser.ReadFromFile(malformedinisample)

		if err == nil {
			t.Errorf("It should've returned an error: %v", err)

		}

	})
}

func TestWrongSection(t *testing.T) {

	t.Run("test_wrong_section_from_string", func(t *testing.T) {
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
		err := parser.ReadFromString(content)

		if err != nil {
			t.Errorf("Error: %v", err)

		}

		_, sectionErr := parser.GetSection("ownerr")

		if sectionErr == nil {
			t.Errorf("ownerr section exists")
		}
	})

	t.Run("test_wrong_section_from_file", func(t *testing.T) {
		parser := NewParser()
		err := parser.ReadFromFile(inisample)

		if err != nil {
			t.Errorf("Error: %v", err)

		}

		_, sectionErr := parser.GetSection("ownerr")

		if sectionErr == nil {
			t.Errorf("ownerr section exists")
		}
	})
}

func TestWrongGet(t *testing.T) {

	t.Run("test_wrong_key_from_string", func(t *testing.T) {
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
		err := parser.ReadFromString(content)

		if err != nil {
			t.Errorf("Error: %v", err)

		}

		_, getErr := parser.Get("Owner", "namee")

		if getErr == nil {
			t.Errorf("Key exists")
		}
	})

	t.Run("test_wrong_key_from_file", func(t *testing.T) {
		parser := NewParser()
		err := parser.ReadFromFile(inisample)

		if err != nil {
			t.Errorf("Error: %v", err)

		}

		_, getErr := parser.Get("owner", "nameee")

		if getErr == nil {
			t.Errorf("Key exists")
		}
	})
}

func TestInvalidWriteToFile(t *testing.T) {
	t.Run("writeto_invalidfile_fromstring", func(t *testing.T) {
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
		got := parser.WriteToFile("./")
		// _, isFile := os.Stat("./")

		if got == nil {
			t.Errorf("File not created")

		}

	})

	t.Run("writetofile_fromfile", func(t *testing.T) {
		parser := NewParser()
		parser.ReadFromFile(inisample)
		got := parser.WriteToFile("./")

		if got == nil {
			t.Errorf("File wasn't created")

		}

	})

}
