@echo off
set PATH=%PATH%;D:\msys64\mingw64\bin
set CGO_CFLAGS="-ID:/works/asm2hex/link/win64/include"
set CGO_LDFLAGS="-LD:/works/asm2hex/link/win64/lib -lcapstone -lkeystone"
fyne package
copy /Y link\win64\bin\capstone.dll .
copy /Y link\win64\bin\keystone.dll .
echo "Build complete."