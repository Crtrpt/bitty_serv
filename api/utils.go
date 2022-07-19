package api

import( 
	"fmt"
	"net/smtp"
	"github.com/jordan-wright/email"
    "os"
	"crypto/tls"
	"crypto/aes"
    "encoding/hex"
)



func  sendMail(sendto string,msg string,subject string)  error {
	from := os.Getenv("email_from")
	account := os.Getenv("email_account")
    password := os.Getenv("email_password")

    host := os.Getenv("email_host")
   	port := os.Getenv("email_port")

	   fmt.Printf(" %s %s %s %s %s",host,port,account,password,from)
	   e := email.NewEmail()
	   e.From =from
	   e.To = []string{sendto}
	   e.Subject = subject
	   e.Text = []byte(msg)
	   tlsconfig := &tls.Config {
        InsecureSkipVerify: true,
        ServerName: host,
    	}
		err:= e.SendWithTLS(host+":"+port, smtp.PlainAuth("", account, password, host),tlsconfig)
		if(err!=nil){
			fmt.Println("")
			fmt.Println(err)
			return err
		}

	return nil;
}

func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}

func EncryptAES(key []byte, plaintext string) string {
	// create cipher
c, err := aes.NewCipher(key)
CheckError(err)
   
	// allocate space for ciphered data
out := make([]byte, len(plaintext))

	// encrypt
c.Encrypt(out, []byte(plaintext))
	// return hex string
return hex.EncodeToString(out)
}

func DecryptAES(key []byte, ct string) string {
	ciphertext, _ := hex.DecodeString(ct)

	c, err := aes.NewCipher(key)
	CheckError(err)

	pt := make([]byte, len(ciphertext))
	c.Decrypt(pt, ciphertext)

	s := string(pt[:])
	return s;
}
