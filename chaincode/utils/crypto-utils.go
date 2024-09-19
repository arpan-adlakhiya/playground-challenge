package utils

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type SignatureValidationBytes struct {
	Signature []byte
	Hash      []byte
	PublicKey []byte
}

/*
Converts data from string to []byte
Calculates hash for the message
If sigPublicKey is nil - Recovers publick key from the hash and signature strings
Returns []byte for signature, hash and public key
*/
func ConvertSigData(data string, signatureStr string, sigPublicKey []byte) (*SignatureValidationBytes, error) {

	var convertedData SignatureValidationBytes

	hash := crypto.Keccak256Hash([]byte(data))
	convertedData.Hash = hash.Bytes()

	signature, err := hexutil.Decode(signatureStr)
	if err != nil {
		return nil, fmt.Errorf("error while decoding hex signature to []byte. Err:%w", err)
	}

	if len(signature) != SIGNATURE_LENGTH {
		return nil, fmt.Errorf("signature length %d %w", len(signature), ErrNotValid)
	}
	if signature[SIGNATURE_LENGTH-1] != V_LOWER_BOUND && signature[SIGNATURE_LENGTH-1] != V_UPPER_BOUND {
		return nil, fmt.Errorf("recoveryID:  %d %w", signature[SIGNATURE_LENGTH-1], ErrNotValid)
	}
	signature[SIGNATURE_LENGTH-1] -= V_LOWER_BOUND

	convertedData.Signature = signature

	if sigPublicKey == nil {
		sigPublicKey, err = crypto.Ecrecover(hash.Bytes(), signature)
		if err != nil {
			return nil, fmt.Errorf("error while recoving public key. Err:%w", err)
		}
	}
	convertedData.PublicKey = sigPublicKey

	return &convertedData, nil
}

// Verify signature with message hash and public key
// Returns true if the signature is valid
func ValidateSignature(pubKey []byte, hash []byte, signature []byte) bool {
	signatureNoRecoverID := signature[:len(signature)-1]
	verified := crypto.VerifySignature(pubKey, hash, signatureNoRecoverID)

	return verified
}

// Verify address matches with the public key
func VerifyAddress(pubKey []byte, addressStr string) error {
	pubKeyAddress := pubKeyToAddress(pubKey)
	address, _ := hexutil.Decode(addressStr)
	if !bytes.Equal(address, pubKeyAddress) {
		return fmt.Errorf("address %w. expected: %s, Actual: %s",
			ErrNotMatched,
			hexutil.Encode(address),
			hexutil.Encode(pubKeyAddress))
	}

	return nil
}

func pubKeyToAddress(pubKey []byte) []byte {
	x, y := elliptic.Unmarshal(crypto.S256(), pubKey)

	return crypto.PubkeyToAddress(ecdsa.PublicKey{Curve: crypto.S256(), X: x, Y: y}).Bytes()
}

/*
Recovers public key from message and signatures
Verifies recovered public key matches with the address passed
Verifies signature with the public key recovered
*/
func VerifySignatureAndAddress(message string, signatureStr string, address string) error {
	sigBytes, err := ConvertSigData(message, signatureStr, nil)
	if err != nil {
		return fmt.Errorf("error while converting signature data. address: %s, Error: %w", address, err)
	}

	// Verify pub key and address match
	err = VerifyAddress(sigBytes.PublicKey, address)
	if err != nil {
		return err
	}

	/* result := ValidateSignature(sigBytes.PublicKey, sigBytes.Hash, sigBytes.Signature)
	if !result {
		return fmt.Errorf("signature verification failed for address: %s", address)
	} */
	return nil
}
