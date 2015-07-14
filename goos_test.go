package goos

import (
	"testing"
)

func Test_readOSRelease_ubuntu_14(t *testing.T) {
	id, version_id, version, err := readOSRelease("./tests/os-release/ubuntu.14.04")
	t.Logf("id: >%s<", id)
	t.Logf("version_id: >%s<", version_id)
	t.Logf("version: >%s<", version)
	t.Logf("err: >%v<", err)
	if err != nil {
		t.Errorf("err: %v", err)
	}
	if id != "ubuntu" {
		t.Errorf("expected ubuntu: %s", id)
	}

}

func Test_centos_readRedhatRelease(t *testing.T) {
	id, version_id, version, err := readRedhatRelease("./tests/redhat-release/centos.6.4Final")
	t.Logf("id: >%s<", id)
	t.Logf("version_id: >%s<", version_id)
	t.Logf("version: >%s<", version)
	t.Logf("err: >%v<", err)

	if err != nil {
		t.Errorf("err: %v", err)
	}
	if id != "centos" {
		t.Errorf("expected centos: %s", id)
	}
}

func Test_busybox(t *testing.T) {
	id, version_id, version, err := busybox_version()
	t.Logf("id: >%s<", id)
	t.Logf("version_id: >%s<", version_id)
	t.Logf("version: >%s<", version)
	t.Logf("err: >%v<", err)

	if err != nil {
		t.Errorf("err: %v", err)
	}
	if id != "busybox" {
		t.Errorf("expected centos: %s", id)
	}
}
