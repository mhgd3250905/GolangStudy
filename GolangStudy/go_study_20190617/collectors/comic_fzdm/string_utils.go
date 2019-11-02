package comic_ikk

import (
	"encoding/base64"
	"github.com/axgle/mahonia"
	"golang.org/x/text/encoding/simplifiedchinese"
)


type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
	GBK = Charset("GBK")
)

func ConverString2Base64(input string) (output string){
	strbytes := []byte(input)
	return base64.StdEncoding.EncodeToString(strbytes)
}

func ConvertByte2String(byte []byte, charset Charset) string {

	var str string
	switch charset {
	case GB18030:
		var decodeBytes,_=simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str= string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}

	return str
}

func ConvertToString(src string, srcCode string, tagCode string) string {

	srcCoder := mahonia.NewDecoder(srcCode)

	srcResult := srcCoder.ConvertString(src)

	tagCoder := mahonia.NewDecoder(tagCode)

	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)

	result := string(cdata)

	return result

}
