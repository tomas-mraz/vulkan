//go:build android

package vulkan

import "unsafe"

// CreateWindowSurface creates a Vulkan surface (VK_KHR_android_surface) for ANativeWindow from Android NDK.
func CreateWindowSurface(instance Instance, nativeWindow uintptr, pAllocator *AllocationCallbacks, pSurface *Surface) Result {
	pCreateInfo := &AndroidSurfaceCreateInfo{
		SType:  StructureTypeAndroidSurfaceCreateInfo,
		Window: (*ANativeWindow)(unsafe.Pointer(nativeWindow)),
	}
	return CreateAndroidSurface(instance, pCreateInfo, pAllocator, pSurface)
}
