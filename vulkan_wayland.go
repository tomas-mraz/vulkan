//go:build wayland

package vulkan

/*
#cgo LDFLAGS: -ldl
#cgo CFLAGS: -DVK_USE_PLATFORM_WAYLAND_KHR -Iwayland

#include <wayland-client.h>

#include "headers/vulkan.h"
#include "vk_wrapper.h"
#include "vk_bridge.h"
#include "cgo_helpers.h"
*/
import "C"
import "unsafe"

const (
	// KhrWaylandSurface as defined in headers/vulkan_wayland.h
	KhrWaylandSurface = 1
	// KhrWaylandSurfaceSpecVersion as defined in headers/vulkan_wayland.h
	KhrWaylandSurfaceSpecVersion = 6
	// KhrWaylandSurfaceExtensionName as defined in headers/vulkan_wayland.h
	KhrWaylandSurfaceExtensionName = "VK_KHR_wayland_surface"
)

// WaylandSurfaceCreateFlags type as declared in vulkan_wayland.h
type WaylandSurfaceCreateFlags uint32

// WaylandSurfaceCreateInfo as declared in vulkan_wayland.h
type WaylandSurfaceCreateInfo struct {
	SType   StructureType
	PNext   unsafe.Pointer
	Flags   WaylandSurfaceCreateFlags
	Display uintptr // *wl_display
	Surface uintptr // *wl_surface
}
