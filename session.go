package kaltura

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"time"
)

//Session - Definition for Session object.
type Session struct {
	UserID     string
	Type       int
	Duration   int64
	Privileges string
	Value      string
}

//SessionStart - Start a new session
func (k *Kaltura) SessionStart() {
	randomNumber := strconv.Itoa(rand.Intn(32000))
	expiry := strconv.FormatInt(k.Session.Duration+time.Now().Unix(), 10)

	sessionString := k.PartnerID + ";" +
		k.PartnerID + ";" +
		expiry + ";" +
		strconv.Itoa(k.Session.Type) + ";" +
		randomNumber + ";" +
		k.Session.UserID + ";"

	if k.Session.Privileges != "" {
		sessionString += k.Session.Privileges + ";"
	}

	hash := sha1.New()
	io.WriteString(hash, k.Secret+sessionString)
	signature := fmt.Sprintf("%x", hash.Sum(nil))

	k.Session.Value = base64.StdEncoding.EncodeToString([]byte(signature + "|" + sessionString))
}

//SessionEnd - End kaltura session
func (k *Kaltura) SessionEnd() {
	k.DoRequest("GET", "session", "get", nil)
}

//SessionInfo - Get the kaltura Session information
func (k *Kaltura) SessionInfo() string {
	return k.DoRequest("GET", "session", "get", nil)
}
