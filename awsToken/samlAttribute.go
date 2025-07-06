package awsToken

import (
	"encoding/xml"
	"fmt"
)

type AttributeValue struct {
	XMLName xml.Name `xml:"saml2:AttributeValue"`
	Value   string   `xml:",chardata"`
}

type Attribute struct {
	XMLName xml.Name         `xml:"saml2:Attribute"`
	Name    string           `xml:"Name,attr"`
	Values  []AttributeValue `xml:"saml2:AttributeValue"`
}

type AttributeStatement struct {
	XMLName    xml.Name    `xml:"saml2:AttributeStatement"`
	Attributes []Attribute `xml:"saml2:Attribute"`
}

type Assertion struct {
	XMLName            xml.Name           `xml:"saml2:Assertion"`
	AttributeStatement AttributeStatement `xml:"saml2:AttributeStatement"`
}

type Response struct {
	XMLName   xml.Name  `xml:"saml2p:Response"`
	Assertion Assertion `xml:"saml2:Assertion"`
}

func getAttribute(samlXML []byte, attribute string) ([]string, error) {
	var xmlResp Response
	err := xml.Unmarshal([]byte(samlXML), &xmlResp)
	if err != nil {
		return nil, fmt.Errorf("samlAssertion[38] failed to parse XML: %w", err)
	}

	var listAttr []string
	for _, attr := range xmlResp.Assertion.AttributeStatement.Attributes {
		if attr.Name == "https://aws.amazon.com/SAML/Attributes/"+attribute {
			for _, v := range attr.Values {
				listAttr = append(listAttr, v.Value)
			}
		}
	}

	return listAttr, nil
}
