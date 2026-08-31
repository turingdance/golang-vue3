// Package appkit 提供应用主目录路径管理
//
// 目录结构:
//
//	{apphome}/           ← 用户数据目录/.insighthub/
//	├── models/          ← 模型文件存放
//	├── ffmpeg/          ← FFmpeg 二进制存放
//	└── ...
//
// Windows: C:\Users\{user}\.insighthub\
// Linux:   ~/.insighthub/
// macOS:   ~/.insighthub/
package appkit

import (
	"os"
	"path/filepath"
	"sync"
)

var (
	appHomeOnce sync.Once
	appHome     string
	appHomeErr  error
)

// AppHome 获取应用主目录
// Windows: {UserProfile}/.insighthub/
// Linux/macOS: ~/.insighthub/
func AppHome() (string, error) {
	appHomeOnce.Do(func() {
		home, err := os.UserHomeDir()
		if err != nil {
			home = "."
		}
		appHome = filepath.Join(home, ".insighthub")
	})
	return appHome, appHomeErr
}

// MustAppHome 获取应用主目录, 失败时 panic
func MustAppHome() string {
	dir, err := AppHome()
	if err != nil {
		panic(err)
	}
	return dir
}

// ModelsDir 获取模型根目录
// {apphome}/models/
func ModelsDir() (string, error) {
	home, err := AppHome()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, "models"), nil
}

// MustModelsDir 获取模型根目录, 失败时 panic
func MustModelsDir() string {
	dir, err := ModelsDir()
	if err != nil {
		panic(err)
	}
	return dir
}

// ModelDir 获取指定模型的缓存目录
// {apphome}/models/{modelID}/
func ModelDir(modelID string) (string, error) {
	modelsDir, err := ModelsDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(modelsDir, modelID), nil
}

// MustModelDir 获取指定模型的缓存目录, 失败时 panic
func MustModelDir(modelID string) string {
	dir, err := ModelDir(modelID)
	if err != nil {
		panic(err)
	}
	return dir
}

// FFmpegDir 获取 FFmpeg 存放目录
// {apphome}/ffmpeg/
func FFmpegDir() (string, error) {
	home, err := AppHome()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, "ffmpeg"), nil
}

// MustFFmpegDir 获取 FFmpeg 存放目录, 失败时 panic
func MustFFmpegDir() string {
	dir, err := FFmpegDir()
	if err != nil {
		panic(err)
	}
	return dir
}

// EnsureDir 确保目录存在
func EnsureDir(dir string) error {
	return os.MkdirAll(dir, 0o755)
}
