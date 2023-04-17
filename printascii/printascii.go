package printascii

import (
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"unicode"
)

func AsciiWeb(inputStr, font string) (string, error) {
	myStr := inputStr
	fileName := ""
	result := ""
	var err error
	fileName = fileNameCheck(font)
	if fileName == "" {
		return "", ErrFont
	}
	if !txtFileCheck(fileName) {
		return "", ErrTxtFile
	}
	if !isASCII(myStr) {
		return "", ErrNonAscii
	}
	if len(myStr) == 0 {
		return "", ErrString
	}

	result, err = splitWord(myStr, fileName)
	if err != nil {
		return "", ErrRead
	}
	return result, err
}

// splits the input string by newlines
func splitWord(myStr, myFile string) (string, error) {
	charMap, err := createMaps(myFile)
	if err != nil {
		return "", err
	}
	var finalStr string
	re := regexp.MustCompile(`\r`)
	newStr := re.Split(myStr, -1)
	for i := 0; i < len(newStr); i++ {
		if len(newStr[i]) > 0 {
			finalStr += printWord(newStr[i], myFile, charMap)
		}
		if newStr[i] == "" {
			finalStr += "\n"
		}
	}
	return finalStr, nil
}

// checks whether the given txt fileName exists
func fileNameCheck(fName string) string {
	myFile := ""
	switch fName {
	case "standard":
		myFile = "standard.txt"
	case "shadow":
		myFile = "shadow.txt"
	case "thinkertoy":
		myFile = "thinkertoy.txt"
	}

	return myFile
}

// check whether the changes have been made to the existing txt files
func txtFileCheck(fileName string) bool {
	hashStandard := []byte{225, 148, 241, 3, 52, 66, 97, 122, 184, 167, 142, 28, 166, 58, 32, 97, 245, 204, 7, 163, 240, 90, 194, 38, 237, 50, 235, 157, 253, 34, 166, 191}
	hashShadow := []byte{184, 17, 37, 168, 183, 46, 207, 226, 35, 69, 169, 190, 218, 184, 99, 86, 141, 179, 152, 16, 96, 21, 242, 206, 76, 172, 130, 232, 162, 21, 7, 76}
	hashThinkertoy := []byte{236, 241, 252, 123, 255, 114, 166, 211, 68, 247, 17, 86, 18, 3, 196, 224, 126, 132, 206, 58, 147, 120, 23, 16, 71, 60, 235, 235, 128, 88, 253, 28}
	file, err := os.Open("banners/" + fileName)
	if err != nil {
		log.Println(err)
		return false
	}
	defer file.Close()
	buf := make([]byte, 30*1024)
	sha256 := sha256.New()
	for {
		n, err := file.Read(buf)
		if n > 0 {
			_, err := sha256.Write(buf[:n])
			if err != nil {
				log.Fatal(err)
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Read %d bytes: %v", n, err)
			break
		}
	}
	sum := sha256.Sum(nil)
	switch fileName {
	case "shadow.txt":
		if string(sum) == string(hashShadow) {
			return true
		}
	case "standard.txt":
		if string(sum) == string(hashStandard) {
			return true
		}
	case "thinkertoy.txt":
		if string(sum) == string(hashThinkertoy) {
			return true
		}
	}
	return false
}

// checks whether the input consists of only the ascii characters
func isASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}

// returns the given string as characters given in the txt file format
func printWord(s, fileName string, charMap map[rune][8]string) string {
	myWord := ""
	for i := 0; i < 8; i++ {
		for _, char := range s {
			myWord += charMap[char][i]
		}
		myWord += "\n"
	}
	return myWord
}

// creates map of ascii characters given in the txt file format
func createMaps(fileName string) (map[rune][8]string, error) {
	asciiSign := make(map[rune][8]string)
	var line int = 1
	var myword [8]string
	content, err := ioutil.ReadFile("banners/" + fileName)
	if err != nil {
		fmt.Println(err)
		return asciiSign, err
	}
	text := string(content)
	reFile := regexp.MustCompile(`\n`)
	myAscii := reFile.Split(text, -1)
	for i := ' '; i <= '~'; i++ {
		for j := line; j < line+8; j++ {
			myword[j-line] = myAscii[j]
		}
		asciiSign[i] = myword
		line += 9
	}
	return asciiSign, nil
}
