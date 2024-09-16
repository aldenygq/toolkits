package toolkits
import (
        "crypto/tls"
        "net/smtp"
        "strings"
        "fmt"
        "errors"
        "bytes"
        "time"
        "io/ioutil"
        "encoding/base64"
)
type Mail interface {
	NewMailClient(mail MailConf) (*smtp.Client,error)
	SendMail(client *smtp.Client,message Message) error
}
type Attachment struct {
	name        string
	contentType string
	withFile    bool
}

type Message struct {
	from        string
	to          []string
	cc          []string
	bcc         []string
	subject     string
	body        string
	contentType string
	attachment  Attachment
}
type MailConf struct {
	user     string
	passwd string
	host     string
	port     string
}

func NewMailClient(mail MailConf) (*smtp.Client,error) {
    auth := smtp.PlainAuth("", mail.user, mail.passwd, mail.host)
    conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%s", mail.host, mail.port), &tls.Config{})
    if err != nil {
        fmt.Printf("connect mail failed:%v\n",err)
        return nil,err
    }
    c, err := smtp.NewClient(conn, mail.host)
    if err != nil {
        fmt.Printf("new mail client  failed:%v\n",err)
        return nil,err
    }
    if auth != nil {
        if ok, _ := c.Extension("AUTH"); ok {
            if err = c.Auth(auth); err != nil {
                fmt.Printf("mail auth failed:%v\n",err)
                return nil,err
            }
        } else {
            fmt.Printf("auth not supported")
            return nil,errors.New("auth not supported")
        }
    }
    return c,nil
}

func SendMail(c *smtp.Client,content Message) error {
	buffer := bytes.NewBuffer(nil)
	boundary := "GoBoundary"
	Header := make(map[string]string)
	Header["From"] = content.from
	Header["To"] = strings.Join(content.to, ";")
	Header["Cc"] = strings.Join(content.cc, ";")
	Header["Bcc"] = strings.Join(content.bcc, ";")
	Header["Subject"] = content.subject
	Header["Content-Type"] = "multipart/mixed;boundary=" + boundary
	Header["Mime-Version"] = "1.0"
	Header["Date"] = time.Now().String()
	writeHeader(buffer, Header)

	body := "\r\n--" + boundary + "\r\n"
	body += "Content-Type:" + content.contentType + "\r\n"
	body += "\r\n" + content.body + "\r\n"
	buffer.WriteString(body)
    if content.attachment.withFile {
		attachment := "\r\n--" + boundary + "\r\n"
		attachment += "Content-Transfer-Encoding:base64\r\n"
		attachment += "Content-Disposition:attachment\r\n"
		attachment += "Content-Type:" + content.attachment.contentType + ";name=\"" + content.attachment.name + "\"\r\n"
		buffer.WriteString(attachment)
		defer func() {
			if err := recover(); err != nil {
				panic(err)
			}
		}()
		writeFile(buffer, content.attachment.name)
	}

	buffer.WriteString("\r\n--" + boundary + "--")

       // 发送邮件
    if err := c.Mail(content.from); err != nil {
        return err
    }
    for _, addr := range content.to {
        if err := c.Rcpt(addr); err != nil {
            return err
        }
    }
    w, err := c.Data()
    if err != nil {
        return err
    }

    _, err = w.Write([]byte(buffer.Bytes()))
    if err != nil {
        return err
    }
    err = w.Close()
    if err != nil {
        return err
    }
    // 发送结束
    if err = c.Quit(); err != nil {
        return err
    }

    return nil
}

func writeFile(buffer *bytes.Buffer, fileName string) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err.Error())
	}
	payload := make([]byte, base64.StdEncoding.EncodedLen(len(file)))
	base64.StdEncoding.Encode(payload, file)
	buffer.WriteString("\r\n")
	for index, line := 0, len(payload); index < line; index++ {
		buffer.WriteByte(payload[index])
		if (index+1)%76 == 0 {
			buffer.WriteString("\r\n")
		}
	}
}
func writeHeader(buffer *bytes.Buffer, Header map[string]string) string {
	header := ""
	for key, value := range Header {
		header += key + ":" + value + "\r\n"
	}
	header += "\r\n"
	buffer.WriteString(header)
	return header
}
