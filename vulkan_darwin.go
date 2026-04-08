//go:build darwin && !ios

package vulkan

/*
#cgo darwin CFLAGS: -DVK_USE_PLATFORM_METAL_EXT -Wno-deprecated-declarations
#cgo darwin LDFLAGS: -Wl,-rpath,/usr/local/lib -F/Library/Frameworks -framework Cocoa -framework IOKit -framework IOSurface -framework QuartzCore -framework Metal -lvulkan

#include "headers/vulkan.h"
#include "vk_wrapper.h"
#include "vk_bridge.h"
*/
import "C"

const (
	// UsePlatformMacos means enabled support of Vulkan on macOS via Metal.
	UsePlatformMacos = 1
	// ExtMetalSurface means that VK_EXT_metal_surface is available.
	ExtMetalSurface = 1
	// ExtMetalSurfaceSpecVersion
	ExtMetalSurfaceSpecVersion = 1
	// ExtMetalSurfaceExtensionName
	ExtMetalSurfaceExtensionName = "VK_EXT_metal_surface"
)
