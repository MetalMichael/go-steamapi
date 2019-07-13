package steamapi

import "net/url"

type gameServerInfoResponse struct {
	Response struct {
		Success bool
		Servers []GameServerKeyInfo
	}
}

// GameServerKeyInfo contains the information about a GameServerLoginToken
type GameServerKeyInfo struct {
	SteamID    string
	AppID      uint16
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
