<img src="https://cl.ly/2H2E3c0T1X16/Vulkan_500px_Mar15.png" width="200">

# Golang Bindings for Vulkan API ![version-1.3.241](https://img.shields.io/badge/version-1.3.241-red.svg) [![GoDoc](https://pkg.go.dev/badge/github.com/tomas-mraz/vulkan.svg)](https://pkg.go.dev/github.com/tomas-mraz/vulkan)

Package provides Go bindings for [Vulkan API](https://www.khronos.org/vulkan/) a low-overhead, cross-platform 3D graphics and compute API.  
The Vulkan API is a cross-platform industry standard enabling developers to target a wide range of devices with the same graphics API.

Original repository - https://github.com/vulkan-go/vulkan (from https://github.com/xlab)  
Forked and upgraded - https://github.com/goki/vulkan (from https://github.com/cogentcore)  
Forked goki fork - https://github.com/tomas-mraz/vulkan (from me :-)

See [UPDATING](UPDATING.md) for extensive notes on how to update to newer vulkan versions as they are released.  
See [README_Apple](README_Apple.md) for information about support for Vulkan on Apple platforms.

## Introduction


This Go binding allows one to use Vulkan API directly within Go code, avoiding adding lots of C/C++ in the projects.  
The original version is at https://github.com/vulkan-go/vulkan (still on 1.1.88 from 2018),  
and a fork at https://github.com/goki/vulkan (still on 1.3.239 from 2023) was abandoned in favor of the WebGPU wrapper in Rust.


## Examples and usage

The original author, `xlab`, has examples at: https://github.com/vulkan-go/demos and the beginnings of a toolkit at: https://github.com/vulkan-go/asche.

The updated version is being used extensively in the [goki](https://github.com/goki) framework, powering the [GoGi](https://github.com/goki/gi) 2D and 3D GUI framework, based on the [VGPU](https://github.com/goki/vgpu) toolkit that manages the considerable complexity of dealing with Vulkan.  
VGPU is also used as a GPU compute engine framework in the emergent neural network modeling framework [axon](https://github.com/emer/axon).

## How to use

Usage of this project is straightforward due to the stateless nature of Vulkan API.
Import the package like this:

```
import vk "github.com/tomas-mraz/vulkan"
```

Set the GetProcAddress pointer (used to look up Vulkan functions) using SetGetInstanceProcAddr or SetDefaultGetInstanceProcAddr. After that you can call Init to initialise the library. For example:

```
// Using SDL2:
vk.SetGetInstanceProcAddr(sdl.VulkanGetVkGetInstanceProcAddr())

// OR using GLFW:
vk.SetGetInstanceProcAddr(glfw.GetVulkanGetInstanceProcAddress())

// OR without using a windowing library (Linux only, recommended for compute-only tasks)
if err := vk.SetDefaultGetInstanceProcAddr(); err != nil {
    panic(err)
}

if err := vk.Init(); err != nil {
    panic(err)
}
```

And you're set. I must warn you that using the API properly is not an easy task at all, so beware and follow the official documentation: https://www.khronos.org/registry/vulkan/specs/1.0/html/vkspec.html

To simplify development, I created a high-level framework that manages Vulkan platform state and initialization. It is called [asche](https://github.com/vulkan-go/asche) because when you throw a gopher into volcano you get a pile of ash. Currently it's used in [VulkanCube](https://github.com/vulkan-go/demos/blob/master/vulkancube/vulkancube_android/main.go) demo app.



## Validation Layers

A good brief of the current state of Vulkan validation layers: [Explore the Vulkan Loader and Validation Layers](https://lunarg.com/wp-content/uploads/2016/07/lunarg-birds-feather-session-siggraph-july-26-2016.pdf) (PDF).

There is full support of validation layers with custom callbacks in Go. For my Android experiments I got the standard pack of layers from https://github.com/LunarG/VulkanTools and built them like this:

```
$ cd build-android
$ ./update_external_sources_android.sh
$ ./android-generate.sh
$ ndk-build
```

After that you'd copy the files to `android/jni/libs` in your project and activate the `ValidationLayers.mk` in your `Android.mk` so when building APK they will be copied alongside with your shared library. It just works then:

```
[INFO] Instance extensions: [VK_KHR_surface VK_KHR_android_surface]
[INFO] Instance layers: [VK_LAYER_LUNARG_screenshot VK_LAYER_GOOGLE_unique_objects VK_LAYER_LUNARG_api_dump VK_LAYER_LUNARG_image VK_LAYER_LUNARG_core_validation VK_LAYER_LUNARG_object_tracker VK_LAYER_GOOGLE_threading VK_LAYER_LUNARG_parameter_validation VK_LAYER_LUNARG_swapchain]

[Layer Swapchain][ERROR 4] The surface in pCreateInfo->surface, that was given to vkCreateSwapchainKHR(), must be a surface that is supported by the device as determined by vkGetPhysicalDeviceSurfaceSupportKHR().  However, vkGetPhysicalDeviceSurfaceSupportKHR() was never called with this surface.

[Layer Swapchain][ERROR 10] vkCreateSwapchainKHR() called with a non-supported pCreateInfo->compositeAlpha (i.e. VK_COMPOSITE_ALPHA_OPAQUE_BIT_KHR).  Supported values are:
     VK_COMPOSITE_ALPHA_INHERIT_BIT_KHR

[Layer DS][ERROR 8] Attempt to set lineWidth to 0.000000 but physical device wideLines feature not supported/enabled so lineWidth must be 1.0f!

[Layer DS][ERROR 22] Unable to allocate 2 descriptors of type VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER from pool 0x1c. This pool only has 1 descriptors of this type remaining.
```

## Useful links

* [vulkanGo.com](https://vulkanGo.com)
* [Sascha Willem's demos (C++)](https://github.com/SaschaWillems/Vulkan)
* [LunarG Vulkan Samples](https://github.com/LunarG/VulkanSamples) (archived)
* [Vulkan Samples from Khronos Group](https://github.com/KhronosGroup/Vulkan-Samples)
* [Official list of Vulkan resources](https://www.khronos.org/vulkan/resources)
* [API description](https://registry.khronos.org/vulkan/)
* [Vulkan API quick reference](https://www.khronos.org/registry/vulkan/specs/1.0/refguide/Vulkan-1.0-web.pdf)

## License

MIT
