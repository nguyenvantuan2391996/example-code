import (
	"crypto/rand"
	"crypto/rsa"
	"crypto"
	"crypto/sha256"
)

func GenerateRSAKey(bits int) (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, bits)
}

func RSAEncrypt(pub *rsa.PublicKey, msg []byte) ([]byte, error) {
	return rsa.EncryptOAEP(sha256.New(), rand.Reader, pub, msg, nil)
}

func RSADecrypt(priv *rsa.PrivateKey, ciphertext []byte) ([]byte, error) {
	return rsa.DecryptOAEP(sha256.New(), rand.Reader, priv, ciphertext, nil)
}

func RSASign(priv *rsa.PrivateKey, msg []byte) ([]byte, error) {
	h := sha256.Sum256(msg)
	return rsa.SignPSS(rand.Reader, priv, crypto.SHA256, h[:], nil)
}

func RSAVerify(pub *rsa.PublicKey, msg, signature []byte) error {
	h := sha256.Sum256(msg)
	return rsa.VerifyPSS(pub, crypto.SHA256, h[:], signature, nil)
}
