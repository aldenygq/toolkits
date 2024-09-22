package toolkits
import (
    "testing"
    "fmt"
)
//go test -v -test.run Test_Crypt
func Test_Crypt(t *testing.T) {
	key := "HXl8G2+60gnxdvi.oP5?L94*F=VrDM3_"
	println("key:", key)
	pwd := "test123"
	encrypted, err := Encrypt([]byte(pwd), key)
	if err != nil {
		fmt.Printf("encrypt error: %v\n", err)
		return
	}
	fmt.Printf("encrypt:%v\n",encrypted)
	decrypted, err := Decrypt(encrypted, key)
	if err != nil {
		fmt.Printf("decrypt error: %v\n", err)
		return
	}
	if string(decrypted) != pwd {
		fmt.Printf("Decrypted text does not match original text")
	}
	fmt.Printf("decrypted:%v\n",string(decrypted))
}