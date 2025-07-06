package awsToken

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/aws-sdk-go-v2/service/sts/types"

	"github.com/PuerkitoBio/goquery"
	"gopkg.in/ini.v1"
)

const (
	oktaBaseURL     = "https://yourcompany.okta.com"
	oktaAppURL      = "https://yourcompany.okta.com/app/aws/abc123/sso/saml"
	sessionName     = "TempSession"
	profileName     = "saml"
	durationSeconds = 3600
	writeToFile     = true
)

type OktaAuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Options  struct {
		MultiOptionalFactorEnroll bool `json:"multiOptionalFactorEnroll"`
		WarnBeforePasswordExpired bool `json:"warnBeforePasswordExpired"`
	} `json:"options"`
}

type OktaAuthResponse struct {
	SessionToken string `json:"sessionToken"`
	ExpiresAt    string `json:"expiresAt"`
	Status       string `json:"status"`
}

func getSAMLAssertion(un string, pw string) (string, error) {
	// New session
	client := &http.Client{}

	// Okta Auth request
	reqBody := OktaAuthRequest{
		Username: un,
		Password: pw,
	}
	reqBody.Options.MultiOptionalFactorEnroll = false
	reqBody.Options.WarnBeforePasswordExpired = false

	// Jsonify the request
	bodyBytes, _ := json.Marshal(reqBody)
	authReq, err := http.NewRequest("POST", oktaBaseURL+"/api/v1/authn", bytes.NewReader(bodyBytes))
	if err != nil {
		return "awsToken[66] Request Error", err
	}
	// Set headers
	authReq.Header.Set("Content-Type", "application/json")

	// Send request
	response, err := client.Do(authReq)
	if err != nil {
		return "awsToken[74] Request Error", err
	}
	// Close the stream
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("awsToken[80] Authentication failed: %s", response.Status)
	}

	var authResp OktaAuthResponse
	if err := json.NewDecoder(response.Body).Decode(&authResp); err != nil {
		return "awsToken[85]", err
	}

	if authResp.Status != "SUCCESS" || authResp.SessionToken == "" {
		return "awsToken[89]", errors.New("awsToken[89] Authentication Unsuccessful")
	}

	// Fetch SAML response with one-time session token
	samlURL := fmt.Sprintf("%s?onetimetoken=%s", oktaAppURL, authResp.SessionToken)
	samlResp, err := client.Get(samlURL)
	if err != nil {
		return "awsToken[96] Request Error", err
	}
	// Close response stream
	defer samlResp.Body.Close()

	if samlResp.StatusCode != http.StatusOK {
		return "awsToken[102]", fmt.Errorf("awsToken[102] Failed to get SAML response: %s", samlResp.Status)
	}

	// Parse HTML
	doc, err := goquery.NewDocumentFromReader(samlResp.Body)
	if err != nil {
		return "awsToken[108]", err
	}

	samlValue := ""
	doc.Find("input[name='SAMLResponse']").Each(func(i int, s *goquery.Selection) {
		val, exists := s.Attr("value")
		if exists {
			samlValue = val
		}
	})

	if samlValue == "" {
		return "awsToken[120]", errors.New("awsToken[120] SAMLResponse not found in HTML")
	}

	return samlValue, nil
}

func ListRoles(un string, pw string) {
	// Get saml assertion
	samlAssertion, _ := getSAMLAssertion(un, pw)
	decodeAssertion, err := base64.StdEncoding.DecodeString(samlAssertion)
	if err != nil {
		fmt.Println("awsToken[131]", err)
		return
	}
	roles, err := getAttribute(decodeAssertion, "Role")
	if err != nil {
		log.Fatalf("awsToken[136] %v", err)
	}
	for _, role := range roles {
		fmt.Println(role)
	}

}

func genTokenforAssumeRoleWithSAML(un string, pw string, roleArn string, duration_optional ...int) (*types.Credentials, error) {
	duration := durationSeconds
	if len(duration_optional) > 0 {
		duration = duration_optional[0]
	}
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("awsToken[152] failed to load AWS config: %w", err)
	}
	// Get saml assertion
	samlAssertion, _ := getSAMLAssertion(un, pw)
	decodeAssertion, err := base64.StdEncoding.DecodeString(samlAssertion)
	if err != nil {
		return nil, fmt.Errorf("awsToken[158] : %w", err)
	}

	roles, err := getAttribute(decodeAssertion, "Role")
	if err != nil {
		log.Fatalf("awsToken[163] %v", err)
	}
	var principalArn string
	for _, role := range roles {
		roleName := strings.Split(role, ",")
		if roleName[0] == roleArn {
			principalArn = roleName[1]
		}
	}

	stsClient := sts.NewFromConfig(cfg)
	resp, err := stsClient.AssumeRoleWithSAML(ctx, &sts.AssumeRoleWithSAMLInput{
		PrincipalArn:    aws.String(principalArn),
		RoleArn:         aws.String(roleArn),
		SAMLAssertion:   aws.String(samlAssertion),
		DurationSeconds: aws.Int32(int32(duration)),
	})
	if err != nil {
		return nil, fmt.Errorf("awsToken[181] failed to get AWS saml token: %w", err)
	}

	credFile := filepath.Join(os.Getenv("HOME"), ".aws", "credentials")

	creds := resp.Credentials

	if !writeToFile {
		fmt.Println("awsToken[189] Exiting without writing to file...")
		return creds, nil
	}
	err = writeCreds(profileName, creds, credFile)
	if err != nil {
		return nil, fmt.Errorf("awsToken[194] Failed to write credentials: %v", err)
	}
	fmt.Printf("awsToken[196] ✅ Credentials written to [%s] in %s\n", profileName, credFile)
	return creds, nil
}

func genTokenforAssumeRole(roleArn string, writeToFile_optional ...bool) {
	writetofile := writeToFile
	if len(writeToFile_optional) > 0 {
		writetofile = writeToFile_optional[0]
	}
	ctx := context.Background()
	creds, err := getToken(ctx, roleArn, sessionName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "awsToken[208] Failed to get token: %v\n", err)
		return
	}

	credFile := filepath.Join(os.Getenv("HOME"), ".aws", "credentials")

	if !writetofile {
		fmt.Println("awsToken[215] Exiting without writing to file...")
		return
	}
	err = writeCreds(profileName, creds, credFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "awsToken[220] Failed to write credentials: %v\n", err)
		return
	}

	fmt.Printf("awsToken[224] ✅ Credentials written to [%s] in %s\n", profileName, credFile)
}

func getToken(ctx context.Context, roleArn, sessionName string) (*types.Credentials, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("awsToken[230] failed to load AWS config: %w", err)
	}

	stsClient := sts.NewFromConfig(cfg)

	resp, err := stsClient.AssumeRole(ctx, &sts.AssumeRoleInput{
		RoleArn:         aws.String(roleArn),
		RoleSessionName: aws.String(sessionName),
		DurationSeconds: aws.Int32(durationSeconds),
	})
	if err != nil {
		return nil, fmt.Errorf("awsToken[241] failed to assume role: %w", err)
	}

	return resp.Credentials, nil
}

func writeCreds(profile string, creds *types.Credentials, path string) error {
	cfg, err := ini.Load(path)
	if err != nil {
		if os.IsNotExist(err) {
			cfg = ini.Empty()
		} else {
			return fmt.Errorf("awsToken[253] loading credentials file: %w", err)
		}
	}

	section, err := cfg.GetSection(profile)
	if err != nil {
		section, _ = cfg.NewSection(profile)
	}

	section.Key("aws_access_key_id").SetValue(*creds.AccessKeyId)
	section.Key("aws_secret_access_key").SetValue(*creds.SecretAccessKey)
	section.Key("aws_session_token").SetValue(*creds.SessionToken)
	section.Key("expiration").SetValue(creds.Expiration.Format(time.RFC3339))

	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		return fmt.Errorf("awsToken[268] Error creating directory: %w", err)
	}

	if err := cfg.SaveTo(path); err != nil {
		return fmt.Errorf("awsToken[272] Error saving file: %w", err)
	}

	return nil
}
