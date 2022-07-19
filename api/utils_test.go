package api

import(
    //"os"
    "testing"
    // "gotest.tools/v3/assert"
    // "github.com/joho/godotenv"
)



func  TestEncryptAES(t *testing.T){
    // godotenv.Load();
    source:="123456"
    key:=[]byte("thisis32bitlongpassphraseimusing");
    ec:=  EncryptAES(key,source);
    dc:= DecryptAES(key,ec);
    if(source!=dc){
        t.Errorf("加密解密异常")
    }
    
    // return source==dc;
    //assert.Equal(t, source, dc);
}
