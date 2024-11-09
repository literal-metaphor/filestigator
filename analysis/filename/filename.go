// Filename analysis

package filenameanalysis

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

// Generic strange chars
var blacklist_generic = []string{
	"/", "\\", ":", "?", "*", "\"", "'", "<", ">", "|", ",", "^", "~", "[", "]", "{", "}", ";", "(", ")", "\\0", "$", "\u200B", "\uFEFF",
}
// Control characters
// !Untested
var blacklist_control = []string{
	"\x00","\x01","\x02","\x03","\x04","\x05","\x06","\x07","\x08","\x09","\x0A","\x0B","\x0C","\x0D","\x0E","\x0F","\x10","\x11","\x12","\x13","\x14","\x15","\x16","\x17","\x18","\x19","\x1A","\x1B","\x1C","\x1D","\x1E","\x1F",
}
// URL-encoded
var blacklist_url = []string{
	"%00", "%20", "%2E%2E%2F", "%22%20",
}
// Windows system files
var blacklist_windows = []string{
	"CON","PRN","AUX","NUL","COM0","COM1","COM2","COM3","COM4","COM5","COM6","COM7","COM8","COM9","LPT0","LPT1","LPT2","LPT3","LPT4","LPT5","LPT6","LPT7","LPT8","LPT9",
}

// Check if the target is blacklisted
// func blacklisted(arr []string, target string) bool {
// 	for _, val := range arr {
// 		if val == target {
// 			return true
// 		}
// 	}
// 	return false
// }

// Return chars in filename that matches:
// "/", "\\", ":", "?", "*", "\"", "'", "<", ">", "|", ",", "^", "~", "[", "]", "{", "}", ";", "(", ")", "\\0", "$", "\u200B", "\uFEFF"
func TestGeneric(filename string) () {
	fmt.Println("Checking for strange characters in filename...");
	counter := 0;
	for _, char := range strings.Split(filename, "") {
		bad := false;
		for _, bl := range blacklist_generic {
			if (bl == char) {
				bad = true;
			}
		}

		if bad {
			color.Set(color.BgRed);
			counter++;
		}

		fmt.Print(char)

		color.Unset();
	}
	fmt.Printf("\nFound %v strange characters.", counter);
}