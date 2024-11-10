// Filename analysis

package filenameanalysis

import (
	"fmt"
	"strings"
)

// !Yep, I just use AI here, I'll do deeper research on more dangerous characters to add (or less non-dangerous characters to remove from the lists). But for now I'll settle with these lists
// !Also add reasons, for example how %00 can be used to bypass file upload restriction, how control characters can harm the system, etc.
// Generic strange chars
var blacklist = []string{
	// Strange characters
	"/", "\\", ":", "?", "*", "\"", "'", "<", ">", "|", ",", "^", "~", "[", "]", "{", "}", ";", "(", ")", "\\0", "$",

	// Strange unicodes
	"\u200B", // Zero-width space
	"\uFEFF", // Zero-width no-break space

	// URL-encodes
	"%00", // Null byte
	"%01", "%02", "%03", "%04", "%05", "%06", "%07", "%08", "%09", // Control characters
	"%0A", "%0B", "%0C", "%0D", "%0E", "%0F", // Line feed and carriage return
	"%20", // Space
	"%21", "%23", "%24", "%26", "%27", "%28", "%29", "%2F", "%3A", "%3B", "%3C", "%3E", "%3F", "%40", "%5B", "%5D", "%60", // Special characters
	"%7B", "%7D", "%7C", "%5E", "%5C", "%60", "%7A", "%7A", "%7B", "%7D", // Unicode characters
	"%25", // Percent sign
	"%22", "%23", "%24", "%26", "%27", "%28", "%29", "%2F", "%3A", "%3B", "%3C", "%3E", "%3F", "%40", "%5B", "%5D", "%60", "%7B", "%7D", "%7C", "%5E", "%5C", "%60", "%7A", "%7A", "%7B", "%7D", // HTML entities

	// Control characters
	"\x00","\x01","\x02","\x03","\x04","\x05","\x06","\x07","\x08","\x09","\x0A","\x0B","\x0C","\x0D","\x0E","\x0F","\x10","\x11","\x12","\x13","\x14","\x15","\x16","\x17","\x18","\x19","\x1A","\x1B","\x1C","\x1D","\x1E","\x1F",

	// Windows system files
	"CON","PRN","AUX","NUL","COM0","COM1","COM2","COM3","COM4","COM5","COM6","COM7","COM8","COM9","LPT0","LPT1","LPT2","LPT3","LPT4","LPT5","LPT6","LPT7","LPT8","LPT9",
}

// Detect and print suspicious chars
func Check(filename string) () {
	var counter []string = make([]string, 0)
	for _, bl := range blacklist {
		if (strings.Contains(filename, bl)) {
			counter = append(counter, bl)
		}
	}
	if (len(counter) > 0) {
		fmt.Printf("\nFound %v suspicious characters on file %v:\n", len(counter), filename);
		for _, c := range counter {
			fmt.Printf("-> %v\n", c)
		}
		fmt.Printf("\n")
	}
}