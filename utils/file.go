package utils

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// =======================
// File existence / info
// =======================
func FileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// =======================
// Read / Write
// =======================
func ReadFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

func WriteFile(path string, data []byte, perm os.FileMode) error {
	return ioutil.WriteFile(path, data, perm)
}

// Append data to file
func AppendFile(path string, data []byte) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(data)
	return err
}

// =======================
// Copy / Move / Delete
// =======================
func CopyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return out.Sync()
}

func MoveFile(src, dst string) error {
	return os.Rename(src, dst)
}

func DeleteFile(path string) error {
	return os.Remove(path)
}

func DeleteDir(path string) error {
	return os.RemoveAll(path)
}

// =======================
// Temp file / directory
// =======================
func TempFile(prefix string) (*os.File, error) {
	return ioutil.TempFile("", prefix)
}

func TempDir(prefix string) (string, error) {
	return ioutil.TempDir("", prefix)
}

// =======================
// Byte <-> File / io
// =======================
func BytesToReader(data []byte) io.Reader {
	return bytes.NewReader(data)
}

func BytesToReadCloser(data []byte) io.ReadCloser {
	return io.NopCloser(bytes.NewReader(data))
}

func FileToBytes(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

// =======================
// List files in directory
// =======================
func ListFiles(dir string) ([]string, error) {
	var files []string
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		return files, err
	}
	for _, e := range entries {
		if !e.IsDir() {
			files = append(files, filepath.Join(dir, e.Name()))
		}
	}
	return files, nil
}

// =======================
// Helpers
// =======================
func EnsureDir(path string) error {
	if !IsDir(path) {
		return os.MkdirAll(path, 0755)
	}
	return nil
}
