.PHONY: help clean lib install win

help:
	@echo "make help"
	@echo "make clean"
	@echo "make lib"
	@echo "make install"
	@echo "make win"

clean:
	rm -rf build

lib:
	cd link/win64/bin && \
	gendef keystone.dll && \
	gendef capstone.dll && \
	dlltool --as-flags=--64 -m i386:x86-64 -k --output-lib libkeystone.a --input-def keystone.def && \
	dlltool --as-flags=--64 -m i386:x86-64 -k --output-lib libcapstone.a --input-def capstone.def && \
	mv -fv *.a ../lib/

install:
	cp -fRv ./link/win64/include/* /usr/include/
	cp -fv ./link/win64/lib/*.a /usr/lib


win:
	@echo "Preinstall Cygwin or msys64_ucrt64"
	@echo "open shell"

	mkdir -p ./build
	cp -fv ./link/win64/bin/*.dll ./build/

	CGO_CFLAGS="-I/usr/include" CGO_LDFLAGS="-L/usr/lib -lcapstone -lkeystone" fyne package --release --target windows --icon ./theme/icons/asm2hex.png
	mv -fv *.exe ./build