//go:build android

package vulkan

/*
#include "vk_wrapper.h"
#include "vk_bridge.h"
*/
import "C"
import "unsafe"

// CreateWindowSurface creates a Vulkan surface (VK_KHR_android_surface) for ANativeWindow from Android NDK.
func CreateWindowSurface(instance Instance, nativeWindow uintptr, pAllocator *AllocationCallbacks, pSurface *Surface) Result {
	cinstance, _ := *(*C.VkInstance)(unsafe.Pointer(&instance)), cgoAllocsUnknown
	cpAllocator, _ := (*C.VkAllocationCallbacks)(unsafe.Pointer(pAllocator)), cgoAllocsUnknown
	cpSurface, _ := (*C.VkSurfaceKHR)(unsafe.Pointer(pSurface)), cgoAllocsUnknown
	pCreateInfo := &AndroidSurfaceCreateInfo{
		SType:  StructureTypeAndroidSurfaceCreateInfo,
		Window: (*ANativeWindow)(unsafe.Pointer(nativeWindow)),
	}
	cpCreateInfo, _ := pCreateInfo.PassRef()
	__ret := C.callVkCreateAndroidSurfaceKHR(cinstance, cpCreateInfo, cpAllocator, cpSurface)
	__v := (Result)(__ret)
	return __v
}
