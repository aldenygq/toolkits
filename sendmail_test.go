package toolkits
import (
    "fmt"
    "testing"
    //"html/template"
    //"bytes"
)
//go test -v -test.run Test_SendMail
func Test_SendMail(t *testing.T) {
    var mail MailConf
    mail.user = "****************8"
    mail.passwd = "****************"
    mail.host = "smtp.163.com"
    mail.port = "465"
    client,err  := NewMailClient(mail)
    if err != nil {
        fmt.Printf("new mail client failed:%v\n",err)
        return
    }

    if client == nil {
        fmt.Printf("mail client invalid")
        return
    }

    fmt.Printf("new mail client success\n")
    var content Message
    var attachment Attachment
//发送html邮件
/*
const emailTemplate = `<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>中秋节快乐</title>
    <style>
        body {
            font-family: 'Arial', sans-serif;
            background-color: #f8f8f8;
            margin: 0;
            padding: 0;
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            flex-direction: column;
        }
        .container {
            text-align: center;
            background: white;
            padding: 20px;
            border-radius: 8px;
            box-shadow: 0 4px 8px rgba(0,0,0,0.1);
        }
        .moon {
            width: 150px;
            height: 150px;
            border-radius: 50%;
            background: rgba(255, 255, 255, 0.8);
            position: relative;
            margin-bottom: 20px;
        }
        .moon::after {
            content: '';
            position: absolute;
            top: 50%;
            left: 50%;
            transform: translate(-50%, -50%);
            width: 140px;
            height: 140px;
            border-radius: 50%;
            background: #fcd96a;
        }
        h1 {
            color: #d96d46;
            margin: 0 0 10px 0;
        }
        p {
            font-size: 16px;
            color: #666;
            margin: 0 0 20px 0;
        }
        button {
            padding: 10px 20px;
            border: none;
            border-radius: 5px;
            background-color: #d96d46;
            color: white;
            cursor: pointer;
            font-size: 16px;
        }
        button:hover {
            background-color: #c05c40;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="moon"></div>
        <h1>中秋节快乐</h1>
        <p>亲爱的，中秋节到了，祝你幸福美满，团团圆圆！</p>
        <button onclick="showMessage()">点击查看祝福</button>
    </div>

    <script>
        function showMessage() {
            alert('愿我们的爱情像这轮明月一样，圆圆满满，永不褪色。');
        }
    </script>
</body>
</html>`
    tpl := template.Must(template.New("email").Parse(emailTemplate))
    type EmailData struct {
        Name string
    }
    data := EmailData{Name: "Recipient"}
    var buffer bytes.Buffer
    if err := tpl.Execute(&buffer, data); err != nil {
        fmt.Println("Error executing template:", err)
        return
    }
    message := ""
    message += "\r\n" + buffer.String()
    content.contentType = "text/html"
    content.body = message
    attachment.withFile = false
    */

    //普通邮件
    /*
    content.contentType = "text/plain; charset=UTF-8"
    content.body = ""
    attachment.withFile = false
    */

    //发送附件
    /*
    attachment.withFile = true
    attachment.name = "/Users/mingyu/code/src/toolkits/sms.go"
    attachment.contentType = "application/octet-stream"
    */
    content.from = "**************88"
    content.to = []string{"**********************"}
    content.cc = []string{"***********************8"}
    content.bcc = []string{"************************"}
    content.subject = "mail test"
    content.attachment = attachment
    err = SendMail(client,content)
    if err != nil {
        fmt.Printf("send email failed:%v\n",err)
        return
    }

    fmt.Printf("send email success\n")
}
