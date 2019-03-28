package prt

import "fmt"

func Printf(m Mode, f Color, format string, a ...interface{}) (n int, err error) {
	content := fmt.Sprintf(format, a...)
	return prt(m, f, content)
}

func Println(m Mode, f Color, a ...interface{}) (n int, err error) {
	content := fmt.Sprintln(a...)
	return prt(m, f, content)
}

func Print(m Mode, f Color, a ...interface{}) (n int, err error) {
	content := fmt.Sprint(a...)
	return prt(m, f, content)
}

func HighLightRed(a ...interface{}) (n int, err error) {
	return Println(HighLight, Red, a...)
}

func HighLightGreen(a ...interface{}) (n int, err error) {
	return Println(HighLight, Green, a...)
}

func prt(m Mode, f Color, content string) (n int, err error) {
	return fmt.Printf("%c[%d;%dm%s%c[0m", 0x1B, m, f+30, content, 0x1B)
}

func Printf_bg(m Mode, f Color, b Color, format string, a ...interface{}) (n int, err error) {
	var (
		length = len(format) - 1
		count  = 0
	)
	for format[length] == '\n' && length > -1 {
		length--
		count++
	}

	content := fmt.Sprintf(format[:length+1], a...)
	n, err = prtWithBackground(m, f, b, content)
	for count > 0 {
		fmt.Print("\n")
		count--
	}
	return n, err
}

func Println_bg(m Mode, f Color, b Color, a ...interface{}) (n int, err error) {
	content := fmt.Sprint(a...)
	n, err = prtWithBackground(m, f, b, content)
	fmt.Print("\n")
	return n, err
}

func Print_bg(m Mode, f Color, b Color, a ...interface{}) (n int, err error) {
	content := fmt.Sprint(a...)
	return prtWithBackground(m, f, b, content)
}

func prtWithBackground(m Mode, f Color, b Color, content string) (n int, err error) {
	return fmt.Printf("%c[%d;%d;%dm%s%c[0m", 0x1B, m, b+40, f+30, content, 0x1B)
}
