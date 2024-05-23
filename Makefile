.PHONY: help clean lib install win

help:
	@echo "make help"
	@echo "make clean"
	@echo "make win_lib"
	@echo "make win"
	@echo "make mac_lib"
	@echo "make mac"

clean:
	rm -rf build
	rm -rf tmp

win_lib:
	cd link/win64/bin && \
	gendef keystone.dll && \
	gendef capstone.dll && \
	dlltool --as-flags=--64 -m i386:x86-64 -k --output-lib libkeystone.a --input-def keystone.def && \
	dlltool --as-flags=--64 -m i386:x86-64 -k --output-lib libcapstone.a --input-def capstone.def && \
	mv -fv *.a ../lib/
	# install lib
	cp -fRv ./link/win64/include/* /usr/include/
	cp -fv ./link/win64/lib/*.a /usr/lib


win:
	@echo "Preinstall Cygwin or msys64_ucrt64"
	@echo "open shell"

	mkdir -p ./build
	cp -fv ./link/win64/bin/*.dll ./build/

	CGO_CFLAGS="-I/usr/include" CGO_LDFLAGS="-L/usr/lib -lcapstone -lkeystone" fyne package --release --target windows --icon ./theme/icons/asm2hex.png
	mv -fv *.exe ./build

mac_lib:
	rm -rf tmp
	mkdir -p ./tmp
	cd ./tmp && \
	git clone https://github.com/capstone-engine/capstone.git && \
	cd capstone && \
	git checkout 5.0.1 && \
	mkdir build && \
	cd build && \
	cmake .. && \
	sudo cmake --build . --config Release --target install && \
	cd .. && \
	rm -rf capstone && \
	git clone https://github.com/keystone-engine/keystone.git && \
	cd keystone && \
	git checkout 0.9.2 && \
	mkdir build && \
	cd build && \
	../make-lib.sh && \
	sudo make install && \
	cd .. && \
	rm -rf keystone && \
	cd .. && \
	rm -rf tmp

mac:
	mkdir -p ./build
	CGO_CFLAGS="-I/usr/local/include" CGO_LDFLAGS="-L/usr/local/lib -lcapstone -lkeystone" fyne package --release --target darwin --icon ./theme/icons/asm2hex.png
	rm -rf ./build/*.app
	mv -fv *.app ./build
