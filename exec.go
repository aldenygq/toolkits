package toolkits
import (
	"os/exec"
	"bytes"
	"fmt"
)

func RunCmd(script string) (string,error) {
    cmd := exec.Command("/bin/bash", "-c",script)
    var stdin, stdout, stderr bytes.Buffer
    cmd.Stdin = &stdin
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    err := cmd.Run()
    if err != nil {
		fmt.Printf("run command failed:%v\n",err)
       return "",err 
    }
    outStr, _ := string(stdout.Bytes()), string(stderr.Bytes())
    return outStr,nil 
}