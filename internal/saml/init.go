package saml

import (
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"net/url"
)

type InitRequest struct {
	IDPRedirectURL string
	SPEntityID     string
	RelayState     string
}

type InitResponse struct {
	URL string
}

func Init(req *InitRequest) (*InitResponse, error) {
	redirectURL, err := url.Parse(req.IDPRedirectURL)
	if err != nil {
		return nil, fmt.Errorf("parse idp redirect url: %w", err)
	}

	var samlReq samlRequest
	samlReq.Issuer.Name = req.SPEntityID
	samlReqData, err := xml.Marshal(samlReq)

	if err != nil {
		panic(err)
	}

	query := redirectURL.Query()
	query.Set("SAMLRequest", base64.URLEncoding.EncodeToString(samlReqData))
	query.Set("RelayState", req.RelayState) // todo sign this to prevent tampering
	redirectURL.RawQuery = query.Encode()

	return &InitResponse{URL: redirectURL.String()}, nil
}

type samlRequest struct {
	XMLName xml.Name `xml:"urn:oasis:names:tc:SAML:2.0:protocol AuthnRequest"`
	Issuer  struct {
		XMLName xml.Name `xml:"urn:oasis:names:tc:SAML:2.0:assertion Issuer"`
		Name    string   `xml:",chardata"`
	} `xml:"Issuer"`
}