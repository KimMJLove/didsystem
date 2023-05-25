package utils

import (
    "crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    "encoding/pem"
    "os"
)

func GenerateRSAKeyPair(bits int) error {
    // 生成私钥
    privateKey, err := rsa.GenerateKey(rand.Reader, bits)
    if err != nil {
        return err
    }

    // 将私钥编码为PKCS#1格式
    privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)

    // 将私钥存储到文件
    privateKeyFile, err := os.Create("private.pem")
    if err != nil {
        return err
    }
    defer privateKeyFile.Close()

    privateKeyPEM := &pem.Block{
        Type:  "RSA PRIVATE KEY",
        Bytes: privateKeyBytes,
    }

    if err := pem.Encode(privateKeyFile, privateKeyPEM); err != nil {
        return err
    }

    // 生成公钥
    publicKey := privateKey.PublicKey

    // 将公钥编码为PKIX格式
    publicKeyBytes, err := x509.MarshalPKIXPublicKey(&publicKey)
    if err != nil {
        return err
    }

    // 将公钥存储到文件
    publicKeyFile, err := os.Create("public.pem")
    if err != nil {
        return err
    }
    defer publicKeyFile.Close()

    publicKeyPEM := &pem.Block{
        Type:  "RSA PUBLIC KEY",
        Bytes: publicKeyBytes,
    }

    if err := pem.Encode(publicKeyFile, publicKeyPEM); err != nil {
        return err
    }

    return nil
}

// func main() {
//     err := GenerateRSAKeyPair(2048)
//     if err != nil {
//         panic(err)
//     }
// }

//这个函数使用Go语言内置的RSA库生成一个指定位数的RSA公钥和私钥
//然后将它们分别编码为PEM格式，并存储到文件中
//可以通过调用这个函数生成公钥和私钥对，并将它们用于加密和解密数据。