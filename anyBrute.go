package main

import (
	"flag"
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
)

// **  GLOBALS  **//
var verbose bool

// if you ever need the map keys _as_ your characterset ("nums" for instance), randomize order like: "snum" or "usmn" or "msnu"
// i guess i could also fix this by having a double char at the beginning of each charset name (__)
// but even though that solution is PERFECT it _looks bad_ and so i'm not doing it  :]
var charsetMap = map[string]string{
	"nums":                   "1234567890",
	// "asciiNums":              "1234567890", reduce redundancy and retry rext rime
	"hexUpper":               "ABCDEF1234567890",
	"hexLower":               "abcdef1234567890",
	"asciiUpper":             "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	"asciiLower":             "abcdefghijklmnopqrstuvwxyz",
	"asciiUpperLower":        "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
	"asciiUpperAndNums":      "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890",
	"asciiLowerAndNums":      "abcdefghijklmnopqrstuvwxyz1234567890",
	"asciiUpperLowerAndNums": "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890",
	"asciiAllPrintables":     "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890<>?,./:;'{}|[]\"\\~`!@#$%^&*()-_=+ ",

	// i've decided i'm formalizing on camel but hey maybe you like snake i'm not your mom <3
	// "numbers":                    "1234567890",
	// "ascii_nums":                 "1234567890",
	// "hex_upper":                  "ABCDEF1234567890",
	// "hex_lower":                  "abcdef1234567890",
	// "ascii_upper":                "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	// "ascii_lower":                "abcdefghijklmnopqrstuvwxyz",
	// "ascii_upper_lower":          "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
	// "ascii_upper_and_nums":       "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890",
	// "ascii_lower_and_nums":       "abcdefghijklmnopqrstuvwxyz1234567890",
	// "ascii_upper_lower_and_nums": "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890",
	// "ascii_all_printables":       "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890<>?,./:;'{}|[]\"\\~`!@#$%^&*()-_=+ ",
	
	// "asciiNumsMoreRangeTest":     "0123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890", <--- doesn't seem to help, keyspace too small means patterns emerge AF!!
}

// **  HELPERS  **//
func verbosePrint(str string) {
	if verbose {
		fmt.Println("V >>>> " + str + " <<<< _v")
	}
}

func normalPrint(str string) {
	fmt.Println(str)
}

func replaceAtIndex(str string, replacement rune, index int) string {
	return str[:index] + string(replacement) + str[index+1:]
}

// **  BEGINCD  **//
func genRandomCharFromCharset(charset string) rune {
	return rune(charset[rand.IntN((len(charset)))]) // rand.IntN is kinda weird, its argument is a top limit that cannot_ be returned (as such  no need for -1 here for indexing offset!)
}

// / may be unecessary / redundant?  [no we can use it in case of no form!]
func genRandomCharsFromCharset(charset string, length int) string {
	toReturn := ""
	for i := 0; i < length; i++ {
		toReturn += string(genRandomCharFromCharset(charset))
	}
	return toReturn
}

func genFromForm(charset string, form string, controlChar string) string {
	toReplaceInt := strings.Count(form, controlChar)
	// / if no form! \/
	// neededChars := genRandomCharsFromCharset()
	toReplaceIndexes := make([]int, 0, toReplaceInt)
	for i := 0; i < len(form); i++ {
		if string(form[i]) == controlChar {
			toReplaceIndexes = append(toReplaceIndexes, i)
		}
	}

	for _, index := range toReplaceIndexes {
		form = replaceAtIndex(form, genRandomCharFromCharset(charset), index)
	}

	return form
}

func getCharsetFromFlagAndMap(charsetFlagVal string) string {
	if charsetMap[charsetFlagVal] == "" {
		return charsetFlagVal
	}

	return charsetMap[charsetFlagVal]
}

func main() {
	// FLAGTIME
	var countFlag int
	var lengthFlag int
	var formFlag string
	var controlCharFlag string
	var charsetFlag string
	var verboseFlag bool

	flag.IntVar(&countFlag, "count", 8, "Number of strings to generate.  [Default: 8]")
	flag.IntVar(&countFlag, "c", 8, "Number of strings to generate.  [Default: 8]")
	flag.IntVar(&lengthFlag, "length", 10, "Length of strings to generate.  [Default: 10]")
	flag.IntVar(&lengthFlag, "l", 10, "Length of strings to generate.  [Default: 10]")
	flag.StringVar(&formFlag, "form", "XXXX-XXXXXXXX-XXXX", "full word flag, --example, second arg is default")
	flag.StringVar(&formFlag, "f", "XXXX-XXXXXXXX-XXXX", "single letter flag, -e, second arg is default")
	flag.StringVar(&controlCharFlag, "control-char", "X", "full word flag, --example, second arg is default")
	flag.StringVar(&controlCharFlag, "cc", "X", "single letter flag, -e, second arg is default")
	flag.StringVar(&charsetFlag, "charset", "hexLower", "full word flag, --example, second arg is default")
	flag.StringVar(&charsetFlag, "ch", "hexLower", "single letter flag, -e, second arg is default")
	// TODO  output/append to a file, and also encoding options! (b64 mainly)

	flag.BoolVar(&verboseFlag, "verbose", false, "Enable more verbose output.")
	flag.BoolVar(&verboseFlag, "v", false, "Enable more verbose output.")

	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])

		//flag.PrintDefaults()
		flagSet := flag.CommandLine
		order := [][]string{
			[]string{"", "Optional flags:"},
			[]string{"count", "c"},
			[]string{"length", "l"},
			[]string{"form", "f"},
			[]string{"control-char", "cc"},
			[]string{"charset", "ch"},
			[]string{"verbose", "v"}}

		for _, set := range order {
			if set[0] == "" {
				fmt.Println("\n  " + set[1])
				continue
			}
			flag := flagSet.Lookup(set[0])
			fmt.Println("    --" + set[0] + " (-" + set[1] + ")  ->  " + flag.Usage)
		}
		fmt.Println("\n  Available charsets:")
		for name, _ := range charsetMap {
			fmt.Println("     -> " + name)
		}

		fmt.Println("\n If you provide a string which is NOT one of these charsets to the charset flag it will use THAT string as a custom charset.")
	}

	flag.Parse()

	if verboseFlag {
		verbose = true
	} else {
		verbose = false
	}

	// CODETIME
	finalCharset := getCharsetFromFlagAndMap(charsetFlag)
	verbosePrint("Character set to be generated from: " + finalCharset)

	// remember it should count control characters , generate needed chars, THEN replace into control sequence! (otherwise we lose whatever the control char is....)
	// either that or it should scan the string char by char and replace as it goes?
	for range countFlag {
		normalPrint(genFromForm(finalCharset, formFlag, controlCharFlag))
	}
}
