#   --target value, --os value         
#The operating system to target (android, android/arm, android/arm64, android/amd64, android/386, darwin, freebsd, ios, linux, netbsd, openbsd, windows)
release:
	fyne release -os macOS -appID suifei.asm2hex.app -appVersion 1.0 -appBuild 1 -category tools
	fyne release -os iOS -appID suifei.asm2hex.app -appVersion 1.0 -appBuild 1
