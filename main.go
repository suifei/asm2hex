package main

import (
	"fmt"
	"image/color"
	"net/url"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/suifei/asm2hex/archs"
	"github.com/suifei/asm2hex/theme/icons"
)

const (
	ApplicationTitle       = "ASM to HEX Converter"
	ApplicationTitleToggle = "HEX to ASM Converter"
)

type ToggleMode string

const (
	ASM2HEX ToggleMode = "asm2hex"
	HEX2ASM ToggleMode = "hex2asm"
)

var toggle_mode ToggleMode = ASM2HEX
var codes []string = nil
var prefix_hex bool = false
var offset uint64 = 0
var bigEndian bool = false

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow(ApplicationTitle)

	myWindow.SetContent(createMainUI(myWindow))

	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.ShowAndRun()
}

func hexdump(buf []byte) string {
	if len(buf) == 0 {
		return ""
	}
	var hexStr string
	for _, b := range buf {
		// 16进制转换为字符串,保持4位长度，大写，不足补0
		hexStr += fmt.Sprintf("%02X", b)
	}
	if prefix_hex {
		return fmt.Sprintf("0x%s", hexStr)
	}
	return hexStr
}
func hexStringsToBytes(hexStrings []string) ([]byte, error) {
    var result []byte
    for _, hexStr := range hexStrings {
        // 移除字符串中的空格
        hexStr = strings.ReplaceAll(hexStr, " ", "")

        // 检查字符串长度是否为偶数
        if len(hexStr)%2 != 0 {
            return nil, fmt.Errorf("invalid hex string length: %s", hexStr)
        }

        // 解析每个字节
        for i := 0; i < len(hexStr); i += 2 {
            b, err := strconv.ParseUint(hexStr[i:i+2], 16, 8)
            if err != nil {
                return nil, err
            }
            result = append(result, byte(b))
        }
    }
    return result, nil
}
func createMainUI(win fyne.Window) *fyne.Container {
	var status *widget.Label
	var convertBtn *widget.Button
	var toggleBtn *widget.Button

	assemblyLabel := widget.NewLabel("Assembly code")
	offsetLabel := widget.NewLabel("Offset(hex)")

	output1 := widget.NewMultiLineEntry()
	output1.SetPlaceHolder("ARM64")
	output1.SetMinRowsVisible(20)
	output1.TextStyle.Monospace = true

	output2 := widget.NewMultiLineEntry()
	output2.SetPlaceHolder("ARM")
	output2.SetMinRowsVisible(20)
	output2.TextStyle.Monospace = true

	output3 := widget.NewMultiLineEntry()
	output3.SetPlaceHolder("THUMB")
	output3.SetMinRowsVisible(20)
	output3.TextStyle.Monospace = true

	output1_info := widget.NewLabel("Little Endian")
	output2_info := widget.NewLabel("Little Endian")
	output3_info := widget.NewLabel("Little Endian")

	output1_card := container.NewVBox(output1,
		container.NewHBox(
			output1_info,
			layout.NewSpacer(),
			widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {
				win.Clipboard().SetContent(output1.Text)
				fyne.CurrentApp().SendNotification(&fyne.Notification{Title: "Info", Content: "Copied to clipboard"})
			}),
		))
	output2_card := container.NewVBox(output2,
		container.NewHBox(
			output2_info,
			layout.NewSpacer(),
			widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {
				win.Clipboard().SetContent(output2.Text)
				fyne.CurrentApp().SendNotification(&fyne.Notification{Title: "Info", Content: "Copied to clipboard"})
			}),
		))
	output3_card := container.NewVBox(output3,
		container.NewHBox(
			output3_info,
			layout.NewSpacer(),
			widget.NewButtonWithIcon("", theme.ContentCopyIcon(), func() {
				win.Clipboard().SetContent(output3.Text)
				fyne.CurrentApp().SendNotification(&fyne.Notification{Title: "Info", Content: "Copied to clipboard"})
			}),
		))

	if bigEndian {
		output1_info.SetText("Big Endian")
		output2_info.SetText("Big Endian")
		output3_info.SetText("Big Endian")
	} else {
		output1_info.SetText("Little Endian")
		output2_info.SetText("Little Endian")
		output3_info.SetText("Little Endian")
	}
	output1_info.Refresh()
	output2_info.Refresh()
	output3_info.Refresh()

	tabs := container.NewAppTabs(
		container.NewTabItem("ARM64", output1_card),
		container.NewTabItem("ARM", output2_card),
		container.NewTabItem("THUMB", output3_card),
	)

	assembly_code := `; sample code
nop
ret
b #0x1018de444
mov x0, #0x11fe0000
beq #0x10020c
cbnz r0, #0x682c4
`

	assemblyEditor := widget.NewMultiLineEntry()
	assemblyEditor.SetPlaceHolder("Assembly code")
	assemblyEditor.SetText(assembly_code)
	assemblyEditor.AcceptsTab()
	assemblyEditor.SetMinRowsVisible(20)

	offsetInput := widget.NewEntry()
	offsetInput.SetPlaceHolder("Offset(hex)")
	offsetInput.SetText("0")

	assemblyEditor.OnChanged = func(text string) {
		convertBtn.Tapped(nil)
	}
	offsetInput.OnChanged = func(text string) {
		if text == "" {
			offset = 0
		} else {
			num, err := strconv.ParseUint(text, 16, 64)
			if err != nil {
				offset = 0
				status.SetText("Invalid offset")
			} else {
				offset = num
			}
		}
		status.SetText(fmt.Sprintf("Offset: 0x%x", offset))
		status.Refresh()
	}

	left_container := container.New(layout.NewVBoxLayout(),
		assemblyLabel,
		assemblyEditor,
		container.New(layout.NewFormLayout(), offsetLabel, offsetInput),
	)

	right_container := container.New(layout.NewVBoxLayout(),
		tabs,
	)

	grid := container.New(layout.NewGridLayoutWithColumns(2),
		left_container,
		right_container,
	)

	logo_icon := fyne.NewStaticResource("asm2hex.svg", icons.LOGO_ICON_BIN)
	github_icon := fyne.NewStaticResource("github.svg", icons.GITHUB_ICON_BIN)

	background := canvas.NewImageFromResource(logo_icon)
	background.SetMinSize(fyne.NewSize(64, 64))

	app_title := canvas.NewText(ApplicationTitle, color.NRGBA{0, 0x80, 0, 0xff})
	app_title.TextSize = 24

	about_messages := "ASM to HEX Converter\n\n" +
		"Version: 1.0\n" +
		"Author: suifei suifei@gmail.com\n" +
		"License: MIT\n" +
		"Source code: https://github.com/suifei/asm2hex\n\n" +
		"Copyright (c) 2024 suifei"
	// 关注按钮
	openUrlBtn := widget.NewButtonWithIcon("Star", logo_icon, func() {
		uri, _ := url.Parse("https://github.com/suifei/asm2hex")
		fyne.CurrentApp().OpenURL(uri)
	})
	openGithub := widget.NewButtonWithIcon("Github", github_icon, func() {
		uri, _ := url.Parse("https://github.com/suifei")
		fyne.CurrentApp().OpenURL(uri)
	})
	aboutDlg := dialog.NewCustom("About", "Close", container.New(
		layout.NewVBoxLayout(), widget.NewLabel(about_messages), container.NewHBox(openUrlBtn, openGithub),
	), win)
	status = widget.NewLabel("Ready")
	status.Alignment = fyne.TextAlignTrailing
	status.TextStyle = fyne.TextStyle{Bold: true}

	convertBtn = widget.NewButtonWithIcon("Convert", theme.StorageIcon(), func() {
		// Do conversion here
		doConversion(
			status,
			output1,
			output2,
			output3,
			assemblyEditor,
			offsetInput,
		)
	})
	convertBtn.Importance = widget.HighImportance
	toggleBtn = widget.NewButtonWithIcon("Toggle Mode", theme.ViewRefreshIcon(), func() {
		if toggle_mode == ASM2HEX {
			toggle_mode = HEX2ASM
			app_title.Text = ApplicationTitleToggle
			app_title.Color = color.NRGBA{0x80, 0, 0, 0xff}

			toggleBtn.Importance = widget.SuccessImportance
		} else {
			toggle_mode = ASM2HEX
			app_title.Text = ApplicationTitle
			app_title.Color = color.NRGBA{0, 0x80, 0, 0xff}

			toggleBtn.Importance = widget.WarningImportance
		}
		app_title.Refresh()
		SetMode(win, status, assemblyLabel, assemblyEditor)
		convertBtn.Tapped(nil)
		toggleBtn.Refresh()
	})
	toggleBtn.Importance = widget.WarningImportance
	clearBtn :=widget.NewButtonWithIcon("Clear", theme.DeleteIcon(), func() {
			status.SetText("Clear")
			status.Refresh()
			assemblyEditor.SetText("")
			offsetInput.SetText("0")
			output1.SetText("")
			output2.SetText("")
			output3.SetText("")
		})
	clearBtn.Importance = widget.DangerImportance
	aboutBtn :=widget.NewButtonWithIcon("About...", theme.QuestionIcon(), func() {
			status.SetText("About")
			status.Refresh()
			aboutDlg.Resize(fyne.NewSize(400, 300))
			aboutDlg.Refresh()
			aboutDlg.Show()
		})
	aboutBtn.Importance = widget.LowImportance

	status_container := container.New(layout.NewHBoxLayout(),
		status,
		layout.NewSpacer(),
		convertBtn,
		clearBtn,
		toggleBtn,
		aboutBtn,
	)
	theme.DocumentCreateIcon()

	status_container.Resize(fyne.NewSize(100, 24))

	main_layout := container.New(layout.NewVBoxLayout(),
		container.New(layout.NewHBoxLayout(),
			background,
			app_title,
			widget.NewCheck("0x", func(checked bool) {
				status.SetText("Changed")
				v := offsetInput.Text
				if checked {
					if strings.HasPrefix(v, "0x") {
						offsetInput.SetText(v)
					} else {
						offsetInput.SetText("0x" + v)
					}
					prefix_hex = true
				} else {
					if strings.HasPrefix(v, "0x") {
						offsetInput.SetText(strings.TrimPrefix(v, "0x"))
					} else {
						offsetInput.SetText(v)
					}
					prefix_hex = false
				}
				convertBtn.Tapped(nil)
			}),
			widget.NewCheck("GDB/LLDB", func(checked bool) {
				status.SetText("Changed")
				bigEndian = checked
				if bigEndian {
					status.SetText("Big Endian")
					output1_info.SetText("Big Endian")
					output2_info.SetText("Big Endian")
					output3_info.SetText("Big Endian")
				} else {
					status.SetText("Little Endian")
					output1_info.SetText("Little Endian")
					output2_info.SetText("Little Endian")
					output3_info.SetText("Little Endian")
				}
				output1_info.Refresh()
				output2_info.Refresh()
				output3_info.Refresh()
				status.Refresh()
				convertBtn.Tapped(nil)
			})),
		grid,
		layout.NewSpacer(),
		status_container,
	)

	convertBtn.Tapped(nil)

	return main_layout
}

func SetMode(win fyne.Window, status *widget.Label, assemblyLabel *widget.Label, assemblyEditor *widget.Entry) {
	if toggle_mode == ASM2HEX {
		assemblyLabel.SetText("Assembly code")
		assemblyEditor.SetPlaceHolder("Assembly code")
		status.SetText("Toggle to ASM2HEX")
		win.SetTitle(ApplicationTitle)
	} else {
		assemblyLabel.SetText("HEX code")
		assemblyEditor.SetPlaceHolder("HEX code")
		status.SetText("Toggle to HEX2ASM")
		win.SetTitle(ApplicationTitleToggle)
	}

	status.Refresh()
	assemblyEditor.Refresh()
	assemblyLabel.Refresh()
}

func doConversion(status *widget.Label,
	output1 *widget.Entry,
	output2 *widget.Entry,
	output3 *widget.Entry,
	assemblyEditor *widget.Entry,
	offsetInput *widget.Entry) {

	status.SetText("Converting...")
	status.Refresh()

	output1.SetText("")
	output2.SetText("")
	output3.SetText("")

	codes = strings.Split(assemblyEditor.Text, "\n")
	offset, _ := strconv.ParseUint(strings.ReplaceAll(strings.ToLower(offsetInput.Text), "0x", ""), 16, 64)

	pc_1 := offset
	pc_2 := offset
	pc_3 := offset

	status.SetText("Done")
	status.Refresh()

	if toggle_mode == ASM2HEX {
		for _, v := range codes {
			if strings.TrimSpace(v) == "" {
				continue
			}
			//asm to hex
			encoding, _, ok, err := archs.Arm64(v, pc_1, bigEndian)
			process(ok, output1, encoding, status, &pc_1, uint64(len(encoding)), err)
			encoding, _, ok, err = archs.Arm32(v, pc_2, bigEndian)
			process(ok, output2, encoding, status, &pc_2, uint64(len(encoding)), err)
			encoding, _, ok, err = archs.Thumb(v, pc_3, bigEndian)
			process(ok, output3, encoding, status, &pc_3, uint64(len(encoding)), err)
		}
	} else {
		encoding, err := hexStringsToBytes(codes)
		if err != nil {
			status.SetText(err.Error())
			status.Refresh()
			return
		}
		//hex to asm
		asm, _, ok, err := archs.Arm64Disasm(encoding, pc_1, bigEndian)
		process(ok, output1, []byte(asm), status, &pc_1, uint64(len(asm)), err)
		asm, _, ok, err = archs.Arm32Disasm(encoding, pc_2, bigEndian)
		process(ok, output2, []byte(asm), status, &pc_2, uint64(len(asm)), err)
		asm, _, ok, err = archs.ThumbDisasm(encoding, pc_3, bigEndian)
		process(ok, output3, []byte(asm), status, &pc_3, uint64(len(asm)), err)
	}
}

func process(ok bool, output *widget.Entry, encoding []byte, status *widget.Label, pc *uint64, pcSize uint64, err error) {
	if ok {
		if toggle_mode == HEX2ASM {
			output.Append(fmt.Sprintf("%s\n", encoding))
		} else {
			output.Append(hexdump(encoding) + "\n")

			status.SetText("Done")
			status.Refresh()
		}
		output.Refresh()
		*pc += pcSize

	} else {
		errMsg := err.Error()
		if strings.Contains(errMsg, "(KS") {
			errMsg = strings.Split(errMsg, "(KS")[0]
		}

		output.Append(errMsg + "\n")
		output.Refresh()

		status.SetText("Error:" + errMsg)
		status.Refresh()
	}

}
