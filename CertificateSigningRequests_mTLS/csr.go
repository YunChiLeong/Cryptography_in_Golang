package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	pubKey, priKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Print("Failed to generate public/private key pair.")
		os.Exit(1)
	}

	//Write private key to a file named clientKey.pem
	priKeyFile, err := os.Create("clientKey.pem")
	if err != nil {
		log.Fatalln("Error creating empty clientKey.pem file")
	}
	defer priKeyFile.Close()
	pkcs8key, err := x509.MarshalPKCS8PrivateKey(priKey)
	priKeyPem := pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: pkcs8key,
	}
	err = pem.Encode(priKeyFile, &priKeyPem)

	subject := pkix.Name{
		CommonName:         "request.com",
		Organization:       []string{"Request Company"},
		OrganizationalUnit: []string{"IT"},
	}

	extSubjectAltName := pkix.Extension{}
	extSubjectAltName.Id = asn1.ObjectIdentifier{2, 5, 29, 17}
	extSubjectAltName.Critical = false
	extSubjectAltName.Value = []byte(`yleong@uwyo.edu`)

	template := x509.CertificateRequest{
		SignatureAlgorithm: x509.PureEd25519,
		PublicKey:          pubKey,
		Subject:            subject,
		ExtraExtensions:    []pkix.Extension{extSubjectAltName},
	}

	csrBytes, err := x509.CreateCertificateRequest(rand.Reader, &template, priKey)
	if err != nil {
		fmt.Println("Failed to create certificate request.")
		os.Exit(1)
	}

	//Write csr into a file name csr.pem
	csrFile, err := os.Create("csr.pem")
	if err != nil {
		log.Fatalln("Error creating empty csr.pem file")
	}
	defer csrFile.Close()
	pem.Encode(csrFile, &pem.Block{Type: "CERTIFICATE REQUEST", Bytes: csrBytes})

	//Create a http POST request
	request, err := http.NewRequest("POST", "https://crypt.invariant.dev/sign", csrFile)
	if err != nil {
		log.Fatalln("error creating post request:", err)
	}
	request.Header.Add("Csr-Auth-Code", "12345")
	request.Header.Add("Content-Disposition", "csr.pem")

	//Send the POST request
	client1 := &http.Client{}
	res, err := client1.Do(request)
	if err != nil {
		log.Fatalln("Error sending post request", err)
	}
	defer res.Body.Close()

	//Read the returned data from CA into signedCert.pem
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("Error reading response body", err)
	}

	signedFile, err := os.Create("signedCert.pem")
	if err != nil {
		log.Fatalln("Error creating empty signedCert.pem file")
	}
	defer signedFile.Close()
	pem.Encode(signedFile, &pem.Block{Type: "CERTIFICATE", Bytes: body})

	//Verify cert with CA
	cert, err := tls.LoadX509KeyPair("signedCert.pem", "clientKey.pem")
	if err != nil {
		log.Fatalln("Error loading key pair:", err)
	}

	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				Certificates: []tls.Certificate{cert},
			},
		},
	}

	resp, err := client.Get("https://mtls.invariant.dev")
	if err != nil {
		log.Fatalln("error sending get request:", err)
	}
	if resp.StatusCode > 299 {
		log.Fatalln("error sending http request, status code:", resp.StatusCode)
	}

	io.Copy(os.Stdout, resp.Body)

}
