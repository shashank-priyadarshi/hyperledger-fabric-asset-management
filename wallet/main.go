package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Identity represents the structure of an identity in the wallet
type Identity struct {
	CertificatePEM []byte
	PrivateKeyPEM  []byte
	MSPID          string
	Type           string
}

// Wallet represents a collection of identities
type Wallet struct {
	Identities map[string]Identity // Map of identity label to Identity
}

// NewWallet creates a new empty wallet
func NewWallet() *Wallet {
	return &Wallet{
		Identities: make(map[string]Identity),
	}
}

// AddIdentity adds a new identity to the wallet with a given label
func (w *Wallet) AddIdentity(label string, identity Identity) {
	w.Identities[label] = identity
}

// GetIdentity retrieves an identity from the wallet using its label
func (w *Wallet) GetIdentity(label string) (Identity, error) {
	identity, ok := w.Identities[label]
	if !ok {
		return Identity{}, fmt.Errorf("identity not found")
	}
	return identity, nil
}

// SaveWalletToFile saves the wallet to a file
func (w *Wallet) SaveWalletToFile(filePath string) error {
	// Marshal the wallet to JSON or any desired format
	// In this example, we're saving as a simple text file
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	for label, identity := range w.Identities {
		// Write each identity to the file
		file.WriteString(fmt.Sprintf("Label: %s\n", label))
		file.WriteString(fmt.Sprintf("MSPID: %s\n", identity.MSPID))
		file.WriteString(fmt.Sprintf("Type: %s\n", identity.Type))
		file.WriteString(fmt.Sprintf("Certificate: %s\n", string(identity.CertificatePEM)))
		file.WriteString(fmt.Sprintf("PrivateKey: %s\n", string(identity.PrivateKeyPEM)))
		file.WriteString("\n")
	}
	return nil
}

func main() {
	// Create a new empty wallet
	wallet := NewWallet()

	// Example identity data
	certificatePEM, _ := ioutil.ReadFile("/Users/coffeebeans/hyperledger-fabric-asset-management/certificate.pem")
	privateKeyPEM, _ := ioutil.ReadFile("/Users/coffeebeans/hyperledger-fabric-asset-management/private_key.pem")

	// Create an identity
	newIdentity := Identity{
		CertificatePEM: certificatePEM,
		PrivateKeyPEM:  privateKeyPEM,
		MSPID:          "Org1MSP",
		Type:           "X.509",
	}

	// Add the identity to the wallet
	wallet.AddIdentity("User1@org1.example.com", newIdentity)

	// Save the wallet to a file
	err := wallet.SaveWalletToFile("/Users/coffeebeans/hyperledger-fabric-asset-management/wallet.txt")
	if err != nil {
		fmt.Println("Error saving wallet:", err)
		return
	}

	// Retrieve an identity from the wallet
	retrievedIdentity, err := wallet.GetIdentity("User1@org1.example.com")
	if err != nil {
		fmt.Println("Error retrieving identity:", err)
		return
	}

	// Display retrieved identity details
	fmt.Println("Retrieved Identity:")
	fmt.Printf("Label: %s\n", "User1@org1.example.com")
	fmt.Printf("MSPID: %s\n", retrievedIdentity.MSPID)
	fmt.Printf("Type: %s\n", retrievedIdentity.Type)
	fmt.Printf("Certificate: %s\n", string(retrievedIdentity.CertificatePEM))
	fmt.Printf("PrivateKey: %s\n", string(retrievedIdentity.PrivateKeyPEM))
}
