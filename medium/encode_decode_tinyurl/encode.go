package tinyUrl

import (
	"encoding/base64"
)

type Codec struct {
	urlToEncoded map[string]string
	encodedToUrl map[string]string
}

func Constructor() Codec {
	return Codec{urlToEncoded: make(map[string]string), encodedToUrl: make(map[string]string)}
}

func (this *Codec) encode(longUrl string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(longUrl))[:6]
	this.urlToEncoded[longUrl] = encoded
	this.encodedToUrl[encoded] = longUrl

	return this.urlToEncoded[longUrl]
}

func (this *Codec) decode(shortUrl string) string {

	return this.encodedToUrl[shortUrl]
}
