# easy-gomails
- [Copy this link to download][github.com/Parasdeveloper8/easy-gomails.git/mails](github.com/Parasdeveloper8/easy-gomails.git/mails)
### With this package you can send mails easily using sendMail() function 
#### Example Code := 
```go
    package main 

    import (
        easymails "github.com/Parasdeveloper8/easy-gomails.git/mails"
    )
    func main(){
        senderMail := "yourmail@mail.com"//your mail
        emailPassword := "App Password" //your app password
        receivermail := "receiver@mail.com"//Person who will receive mail
        subject := "Testing"
        body := "Hello ! "
    easymails.SendMail(senderMail, emailPassword, receivermail, subject, body string)
    }
```
#### It uses gomail.v2 package internaly
