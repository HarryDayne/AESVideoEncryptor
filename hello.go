package main

import (
	"fmt"
	"os"
	"io"
	"io/ioutil"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

func main() {
	var i int //variable to store the choice
	fmt.Println("************The Encryptor************")
	fmt.Println("Please enter 1 for encrption and 2 for decryption:")
	fmt.Scan(&i)
	if(i==1){
		fmt.Println("Chosen to encrypt ( Please make sure the file is named input.mp4 )")
		encrypt()
	}
	if i==2 {
		fmt.Println("Chosen to Decrypt ( Please make sure the file is named encrypted.txt and the key is provided in key.txt )")
		decrypt()
	}
}

func encrypt() { //function to open a file named input.mp4 , encrypt it and store the result in a new file called out.bin
	data,err := ioutil.ReadFile("input.mp4")
	if(err != nil){
		fmt.Println("Theres an error!\n",err)
	}

	key := make([]byte, 32) // 256-bit key
	if _, err := rand.Read(key); err != nil {
		fmt.Println(err)
	}
	// Create a new AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}

	// Create a new GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println(err);
	}

	// Create a nonce (number used once)
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}

	// Encrypt the data
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	file,err := os.Create("./encrypted.bin")
	if err != nil {
		fmt.Println(err)
	}

	l,err:=file.Write(ciphertext)
	fmt.Println(l)

	file1,err :=os.Create("./key.txt")
	if err != nil {
		fmt.Println(err)
	}
	l1,err := io.WriteString(file1,string(key))
	fmt.Println(l1)
}

func decrypt(){
	ciphertext,err := ioutil.ReadFile("encrypted.bin")
	key1,err := ioutil.ReadFile("key.txt")
	key := []byte(key1)
	if err != nil {
		fmt.Println(err)
    }

   // Create a new cipher block from the key
    block, err := aes.NewCipher(key)
    if err != nil {
		fmt.Println(err)
	}

    // Create a new GCM instance
    gcm, err := cipher.NewGCM(block)
    if err != nil {
		fmt.Println(err)
	}

    // Get the nonce size
    nonceSize := gcm.NonceSize()
    if len(ciphertext) < nonceSize {
		fmt.Println("Cipher Text is too short")
	}

    // Extract the nonce from the ciphertext
    nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

    // Decrypt the data
    plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
    if err != nil {
		fmt.Println(err)
    }
	outputfile ,err := os.Create("./out.mp4")
	if err != nil {
		fmt.Println(err)
	}
	l,err:=outputfile.Write(plaintext)
	fmt.Println(l)
	if err != nil {
		fmt.Println(err)
	}
}
