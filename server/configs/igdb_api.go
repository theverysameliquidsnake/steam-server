package configs

import (
	"fmt"
	"io"
	"os"

	jsoniter "github.com/json-iterator/go"
	"github.com/theverysameliquidsnake/steam-db/pkg/utils"
)

var token string

func InitIGDBToken() error {
	client, err := utils.UseProxyClient()
	if err != nil {
		return err
	}

	response, err := client.Post(fmt.Sprintf("https://id.twitch.tv/oauth2/token?client_id=%s&client_secret=%s&grant_type=client_credentials", os.Getenv("IGDB_ID"), os.Getenv("IGDB_SECRET")), "text/plain", nil)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	token = jsoniter.Get(body, "access_token").ToString()

	return nil
}

func GetIGDBHeaders() map[string]string {
	headers := make(map[string]string)
	headers["Client-ID"] = os.Getenv("IGDB_ID")
	headers["Authorization"] = fmt.Sprintf("Bearer %s", token)

	return headers
}
