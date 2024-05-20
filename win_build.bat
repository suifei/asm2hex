@echo off
set CGO_CFLAGS="-I.\bindings\include -O3 -g -Wall -Werror"
set CGO_LDFLAGS="-L.\bindings\lib -lcapstone -lkeystone -O3 -g"
fyne package
echo "Build complete."