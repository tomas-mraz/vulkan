//go:build android

package vulkan

/*
#cgo android LDFLAGS: -Wl,--no-warn-mismatch
#cgo android CFLAGS: -DVK_USE_PLATFORM_ANDROID_KHR

#include <android/native_window.h>

#include "headers/vulkan.h"
#include "vk_wrapper.h"
#include "vk_bridge.h"
#include "cgo_helpers.h"
*/
import "C"
import (
	"fmt"
	"runtime"
	"unsafe"
)

const (
	// UsePlatformAndroid as defined in https://www.khronos.org/registry/vulkan/specs/1.0-wsi_extensions/xhtml/vkspec.html
	UsePlatformAndroid = 1
	// KhrAndroidSurface as defined in vulkan_android.h:22
	KhrAndroidSurface = 1
	// KhrAndroidSurfaceSpecVersion as defined in vulkan_android.h:24
	KhrAndroidSurfaceSpecVersion = 6
	// KhrAndroidSurfaceExtensionName as defined in vulkan_android.h:25
	KhrAndroidSurfaceExtensionName = "VK_KHR_android_surface"
	// AndroidExternalMemoryAndroidHardwareBuffer as defined in vulkan_android.h:45
	AndroidExternalMemoryAndroidHardwareBuffer = 1
	// AndroidExternalMemoryAndroidHardwareBufferSpecVersion as defined in vulkan_android.h:47
	AndroidExternalMemoryAndroidHardwareBufferSpecVersion = 5
	// AndroidExternalMemoryAndroidHardwareBufferExtensionName as defined in vulkan_android.h:48
	AndroidExternalMemoryAndroidHardwareBufferExtensionName = "VK_ANDROID_external_memory_android_hardware_buffer"
)

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

// GetRequiredInstanceExtensions should be used to query instance extensions required for surface initialization.
func GetRequiredInstanceExtensions() []string {
	return []string{
		"VK_KHR_surface\x00",
		"VK_KHR_android_surface\x00",
	}
}

// CreateAndroidSurface function as declared in https://www.khronos.org/registry/vulkan/specs/1.0-wsi_extensions/xhtml/vkspec.html#vkCreateAndroidSurfaceKHR
func CreateAndroidSurface(instance Instance, pCreateInfo *AndroidSurfaceCreateInfo, pAllocator *AllocationCallbacks, pSurface *Surface) Result {
	cinstance, cinstanceAllocMap := *(*C.VkInstance)(unsafe.Pointer(&instance)), cgoAllocsUnknown
	cpCreateInfo, cpCreateInfoAllocMap := pCreateInfo.PassRef()
	cpAllocator, cpAllocatorAllocMap := (*C.VkAllocationCallbacks)(unsafe.Pointer(pAllocator)), cgoAllocsUnknown
	cpSurface, cpSurfaceAllocMap := (*C.VkSurfaceKHR)(unsafe.Pointer(pSurface)), cgoAllocsUnknown
	__ret := C.callVkCreateAndroidSurfaceKHR(cinstance, cpCreateInfo, cpAllocator, cpSurface)
	runtime.KeepAlive(cpSurfaceAllocMap)
	runtime.KeepAlive(cpAllocatorAllocMap)
	runtime.KeepAlive(cpCreateInfoAllocMap)
	runtime.KeepAlive(cinstanceAllocMap)
	__v := (Result)(__ret)
	return __v
}

// ANativeWindow as declared in vulkan_android.h:23
type ANativeWindow C.ANativeWindow

// AHardwareBuffer as declared in vulkan_android.h:46
type AHardwareBuffer C.struct_AHardwareBuffer

// AndroidSurfaceCreateFlags type as declared in https://www.khronos.org/registry/vulkan/specs/1.0-wsi_extensions/xhtml/vkspec.html#VkAndroidSurfaceCreateFlagsKHR
type AndroidSurfaceCreateFlags uint32

// AndroidSurfaceCreateInfo as declared in https://www.khronos.org/registry/vulkan/specs/1.0-wsi_extensions/xhtml/vkspec.html#VkAndroidSurfaceCreateInfoKHR
type AndroidSurfaceCreateInfo struct {
	SType          StructureType
	PNext          unsafe.Pointer
	Flags          AndroidSurfaceCreateFlags
	Window         *ANativeWindow
	refeca5c35c    *C.VkAndroidSurfaceCreateInfoKHR
	allocseca5c35c interface{}
}

// AndroidHardwareBufferUsageANDROID as declared in https://www.khronos.org/registry/vulkan/specs/1.0/man/html/VkAndroidHardwareBufferUsageANDROID.html
type AndroidHardwareBufferUsageANDROID struct {
	SType                      StructureType
	PNext                      unsafe.Pointer
	AndroidHardwareBufferUsage uint64
	refcbdd253f                *C.VkAndroidHardwareBufferUsageANDROID
	allocscbdd253f             interface{}
}

// AndroidHardwareBufferPropertiesANDROID as declared in https://www.khronos.org/registry/vulkan/specs/1.0/man/html/VkAndroidHardwareBufferPropertiesANDROID.html
type AndroidHardwareBufferPropertiesANDROID struct {
	SType          StructureType
	PNext          unsafe.Pointer
	AllocationSize DeviceSize
	MemoryTypeBits uint32
	ref9506a7d8    *C.VkAndroidHardwareBufferPropertiesANDROID
	allocs9506a7d8 interface{}
}

// AndroidHardwareBufferFormatPropertiesANDROID as declared in https://www.khronos.org/registry/vulkan/specs/1.0/man/html/VkAndroidHardwareBufferFormatPropertiesANDROID.html
type AndroidHardwareBufferFormatPropertiesANDROID struct {
	SType                            StructureType
	PNext                            unsafe.Pointer
	Format                           Format
	ExternalFormat                   uint64
	FormatFeatures                   FormatFeatureFlags
	SamplerYcbcrConversionComponents ComponentMapping
	SuggestedYcbcrModel              SamplerYcbcrModelConversion
	SuggestedYcbcrRange              SamplerYcbcrRange
	SuggestedXChromaOffset           ChromaLocation
	SuggestedYChromaOffset           ChromaLocation
	ref158f0702                      *C.VkAndroidHardwareBufferFormatPropertiesANDROID
	allocs158f0702                   interface{}
}

// ImportAndroidHardwareBufferInfoANDROID as declared in https://www.khronos.org/registry/vulkan/specs/1.0/man/html/VkImportAndroidHardwareBufferInfoANDROID.html
type ImportAndroidHardwareBufferInfoANDROID struct {
	SType         StructureType
	PNext         unsafe.Pointer
	Buffer        *AHardwareBuffer
	ref5d2b47d    *C.VkImportAndroidHardwareBufferInfoANDROID
	allocs5d2b47d interface{}
}

// MemoryGetAndroidHardwareBufferInfoANDROID as declared in https://www.khronos.org/registry/vulkan/specs/1.0/man/html/VkMemoryGetAndroidHardwareBufferInfoANDROID.html
type MemoryGetAndroidHardwareBufferInfoANDROID struct {
	SType          StructureType
	PNext          unsafe.Pointer
	Memory         DeviceMemory
	ref51a5f19d    *C.VkMemoryGetAndroidHardwareBufferInfoANDROID
	allocs51a5f19d interface{}
}

// ExternalFormatANDROID as declared in https://www.khronos.org/registry/vulkan/specs/1.0/man/html/VkExternalFormatANDROID.html
type ExternalFormatANDROID struct {
	SType          StructureType
	PNext          unsafe.Pointer
	ExternalFormat uint64
	ref82bc5095    *C.VkExternalFormatANDROID
	allocs82bc5095 interface{}
}

// AndroidHardwareBufferFormatProperties2ANDROID as declared in https://www.khronos.org/registry/vulkan/specs/1.0/man/html/VkAndroidHardwareBufferFormatProperties2ANDROID.html
type AndroidHardwareBufferFormatProperties2ANDROID struct {
	SType                            StructureType
	PNext                            unsafe.Pointer
	Format                           Format
	ExternalFormat                   uint64
	FormatFeatures                   FormatFeatureFlags2
	SamplerYcbcrConversionComponents ComponentMapping
	SuggestedYcbcrModel              SamplerYcbcrModelConversion
	SuggestedYcbcrRange              SamplerYcbcrRange
	SuggestedXChromaOffset           ChromaLocation
	SuggestedYChromaOffset           ChromaLocation
	refc454ae4c                      *C.VkAndroidHardwareBufferFormatProperties2ANDROID
	allocsc454ae4c                   interface{}
}

// Ref returns a reference to C object as it is.
func (x *ANativeWindow) Ref() *C.ANativeWindow {
	if x == nil {
		return nil
	}
	return (*C.ANativeWindow)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *ANativeWindow) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewANativeWindowRef converts the C object reference into a raw struct reference without wrapping.
func NewANativeWindowRef(ref unsafe.Pointer) *ANativeWindow {
	return (*ANativeWindow)(ref)
}

// NewANativeWindow allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewANativeWindow() *ANativeWindow {
	return (*ANativeWindow)(allocANativeWindowMemory(1))
}

// allocANativeWindowMemory allocates memory for type C.ANativeWindow in C.
// The caller is responsible for freeing the this memory via C.free.
func allocANativeWindowMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfANativeWindowValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfANativeWindowValue = unsafe.Sizeof([1]C.ANativeWindow{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *ANativeWindow) PassRef() *C.ANativeWindow {
	if x == nil {
		x = (*ANativeWindow)(allocANativeWindowMemory(1))
	}
	return (*C.ANativeWindow)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *AHardwareBuffer) Ref() *C.struct_AHardwareBuffer {
	if x == nil {
		return nil
	}
	return (*C.struct_AHardwareBuffer)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *AHardwareBuffer) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewAHardwareBufferRef converts the C object reference into a raw struct reference without wrapping.
func NewAHardwareBufferRef(ref unsafe.Pointer) *AHardwareBuffer {
	return (*AHardwareBuffer)(ref)
}

// NewAHardwareBuffer allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewAHardwareBuffer() *AHardwareBuffer {
	return (*AHardwareBuffer)(allocStruct_AHardwareBufferMemory(1))
}

// allocStruct_AHardwareBufferMemory allocates memory for type C.struct_AHardwareBuffer in C.
// The caller is responsible for freeing the this memory via C.free.
func allocStruct_AHardwareBufferMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfStruct_AHardwareBufferValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfStruct_AHardwareBufferValue = unsafe.Sizeof([1]C.struct_AHardwareBuffer{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *AHardwareBuffer) PassRef() *C.struct_AHardwareBuffer {
	if x == nil {
		x = (*AHardwareBuffer)(allocStruct_AHardwareBufferMemory(1))
	}
	return (*C.struct_AHardwareBuffer)(unsafe.Pointer(x))
}

// allocAndroidSurfaceCreateInfoMemory allocates memory for type C.VkAndroidSurfaceCreateInfoKHR in C.
// The caller is responsible for freeing the this memory via C.free.
func allocAndroidSurfaceCreateInfoMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfAndroidSurfaceCreateInfoValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfAndroidSurfaceCreateInfoValue = unsafe.Sizeof([1]C.VkAndroidSurfaceCreateInfoKHR{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *AndroidSurfaceCreateInfo) Ref() *C.VkAndroidSurfaceCreateInfoKHR {
	if x == nil {
		return nil
	}
	return x.refeca5c35c
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *AndroidSurfaceCreateInfo) Free() {
	if x != nil && x.allocseca5c35c != nil {
		x.allocseca5c35c.(*cgoAllocMap).Free()
		x.refeca5c35c = nil
	}
}

// NewAndroidSurfaceCreateInfoRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewAndroidSurfaceCreateInfoRef(ref unsafe.Pointer) *AndroidSurfaceCreateInfo {
	if ref == nil {
		return nil
	}
	obj := new(AndroidSurfaceCreateInfo)
	obj.refeca5c35c = (*C.VkAndroidSurfaceCreateInfoKHR)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *AndroidSurfaceCreateInfo) PassRef() (*C.VkAndroidSurfaceCreateInfoKHR, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refeca5c35c != nil {
		return x.refeca5c35c, nil
	}
	memeca5c35c := allocAndroidSurfaceCreateInfoMemory(1)
	refeca5c35c := (*C.VkAndroidSurfaceCreateInfoKHR)(memeca5c35c)
	allocseca5c35c := new(cgoAllocMap)
	allocseca5c35c.Add(memeca5c35c)

	var csType_allocs *cgoAllocMap
	refeca5c35c.sType, csType_allocs = (C.VkStructureType)(x.SType), cgoAllocsUnknown
	allocseca5c35c.Borrow(csType_allocs)

	var cpNext_allocs *cgoAllocMap
	refeca5c35c.pNext, cpNext_allocs = *(*unsafe.Pointer)(unsafe.Pointer(&x.PNext)), cgoAllocsUnknown
	allocseca5c35c.Borrow(cpNext_allocs)

	var cflags_allocs *cgoAllocMap
	refeca5c35c.flags, cflags_allocs = (C.VkAndroidSurfaceCreateFlagsKHR)(x.Flags), cgoAllocsUnknown
	allocseca5c35c.Borrow(cflags_allocs)

	var cwindow_allocs *cgoAllocMap
	refeca5c35c.window, cwindow_allocs = *(**C.ANativeWindow)(unsafe.Pointer(&x.Window)), cgoAllocsUnknown
	allocseca5c35c.Borrow(cwindow_allocs)

	x.refeca5c35c = refeca5c35c
	x.allocseca5c35c = allocseca5c35c
	return refeca5c35c, allocseca5c35c

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x AndroidSurfaceCreateInfo) PassValue() (C.VkAndroidSurfaceCreateInfoKHR, *cgoAllocMap) {
	if x.refeca5c35c != nil {
		return *x.refeca5c35c, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *AndroidSurfaceCreateInfo) Deref() {
	if x.refeca5c35c == nil {
		return
	}
	x.SType = (StructureType)(x.refeca5c35c.sType)
	x.PNext = (unsafe.Pointer)(unsafe.Pointer(x.refeca5c35c.pNext))
	x.Flags = (AndroidSurfaceCreateFlags)(x.refeca5c35c.flags)
	x.Window = (*ANativeWindow)(unsafe.Pointer(x.refeca5c35c.window))
}

// allocAndroidHardwareBufferUsageANDROIDMemory allocates memory for type C.VkAndroidHardwareBufferUsageANDROID in C.
// The caller is responsible for freeing the this memory via C.free.
func allocAndroidHardwareBufferUsageANDROIDMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfAndroidHardwareBufferUsageANDROIDValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfAndroidHardwareBufferUsageANDROIDValue = unsafe.Sizeof([1]C.VkAndroidHardwareBufferUsageANDROID{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *AndroidHardwareBufferUsageANDROID) Ref() *C.VkAndroidHardwareBufferUsageANDROID {
	if x == nil {
		return nil
	}
	return x.refcbdd253f
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *AndroidHardwareBufferUsageANDROID) Free() {
	if x != nil && x.allocscbdd253f != nil {
		x.allocscbdd253f.(*cgoAllocMap).Free()
		x.refcbdd253f = nil
	}
}

// NewAndroidHardwareBufferUsageANDROIDRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewAndroidHardwareBufferUsageANDROIDRef(ref unsafe.Pointer) *AndroidHardwareBufferUsageANDROID {
	if ref == nil {
		return nil
	}
	obj := new(AndroidHardwareBufferUsageANDROID)
	obj.refcbdd253f = (*C.VkAndroidHardwareBufferUsageANDROID)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *AndroidHardwareBufferUsageANDROID) PassRef() (*C.VkAndroidHardwareBufferUsageANDROID, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refcbdd253f != nil {
		return x.refcbdd253f, nil
	}
	memcbdd253f := allocAndroidHardwareBufferUsageANDROIDMemory(1)
	refcbdd253f := (*C.VkAndroidHardwareBufferUsageANDROID)(memcbdd253f)
	allocscbdd253f := new(cgoAllocMap)
	allocscbdd253f.Add(memcbdd253f)

	var csType_allocs *cgoAllocMap
	refcbdd253f.sType, csType_allocs = (C.VkStructureType)(x.SType), cgoAllocsUnknown
	allocscbdd253f.Borrow(csType_allocs)

	var cpNext_allocs *cgoAllocMap
	refcbdd253f.pNext, cpNext_allocs = *(*unsafe.Pointer)(unsafe.Pointer(&x.PNext)), cgoAllocsUnknown
	allocscbdd253f.Borrow(cpNext_allocs)

	var candroidHardwareBufferUsage_allocs *cgoAllocMap
	refcbdd253f.androidHardwareBufferUsage, candroidHardwareBufferUsage_allocs = (C.uint64_t)(x.AndroidHardwareBufferUsage), cgoAllocsUnknown
	allocscbdd253f.Borrow(candroidHardwareBufferUsage_allocs)

	x.refcbdd253f = refcbdd253f
	x.allocscbdd253f = allocscbdd253f
	return refcbdd253f, allocscbdd253f

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x AndroidHardwareBufferUsageANDROID) PassValue() (C.VkAndroidHardwareBufferUsageANDROID, *cgoAllocMap) {
	if x.refcbdd253f != nil {
		return *x.refcbdd253f, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *AndroidHardwareBufferUsageANDROID) Deref() {
	if x.refcbdd253f == nil {
		return
	}
	x.SType = (StructureType)(x.refcbdd253f.sType)
	x.PNext = (unsafe.Pointer)(unsafe.Pointer(x.refcbdd253f.pNext))
	x.AndroidHardwareBufferUsage = (uint64)(x.refcbdd253f.androidHardwareBufferUsage)
}

// allocAndroidHardwareBufferPropertiesANDROIDMemory allocates memory for type C.VkAndroidHardwareBufferPropertiesANDROID in C.
// The caller is responsible for freeing the this memory via C.free.
func allocAndroidHardwareBufferPropertiesANDROIDMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfAndroidHardwareBufferPropertiesANDROIDValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfAndroidHardwareBufferPropertiesANDROIDValue = unsafe.Sizeof([1]C.VkAndroidHardwareBufferPropertiesANDROID{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *AndroidHardwareBufferPropertiesANDROID) Ref() *C.VkAndroidHardwareBufferPropertiesANDROID {
	if x == nil {
		return nil
	}
	return x.ref9506a7d8
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *AndroidHardwareBufferPropertiesANDROID) Free() {
	if x != nil && x.allocs9506a7d8 != nil {
		x.allocs9506a7d8.(*cgoAllocMap).Free()
		x.ref9506a7d8 = nil
	}
}

// NewAndroidHardwareBufferPropertiesANDROIDRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewAndroidHardwareBufferPropertiesANDROIDRef(ref unsafe.Pointer) *AndroidHardwareBufferPropertiesANDROID {
	if ref == nil {
		return nil
	}
	obj := new(AndroidHardwareBufferPropertiesANDROID)
	obj.ref9506a7d8 = (*C.VkAndroidHardwareBufferPropertiesANDROID)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *AndroidHardwareBufferPropertiesANDROID) PassRef() (*C.VkAndroidHardwareBufferPropertiesANDROID, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref9506a7d8 != nil {
		return x.ref9506a7d8, nil
	}
	mem9506a7d8 := allocAndroidHardwareBufferPropertiesANDROIDMemory(1)
	ref9506a7d8 := (*C.VkAndroidHardwareBufferPropertiesANDROID)(mem9506a7d8)
	allocs9506a7d8 := new(cgoAllocMap)
	allocs9506a7d8.Add(mem9506a7d8)

	var csType_allocs *cgoAllocMap
	ref9506a7d8.sType, csType_allocs = (C.VkStructureType)(x.SType), cgoAllocsUnknown
	allocs9506a7d8.Borrow(csType_allocs)

	var cpNext_allocs *cgoAllocMap
	ref9506a7d8.pNext, cpNext_allocs = *(*unsafe.Pointer)(unsafe.Pointer(&x.PNext)), cgoAllocsUnknown
	allocs9506a7d8.Borrow(cpNext_allocs)

	var callocationSize_allocs *cgoAllocMap
	ref9506a7d8.allocationSize, callocationSize_allocs = (C.VkDeviceSize)(x.AllocationSize), cgoAllocsUnknown
	allocs9506a7d8.Borrow(callocationSize_allocs)

	var cmemoryTypeBits_allocs *cgoAllocMap
	ref9506a7d8.memoryTypeBits, cmemoryTypeBits_allocs = (C.uint32_t)(x.MemoryTypeBits), cgoAllocsUnknown
	allocs9506a7d8.Borrow(cmemoryTypeBits_allocs)

	x.ref9506a7d8 = ref9506a7d8
	x.allocs9506a7d8 = allocs9506a7d8
	return ref9506a7d8, allocs9506a7d8

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x AndroidHardwareBufferPropertiesANDROID) PassValue() (C.VkAndroidHardwareBufferPropertiesANDROID, *cgoAllocMap) {
	if x.ref9506a7d8 != nil {
		return *x.ref9506a7d8, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *AndroidHardwareBufferPropertiesANDROID) Deref() {
	if x.ref9506a7d8 == nil {
		return
	}
	x.SType = (StructureType)(x.ref9506a7d8.sType)
	x.PNext = (unsafe.Pointer)(unsafe.Pointer(x.ref9506a7d8.pNext))
	x.AllocationSize = (DeviceSize)(x.ref9506a7d8.allocationSize)
	x.MemoryTypeBits = (uint32)(x.ref9506a7d8.memoryTypeBits)
}

// allocAndroidHardwareBufferFormatPropertiesANDROIDMemory allocates memory for type C.VkAndroidHardwareBufferFormatPropertiesANDROID in C.
// The caller is responsible for freeing the this memory via C.free.
func allocAndroidHardwareBufferFormatPropertiesANDROIDMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfAndroidHardwareBufferFormatPropertiesANDROIDValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfAndroidHardwareBufferFormatPropertiesANDROIDValue = unsafe.Sizeof([1]C.VkAndroidHardwareBufferFormatPropertiesANDROID{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *AndroidHardwareBufferFormatPropertiesANDROID) Ref() *C.VkAndroidHardwareBufferFormatPropertiesANDROID {
	if x == nil {
		return nil
	}
	return x.ref158f0702
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *AndroidHardwareBufferFormatPropertiesANDROID) Free() {
	if x != nil && x.allocs158f0702 != nil {
		x.allocs158f0702.(*cgoAllocMap).Free()
		x.ref158f0702 = nil
	}
}

// NewAndroidHardwareBufferFormatPropertiesANDROIDRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewAndroidHardwareBufferFormatPropertiesANDROIDRef(ref unsafe.Pointer) *AndroidHardwareBufferFormatPropertiesANDROID {
	if ref == nil {
		return nil
	}
	obj := new(AndroidHardwareBufferFormatPropertiesANDROID)
	obj.ref158f0702 = (*C.VkAndroidHardwareBufferFormatPropertiesANDROID)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *AndroidHardwareBufferFormatPropertiesANDROID) PassRef() (*C.VkAndroidHardwareBufferFormatPropertiesANDROID, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref158f0702 != nil {
		return x.ref158f0702, nil
	}
	mem158f0702 := allocAndroidHardwareBufferFormatPropertiesANDROIDMemory(1)
	ref158f0702 := (*C.VkAndroidHardwareBufferFormatPropertiesANDROID)(mem158f0702)
	allocs158f0702 := new(cgoAllocMap)
	allocs158f0702.Add(mem158f0702)

	var csType_allocs *cgoAllocMap
	ref158f0702.sType, csType_allocs = (C.VkStructureType)(x.SType), cgoAllocsUnknown
	allocs158f0702.Borrow(csType_allocs)

	var cpNext_allocs *cgoAllocMap
	ref158f0702.pNext, cpNext_allocs = *(*unsafe.Pointer)(unsafe.Pointer(&x.PNext)), cgoAllocsUnknown
	allocs158f0702.Borrow(cpNext_allocs)

	var cformat_allocs *cgoAllocMap
	ref158f0702.format, cformat_allocs = (C.VkFormat)(x.Format), cgoAllocsUnknown
	allocs158f0702.Borrow(cformat_allocs)

	var cexternalFormat_allocs *cgoAllocMap
	ref158f0702.externalFormat, cexternalFormat_allocs = (C.uint64_t)(x.ExternalFormat), cgoAllocsUnknown
	allocs158f0702.Borrow(cexternalFormat_allocs)

	var cformatFeatures_allocs *cgoAllocMap
	ref158f0702.formatFeatures, cformatFeatures_allocs = (C.VkFormatFeatureFlags)(x.FormatFeatures), cgoAllocsUnknown
	allocs158f0702.Borrow(cformatFeatures_allocs)

	var csamplerYcbcrConversionComponents_allocs *cgoAllocMap
	ref158f0702.samplerYcbcrConversionComponents, csamplerYcbcrConversionComponents_allocs = x.SamplerYcbcrConversionComponents.PassValue()
	allocs158f0702.Borrow(csamplerYcbcrConversionComponents_allocs)

	var csuggestedYcbcrModel_allocs *cgoAllocMap
	ref158f0702.suggestedYcbcrModel, csuggestedYcbcrModel_allocs = (C.VkSamplerYcbcrModelConversion)(x.SuggestedYcbcrModel), cgoAllocsUnknown
	allocs158f0702.Borrow(csuggestedYcbcrModel_allocs)

	var csuggestedYcbcrRange_allocs *cgoAllocMap
	ref158f0702.suggestedYcbcrRange, csuggestedYcbcrRange_allocs = (C.VkSamplerYcbcrRange)(x.SuggestedYcbcrRange), cgoAllocsUnknown
	allocs158f0702.Borrow(csuggestedYcbcrRange_allocs)

	var csuggestedXChromaOffset_allocs *cgoAllocMap
	ref158f0702.suggestedXChromaOffset, csuggestedXChromaOffset_allocs = (C.VkChromaLocation)(x.SuggestedXChromaOffset), cgoAllocsUnknown
	allocs158f0702.Borrow(csuggestedXChromaOffset_allocs)

	var csuggestedYChromaOffset_allocs *cgoAllocMap
	ref158f0702.suggestedYChromaOffset, csuggestedYChromaOffset_allocs = (C.VkChromaLocation)(x.SuggestedYChromaOffset), cgoAllocsUnknown
	allocs158f0702.Borrow(csuggestedYChromaOffset_allocs)

	x.ref158f0702 = ref158f0702
	x.allocs158f0702 = allocs158f0702
	return ref158f0702, allocs158f0702

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x AndroidHardwareBufferFormatPropertiesANDROID) PassValue() (C.VkAndroidHardwareBufferFormatPropertiesANDROID, *cgoAllocMap) {
	if x.ref158f0702 != nil {
		return *x.ref158f0702, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *AndroidHardwareBufferFormatPropertiesANDROID) Deref() {
	if x.ref158f0702 == nil {
		return
	}
	x.SType = (StructureType)(x.ref158f0702.sType)
	x.PNext = (unsafe.Pointer)(unsafe.Pointer(x.ref158f0702.pNext))
	x.Format = (Format)(x.ref158f0702.format)
	x.ExternalFormat = (uint64)(x.ref158f0702.externalFormat)
	x.FormatFeatures = (FormatFeatureFlags)(x.ref158f0702.formatFeatures)
	x.SamplerYcbcrConversionComponents = *NewComponentMappingRef(unsafe.Pointer(&x.ref158f0702.samplerYcbcrConversionComponents))
	x.SuggestedYcbcrModel = (SamplerYcbcrModelConversion)(x.ref158f0702.suggestedYcbcrModel)
	x.SuggestedYcbcrRange = (SamplerYcbcrRange)(x.ref158f0702.suggestedYcbcrRange)
	x.SuggestedXChromaOffset = (ChromaLocation)(x.ref158f0702.suggestedXChromaOffset)
	x.SuggestedYChromaOffset = (ChromaLocation)(x.ref158f0702.suggestedYChromaOffset)
}

// allocImportAndroidHardwareBufferInfoANDROIDMemory allocates memory for type C.VkImportAndroidHardwareBufferInfoANDROID in C.
// The caller is responsible for freeing the this memory via C.free.
func allocImportAndroidHardwareBufferInfoANDROIDMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfImportAndroidHardwareBufferInfoANDROIDValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfImportAndroidHardwareBufferInfoANDROIDValue = unsafe.Sizeof([1]C.VkImportAndroidHardwareBufferInfoANDROID{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *ImportAndroidHardwareBufferInfoANDROID) Ref() *C.VkImportAndroidHardwareBufferInfoANDROID {
	if x == nil {
		return nil
	}
	return x.ref5d2b47d
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *ImportAndroidHardwareBufferInfoANDROID) Free() {
	if x != nil && x.allocs5d2b47d != nil {
		x.allocs5d2b47d.(*cgoAllocMap).Free()
		x.ref5d2b47d = nil
	}
}

// NewImportAndroidHardwareBufferInfoANDROIDRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewImportAndroidHardwareBufferInfoANDROIDRef(ref unsafe.Pointer) *ImportAndroidHardwareBufferInfoANDROID {
	if ref == nil {
		return nil
	}
	obj := new(ImportAndroidHardwareBufferInfoANDROID)
	obj.ref5d2b47d = (*C.VkImportAndroidHardwareBufferInfoANDROID)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *ImportAndroidHardwareBufferInfoANDROID) PassRef() (*C.VkImportAndroidHardwareBufferInfoANDROID, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref5d2b47d != nil {
		return x.ref5d2b47d, nil
	}
	mem5d2b47d := allocImportAndroidHardwareBufferInfoANDROIDMemory(1)
	ref5d2b47d := (*C.VkImportAndroidHardwareBufferInfoANDROID)(mem5d2b47d)
	allocs5d2b47d := new(cgoAllocMap)
	allocs5d2b47d.Add(mem5d2b47d)

	var csType_allocs *cgoAllocMap
	ref5d2b47d.sType, csType_allocs = (C.VkStructureType)(x.SType), cgoAllocsUnknown
	allocs5d2b47d.Borrow(csType_allocs)

	var cpNext_allocs *cgoAllocMap
	ref5d2b47d.pNext, cpNext_allocs = *(*unsafe.Pointer)(unsafe.Pointer(&x.PNext)), cgoAllocsUnknown
	allocs5d2b47d.Borrow(cpNext_allocs)

	var cbuffer_allocs *cgoAllocMap
	ref5d2b47d.buffer, cbuffer_allocs = *(**C.struct_AHardwareBuffer)(unsafe.Pointer(&x.Buffer)), cgoAllocsUnknown
	allocs5d2b47d.Borrow(cbuffer_allocs)

	x.ref5d2b47d = ref5d2b47d
	x.allocs5d2b47d = allocs5d2b47d
	return ref5d2b47d, allocs5d2b47d

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x ImportAndroidHardwareBufferInfoANDROID) PassValue() (C.VkImportAndroidHardwareBufferInfoANDROID, *cgoAllocMap) {
	if x.ref5d2b47d != nil {
		return *x.ref5d2b47d, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *ImportAndroidHardwareBufferInfoANDROID) Deref() {
	if x.ref5d2b47d == nil {
		return
	}
	x.SType = (StructureType)(x.ref5d2b47d.sType)
	x.PNext = (unsafe.Pointer)(unsafe.Pointer(x.ref5d2b47d.pNext))
	x.Buffer = (*AHardwareBuffer)(unsafe.Pointer(x.ref5d2b47d.buffer))
}

// allocMemoryGetAndroidHardwareBufferInfoANDROIDMemory allocates memory for type C.VkMemoryGetAndroidHardwareBufferInfoANDROID in C.
// The caller is responsible for freeing the this memory via C.free.
func allocMemoryGetAndroidHardwareBufferInfoANDROIDMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfMemoryGetAndroidHardwareBufferInfoANDROIDValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfMemoryGetAndroidHardwareBufferInfoANDROIDValue = unsafe.Sizeof([1]C.VkMemoryGetAndroidHardwareBufferInfoANDROID{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *MemoryGetAndroidHardwareBufferInfoANDROID) Ref() *C.VkMemoryGetAndroidHardwareBufferInfoANDROID {
	if x == nil {
		return nil
	}
	return x.ref51a5f19d
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *MemoryGetAndroidHardwareBufferInfoANDROID) Free() {
	if x != nil && x.allocs51a5f19d != nil {
		x.allocs51a5f19d.(*cgoAllocMap).Free()
		x.ref51a5f19d = nil
	}
}

// NewMemoryGetAndroidHardwareBufferInfoANDROIDRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewMemoryGetAndroidHardwareBufferInfoANDROIDRef(ref unsafe.Pointer) *MemoryGetAndroidHardwareBufferInfoANDROID {
	if ref == nil {
		return nil
	}
	obj := new(MemoryGetAndroidHardwareBufferInfoANDROID)
	obj.ref51a5f19d = (*C.VkMemoryGetAndroidHardwareBufferInfoANDROID)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *MemoryGetAndroidHardwareBufferInfoANDROID) PassRef() (*C.VkMemoryGetAndroidHardwareBufferInfoANDROID, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref51a5f19d != nil {
		return x.ref51a5f19d, nil
	}
	mem51a5f19d := allocMemoryGetAndroidHardwareBufferInfoANDROIDMemory(1)
	ref51a5f19d := (*C.VkMemoryGetAndroidHardwareBufferInfoANDROID)(mem51a5f19d)
	allocs51a5f19d := new(cgoAllocMap)
	allocs51a5f19d.Add(mem51a5f19d)

	var csType_allocs *cgoAllocMap
	ref51a5f19d.sType, csType_allocs = (C.VkStructureType)(x.SType), cgoAllocsUnknown
	allocs51a5f19d.Borrow(csType_allocs)

	var cpNext_allocs *cgoAllocMap
	ref51a5f19d.pNext, cpNext_allocs = *(*unsafe.Pointer)(unsafe.Pointer(&x.PNext)), cgoAllocsUnknown
	allocs51a5f19d.Borrow(cpNext_allocs)

	var cmemory_allocs *cgoAllocMap
	ref51a5f19d.memory, cmemory_allocs = *(*C.VkDeviceMemory)(unsafe.Pointer(&x.Memory)), cgoAllocsUnknown
	allocs51a5f19d.Borrow(cmemory_allocs)

	x.ref51a5f19d = ref51a5f19d
	x.allocs51a5f19d = allocs51a5f19d
	return ref51a5f19d, allocs51a5f19d

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x MemoryGetAndroidHardwareBufferInfoANDROID) PassValue() (C.VkMemoryGetAndroidHardwareBufferInfoANDROID, *cgoAllocMap) {
	if x.ref51a5f19d != nil {
		return *x.ref51a5f19d, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *MemoryGetAndroidHardwareBufferInfoANDROID) Deref() {
	if x.ref51a5f19d == nil {
		return
	}
	x.SType = (StructureType)(x.ref51a5f19d.sType)
	x.PNext = (unsafe.Pointer)(unsafe.Pointer(x.ref51a5f19d.pNext))
	x.Memory = *(*DeviceMemory)(unsafe.Pointer(&x.ref51a5f19d.memory))
}

// allocExternalFormatANDROIDMemory allocates memory for type C.VkExternalFormatANDROID in C.
// The caller is responsible for freeing the this memory via C.free.
func allocExternalFormatANDROIDMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfExternalFormatANDROIDValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfExternalFormatANDROIDValue = unsafe.Sizeof([1]C.VkExternalFormatANDROID{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *ExternalFormatANDROID) Ref() *C.VkExternalFormatANDROID {
	if x == nil {
		return nil
	}
	return x.ref82bc5095
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *ExternalFormatANDROID) Free() {
	if x != nil && x.allocs82bc5095 != nil {
		x.allocs82bc5095.(*cgoAllocMap).Free()
		x.ref82bc5095 = nil
	}
}

// NewExternalFormatANDROIDRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewExternalFormatANDROIDRef(ref unsafe.Pointer) *ExternalFormatANDROID {
	if ref == nil {
		return nil
	}
	obj := new(ExternalFormatANDROID)
	obj.ref82bc5095 = (*C.VkExternalFormatANDROID)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *ExternalFormatANDROID) PassRef() (*C.VkExternalFormatANDROID, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref82bc5095 != nil {
		return x.ref82bc5095, nil
	}
	mem82bc5095 := allocExternalFormatANDROIDMemory(1)
	ref82bc5095 := (*C.VkExternalFormatANDROID)(mem82bc5095)
	allocs82bc5095 := new(cgoAllocMap)
	allocs82bc5095.Add(mem82bc5095)

	var csType_allocs *cgoAllocMap
	ref82bc5095.sType, csType_allocs = (C.VkStructureType)(x.SType), cgoAllocsUnknown
	allocs82bc5095.Borrow(csType_allocs)

	var cpNext_allocs *cgoAllocMap
	ref82bc5095.pNext, cpNext_allocs = *(*unsafe.Pointer)(unsafe.Pointer(&x.PNext)), cgoAllocsUnknown
	allocs82bc5095.Borrow(cpNext_allocs)

	var cexternalFormat_allocs *cgoAllocMap
	ref82bc5095.externalFormat, cexternalFormat_allocs = (C.uint64_t)(x.ExternalFormat), cgoAllocsUnknown
	allocs82bc5095.Borrow(cexternalFormat_allocs)

	x.ref82bc5095 = ref82bc5095
	x.allocs82bc5095 = allocs82bc5095
	return ref82bc5095, allocs82bc5095

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x ExternalFormatANDROID) PassValue() (C.VkExternalFormatANDROID, *cgoAllocMap) {
	if x.ref82bc5095 != nil {
		return *x.ref82bc5095, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *ExternalFormatANDROID) Deref() {
	if x.ref82bc5095 == nil {
		return
	}
	x.SType = (StructureType)(x.ref82bc5095.sType)
	x.PNext = (unsafe.Pointer)(unsafe.Pointer(x.ref82bc5095.pNext))
	x.ExternalFormat = (uint64)(x.ref82bc5095.externalFormat)
}

// allocAndroidHardwareBufferFormatProperties2ANDROIDMemory allocates memory for type C.VkAndroidHardwareBufferFormatProperties2ANDROID in C.
// The caller is responsible for freeing the this memory via C.free.
func allocAndroidHardwareBufferFormatProperties2ANDROIDMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfAndroidHardwareBufferFormatProperties2ANDROIDValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfAndroidHardwareBufferFormatProperties2ANDROIDValue = unsafe.Sizeof([1]C.VkAndroidHardwareBufferFormatProperties2ANDROID{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *AndroidHardwareBufferFormatProperties2ANDROID) Ref() *C.VkAndroidHardwareBufferFormatProperties2ANDROID {
	if x == nil {
		return nil
	}
	return x.refc454ae4c
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *AndroidHardwareBufferFormatProperties2ANDROID) Free() {
	if x != nil && x.allocsc454ae4c != nil {
		x.allocsc454ae4c.(*cgoAllocMap).Free()
		x.refc454ae4c = nil
	}
}

// NewAndroidHardwareBufferFormatProperties2ANDROIDRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewAndroidHardwareBufferFormatProperties2ANDROIDRef(ref unsafe.Pointer) *AndroidHardwareBufferFormatProperties2ANDROID {
	if ref == nil {
		return nil
	}
	obj := new(AndroidHardwareBufferFormatProperties2ANDROID)
	obj.refc454ae4c = (*C.VkAndroidHardwareBufferFormatProperties2ANDROID)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *AndroidHardwareBufferFormatProperties2ANDROID) PassRef() (*C.VkAndroidHardwareBufferFormatProperties2ANDROID, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refc454ae4c != nil {
		return x.refc454ae4c, nil
	}
	memc454ae4c := allocAndroidHardwareBufferFormatProperties2ANDROIDMemory(1)
	refc454ae4c := (*C.VkAndroidHardwareBufferFormatProperties2ANDROID)(memc454ae4c)
	allocsc454ae4c := new(cgoAllocMap)
	allocsc454ae4c.Add(memc454ae4c)

	var csType_allocs *cgoAllocMap
	refc454ae4c.sType, csType_allocs = (C.VkStructureType)(x.SType), cgoAllocsUnknown
	allocsc454ae4c.Borrow(csType_allocs)

	var cpNext_allocs *cgoAllocMap
	refc454ae4c.pNext, cpNext_allocs = *(*unsafe.Pointer)(unsafe.Pointer(&x.PNext)), cgoAllocsUnknown
	allocsc454ae4c.Borrow(cpNext_allocs)

	var cformat_allocs *cgoAllocMap
	refc454ae4c.format, cformat_allocs = (C.VkFormat)(x.Format), cgoAllocsUnknown
	allocsc454ae4c.Borrow(cformat_allocs)

	var cexternalFormat_allocs *cgoAllocMap
	refc454ae4c.externalFormat, cexternalFormat_allocs = (C.uint64_t)(x.ExternalFormat), cgoAllocsUnknown
	allocsc454ae4c.Borrow(cexternalFormat_allocs)

	var cformatFeatures_allocs *cgoAllocMap
	refc454ae4c.formatFeatures, cformatFeatures_allocs = (C.VkFormatFeatureFlags2)(x.FormatFeatures), cgoAllocsUnknown
	allocsc454ae4c.Borrow(cformatFeatures_allocs)

	var csamplerYcbcrConversionComponents_allocs *cgoAllocMap
	refc454ae4c.samplerYcbcrConversionComponents, csamplerYcbcrConversionComponents_allocs = x.SamplerYcbcrConversionComponents.PassValue()
	allocsc454ae4c.Borrow(csamplerYcbcrConversionComponents_allocs)

	var csuggestedYcbcrModel_allocs *cgoAllocMap
	refc454ae4c.suggestedYcbcrModel, csuggestedYcbcrModel_allocs = (C.VkSamplerYcbcrModelConversion)(x.SuggestedYcbcrModel), cgoAllocsUnknown
	allocsc454ae4c.Borrow(csuggestedYcbcrModel_allocs)

	var csuggestedYcbcrRange_allocs *cgoAllocMap
	refc454ae4c.suggestedYcbcrRange, csuggestedYcbcrRange_allocs = (C.VkSamplerYcbcrRange)(x.SuggestedYcbcrRange), cgoAllocsUnknown
	allocsc454ae4c.Borrow(csuggestedYcbcrRange_allocs)

	var csuggestedXChromaOffset_allocs *cgoAllocMap
	refc454ae4c.suggestedXChromaOffset, csuggestedXChromaOffset_allocs = (C.VkChromaLocation)(x.SuggestedXChromaOffset), cgoAllocsUnknown
	allocsc454ae4c.Borrow(csuggestedXChromaOffset_allocs)

	var csuggestedYChromaOffset_allocs *cgoAllocMap
	refc454ae4c.suggestedYChromaOffset, csuggestedYChromaOffset_allocs = (C.VkChromaLocation)(x.SuggestedYChromaOffset), cgoAllocsUnknown
	allocsc454ae4c.Borrow(csuggestedYChromaOffset_allocs)

	x.refc454ae4c = refc454ae4c
	x.allocsc454ae4c = allocsc454ae4c
	return refc454ae4c, allocsc454ae4c

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x AndroidHardwareBufferFormatProperties2ANDROID) PassValue() (C.VkAndroidHardwareBufferFormatProperties2ANDROID, *cgoAllocMap) {
	if x.refc454ae4c != nil {
		return *x.refc454ae4c, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *AndroidHardwareBufferFormatProperties2ANDROID) Deref() {
	if x.refc454ae4c == nil {
		return
	}
	x.SType = (StructureType)(x.refc454ae4c.sType)
	x.PNext = (unsafe.Pointer)(unsafe.Pointer(x.refc454ae4c.pNext))
	x.Format = (Format)(x.refc454ae4c.format)
	x.ExternalFormat = (uint64)(x.refc454ae4c.externalFormat)
	x.FormatFeatures = (FormatFeatureFlags2)(x.refc454ae4c.formatFeatures)
	x.SamplerYcbcrConversionComponents = *NewComponentMappingRef(unsafe.Pointer(&x.refc454ae4c.samplerYcbcrConversionComponents))
	x.SuggestedYcbcrModel = (SamplerYcbcrModelConversion)(x.refc454ae4c.suggestedYcbcrModel)
	x.SuggestedYcbcrRange = (SamplerYcbcrRange)(x.refc454ae4c.suggestedYcbcrRange)
	x.SuggestedXChromaOffset = (ChromaLocation)(x.refc454ae4c.suggestedXChromaOffset)
	x.SuggestedYChromaOffset = (ChromaLocation)(x.refc454ae4c.suggestedYChromaOffset)
}
