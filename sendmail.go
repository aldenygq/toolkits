package toolkits
import (
        "crypto/tls"
        "net/smtp"
        "strings"
)
func NewMailClient(from,passwd,stmpHost,stmpPort string) (*stmp.Client,error) {
    conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%s", smtpHost, smtpPort), &tls.Config{})
    if err != nil {
        return nil,err
    }
    c, err := smtp.NewClient(conn, smtpHost)
    if err != nil {
        return nil,err
    }
    err = StmpAuth(c,from,passwd,smtpHost)
    if err != nil {
        return nil,err
    }
    return c,nil
}
func StmpAuth(c *smtp.Client,from,passwd,smtpHost string) error {
    auth := smtp.PlainAuth("", from, password, smtpHost)
    // 认证
    if auth != nil {
        if ok, _ := c.Extension("AUTH"); ok {
            if err = c.Auth(auth); err != nil {
                return err
            }
        } else {
            return errors.New("auth not supported")
        }
    }
    return nil
}

func SendMail(c *smtp.Client,form,message string,to []string]) error {
       // 发送邮件
    if err = c.Mail(from); err != nil {
        return err
    }
    for _, addr := range to {
        if err = c.Rcpt(addr); err != nil {
            return err
        }
    }
    w, err := c.Data()
    if err != nil {
        return err
    }

    _, err = w.Write([]byte(message))
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
