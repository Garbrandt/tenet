package utlis

import (
	"bufio"
	"bytes"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io"
	"io/ioutil"
	"log"
)

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
	GB2312  = Charset("gb2312")
)

// 读取文件文件，返回非乱码文件
func Convert(path string) []byte {
	fileBytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println(err)
	}

	cs, _, _ := autoDetectCharset(fileBytes)

	content := getContent(path)
	switch cs {
	case "gbk":
		r := bytes.NewReader(content)
		d, err := charset.NewReader(r, "gb2312")
		if err != nil {
			return content
		}
		content, err = ioutil.ReadAll(d)
		if err != nil {
			return content
		}
		return content
	case "gb-18030":
		return []byte(convertByte2String(content, "GB18030"))
	}

	return content
}

func determineEncodingFromReader(r io.Reader) (e encoding.Encoding, name string, certain bool, err error) {
	b, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		return
	}

	e, name, certain = charset.DetermineEncoding(b, "")
	return
}

func autoDetectCharset(b []byte) (string, bool, error) {
	_, name, certain, err := determineEncodingFromReader(bytes.NewReader(b))
	return name, certain, err
}

func convertByte2String(byte []byte, charset Charset) string {
	var str string
	switch charset {
	case GB18030:
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	case GB2312:
		var decodeBytes, _ = simplifiedchinese.HZGB2312.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	default:
		str = string(byte)
	}

	return str
}
