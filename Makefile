#   --target value, --os value         
#The operating system to target (android, android/arm, android/arm64, android/amd64, android/386, darwin, freebsd, ios, linux, netbsd, openbsd, windows)
release:
fyne package
hdiutil create -srcfolder "ASM to HEX Converter.app" -volname "ASM to HEX Converter" -fs HFS+ -fsargs "-c c=64,a=16,e=16" -format UDRW -size 45M "ASM2HEX_MacOS_ARM64.dmg"

win64:
set CGO_CFLAGS="-ID:\works\vcpkg\installed\x64-windows\include" 
set CGO_LDFLAGS="-LD:\works\vcpkg\installed\x64-windows\lib -lcapstone" 
fyne package
