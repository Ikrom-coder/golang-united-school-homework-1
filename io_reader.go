package main

import (
	"strings"
)

/*
SeekTillHalfOfString -  contains a code snippet in Go that defines a function called
"SeekTillHalfOfString". The function takes a string reader as input,
seeks to the middle of the string, reads
half of the remaining string, and returns it as a string.
*/
func SeekTillHalfOfString(strReader *strings.Reader) string {
	lenStr := strReader.Len()
	offSet := int64(lenStr/2 + 1)

	_, err := strReader.Seek(offSet-1, 0)
	if err != nil {
		return ""
	}

	byteSlice := make([]byte, offSet)
	content, err := strReader.Read(byteSlice)
	if err != nil {
		return ""
	}

	return string(byteSlice[:content])
}

/*
ReaderSplit - contains a code snippet written in Go that
defines a function called ReaderSplit.
The function takes a strings.Reader and an integer n as input,
and splits the contents of the reader into chunks of size n.
The function returns a slice of strings containing the chunks
*/
func ReaderSplit(strReader *strings.Reader, n int) []string {
	byteSlice := make([]byte, strReader.Size())
	content, err := strReader.Read(byteSlice)
	if err != nil {
		return nil
	}
	str := string(byteSlice[:content])

	var strList []string
	i := n
	for ; i < len(str); i += n {
		strList = append(strList, str[i-n:i])
	}
	if i-n != len(str) {
		strList = append(strList, str[i-n:])
	}

	return strList
}
