//go:build android && (368 || amd64)

package vulkan

// #cgo android CFLAGS: -DVK_USE_PLATFORM_ANDROID_KHR -arch amd64
import "C"
