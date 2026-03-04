# Vulkan on Apple

Vulkan API is not officially supported by the Apple company as a manufacturer of the devices and operating systems macOS and iOS.  
Apple has its own graphics API called Metal.

There are two projects that are trying to solve this situation by providind layer translating Vulkan to Metal.  
[MoltenVK](https://github.com/KhronosGroup/MoltenVK) as a universal library and **KosmicKrisp** as MESA driver on Apple Silicon only.


## KosmicKrisp

KosmicKrisp is a new open-source Mesa 3D Vulkan-to-Metal driver designed for Apple Silicon (macOS 15+), achieving Vulkan 1.3 conformance.
It provides a high-performance alternative to MoltenVK by translating Vulkan commands to Metal 4.

https://www.lunarg.com/a-vulkan-on-metal-mesa-3d-graphics-driver  
https://vulkan.org/user/pages/09.events/vulkanised-2026/1545-Richard-Wright-LunarG.pdf  
https://www.phoronix.com/news/KosmicKrisp-2026  

## MoltenVK os macOS

MoltenVK provides a `MoltenVK.xcframework` which contains static libraries for all Apple platforms. Unfortuantely, linking with a xcframework outside of XCode is not possible.

Instead vulkan-go expects the dylibs to be present in the library search path.

Follow the [build instructions](https://github.com/KhronosGroup/MoltenVK#building), but instead of `make install` manually copy `./Package/Latest/MoltenVK/dylib/macOS/libMoltenVK.dylib` to `/usr/local/lib`

**IMPORTANT:** be sure to remove any existing `libMoltenVK.dylib` file *before* copying a new one, otherwise you'll have to reboot your computer due to the way the gatekeeper mechanism works!

## MoltenVK on iOS

The following steps are needed when developing for iOS and **not** using the `goki` tool. When using the `goki` tool, it will do all of these steps for you; you just need to run `goki setup ios` once to create the framework and then `goki build` will always copy the framework and set the environment variables for you. This information only exists for reference if you are not using the `goki` tool, and should not be relevant for most people.

Download the MoltenVK iOS asset from [the MoltenVK GitHub releases](https://github.com/KhronosGroup/MoltenVK/releases/latest/download/MoltenVK-ios.tar). Then, copy it to your `~/Library/goki` directory, and make a `.framework` by running:

```sh
install_name_tool -id @executable_path/MoltenVK.framework/MoltenVK libMoltenVK.dylib
lipo -create libMoltenVK.dylib -output MoltenVK
mkdir MoltenVK.framework
mv MoltenVK MoltenVK.framework
# now copy the Info.plist for MoltenVK.framework below into MoltenVK.framework/Info.plist
codesign --force --deep --verbose=2 --sign "rcoreilly@me.com" MoltenVK.framework
codesign -vvvv MoltenVK.framework
```

When building apps, build the app with the environment variable `CGO_LDFLAGS=-F/Users/{{you}}/Library/goki`. Then, after you build the app, run:
```
cp -r ~/Library/goki/MoltenVK.framework {{appname}}.app
```
For example:
```
cp -r ~/Library/goki/MoltenVK.framework drawtri.app
```

Info.plist for `MoltenVK.framework` (needs to be copied when making a framework above)

```
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>BuildMachineOSBuild</key>
	<string>22F82</string>
	<key>CFBundleDevelopmentRegion</key>
	<string>en</string>
	<key>CFBundleExecutable</key>
	<string>MoltenVK</string>
	<key>CFBundleIdentifier</key>
	<string>com.goki.MoltenVK</string>
	<key>CFBundleInfoDictionaryVersion</key>
	<string>6.0</string>
	<key>CFBundleName</key>
	<string>MoltenVK</string>
	<key>CFBundlePackageType</key>
	<string>FMWK</string>
	<key>CFBundleShortVersionString</key>
	<string>1.0</string>
	<key>CFBundleSupportedPlatforms</key>
	<array>
		<string>iPhoneOS</string>
	</array>
	<key>CFBundleVersion</key>
	<string>1</string>
</dict>
</plist>
```

main Info.plist for reference:

```
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>BuildMachineOSBuild</key>
	<string>22F82</string>
	<key>CFBundleDevelopmentRegion</key>
	<string>en</string>
	<key>CFBundleExecutable</key>
	<string>main</string>
	<key>CFBundleIdentifier</key>
	<string>com.example.test.widgets</string>
	<key>CFBundleInfoDictionaryVersion</key>
	<string>6.0</string>
	<key>CFBundleName</key>
	<string>Widgets</string>
	<key>CFBundlePackageType</key>
	<string>APPL</string>
	<key>CFBundleShortVersionString</key>
	<string>1.0</string>
	<key>CFBundleSignature</key>
	<string>????</string>
	<key>CFBundleSupportedPlatforms</key>
	<array>
		<string>iPhoneOS</string>
	</array>
	<key>CFBundleVersion</key>
	<string>1</string>
	<key>DTCompiler</key>
	<string>com.apple.compilers.llvm.clang.1_0</string>
	<key>DTPlatformBuild</key>
	<string>20E238</string>
	<key>DTPlatformName</key>
	<string>iphoneos</string>
	<key>DTPlatformVersion</key>
	<string>16.4</string>
	<key>DTSDKBuild</key>
	<string>20E238</string>
	<key>DTSDKName</key>
	<string>iphoneos16.4</string>
	<key>DTXcode</key>
	<string>1431</string>
	<key>DTXcodeBuild</key>
	<string>14E300c</string>
	<key>LSRequiresIPhoneOS</key>
	<true/>
	<key>MinimumOSVersion</key>
	<string>16.4</string>
	<key>UIDeviceFamily</key>
	<array>
		<integer>1</integer>
		<integer>2</integer>
	</array>
	<key>UILaunchStoryboardName</key>
	<string>LaunchScreen</string>
	<key>UIRequiredDeviceCapabilities</key>
	<array>
		<string>arm64</string>
	</array>
	<key>UISupportedInterfaceOrientations</key>
	<array>
		<string>UIInterfaceOrientationPortrait</string>
		<string>UIInterfaceOrientationLandscapeLeft</string>
		<string>UIInterfaceOrientationLandscapeRight</string>
	</array>
	<key>UISupportedInterfaceOrientations~ipad</key>
	<array>
		<string>UIInterfaceOrientationPortrait</string>
		<string>UIInterfaceOrientationPortraitUpsideDown</string>
		<string>UIInterfaceOrientationLandscapeLeft</string>
		<string>UIInterfaceOrientationLandscapeRight</string>
	</array>
</dict>
</plist>
```

