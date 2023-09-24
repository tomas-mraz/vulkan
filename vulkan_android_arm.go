//go:build android && (arm || arm64)

package vulkan

// #cgo android CFLAGS: -DVK_USE_PLATFORM_ANDROID_KHR -arch arm64
import "C"
