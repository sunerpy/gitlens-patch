package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetVSCodeExtensionsDirectory_CursorServer(t *testing.T) {
	dir := GetVSCodeExtensionsDirectory(EditionCursorServer)
	assert.Equal(t, "/config/.cursor-server/extensions", dir)
}

func TestDetectAllVSCodeEnvironments_CursorServer(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("skip cursor-server symlink test in CI environment")
	}
	if err := os.MkdirAll("/config", 0755); err != nil {
		t.Fatalf("failed to create /config: %v", err)
	}
	tmpDir := "/tmp/test-cursor-server"
	if err := os.MkdirAll(tmpDir+"/extensions", 0755); err != nil {
		t.Fatalf("failed to create dir: %v", err)
	}
	defer func() {
		if err := os.RemoveAll(tmpDir); err != nil {
			t.Logf("cleanup failed: %v", err)
		}
	}()

	_ = os.Setenv("HOME", "/tmp")
	// 先删除已存在的软链接或目录
	if _, err := os.Lstat("/config/.cursor-server"); err == nil {
		_ = os.RemoveAll("/config/.cursor-server")
	}
	if err := os.Symlink(tmpDir, "/config/.cursor-server"); err != nil {
		t.Fatalf("failed to symlink: %v", err)
	}
	defer func() {
		if err := os.RemoveAll("/config/.cursor-server"); err != nil {
			t.Logf("cleanup symlink failed: %v", err)
		}
	}()

	envs := DetectAllVSCodeEnvironments()
	found := false
	for _, e := range envs {
		if e == EditionCursorServer {
			found = true
		}
	}
	assert.True(t, found, "should detect cursor-server edition")
}

func TestGetVSCodeExtensionsDirectory_Cursor(t *testing.T) {
	home := os.Getenv("HOME")
	if err := os.MkdirAll(home+"/.cursor/extensions", 0755); err != nil {
		t.Fatalf("failed to create dir: %v", err)
	}
	defer func() {
		if err := os.RemoveAll(home + "/.cursor"); err != nil {
			t.Logf("cleanup failed: %v", err)
		}
	}()
	dir := GetVSCodeExtensionsDirectory(EditionCursor)
	assert.Contains(t, dir, ".cursor/extensions")
}

func TestDetectAllVSCodeEnvironments_RemoteSSH(t *testing.T) {
	home := os.Getenv("HOME")
	if err := os.MkdirAll(home+"/.vscode-server/extensions", 0755); err != nil {
		t.Fatalf("failed to create dir: %v", err)
	}
	defer func() {
		if err := os.RemoveAll(home + "/.vscode-server"); err != nil {
			t.Logf("cleanup failed: %v", err)
		}
	}()
	envs := DetectAllVSCodeEnvironments()
	found := false
	for _, e := range envs {
		if e == EditionRemoteSSH {
			found = true
		}
	}
	assert.True(t, found, "should detect remote-ssh edition")
}
