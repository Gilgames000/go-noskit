package gfclient

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gilgames000/go-noskit/actions"
	"github.com/gilgames000/go-noskit/errors"

	"github.com/parnurzeal/gorequest"
)

type GFClient struct {
}

// Authenticate proceeds to authenticate the user to the Gameforge servers.
// An error will be returned if the connection can't be established, the
// server response's status code is not 201 (session created), or the response
// body (JSON) can't be parsed.
// The user needs to be authenticated in order to request a login code.
func (c *GFClient) Authenticate(user actions.User, serverLang string) (string, string, error) {
	url := "https://spark.gameforge.com/api/v1/auth/thin/sessions"
	reqBody := fmt.Sprintf(
		`{"gfLang":"%s","identity":"%s","locale":"%s","password":"%s","platformGameId":"dd4e22d6-00d1-44b9-8126-d8b40e0cd7c9"}`,
		serverLang,
		user.Email,
		user.Locale,
		user.Password,
	)

	request := gorequest.New()
	resp, _, errs := request.Post(url).
		Set("Content-Type", "application/json").
		Set("User-Agent", "Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36").
		Set("TNT-Installation-Id", user.InstallationUUID).
		Set("Origin", "spark://www.gameforge.com").
		Send(reqBody).
		End()
	defer resp.Body.Close()
	if errs != nil {
		var e strings.Builder
		for _, err := range errs {
			e.WriteString(err.Error())
			e.WriteString("\n")
		}
		return "", "", errors.New(e.String())
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}

	if resp.StatusCode != 201 {
		return "", "", errors.New(fmt.Sprintf(
			"failed to authenticate user\nstatus code: %d\nresponse body: %s",
			resp.StatusCode,
			body,
		))
	}

	b := make(map[string]interface{})
	err = json.Unmarshal(body, &b)
	if err != nil {
		return "", "", err
	}

	token, ok := b["token"].(string)
	if !ok {
		return "", "", errors.New("couldn't fetch token from server response")
	}

	accountID, ok := b["platformGameAccountId"].(string)
	if !ok {
		return "", "", errors.New("couldn't fetch accountID from server response")
	}

	return token, accountID, nil
}

// GetLoginCode returns the code (UUID) used to log in to NosTale login server.
// This code needs to be encoded to Base64 before it can be used in the login packet.
// The empty string and an error will be returned if the connection can't be
// established, the server response's status code is not 201 (login code
// generated), or the response body (JSON) can't be parsed.
// The user needs to be authenticated in order to request a login code. If the
// session has expired or the user hasn't been authenticated yet, the server
// will reply with a status code 503 (could not fetch session).
func (c *GFClient) GetLoginCode(user actions.User, token, accountID string) (string, error) {
	url := "https://spark.gameforge.com/api/v1/auth/thin/codes"
	reqBody := fmt.Sprintf(
		`{"platformGameAccountId":"%s"}`,
		accountID,
	)

	request := gorequest.New()
	resp, _, errs := request.Post(url).
		Set("Content-Type", "application/json").
		Set("User-Agent", "TNTClientMS2/1.3.39").
		Set("TNT-Installation-Id", user.InstallationUUID).
		Set("Origin", "spark://www.gameforge.com").
		Set("Connection", "Keep-Alive").
		Set("Authorization", "Bearer "+token).
		Send(reqBody).
		End()
	defer resp.Body.Close()
	if errs != nil {
		var e strings.Builder
		for _, err := range errs {
			e.WriteString(err.Error())
			e.WriteString("\n")
		}
		return "", errors.New(e.String())
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 201 {
		return "", errors.New(fmt.Sprintf("failed to retrieve login code\nstatus code: %d\nresponse body: %s", resp.StatusCode, body))
	}

	b := make(map[string]interface{})
	err = json.Unmarshal(body, &b)
	if err != nil {
		return "", err
	}
	code, _ := b["code"].(string)

	return code, nil
}

// GetLoginCodeHex returns the Base64-encoded code (UUID) used to log in to
// NosTale login server.
// The empty string and an error will be returned if the connection can't be
// established, the server response's status code is not 201 (login code
// generated), or the response body (JSON) can't be parsed.
// The user needs to be authenticated in order to request a login code. If the
// session has expired or the user hasn't been authenticated yet, the server
// will reply with a status code 503 (could not fetch session).
func (c *GFClient) GetLoginCodeHex(user actions.User, token, accountID string) (string, error) {
	code, err := c.GetLoginCode(user, token, accountID)
	if err != nil {
		return "", err
	}

	return strings.ToUpper(hex.EncodeToString([]byte(code))), nil
}
