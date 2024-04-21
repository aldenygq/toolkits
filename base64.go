package toolkits

import (
	b64 "encoding/base64"
)

func B64EncodeToString(str string) string {
	return b64.StdEncoding.EncodeToString([]byte(str))
}
func B64DecodeString(str string) (string,error) {
	s,err := b64.StdEncoding.DecodeString(str)
	if err != nil {
		return "",err
	}
	
	return string(s),nil
}