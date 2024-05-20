@echo off
set CGO_CFLAGS="-ID:\works\asm2hex\bindings\include"
set CGO_LDFLAGS="-LD:\works\asm2hex\bindings\lib -lcapstone"
fyne package
echo "Build complete."