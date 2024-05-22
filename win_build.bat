@echo off
rem Set the GCC path to the mingw64 bin directory.
set PATH=%PATH%;D:\msys64\mingw64\bin
rem Set the CGO_CFLAGS and CGO_LDFLAGS to the capstone and keystone include and lib directories.
set CGO_CFLAGS="-ID:/works/asm2hex/link/win64/include"
set CGO_LDFLAGS="-LD:/works/asm2hex/link/win64/shared -lcapstone_dll -lkeystone_dll"
fyne package
rem If the build with Dnyamic Linking, copy the capstone.dll and keystone.dll to the output directory.
rem copy /Y link\win64\bin\capstone.dll .
rem copy /Y link\win64\bin\keystone.dll .
echo "Build complete."