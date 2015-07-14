package goos

import (
	"fmt"
	"regexp"
	"runtime"
	"strings"
)

func GetOSVersion() (ID string, VERSION_ID string, VERSION string, err error) {
	if runtime.GOOS == "windows" {
		return "windows", "unknown", "unknown", nil
	}

	ID, VERSION_ID, VERSION, err = readRedhatRelease("")
	if err == nil {
		return
	}

	ID, VERSION_ID, VERSION, err = readOSRelease("")
	if ID == "buildroot" {
		// maybe busybox
		ID, VERSION_ID, VERSION, err = busybox_version()
	}
	if err == nil {
		return
	}

	return "unknown", "unknown", "unknown", fmt.Errorf("unsupported os")
}

func readOSRelease(fpath string) (string, string, string, error) {
	var (
		ID         string
		VERSION_ID string
		VERSION    string
	)
	if fpath == "" {
		fpath = "/etc/os-release"
	}
	err := readLine(fpath, func(line string) error {
		if strings.HasPrefix(line, "ID=") {
			ID = strings.ToLower(line[3:])
		} else if strings.HasPrefix(line, "VERSION_ID=") {
			VERSION_ID = line[12 : len(line)-1]
		} else if strings.HasPrefix(line, "VERSION=") {
			VERSION = line[9 : len(line)-1]
		}
		return nil
	})
	return ID, VERSION_ID, VERSION, err
}

var (
	pat_ver_id = regexp.MustCompile(`\d+(.\d+)*`)
	pat_ver    = regexp.MustCompile(`\d+.*`)
)

func readRedhatRelease(fpath string) (string, string, string, error) {
	var (
		ID         string
		VERSION_ID string
		VERSION    string
	)
	if fpath == "" {
		fpath = "/etc/redhat-release"
	}
	err := readLine(fpath, func(line string) error {
		lines := strings.SplitN(line, " ", 2)
		if len(lines) == 2 {
			ID = strings.ToLower(lines[0])
			VERSION_ID = pat_ver_id.FindString(lines[1])
			VERSION = pat_ver.FindString(lines[1])
			return nil
		} else {
			return fmt.Errorf("invalid file format: %v, %d", lines, len(lines))
		}
	})
	return ID, VERSION_ID, VERSION, err
}

// # busybox | head -1
// BusyBox v1.22.1 (2014-05-22 23:22:11 UTC) multi-call binary.
func busybox_version() (string, string, string, error) {
	var (
		ID         string
		VERSION_ID string
		VERSION    string
	)
	first_line := ""
	err := ReadCommand(func(line string) error {
		if first_line == "" {
			first_line = line
			lines := strings.SplitN(first_line, " ", 3)
			if len(lines) == 3 {
				ID = strings.ToLower(lines[0])
				VERSION_ID = lines[1]
				VERSION = strings.Join(lines[1:], " ")
			}
		}
		return nil
	}, "busybox")
	return ID, VERSION_ID, VERSION, err
}
