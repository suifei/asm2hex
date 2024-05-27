.PHONY: help clean lib build

# 检测当前操作系统
ifeq ($(OS),Windows_NT)
    PLATFORM := Windows
else
    PLATFORM := $(shell uname)
endif

# 定义变量
CAPSTONE_VERSION := 5.0.1
KEYSTONE_VERSION := 0.9.2
INSTALL_PREFIX := /usr/local
CMAKE_FLAGS := -DCMAKE_INSTALL_PREFIX="$(INSTALL_PREFIX)"
MAKE_FLAGS := 
SUDO := 

ifeq ($(PLATFORM),Windows)
    CMAKE_FLAGS += -DBUILD_LIBS_ONLY=1 \
                   -DLLVM_BUILD_32_BITS=0 \
                   -DCMAKE_OSX_ARCHITECTURES="x86_64" \
                   -DCMAKE_BUILD_TYPE="Release" \
                   -DBUILD_SHARED_LIBS=0 \
                   -DLLVM_TARGETS_TO_BUILD="all" \
                   -G "Unix Makefiles"
    MAKE_FLAGS += -j8
    CGO_CFLAGS := -I/usr/local/include -O2 -Wall
    CGO_LDFLAGS := -L/usr/local/lib -static -lcapstone -lkeystone -lole32 -lshell32 -lkernel32 -lversion -luuid
    TARGET := windows
    KEYSTONE_BUILD_CMD := cmake $(CMAKE_FLAGS) .. && time make $(MAKE_FLAGS) && make install
else ifeq ($(PLATFORM),Darwin)
    SUDO := sudo
    CGO_CFLAGS := -I/usr/local/include -O2 -Wall
    CGO_LDFLAGS := -L/usr/local/lib -lcapstone -lkeystone
    TARGET := darwin
    KEYSTONE_BUILD_CMD := ../make-lib.sh && $(SUDO) make install
else
    $(error Unsupported platform: $(PLATFORM))
endif

help:
	@echo "Available targets:"
	@echo "  help        - Show this help message"
	@echo "  clean       - Clean up temporary files"
	@echo "  lib         - Build and install Capstone and Keystone libraries"
	@echo "  build       - Build the project"
	@echo "  lib_riscv   - Build and install Capstone and Keystone libraries for RISC-V"
	@echo "  build_riscv - Build the project for RISC-V"

clean:
	@rm -rf tmp build && \
	rm -rf $(INSTALL_PREFIX)/include/capstone && \
	rm -rf $(INSTALL_PREFIX)/include/keystone && \
	rm -f $(INSTALL_PREFIX)/lib/libcapstone* && \
	rm -f $(INSTALL_PREFIX)/lib/libkeystone* && \
	go clean -cache && \
	echo "Cleaned up successfully"

lib:
	@if [ -d "$(INSTALL_PREFIX)/lib" ] && \
	   [ -f "$(INSTALL_PREFIX)/lib/libcapstone.a" ] && \
	   [ -f "$(INSTALL_PREFIX)/lib/libkeystone.a" ]; then \
		echo "Capstone and Keystone libraries are already installed"; \
	else \
		mkdir -p ./tmp && \
		cd ./tmp && \
		if [ ! -d "capstone" ]; then \
			git clone https://github.com/capstone-engine/capstone.git && \
			cd capstone && \
			git checkout $(CAPSTONE_VERSION) && \
			mkdir build && \
			cd build && \
			cmake $(CMAKE_FLAGS) .. && \
			$(SUDO) cmake --build . --config Release --target install; \
		else \
			echo "Capstone library is already built, skipping build"; \
		fi && \
		cd ../.. && \
		if [ ! -d "keystone" ]; then \
			git clone https://github.com/keystone-engine/keystone.git && \
			cd keystone && \
			git checkout $(KEYSTONE_VERSION) && \
			mkdir build && \
			cd build && \
			$(KEYSTONE_BUILD_CMD); \
		else \
			echo "Keystone library is already built, skipping build"; \
		fi && \
		cd ../.. && \
		echo "Libraries installed successfully for $(PLATFORM)"; \
	fi


lib_riscv:
	@if [ -d "$(INSTALL_PREFIX)/lib" ] && \
	   [ -f "$(INSTALL_PREFIX)/lib/libcapstone.a" ] && \
	   [ -f "$(INSTALL_PREFIX)/lib/libkeystone.a" ]; then \
		echo "Capstone and Keystone libraries are already installed"; \
	else \
		mkdir -p ./tmp && \
		cd ./tmp && \
		if [ ! -d "capstone" ]; then \
			git clone https://github.com/capstone-engine/capstone.git && \
			cd capstone && \
			git checkout $(CAPSTONE_VERSION) && \
			mkdir build && \
			cd build && \
			cmake $(CMAKE_FLAGS) .. && \
			$(SUDO) cmake --build . --config Release --target install; \
		else \
			echo "Capstone library is already built, skipping build"; \
		fi && \
		cd ../.. && \
		if [ ! -d "keystone" ]; then \
			git clone https://github.com/null-cell/keystone.git keystone-riscv && \
			cd keystone-riscv && \
			git checkout 0.9.3.dev2 && \
			mkdir build && \
			cd build && \
			$(KEYSTONE_BUILD_CMD); \
		else \
			echo "Keystone library is already built, skipping build"; \
		fi && \
		cd ../.. && \
		echo "Libraries installed successfully for $(PLATFORM)"; \
	fi

build_riscv:
	@mkdir -p ./build-riscv && \
	CGO_ENABLED=1 \
	CGO_CFLAGS="$(CGO_CFLAGS)" \
	CGO_LDFLAGS="$(CGO_LDFLAGS)" \
	fyne package --release --target $(TARGET) --icon ./theme/icons/asm2hex.png -tags build_riscv && \
	rm -rf ./build-riscv/$(if $(filter $(PLATFORM),Windows),*.exe,*.app) && \
	mv -fv $(if $(filter $(PLATFORM),Windows),*.exe,*.app) ./build-riscv && \
	echo "Build completed for $(PLATFORM)"


build:
	@mkdir -p ./build && \
	CGO_ENABLED=1 \
	CGO_CFLAGS="$(CGO_CFLAGS)" \
	CGO_LDFLAGS="$(CGO_LDFLAGS)" \
	fyne package --release --target $(TARGET) --icon ./theme/icons/asm2hex.png && \
	rm -rf ./build/$(if $(filter $(PLATFORM),Windows),*.exe,*.app) && \
	mv -fv $(if $(filter $(PLATFORM),Windows),*.exe,*.app) ./build && \
	echo "Build completed for $(PLATFORM)"
