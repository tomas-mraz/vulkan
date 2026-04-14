//go:build android

#include "vk_wrapper.h"
#include <dlfcn.h>
#include <android/log.h>
#define VKLOG(...) __android_log_print(ANDROID_LOG_INFO, "VkWrapper", __VA_ARGS__)

static PFN_vkGetInstanceProcAddr getInstanceProcAddress = NULL;

void setProcAddr(void* getProcAddr) {
    getInstanceProcAddress = (PFN_vkGetInstanceProcAddr)getProcAddr;
}

void setDefaultProcAddr() {
    void* libvulkan = dlopen("libvulkan.so", RTLD_NOW | RTLD_LOCAL);
    if (!libvulkan) {
        return;
    }
    getInstanceProcAddress = (PFN_vkGetInstanceProcAddr)(dlsym(libvulkan, "vkGetInstanceProcAddr"));
}

int isProcAddrSet() {
    return getInstanceProcAddress == NULL ? 0 : 1;
}

int vkInit(void) {
    if (!getInstanceProcAddress) {
        // Auto-init if not set explicitly
        setDefaultProcAddr();
        if (!getInstanceProcAddress) return -1;
    }

    // Global functions (instance = NULL)
    vgo_vkCreateInstance = (PFN_vkCreateInstance)(getInstanceProcAddress(NULL, "vkCreateInstance"));
    vgo_vkEnumerateInstanceVersion = (PFN_vkEnumerateInstanceVersion)(getInstanceProcAddress(NULL, "vkEnumerateInstanceVersion"));
    vgo_vkEnumerateInstanceExtensionProperties = (PFN_vkEnumerateInstanceExtensionProperties)(getInstanceProcAddress(NULL, "vkEnumerateInstanceExtensionProperties"));
    vgo_vkEnumerateInstanceLayerProperties = (PFN_vkEnumerateInstanceLayerProperties)(getInstanceProcAddress(NULL, "vkEnumerateInstanceLayerProperties"));
    vgo_vkGetInstanceProcAddr = getInstanceProcAddress;

    return 0;
}

int vkInitInstance(VkInstance instance) {
    PFN_vkGetInstanceProcAddr gpa = getInstanceProcAddress;
    VKLOG("vkInitInstance called, gpa=%p, instance=%p", gpa, instance);
    if (!gpa) return -1;

    // Core instance functions
    vgo_vkDestroyInstance = (PFN_vkDestroyInstance)(gpa(instance, "vkDestroyInstance"));
    vgo_vkEnumeratePhysicalDevices = (PFN_vkEnumeratePhysicalDevices)(gpa(instance, "vkEnumeratePhysicalDevices"));
    vgo_vkGetPhysicalDeviceFeatures = (PFN_vkGetPhysicalDeviceFeatures)(gpa(instance, "vkGetPhysicalDeviceFeatures"));
    vgo_vkGetPhysicalDeviceFormatProperties = (PFN_vkGetPhysicalDeviceFormatProperties)(gpa(instance, "vkGetPhysicalDeviceFormatProperties"));
    vgo_vkGetPhysicalDeviceImageFormatProperties = (PFN_vkGetPhysicalDeviceImageFormatProperties)(gpa(instance, "vkGetPhysicalDeviceImageFormatProperties"));
    vgo_vkGetPhysicalDeviceProperties = (PFN_vkGetPhysicalDeviceProperties)(gpa(instance, "vkGetPhysicalDeviceProperties"));
    vgo_vkGetPhysicalDeviceQueueFamilyProperties = (PFN_vkGetPhysicalDeviceQueueFamilyProperties)(gpa(instance, "vkGetPhysicalDeviceQueueFamilyProperties"));
    vgo_vkGetPhysicalDeviceMemoryProperties = (PFN_vkGetPhysicalDeviceMemoryProperties)(gpa(instance, "vkGetPhysicalDeviceMemoryProperties"));
    vgo_vkGetDeviceProcAddr = (PFN_vkGetDeviceProcAddr)(gpa(instance, "vkGetDeviceProcAddr"));
    vgo_vkCreateDevice = (PFN_vkCreateDevice)(gpa(instance, "vkCreateDevice"));
    vgo_vkDestroyDevice = (PFN_vkDestroyDevice)(gpa(instance, "vkDestroyDevice"));
    vgo_vkEnumerateDeviceExtensionProperties = (PFN_vkEnumerateDeviceExtensionProperties)(gpa(instance, "vkEnumerateDeviceExtensionProperties"));
    vgo_vkEnumerateDeviceLayerProperties = (PFN_vkEnumerateDeviceLayerProperties)(gpa(instance, "vkEnumerateDeviceLayerProperties"));
    vgo_vkGetDeviceQueue = (PFN_vkGetDeviceQueue)(gpa(instance, "vkGetDeviceQueue"));
    vgo_vkQueueSubmit = (PFN_vkQueueSubmit)(gpa(instance, "vkQueueSubmit"));
    vgo_vkQueueWaitIdle = (PFN_vkQueueWaitIdle)(gpa(instance, "vkQueueWaitIdle"));
    vgo_vkDeviceWaitIdle = (PFN_vkDeviceWaitIdle)(gpa(instance, "vkDeviceWaitIdle"));
    vgo_vkAllocateMemory = (PFN_vkAllocateMemory)(gpa(instance, "vkAllocateMemory"));
    vgo_vkFreeMemory = (PFN_vkFreeMemory)(gpa(instance, "vkFreeMemory"));
    vgo_vkMapMemory = (PFN_vkMapMemory)(gpa(instance, "vkMapMemory"));
    vgo_vkUnmapMemory = (PFN_vkUnmapMemory)(gpa(instance, "vkUnmapMemory"));
    vgo_vkFlushMappedMemoryRanges = (PFN_vkFlushMappedMemoryRanges)(gpa(instance, "vkFlushMappedMemoryRanges"));
    vgo_vkInvalidateMappedMemoryRanges = (PFN_vkInvalidateMappedMemoryRanges)(gpa(instance, "vkInvalidateMappedMemoryRanges"));
    vgo_vkGetDeviceMemoryCommitment = (PFN_vkGetDeviceMemoryCommitment)(gpa(instance, "vkGetDeviceMemoryCommitment"));
    vgo_vkBindBufferMemory = (PFN_vkBindBufferMemory)(gpa(instance, "vkBindBufferMemory"));
    vgo_vkBindImageMemory = (PFN_vkBindImageMemory)(gpa(instance, "vkBindImageMemory"));
    vgo_vkGetBufferMemoryRequirements = (PFN_vkGetBufferMemoryRequirements)(gpa(instance, "vkGetBufferMemoryRequirements"));
    vgo_vkGetImageMemoryRequirements = (PFN_vkGetImageMemoryRequirements)(gpa(instance, "vkGetImageMemoryRequirements"));
    vgo_vkGetImageSparseMemoryRequirements = (PFN_vkGetImageSparseMemoryRequirements)(gpa(instance, "vkGetImageSparseMemoryRequirements"));
    vgo_vkGetPhysicalDeviceSparseImageFormatProperties = (PFN_vkGetPhysicalDeviceSparseImageFormatProperties)(gpa(instance, "vkGetPhysicalDeviceSparseImageFormatProperties"));
    vgo_vkQueueBindSparse = (PFN_vkQueueBindSparse)(gpa(instance, "vkQueueBindSparse"));
    vgo_vkCreateFence = (PFN_vkCreateFence)(gpa(instance, "vkCreateFence"));
    vgo_vkDestroyFence = (PFN_vkDestroyFence)(gpa(instance, "vkDestroyFence"));
    vgo_vkResetFences = (PFN_vkResetFences)(gpa(instance, "vkResetFences"));
    vgo_vkGetFenceStatus = (PFN_vkGetFenceStatus)(gpa(instance, "vkGetFenceStatus"));
    vgo_vkWaitForFences = (PFN_vkWaitForFences)(gpa(instance, "vkWaitForFences"));
    vgo_vkCreateSemaphore = (PFN_vkCreateSemaphore)(gpa(instance, "vkCreateSemaphore"));
    vgo_vkDestroySemaphore = (PFN_vkDestroySemaphore)(gpa(instance, "vkDestroySemaphore"));
    vgo_vkCreateEvent = (PFN_vkCreateEvent)(gpa(instance, "vkCreateEvent"));
    vgo_vkDestroyEvent = (PFN_vkDestroyEvent)(gpa(instance, "vkDestroyEvent"));
    vgo_vkGetEventStatus = (PFN_vkGetEventStatus)(gpa(instance, "vkGetEventStatus"));
    vgo_vkSetEvent = (PFN_vkSetEvent)(gpa(instance, "vkSetEvent"));
    vgo_vkResetEvent = (PFN_vkResetEvent)(gpa(instance, "vkResetEvent"));
    vgo_vkCreateQueryPool = (PFN_vkCreateQueryPool)(gpa(instance, "vkCreateQueryPool"));
    vgo_vkDestroyQueryPool = (PFN_vkDestroyQueryPool)(gpa(instance, "vkDestroyQueryPool"));
    vgo_vkGetQueryPoolResults = (PFN_vkGetQueryPoolResults)(gpa(instance, "vkGetQueryPoolResults"));
    vgo_vkCreateBuffer = (PFN_vkCreateBuffer)(gpa(instance, "vkCreateBuffer"));
    vgo_vkDestroyBuffer = (PFN_vkDestroyBuffer)(gpa(instance, "vkDestroyBuffer"));
    vgo_vkCreateBufferView = (PFN_vkCreateBufferView)(gpa(instance, "vkCreateBufferView"));
    vgo_vkDestroyBufferView = (PFN_vkDestroyBufferView)(gpa(instance, "vkDestroyBufferView"));
    vgo_vkCreateImage = (PFN_vkCreateImage)(gpa(instance, "vkCreateImage"));
    vgo_vkDestroyImage = (PFN_vkDestroyImage)(gpa(instance, "vkDestroyImage"));
    vgo_vkGetImageSubresourceLayout = (PFN_vkGetImageSubresourceLayout)(gpa(instance, "vkGetImageSubresourceLayout"));
    vgo_vkCreateImageView = (PFN_vkCreateImageView)(gpa(instance, "vkCreateImageView"));
    vgo_vkDestroyImageView = (PFN_vkDestroyImageView)(gpa(instance, "vkDestroyImageView"));
    vgo_vkCreateShaderModule = (PFN_vkCreateShaderModule)(gpa(instance, "vkCreateShaderModule"));
    vgo_vkDestroyShaderModule = (PFN_vkDestroyShaderModule)(gpa(instance, "vkDestroyShaderModule"));
    vgo_vkCreatePipelineCache = (PFN_vkCreatePipelineCache)(gpa(instance, "vkCreatePipelineCache"));
    vgo_vkDestroyPipelineCache = (PFN_vkDestroyPipelineCache)(gpa(instance, "vkDestroyPipelineCache"));
    vgo_vkGetPipelineCacheData = (PFN_vkGetPipelineCacheData)(gpa(instance, "vkGetPipelineCacheData"));
    vgo_vkMergePipelineCaches = (PFN_vkMergePipelineCaches)(gpa(instance, "vkMergePipelineCaches"));
    vgo_vkCreateGraphicsPipelines = (PFN_vkCreateGraphicsPipelines)(gpa(instance, "vkCreateGraphicsPipelines"));
    vgo_vkCreateComputePipelines = (PFN_vkCreateComputePipelines)(gpa(instance, "vkCreateComputePipelines"));
    vgo_vkDestroyPipeline = (PFN_vkDestroyPipeline)(gpa(instance, "vkDestroyPipeline"));
    vgo_vkCreatePipelineLayout = (PFN_vkCreatePipelineLayout)(gpa(instance, "vkCreatePipelineLayout"));
    vgo_vkDestroyPipelineLayout = (PFN_vkDestroyPipelineLayout)(gpa(instance, "vkDestroyPipelineLayout"));
    vgo_vkCreateSampler = (PFN_vkCreateSampler)(gpa(instance, "vkCreateSampler"));
    vgo_vkDestroySampler = (PFN_vkDestroySampler)(gpa(instance, "vkDestroySampler"));
    vgo_vkCreateDescriptorSetLayout = (PFN_vkCreateDescriptorSetLayout)(gpa(instance, "vkCreateDescriptorSetLayout"));
    vgo_vkDestroyDescriptorSetLayout = (PFN_vkDestroyDescriptorSetLayout)(gpa(instance, "vkDestroyDescriptorSetLayout"));
    vgo_vkCreateDescriptorPool = (PFN_vkCreateDescriptorPool)(gpa(instance, "vkCreateDescriptorPool"));
    vgo_vkDestroyDescriptorPool = (PFN_vkDestroyDescriptorPool)(gpa(instance, "vkDestroyDescriptorPool"));
    vgo_vkResetDescriptorPool = (PFN_vkResetDescriptorPool)(gpa(instance, "vkResetDescriptorPool"));
    vgo_vkAllocateDescriptorSets = (PFN_vkAllocateDescriptorSets)(gpa(instance, "vkAllocateDescriptorSets"));
    vgo_vkFreeDescriptorSets = (PFN_vkFreeDescriptorSets)(gpa(instance, "vkFreeDescriptorSets"));
    vgo_vkUpdateDescriptorSets = (PFN_vkUpdateDescriptorSets)(gpa(instance, "vkUpdateDescriptorSets"));
    vgo_vkCreateFramebuffer = (PFN_vkCreateFramebuffer)(gpa(instance, "vkCreateFramebuffer"));
    vgo_vkDestroyFramebuffer = (PFN_vkDestroyFramebuffer)(gpa(instance, "vkDestroyFramebuffer"));
    vgo_vkCreateRenderPass = (PFN_vkCreateRenderPass)(gpa(instance, "vkCreateRenderPass"));
    vgo_vkDestroyRenderPass = (PFN_vkDestroyRenderPass)(gpa(instance, "vkDestroyRenderPass"));
    vgo_vkGetRenderAreaGranularity = (PFN_vkGetRenderAreaGranularity)(gpa(instance, "vkGetRenderAreaGranularity"));
    vgo_vkCreateCommandPool = (PFN_vkCreateCommandPool)(gpa(instance, "vkCreateCommandPool"));
    vgo_vkDestroyCommandPool = (PFN_vkDestroyCommandPool)(gpa(instance, "vkDestroyCommandPool"));
    vgo_vkResetCommandPool = (PFN_vkResetCommandPool)(gpa(instance, "vkResetCommandPool"));
    vgo_vkAllocateCommandBuffers = (PFN_vkAllocateCommandBuffers)(gpa(instance, "vkAllocateCommandBuffers"));
    vgo_vkFreeCommandBuffers = (PFN_vkFreeCommandBuffers)(gpa(instance, "vkFreeCommandBuffers"));
    vgo_vkBeginCommandBuffer = (PFN_vkBeginCommandBuffer)(gpa(instance, "vkBeginCommandBuffer"));
    vgo_vkEndCommandBuffer = (PFN_vkEndCommandBuffer)(gpa(instance, "vkEndCommandBuffer"));
    vgo_vkResetCommandBuffer = (PFN_vkResetCommandBuffer)(gpa(instance, "vkResetCommandBuffer"));
    vgo_vkCmdBindPipeline = (PFN_vkCmdBindPipeline)(gpa(instance, "vkCmdBindPipeline"));
    vgo_vkCmdSetViewport = (PFN_vkCmdSetViewport)(gpa(instance, "vkCmdSetViewport"));
    vgo_vkCmdSetScissor = (PFN_vkCmdSetScissor)(gpa(instance, "vkCmdSetScissor"));
    vgo_vkCmdSetLineWidth = (PFN_vkCmdSetLineWidth)(gpa(instance, "vkCmdSetLineWidth"));
    vgo_vkCmdSetDepthBias = (PFN_vkCmdSetDepthBias)(gpa(instance, "vkCmdSetDepthBias"));
    vgo_vkCmdSetBlendConstants = (PFN_vkCmdSetBlendConstants)(gpa(instance, "vkCmdSetBlendConstants"));
    vgo_vkCmdSetDepthBounds = (PFN_vkCmdSetDepthBounds)(gpa(instance, "vkCmdSetDepthBounds"));
    vgo_vkCmdSetStencilCompareMask = (PFN_vkCmdSetStencilCompareMask)(gpa(instance, "vkCmdSetStencilCompareMask"));
    vgo_vkCmdSetStencilWriteMask = (PFN_vkCmdSetStencilWriteMask)(gpa(instance, "vkCmdSetStencilWriteMask"));
    vgo_vkCmdSetStencilReference = (PFN_vkCmdSetStencilReference)(gpa(instance, "vkCmdSetStencilReference"));
    vgo_vkCmdBindDescriptorSets = (PFN_vkCmdBindDescriptorSets)(gpa(instance, "vkCmdBindDescriptorSets"));
    vgo_vkCmdBindIndexBuffer = (PFN_vkCmdBindIndexBuffer)(gpa(instance, "vkCmdBindIndexBuffer"));
    vgo_vkCmdBindVertexBuffers = (PFN_vkCmdBindVertexBuffers)(gpa(instance, "vkCmdBindVertexBuffers"));
    vgo_vkCmdDraw = (PFN_vkCmdDraw)(gpa(instance, "vkCmdDraw"));
    vgo_vkCmdDrawIndexed = (PFN_vkCmdDrawIndexed)(gpa(instance, "vkCmdDrawIndexed"));
    vgo_vkCmdDrawIndirect = (PFN_vkCmdDrawIndirect)(gpa(instance, "vkCmdDrawIndirect"));
    vgo_vkCmdDrawIndexedIndirect = (PFN_vkCmdDrawIndexedIndirect)(gpa(instance, "vkCmdDrawIndexedIndirect"));
    vgo_vkCmdDispatch = (PFN_vkCmdDispatch)(gpa(instance, "vkCmdDispatch"));
    vgo_vkCmdDispatchIndirect = (PFN_vkCmdDispatchIndirect)(gpa(instance, "vkCmdDispatchIndirect"));
    vgo_vkCmdCopyBuffer = (PFN_vkCmdCopyBuffer)(gpa(instance, "vkCmdCopyBuffer"));
    vgo_vkCmdCopyImage = (PFN_vkCmdCopyImage)(gpa(instance, "vkCmdCopyImage"));
    vgo_vkCmdBlitImage = (PFN_vkCmdBlitImage)(gpa(instance, "vkCmdBlitImage"));
    vgo_vkCmdCopyBufferToImage = (PFN_vkCmdCopyBufferToImage)(gpa(instance, "vkCmdCopyBufferToImage"));
    vgo_vkCmdCopyImageToBuffer = (PFN_vkCmdCopyImageToBuffer)(gpa(instance, "vkCmdCopyImageToBuffer"));
    vgo_vkCmdUpdateBuffer = (PFN_vkCmdUpdateBuffer)(gpa(instance, "vkCmdUpdateBuffer"));
    vgo_vkCmdFillBuffer = (PFN_vkCmdFillBuffer)(gpa(instance, "vkCmdFillBuffer"));
    vgo_vkCmdClearColorImage = (PFN_vkCmdClearColorImage)(gpa(instance, "vkCmdClearColorImage"));
    vgo_vkCmdClearDepthStencilImage = (PFN_vkCmdClearDepthStencilImage)(gpa(instance, "vkCmdClearDepthStencilImage"));
    vgo_vkCmdClearAttachments = (PFN_vkCmdClearAttachments)(gpa(instance, "vkCmdClearAttachments"));
    vgo_vkCmdResolveImage = (PFN_vkCmdResolveImage)(gpa(instance, "vkCmdResolveImage"));
    vgo_vkCmdSetEvent = (PFN_vkCmdSetEvent)(gpa(instance, "vkCmdSetEvent"));
    vgo_vkCmdResetEvent = (PFN_vkCmdResetEvent)(gpa(instance, "vkCmdResetEvent"));
    vgo_vkCmdWaitEvents = (PFN_vkCmdWaitEvents)(gpa(instance, "vkCmdWaitEvents"));
    vgo_vkCmdPipelineBarrier = (PFN_vkCmdPipelineBarrier)(gpa(instance, "vkCmdPipelineBarrier"));
    vgo_vkCmdBeginQuery = (PFN_vkCmdBeginQuery)(gpa(instance, "vkCmdBeginQuery"));
    vgo_vkCmdEndQuery = (PFN_vkCmdEndQuery)(gpa(instance, "vkCmdEndQuery"));
    vgo_vkCmdResetQueryPool = (PFN_vkCmdResetQueryPool)(gpa(instance, "vkCmdResetQueryPool"));
    vgo_vkCmdWriteTimestamp = (PFN_vkCmdWriteTimestamp)(gpa(instance, "vkCmdWriteTimestamp"));
    vgo_vkCmdCopyQueryPoolResults = (PFN_vkCmdCopyQueryPoolResults)(gpa(instance, "vkCmdCopyQueryPoolResults"));
    vgo_vkCmdPushConstants = (PFN_vkCmdPushConstants)(gpa(instance, "vkCmdPushConstants"));
    vgo_vkCmdBeginRenderPass = (PFN_vkCmdBeginRenderPass)(gpa(instance, "vkCmdBeginRenderPass"));
    vgo_vkCmdNextSubpass = (PFN_vkCmdNextSubpass)(gpa(instance, "vkCmdNextSubpass"));
    vgo_vkCmdEndRenderPass = (PFN_vkCmdEndRenderPass)(gpa(instance, "vkCmdEndRenderPass"));
    vgo_vkCmdExecuteCommands = (PFN_vkCmdExecuteCommands)(gpa(instance, "vkCmdExecuteCommands"));

    // KHR surface / swapchain
    vgo_vkDestroySurfaceKHR = (PFN_vkDestroySurfaceKHR)(gpa(instance, "vkDestroySurfaceKHR"));
    vgo_vkGetPhysicalDeviceSurfaceSupportKHR = (PFN_vkGetPhysicalDeviceSurfaceSupportKHR)(gpa(instance, "vkGetPhysicalDeviceSurfaceSupportKHR"));
    vgo_vkGetPhysicalDeviceSurfaceCapabilitiesKHR = (PFN_vkGetPhysicalDeviceSurfaceCapabilitiesKHR)(gpa(instance, "vkGetPhysicalDeviceSurfaceCapabilitiesKHR"));
    VKLOG("vkGetPhysicalDeviceSurfaceCapabilitiesKHR=%p", vgo_vkGetPhysicalDeviceSurfaceCapabilitiesKHR);
    vgo_vkGetPhysicalDeviceSurfaceFormatsKHR = (PFN_vkGetPhysicalDeviceSurfaceFormatsKHR)(gpa(instance, "vkGetPhysicalDeviceSurfaceFormatsKHR"));
    vgo_vkGetPhysicalDeviceSurfacePresentModesKHR = (PFN_vkGetPhysicalDeviceSurfacePresentModesKHR)(gpa(instance, "vkGetPhysicalDeviceSurfacePresentModesKHR"));
    vgo_vkCreateSwapchainKHR = (PFN_vkCreateSwapchainKHR)(gpa(instance, "vkCreateSwapchainKHR"));
    vgo_vkDestroySwapchainKHR = (PFN_vkDestroySwapchainKHR)(gpa(instance, "vkDestroySwapchainKHR"));
    vgo_vkGetSwapchainImagesKHR = (PFN_vkGetSwapchainImagesKHR)(gpa(instance, "vkGetSwapchainImagesKHR"));
    vgo_vkAcquireNextImageKHR = (PFN_vkAcquireNextImageKHR)(gpa(instance, "vkAcquireNextImageKHR"));
    vgo_vkQueuePresentKHR = (PFN_vkQueuePresentKHR)(gpa(instance, "vkQueuePresentKHR"));

    // KHR display
    vgo_vkGetPhysicalDeviceDisplayPropertiesKHR = (PFN_vkGetPhysicalDeviceDisplayPropertiesKHR)(gpa(instance, "vkGetPhysicalDeviceDisplayPropertiesKHR"));
    vgo_vkGetPhysicalDeviceDisplayPlanePropertiesKHR = (PFN_vkGetPhysicalDeviceDisplayPlanePropertiesKHR)(gpa(instance, "vkGetPhysicalDeviceDisplayPlanePropertiesKHR"));
    vgo_vkGetDisplayPlaneSupportedDisplaysKHR = (PFN_vkGetDisplayPlaneSupportedDisplaysKHR)(gpa(instance, "vkGetDisplayPlaneSupportedDisplaysKHR"));
    vgo_vkGetDisplayModePropertiesKHR = (PFN_vkGetDisplayModePropertiesKHR)(gpa(instance, "vkGetDisplayModePropertiesKHR"));
    vgo_vkCreateDisplayModeKHR = (PFN_vkCreateDisplayModeKHR)(gpa(instance, "vkCreateDisplayModeKHR"));
    vgo_vkGetDisplayPlaneCapabilitiesKHR = (PFN_vkGetDisplayPlaneCapabilitiesKHR)(gpa(instance, "vkGetDisplayPlaneCapabilitiesKHR"));
    vgo_vkCreateDisplayPlaneSurfaceKHR = (PFN_vkCreateDisplayPlaneSurfaceKHR)(gpa(instance, "vkCreateDisplayPlaneSurfaceKHR"));
    vgo_vkCreateSharedSwapchainsKHR = (PFN_vkCreateSharedSwapchainsKHR)(gpa(instance, "vkCreateSharedSwapchainsKHR"));

#ifdef VK_USE_PLATFORM_ANDROID_KHR
    vgo_vkCreateAndroidSurfaceKHR = (PFN_vkCreateAndroidSurfaceKHR)(gpa(instance, "vkCreateAndroidSurfaceKHR"));
#endif

    // EXT debug
    vgo_vkCreateDebugReportCallbackEXT = (PFN_vkCreateDebugReportCallbackEXT)(gpa(instance, "vkCreateDebugReportCallbackEXT"));
    vgo_vkDestroyDebugReportCallbackEXT = (PFN_vkDestroyDebugReportCallbackEXT)(gpa(instance, "vkDestroyDebugReportCallbackEXT"));
    vgo_vkDebugReportMessageEXT = (PFN_vkDebugReportMessageEXT)(gpa(instance, "vkDebugReportMessageEXT"));

    // GOOGLE display timing
    vgo_vkGetRefreshCycleDurationGOOGLE = (PFN_vkGetRefreshCycleDurationGOOGLE)(gpa(instance, "vkGetRefreshCycleDurationGOOGLE"));
    vgo_vkGetPastPresentationTimingGOOGLE = (PFN_vkGetPastPresentationTimingGOOGLE)(gpa(instance, "vkGetPastPresentationTimingGOOGLE"));

    // VK_KHR_buffer_device_address
    vgo_vkGetBufferDeviceAddress = (PFN_vkGetBufferDeviceAddress)(gpa(instance, "vkGetBufferDeviceAddress"));

    // VK_KHR_acceleration_structure
    vgo_vkCreateAccelerationStructureKHR = (PFN_vkCreateAccelerationStructureKHR)(gpa(instance, "vkCreateAccelerationStructureKHR"));
    vgo_vkDestroyAccelerationStructureKHR = (PFN_vkDestroyAccelerationStructureKHR)(gpa(instance, "vkDestroyAccelerationStructureKHR"));
    vgo_vkGetAccelerationStructureBuildSizesKHR = (PFN_vkGetAccelerationStructureBuildSizesKHR)(gpa(instance, "vkGetAccelerationStructureBuildSizesKHR"));
    vgo_vkGetAccelerationStructureDeviceAddressKHR = (PFN_vkGetAccelerationStructureDeviceAddressKHR)(gpa(instance, "vkGetAccelerationStructureDeviceAddressKHR"));
    vgo_vkCmdBuildAccelerationStructuresKHR = (PFN_vkCmdBuildAccelerationStructuresKHR)(gpa(instance, "vkCmdBuildAccelerationStructuresKHR"));

    // VK_KHR_ray_tracing_pipeline
    vgo_vkCreateRayTracingPipelinesKHR = (PFN_vkCreateRayTracingPipelinesKHR)(gpa(instance, "vkCreateRayTracingPipelinesKHR"));
    vgo_vkGetRayTracingShaderGroupHandlesKHR = (PFN_vkGetRayTracingShaderGroupHandlesKHR)(gpa(instance, "vkGetRayTracingShaderGroupHandlesKHR"));
    vgo_vkCmdTraceRaysKHR = (PFN_vkCmdTraceRaysKHR)(gpa(instance, "vkCmdTraceRaysKHR"));

    return 0;
}

// Global variable definitions
PFN_vkCreateInstance vgo_vkCreateInstance;
PFN_vkDestroyInstance vgo_vkDestroyInstance;
PFN_vkEnumerateInstanceVersion vgo_vkEnumerateInstanceVersion;
PFN_vkEnumeratePhysicalDevices vgo_vkEnumeratePhysicalDevices;
PFN_vkGetPhysicalDeviceFeatures vgo_vkGetPhysicalDeviceFeatures;
PFN_vkGetPhysicalDeviceFormatProperties vgo_vkGetPhysicalDeviceFormatProperties;
PFN_vkGetPhysicalDeviceImageFormatProperties vgo_vkGetPhysicalDeviceImageFormatProperties;
PFN_vkGetPhysicalDeviceProperties vgo_vkGetPhysicalDeviceProperties;
PFN_vkGetPhysicalDeviceQueueFamilyProperties vgo_vkGetPhysicalDeviceQueueFamilyProperties;
PFN_vkGetPhysicalDeviceMemoryProperties vgo_vkGetPhysicalDeviceMemoryProperties;
PFN_vkGetInstanceProcAddr vgo_vkGetInstanceProcAddr;
PFN_vkGetDeviceProcAddr vgo_vkGetDeviceProcAddr;
PFN_vkCreateDevice vgo_vkCreateDevice;
PFN_vkDestroyDevice vgo_vkDestroyDevice;
PFN_vkEnumerateInstanceExtensionProperties vgo_vkEnumerateInstanceExtensionProperties;
PFN_vkEnumerateDeviceExtensionProperties vgo_vkEnumerateDeviceExtensionProperties;
PFN_vkEnumerateInstanceLayerProperties vgo_vkEnumerateInstanceLayerProperties;
PFN_vkEnumerateDeviceLayerProperties vgo_vkEnumerateDeviceLayerProperties;
PFN_vkGetDeviceQueue vgo_vkGetDeviceQueue;
PFN_vkQueueSubmit vgo_vkQueueSubmit;
PFN_vkQueueWaitIdle vgo_vkQueueWaitIdle;
PFN_vkDeviceWaitIdle vgo_vkDeviceWaitIdle;
PFN_vkAllocateMemory vgo_vkAllocateMemory;
PFN_vkFreeMemory vgo_vkFreeMemory;
PFN_vkMapMemory vgo_vkMapMemory;
PFN_vkUnmapMemory vgo_vkUnmapMemory;
PFN_vkFlushMappedMemoryRanges vgo_vkFlushMappedMemoryRanges;
PFN_vkInvalidateMappedMemoryRanges vgo_vkInvalidateMappedMemoryRanges;
PFN_vkGetDeviceMemoryCommitment vgo_vkGetDeviceMemoryCommitment;
PFN_vkBindBufferMemory vgo_vkBindBufferMemory;
PFN_vkBindImageMemory vgo_vkBindImageMemory;
PFN_vkGetBufferMemoryRequirements vgo_vkGetBufferMemoryRequirements;
PFN_vkGetImageMemoryRequirements vgo_vkGetImageMemoryRequirements;
PFN_vkGetImageSparseMemoryRequirements vgo_vkGetImageSparseMemoryRequirements;
PFN_vkGetPhysicalDeviceSparseImageFormatProperties vgo_vkGetPhysicalDeviceSparseImageFormatProperties;
PFN_vkQueueBindSparse vgo_vkQueueBindSparse;
PFN_vkCreateFence vgo_vkCreateFence;
PFN_vkDestroyFence vgo_vkDestroyFence;
PFN_vkResetFences vgo_vkResetFences;
PFN_vkGetFenceStatus vgo_vkGetFenceStatus;
PFN_vkWaitForFences vgo_vkWaitForFences;
PFN_vkCreateSemaphore vgo_vkCreateSemaphore;
PFN_vkDestroySemaphore vgo_vkDestroySemaphore;
PFN_vkCreateEvent vgo_vkCreateEvent;
PFN_vkDestroyEvent vgo_vkDestroyEvent;
PFN_vkGetEventStatus vgo_vkGetEventStatus;
PFN_vkSetEvent vgo_vkSetEvent;
PFN_vkResetEvent vgo_vkResetEvent;
PFN_vkCreateQueryPool vgo_vkCreateQueryPool;
PFN_vkDestroyQueryPool vgo_vkDestroyQueryPool;
PFN_vkGetQueryPoolResults vgo_vkGetQueryPoolResults;
PFN_vkCreateBuffer vgo_vkCreateBuffer;
PFN_vkDestroyBuffer vgo_vkDestroyBuffer;
PFN_vkCreateBufferView vgo_vkCreateBufferView;
PFN_vkDestroyBufferView vgo_vkDestroyBufferView;
PFN_vkCreateImage vgo_vkCreateImage;
PFN_vkDestroyImage vgo_vkDestroyImage;
PFN_vkGetImageSubresourceLayout vgo_vkGetImageSubresourceLayout;
PFN_vkCreateImageView vgo_vkCreateImageView;
PFN_vkDestroyImageView vgo_vkDestroyImageView;
PFN_vkCreateShaderModule vgo_vkCreateShaderModule;
PFN_vkDestroyShaderModule vgo_vkDestroyShaderModule;
PFN_vkCreatePipelineCache vgo_vkCreatePipelineCache;
PFN_vkDestroyPipelineCache vgo_vkDestroyPipelineCache;
PFN_vkGetPipelineCacheData vgo_vkGetPipelineCacheData;
PFN_vkMergePipelineCaches vgo_vkMergePipelineCaches;
PFN_vkCreateGraphicsPipelines vgo_vkCreateGraphicsPipelines;
PFN_vkCreateComputePipelines vgo_vkCreateComputePipelines;
PFN_vkDestroyPipeline vgo_vkDestroyPipeline;
PFN_vkCreatePipelineLayout vgo_vkCreatePipelineLayout;
PFN_vkDestroyPipelineLayout vgo_vkDestroyPipelineLayout;
PFN_vkCreateSampler vgo_vkCreateSampler;
PFN_vkDestroySampler vgo_vkDestroySampler;
PFN_vkCreateDescriptorSetLayout vgo_vkCreateDescriptorSetLayout;
PFN_vkDestroyDescriptorSetLayout vgo_vkDestroyDescriptorSetLayout;
PFN_vkCreateDescriptorPool vgo_vkCreateDescriptorPool;
PFN_vkDestroyDescriptorPool vgo_vkDestroyDescriptorPool;
PFN_vkResetDescriptorPool vgo_vkResetDescriptorPool;
PFN_vkAllocateDescriptorSets vgo_vkAllocateDescriptorSets;
PFN_vkFreeDescriptorSets vgo_vkFreeDescriptorSets;
PFN_vkUpdateDescriptorSets vgo_vkUpdateDescriptorSets;
PFN_vkCreateFramebuffer vgo_vkCreateFramebuffer;
PFN_vkDestroyFramebuffer vgo_vkDestroyFramebuffer;
PFN_vkCreateRenderPass vgo_vkCreateRenderPass;
PFN_vkDestroyRenderPass vgo_vkDestroyRenderPass;
PFN_vkGetRenderAreaGranularity vgo_vkGetRenderAreaGranularity;
PFN_vkCreateCommandPool vgo_vkCreateCommandPool;
PFN_vkDestroyCommandPool vgo_vkDestroyCommandPool;
PFN_vkResetCommandPool vgo_vkResetCommandPool;
PFN_vkAllocateCommandBuffers vgo_vkAllocateCommandBuffers;
PFN_vkFreeCommandBuffers vgo_vkFreeCommandBuffers;
PFN_vkBeginCommandBuffer vgo_vkBeginCommandBuffer;
PFN_vkEndCommandBuffer vgo_vkEndCommandBuffer;
PFN_vkResetCommandBuffer vgo_vkResetCommandBuffer;
PFN_vkCmdBindPipeline vgo_vkCmdBindPipeline;
PFN_vkCmdSetViewport vgo_vkCmdSetViewport;
PFN_vkCmdSetScissor vgo_vkCmdSetScissor;
PFN_vkCmdSetLineWidth vgo_vkCmdSetLineWidth;
PFN_vkCmdSetDepthBias vgo_vkCmdSetDepthBias;
PFN_vkCmdSetBlendConstants vgo_vkCmdSetBlendConstants;
PFN_vkCmdSetDepthBounds vgo_vkCmdSetDepthBounds;
PFN_vkCmdSetStencilCompareMask vgo_vkCmdSetStencilCompareMask;
PFN_vkCmdSetStencilWriteMask vgo_vkCmdSetStencilWriteMask;
PFN_vkCmdSetStencilReference vgo_vkCmdSetStencilReference;
PFN_vkCmdBindDescriptorSets vgo_vkCmdBindDescriptorSets;
PFN_vkCmdBindIndexBuffer vgo_vkCmdBindIndexBuffer;
PFN_vkCmdBindVertexBuffers vgo_vkCmdBindVertexBuffers;
PFN_vkCmdDraw vgo_vkCmdDraw;
PFN_vkCmdDrawIndexed vgo_vkCmdDrawIndexed;
PFN_vkCmdDrawIndirect vgo_vkCmdDrawIndirect;
PFN_vkCmdDrawIndexedIndirect vgo_vkCmdDrawIndexedIndirect;
PFN_vkCmdDispatch vgo_vkCmdDispatch;
PFN_vkCmdDispatchIndirect vgo_vkCmdDispatchIndirect;
PFN_vkCmdCopyBuffer vgo_vkCmdCopyBuffer;
PFN_vkCmdCopyImage vgo_vkCmdCopyImage;
PFN_vkCmdBlitImage vgo_vkCmdBlitImage;
PFN_vkCmdCopyBufferToImage vgo_vkCmdCopyBufferToImage;
PFN_vkCmdCopyImageToBuffer vgo_vkCmdCopyImageToBuffer;
PFN_vkCmdUpdateBuffer vgo_vkCmdUpdateBuffer;
PFN_vkCmdFillBuffer vgo_vkCmdFillBuffer;
PFN_vkCmdClearColorImage vgo_vkCmdClearColorImage;
PFN_vkCmdClearDepthStencilImage vgo_vkCmdClearDepthStencilImage;
PFN_vkCmdClearAttachments vgo_vkCmdClearAttachments;
PFN_vkCmdResolveImage vgo_vkCmdResolveImage;
PFN_vkCmdSetEvent vgo_vkCmdSetEvent;
PFN_vkCmdResetEvent vgo_vkCmdResetEvent;
PFN_vkCmdWaitEvents vgo_vkCmdWaitEvents;
PFN_vkCmdPipelineBarrier vgo_vkCmdPipelineBarrier;
PFN_vkCmdBeginQuery vgo_vkCmdBeginQuery;
PFN_vkCmdEndQuery vgo_vkCmdEndQuery;
PFN_vkCmdResetQueryPool vgo_vkCmdResetQueryPool;
PFN_vkCmdWriteTimestamp vgo_vkCmdWriteTimestamp;
PFN_vkCmdCopyQueryPoolResults vgo_vkCmdCopyQueryPoolResults;
PFN_vkCmdPushConstants vgo_vkCmdPushConstants;
PFN_vkCmdBeginRenderPass vgo_vkCmdBeginRenderPass;
PFN_vkCmdNextSubpass vgo_vkCmdNextSubpass;
PFN_vkCmdEndRenderPass vgo_vkCmdEndRenderPass;
PFN_vkCmdExecuteCommands vgo_vkCmdExecuteCommands;
PFN_vkDestroySurfaceKHR vgo_vkDestroySurfaceKHR;
PFN_vkGetPhysicalDeviceSurfaceSupportKHR vgo_vkGetPhysicalDeviceSurfaceSupportKHR;
PFN_vkGetPhysicalDeviceSurfaceCapabilitiesKHR vgo_vkGetPhysicalDeviceSurfaceCapabilitiesKHR;
PFN_vkGetPhysicalDeviceSurfaceFormatsKHR vgo_vkGetPhysicalDeviceSurfaceFormatsKHR;
PFN_vkGetPhysicalDeviceSurfacePresentModesKHR vgo_vkGetPhysicalDeviceSurfacePresentModesKHR;
PFN_vkCreateSwapchainKHR vgo_vkCreateSwapchainKHR;
PFN_vkDestroySwapchainKHR vgo_vkDestroySwapchainKHR;
PFN_vkGetSwapchainImagesKHR vgo_vkGetSwapchainImagesKHR;
PFN_vkAcquireNextImageKHR vgo_vkAcquireNextImageKHR;
PFN_vkQueuePresentKHR vgo_vkQueuePresentKHR;
PFN_vkGetPhysicalDeviceDisplayPropertiesKHR vgo_vkGetPhysicalDeviceDisplayPropertiesKHR;
PFN_vkGetPhysicalDeviceDisplayPlanePropertiesKHR vgo_vkGetPhysicalDeviceDisplayPlanePropertiesKHR;
PFN_vkGetDisplayPlaneSupportedDisplaysKHR vgo_vkGetDisplayPlaneSupportedDisplaysKHR;
PFN_vkGetDisplayModePropertiesKHR vgo_vkGetDisplayModePropertiesKHR;
PFN_vkCreateDisplayModeKHR vgo_vkCreateDisplayModeKHR;
PFN_vkGetDisplayPlaneCapabilitiesKHR vgo_vkGetDisplayPlaneCapabilitiesKHR;
PFN_vkCreateDisplayPlaneSurfaceKHR vgo_vkCreateDisplayPlaneSurfaceKHR;
PFN_vkCreateSharedSwapchainsKHR vgo_vkCreateSharedSwapchainsKHR;

#ifdef VK_USE_PLATFORM_ANDROID_KHR
PFN_vkCreateAndroidSurfaceKHR vgo_vkCreateAndroidSurfaceKHR;
#endif

PFN_vkCreateDebugReportCallbackEXT vgo_vkCreateDebugReportCallbackEXT;
PFN_vkDestroyDebugReportCallbackEXT vgo_vkDestroyDebugReportCallbackEXT;
PFN_vkDebugReportMessageEXT vgo_vkDebugReportMessageEXT;

PFN_vkGetRefreshCycleDurationGOOGLE vgo_vkGetRefreshCycleDurationGOOGLE;
PFN_vkGetPastPresentationTimingGOOGLE vgo_vkGetPastPresentationTimingGOOGLE;

// VK_KHR_buffer_device_address
PFN_vkGetBufferDeviceAddress vgo_vkGetBufferDeviceAddress;

// VK_KHR_acceleration_structure
PFN_vkCreateAccelerationStructureKHR vgo_vkCreateAccelerationStructureKHR;
PFN_vkDestroyAccelerationStructureKHR vgo_vkDestroyAccelerationStructureKHR;
PFN_vkGetAccelerationStructureBuildSizesKHR vgo_vkGetAccelerationStructureBuildSizesKHR;
PFN_vkGetAccelerationStructureDeviceAddressKHR vgo_vkGetAccelerationStructureDeviceAddressKHR;
PFN_vkCmdBuildAccelerationStructuresKHR vgo_vkCmdBuildAccelerationStructuresKHR;

// VK_KHR_ray_tracing_pipeline
PFN_vkCreateRayTracingPipelinesKHR vgo_vkCreateRayTracingPipelinesKHR;
PFN_vkGetRayTracingShaderGroupHandlesKHR vgo_vkGetRayTracingShaderGroupHandlesKHR;
PFN_vkCmdTraceRaysKHR vgo_vkCmdTraceRaysKHR;
