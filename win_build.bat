@echo off
set CGO_CFLAGS="-ID:/works/asm2hex/link/win64/include"
set CGO_LDFLAGS="-LD:/works/asm2hex/link/win64/lib -lcapstone -lkeystone"
fyne package
echo "Build complete."