package test

import (
	"testing"
	"io/ioutil"
	"log"	
	"bytes"
	"strings"

	"github.com/gruntwork-io/terratest/modules/terraform"	
	"golang.org/x/crypto/ssh"	
	//"golang.org/x/crypto/ssh/knownhosts"

)

func TestTerraformInstance(t *testing.T) {
	terraformOptions := &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		TerraformDir: "../instance",
	}

	// Clean up resources with "terraform destroy" at the end of the test.
	defer terraform.Destroy(t, terraformOptions)

	// Run "terraform init" and "terraform apply". Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables and check they have the expected values.
	testSSH(t, terraformOptions)	
}


func testSSH(t *testing.T, terraformOptions *terraform.Options){			
	out := terraform.Output(t, terraformOptions, "public_ip")
	ips := strings.Split(out, "\n")
	ips1 := ips[1]
	ips2 := strings.Split(ips1, "\"")
	publicInstanceIP := ips2[1]
	log.Printf(publicInstanceIP)
	// A public key may be used to authenticate against the remote
	// server by using an unencrypted PEM-encoded private key file.
	//
	// If you have an encrypted private key, the crypto/x509 package
	// can be used to decrypt it.
	key, err := ioutil.ReadFile("C:\\Users\\matias.araya.cohen\\OneDrive - Accenture\\Documents\\SSH Keys\\vmInstanceOracleCloud.pem")
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}		

	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}

	/*
	hostKeyCallback, err := knownhosts.New("C:\\Users\\matias.araya.cohen\\.ssh\\known_hosts")
	if err != nil {
		log.Fatal(err)
	}
	*/

	config := &ssh.ClientConfig{
		User: "opc",
		Auth: []ssh.AuthMethod{
			// Use the PublicKeys method for remote authentication.
			ssh.PublicKeys(signer),
		},
		//HostKeyCallback: hostKeyCallback,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect to the remote server and perform the SSH handshake.
	client, err := ssh.Dial("tcp", publicInstanceIP+":22", config)
	if err != nil {
		log.Fatalf("unable to connect: %v", err)
	}
	defer client.Close()

	session, err := client.NewSession()
    if err != nil {
        log.Fatalf("session failed:%v", err)
    }
    defer session.Close()
    var stdoutBuf bytes.Buffer
    session.Stdout = &stdoutBuf
    err = session.Run("ls -l")
    if err != nil {
        log.Fatalf("Run failed:%v", err)
    }
    log.Printf(">%s", stdoutBuf.String())
}