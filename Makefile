# default command build_win

clean:
	rm -rf build

lib:
	cd link/win64/bin && \
	gendef keystone.dll && \
	gendef capstone.dll && \
	x86_64-w64-mingw32-dlltool --as-flags=--64 -m i386:x86-64 -k --output-lib libkeystone.a --input-def keystone.def && \
	x86_64-w64-mingw32-dlltool --as-flags=--64 -m i386:x86-64 -k --output-lib libcapstone.a --input-def capstone.def && \
	mv -fv *.a ../lib

win:
	@echo "Preinstall Cygwin or msys64_ucrt64"
	@echo "open shell"

	cp -fv ./link/win64/lib/*.a .
	mkdir -p ./build
	mv -fv *.a /usr/lib

	cp -fRv ./link/win64/include/* /usr/include/
	cp -fv ./link/win64/bin/*.dll ./build/

	fyne package --release --target windows --icon ./theme/icons/asm2hex.png
	mv -fv *.exe ./build