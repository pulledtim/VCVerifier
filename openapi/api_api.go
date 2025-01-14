/*
 * vcverifier
 *
 * Backend component to verify credentials
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fiware/VCVerifier/logging"
	"github.com/fiware/VCVerifier/verifier"

	"github.com/gin-gonic/gin"
)

var apiVerifier verifier.Verifier

var ErrorMessagNoGrantType = ErrorMessage{"no_grant_type_provided", "Token requests require a grant_type."}
var ErrorMessageNoCode = ErrorMessage{"no_code_provided", "Token requests require a code."}
var ErrorMessageNoRedircetUri = ErrorMessage{"no_redirect_uri_provided", "Token requests require a redirect_uri."}
var ErrorMessageNoState = ErrorMessage{"no_state_provided", "Authentication requires a state provided as query parameter."}
var ErrorMessageNoToken = ErrorMessage{"no_token_provided", "Authentication requires a token provided as a form parameter."}
var ErrorMessageNoCallback = ErrorMessage{"NoCallbackProvided", "A callback address has to be provided as query-parameter."}
var ErrorMessageUnableToDecodeToken = ErrorMessage{"invalid_token", "Token could not be decoded."}
var ErrorMessageUnableToDecodeCredential = ErrorMessage{"invalid_token", "Could not read the credential(s) inside the token."}
var ErrorMessageUnableToDecodeHolder = ErrorMessage{"invalid_token", "Could not read the holder inside the token."}

func getApiVerifier() verifier.Verifier {
	if apiVerifier == nil {
		apiVerifier = verifier.GetVerifier()
	}
	return apiVerifier
}

// GetToken - Token endpoint to exchange the authorization code with the actual JWT.
func GetToken(c *gin.Context) {

	logging.Log().Debugf("%v", c.Request)
	grantType, grantTypeExists := c.GetPostForm("grant_type")
	if !grantTypeExists {
		logging.Log().Debug("No grant_type present in the request.")
		c.AbortWithStatusJSON(400, ErrorMessagNoGrantType)
		return
	}
	code, codeExists := c.GetPostForm("code")
	if !codeExists {
		logging.Log().Debug("No code present in the request.")
		c.AbortWithStatusJSON(400, ErrorMessageNoCode)
		return
	}
	redirectUri, redirectUriExists := c.GetPostForm("redirect_uri")
	if !redirectUriExists {
		logging.Log().Debug("No redircet_uri present in the request.")
		c.AbortWithStatusJSON(400, ErrorMessageNoRedircetUri)
		return
	}
	jwt, expiration, err := getApiVerifier().GetToken(grantType, code, redirectUri)

	if err != nil {
		c.AbortWithStatusJSON(403, ErrorMessage{Summary: err.Error()})
		return
	}

	c.JSON(http.StatusOK, TokenResponse{"Bearer", float32(expiration), jwt})
}

// StartSIOPSameDevice - Starts the siop flow for credentials hold by the same device
func StartSIOPSameDevice(c *gin.Context) {
	state, stateExists := c.GetQuery("state")
	if !stateExists {
		logging.Log().Debugf("No state was provided.")
		c.AbortWithStatusJSON(400, ErrorMessage{"no_state_provided", "Authentication requires a state provided as query parameter."})
		return
	}
	redirectPath, redirectPathExists := c.GetQuery("redirect_path")
	if !redirectPathExists {
		redirectPath = "/"
	}

	protocol := "https"
	if c.Request.TLS == nil {
		protocol = "http"
	}

	redirect, err := getApiVerifier().StartSameDeviceFlow(c.Request.Host, protocol, state, redirectPath)
	if err != nil {
		logging.Log().Warnf("Error starting the same-device flow. Err: %v", err)
		c.AbortWithStatusJSON(500, ErrorMessage{err.Error(), "Was not able to start the same device flow."})
		return
	}
	c.Redirect(302, redirect)
}

// VerifierAPIAuthenticationResponse - Stores the credential for the given session
func VerifierAPIAuthenticationResponse(c *gin.Context) {
	state, stateExists := c.GetQuery("state")
	if !stateExists {
		c.AbortWithStatusJSON(400, ErrorMessageNoState)
		return
	}

	base64Token, tokenExists := c.GetPostForm("vp_token")
	if !tokenExists {
		logging.Log().Info("No token was provided.")
		c.AbortWithStatusJSON(400, ErrorMessageNoToken)
		return
	}

	handleAuthenticationResponse(c, state, base64Token)
}

// GetVerifierAPIAuthenticationResponse - Stores the credential for the given session
func GetVerifierAPIAuthenticationResponse(c *gin.Context) {
	state, stateExists := c.GetQuery("state")
	if !stateExists {
		c.AbortWithStatusJSON(400, ErrorMessageNoState)
		return
	}

	base64Token, tokenExists := c.GetQuery("vp_token")
	if !tokenExists {
		logging.Log().Info("No token was provided.")
		c.AbortWithStatusJSON(400, ErrorMessageNoToken)
		return
	}

	handleAuthenticationResponse(c, state, base64Token)
}

func handleAuthenticationResponse(c *gin.Context, state string, vpToken string) {
	bytes, err := base64.RawURLEncoding.DecodeString(vpToken)
	if err != nil {
		logging.Log().Infof("Was not able to decode the form string %s. Err: %v", vpToken, err)
		c.AbortWithStatusJSON(400, ErrorMessageUnableToDecodeToken)
		return
	}
	var vpObjectMap map[string]json.RawMessage
	json.Unmarshal(bytes, &vpObjectMap)

	verifiableCredentials := vpObjectMap["verifiableCredential"]
	rawHolder := vpObjectMap["holder"]

	var rawCredentials []map[string]interface{}
	var holder string

	err = json.Unmarshal(verifiableCredentials, &rawCredentials)
	if err != nil {
		logging.Log().Infof("Was not able to decode the credentials from the token %s. Err: %v", vpToken, err)
		c.AbortWithStatusJSON(400, ErrorMessageUnableToDecodeCredential)
		return
	}
	err = json.Unmarshal(rawHolder, &holder)
	if err != nil || holder == "" {
		logging.Log().Infof("Was not able to decode the holder from the token %s. Err: %v", vpToken, err)
		c.AbortWithStatusJSON(400, ErrorMessageUnableToDecodeHolder)
		return
	}

	sameDeviceResponse, err := getApiVerifier().AuthenticationResponse(state, rawCredentials, holder)
	if err != nil {
		logging.Log().Warnf("Was not able to get fullfil the authentication response. Err: %v", err)
		c.AbortWithStatusJSON(400, ErrorMessage{Summary: err.Error()})
		return
	}
	if sameDeviceResponse != (verifier.SameDeviceResponse{}) {
		c.Redirect(302, fmt.Sprintf("%s?state=%s&code=%s", sameDeviceResponse.RedirectTarget, sameDeviceResponse.SessionId, sameDeviceResponse.Code))
		return
	}
	logging.Log().Debugf("Successfully authenticated %s.", state)
	c.JSON(http.StatusOK, gin.H{})
}

// VerifierAPIJWKS - Provides the public keys for the given verifier, to be used for verifing the JWTs
func VerifierAPIJWKS(c *gin.Context) {
	c.JSON(http.StatusOK, getApiVerifier().GetJWKS())
}

// VerifierAPIStartSIOP - Initiates the siop flow and returns the 'openid://...' connection string
func VerifierAPIStartSIOP(c *gin.Context) {
	state, stateExists := c.GetQuery("state")
	if !stateExists {
		c.AbortWithStatusJSON(400, ErrorMessageNoState)
		// early exit
		return
	}

	callback, callbackExists := c.GetQuery("client_callback")
	if !callbackExists {
		c.AbortWithStatusJSON(400, ErrorMessageNoCallback)
		// early exit
		return
	}
	protocol := "https"
	if c.Request.TLS == nil {
		protocol = "http"
	}
	connectionString, err := getApiVerifier().StartSiopFlow(c.Request.Host, protocol, callback, state)
	if err != nil {
		c.AbortWithStatusJSON(500, ErrorMessage{err.Error(), "Was not able to generate the connection string."})
		return
	}
	c.String(http.StatusOK, connectionString)
}
