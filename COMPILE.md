# 编译指南

本指南将帮助你在 Windows 和 macOS 平台上编译 ASM2HEX 项目。

## Windows 平台

### 环境要求

- Windows 操作系统(Windows 7 或更高版本)
- MSYS2 或 Cygwin 环境
- MinGW-w64 工具链

**快速搭建环境：[Release V1.1](https://github.com/suifei/asm2hex/releases/tag/v1.1) 中下载 `msys64.7z`（Windows 10 的编译环境，请解压到 D 盘根目录，运行D:\msys64\mingw64.exe进入编译环境，切换到源码目录内，运行'make'，先编译lib，再build）**

国内下载网盘分流：
- 论坛: [pediy](https://bbs.kanxue.com/thread-281871.htm) [52pojie](https://www.52pojie.cn/thread-1927199-1-1.html)
- 下载: [国内地址1](https://pan.baidu.com/s/1EiXuE9UDfQrAtf4heFINHQ?pwd=52pj)
- 分流: [国内地址1](https://pan.baidu.com/s/1TgSNXi3-DZxg5lqaJiBeyA?pwd=8888)

### 配置步骤

1. 安装 MSYS2 或 Cygwin 环境。可以从官网下载安装包并按照提示进行安装。

2. 打开 MSYS2 或 Cygwin 终端。

3. 确保已安装 MinGW-w64 工具链。可以运行以下命令进行安装:
   ```
   pacman -S mingw-w64-x86_64-toolchain
   ```

4. 安装 Go 语言编译器。可以从 Go 官网下载安装包并按照提示进行安装。

5. 安装 Fyne 命令行工具。在终端中运行以下命令:
   ```
   go get fyne.io/fyne/v2/cmd/fyne
   ```

6. 克隆 ASM2HEX 项目到本地:
   ```
   git clone https://github.com/suifei/asm2hex.git
   ```

7. 进入项目目录:
   ```
   cd asm2hex
   ```

8. 编译 Keystone 和 Capstone 库:
   ```
   make clean
   make lib
   ```

9. 编译 ASM2HEX:
   ```
   make build
   ```

编译完成后,可执行文件将生成在 `build` 目录下。

## macOS 平台

### 环境要求

- macOS 操作系统(10.13 或更高版本)
- Xcode 命令行工具
- Homebrew 包管理器

### 配置步骤

1. 安装 Xcode 命令行工具。可以在终端中运行以下命令:
   ```
   xcode-select --install
   ```

2. 安装 Homebrew 包管理器。可以在终端中运行以下命令:
   ```
   /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
   ```

3. 安装 Go 语言编译器。在终端中运行以下命令:
   ```
   brew install go
   ```

4. 安装 Fyne 命令行工具。在终端中运行以下命令:
   ```
   go get fyne.io/fyne/v2/cmd/fyne
   ```

5. 克隆 ASM2HEX 项目到本地:
   ```
   git clone https://github.com/suifei/asm2hex.git
   ```

6. 进入项目目录:
   ```
   cd asm2hex
   ```

7. 编译 Keystone 和 Capstone 库:
   ```
   make clean
   make lib
   ```

8. 编译 ASM2HEX:
   ```
   make build
   ```

编译完成后,应用程序包将生成在 `build` 目录下。

## 注意事项

- 确保在编译前已正确安装和配置所需的环境和工具链。
- 如果在编译过程中遇到问题,请检查环境变量和路径设置是否正确。
- 如果编译失败,请检查错误信息并尝试解决相应的问题。

如果你在编译过程中遇到任何困难或有任何疑问,欢迎在 GitHub 仓库的 Issues 页面提出。
