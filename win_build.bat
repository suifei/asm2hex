@echo off
rem Set the GCC path to the mingw64 bin directory.
set PATH=D:\cygwin64\bin;%PATH%

rem Set the STATIC_LIB to "on" to build with static linking, or "off" to build with dynamic linking.
set STATIC_LIB="off"

set CGO_CFLAGS="-ID:\works\asm2hex\link\win64\include"
set CGO_CPPFLAGS="-ID:\works\asm2hex\link\win64\include"

rem Set the CGO_CFLAGS and CGO_LDFLAGS to the capstone and keystone include and lib directories.
if %STATIC_LIB% == "on" (
    set CGO_LDFLAGS="-LD:\works\asm2hex\link\win64\lib -lcapstone -lkeystone"
) else (
    set CGO_LDFLAGS="-LD:\works\asm2hex\link\win64\shared -lcapstone -lkeystone"
)
go build bindings\keystone\samples\main.go
fyne package
rem If the build with Dnyamic Linking, copy the capstone.dll and keystone.dll to the output directory.

if %STATIC_LIB% == "off" (
    copy /Y link\win64\bin\capstone.dll .
    copy /Y link\win64\bin\keystone.dll .
)

echo "Build complete."