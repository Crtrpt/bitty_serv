package api

import( 
	"fmt"
	"net/smtp"
	"github.com/jordan-wright/email"
    "os"
	 "bytes"
	"crypto/tls"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
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


func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func EncryptAES(key []byte, orig string) string {
	origData := []byte(orig)
	k := []byte(key)

	block, _ := aes.NewCipher(k)

	blockSize := block.BlockSize()

	origData = PKCS7Padding(origData, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])

	cryted := make([]byte, len(origData))

	blockMode.CryptBlocks(cryted, origData)
	return base64.StdEncoding.EncodeToString(cryted)
}

func DecryptAES(key []byte, cryted string) string {
	crytedByte, _ := base64.StdEncoding.DecodeString(cryted)
	k := []byte(key)

	block, _ := aes.NewCipher(k)

	blockSize := block.BlockSize()

	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])

	orig := make([]byte, len(crytedByte))

	blockMode.CryptBlocks(orig, crytedByte)

	orig = PKCS7UnPadding(orig)
	return string(orig)
}
