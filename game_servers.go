package steamapi

import (
	"net/url"
	"strconv"
)

type gameServerInfoResponse struct {
	Response struct {
		Success bool
		Servers []GameServerKeyInfo
	}
}

// CreateResponse contains the information about the created GameServerLoginToken
type CreateResponse struct {
	SteamID    string
	LoginToken string `json:"login_token"`
}

type gameServerCreateResponse struct {
	Response CreateResponse
}

type gameServerResetResponse struct {
	Response struct {
		LoginToken string `json:"login_token"`
	}
}

// GameServerKeyInfo contains the information about a GameServerLoginToken
type GameServerKeyInfo struct {
	SteamID    string
	AppID      uint32
	LoginToken string `json:"login_token"`
	Memo       string
	IsDeleted  bool   `json:"is_deleted"`
	IsExpired  bool   `json:"is_expired"`
	LastLogin  uint32 `json:"rt_last_logon"`
}

// GetGameServerInfo retrieves all GameServerLoginTokens and their info for a given key (account)
func GetGameServerInfo(apiKey string) ([]GameServerKeyInfo, error) {
	getServerInfo := NewSteamMethod("IGameServersService", "GetAccountList", 1)

	data := url.Values{}
	data.Add("key", apiKey)

	var resp gameServerInfoResponse
	err := getServerInfo.Request(data, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Response.Servers, nil
}

func CreateGameServerKey(apiKey string, appId uint32, memo string) (*CreateResponse, error) {
	createServerInfo := NewSteamMethod("IGameServersService", "CreateAccount", 1)

	data := url.Values{}
	data.Add("key", apiKey)
	data.Add("appid", strconv.FormatUint(uint64(appId), 10))
	data.Add("memo", memo)

	var resp gameServerCreateResponse
	err := createServerInfo.Request(data, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Response, nil
}

func ResetGameServerKey(apiKey string, steamId string) (string, error) {
	createServerInfo := NewSteamMethod("IGameServersService", "ResetLoginToken", 1)

	data := url.Values{}
	data.Add("key", apiKey)
	data.Add("steamId", steamId)

	var resp gameServerResetResponse
	err := createServerInfo.Post(data, &resp)
	if err != nil {
		return "", err
	}
	return resp.Response.LoginToken, nil
}

func DeleteGameServerKey(apiKey string, steamId string) error {
	createServerInfo := NewSteamMethod("IGameServersService", "DeleteAccount", 1)

	data := url.Values{}
	data.Add("key", apiKey)
	data.Add("steamId", steamId)

	var resp interface{}
	err := createServerInfo.Post(data, &resp)
	if err != nil {
		return err
	}
	return nil
}
