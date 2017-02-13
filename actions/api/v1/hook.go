package v1

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

const signaturePrefix = "sha1="
const signatureLength = 45 // len(SignaturePrefix) + len(hex(sha1))

// Hook is an inbound github webhook
type hook struct {
	Signature string
	Payload   []byte
}

func signBody(secret, body []byte) []byte {
	computed := hmac.New(sha1.New, secret)
	computed.Write(body)
	return []byte(computed.Sum(nil))
}

func (h *hook) SignedBy(secret []byte) bool {
	if len(h.Signature) != signatureLength || !strings.HasPrefix(h.Signature, signaturePrefix) {
		return false
	}

	actual := make([]byte, 20)
	hex.Decode(actual, []byte(h.Signature[5:]))

	return hmac.Equal(signBody(secret, h.Payload), actual)
}

// New reads a Hook from an incoming HTTP Request.
func newHook(req *http.Request) (*hook, error) {
	h := &hook{}
	if !strings.EqualFold(req.Method, "POST") {
		return nil, errors.New("Unknown method!")
	}

	if h.Signature = req.Header.Get("x-hub-signature"); len(h.Signature) == 0 {
		return nil, errors.New("No signature!")
	}

	var err error
	h.Payload, err = ioutil.ReadAll(req.Body)
	return h, err
}

// Parse reads and verifies the hook in an inbound request.
func parseHook(secret []byte, req *http.Request) (*hook, error) {
	h, err := newHook(req)
	if err == nil && !h.SignedBy(secret) {
		return h, errors.New("Invalid signature")
	}
	return h, nil
}
