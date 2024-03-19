package configbase_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mtechnavi/sharelib/configbase"
	"os"
	"time"
)

type Config struct {
	LogLevel string

	LogFile string `env:",default=-"`

	Sendgrid *SendgridConfig

	OAuth *OAuthConfig `env:"OAUTH"`
}

type SendgridConfig struct {
	APIKey string `env:"APIKEY"`

	RetryCount int
}

type OAuthConfig struct {
	EnableSingleSignOn bool `env:"ENABLE_SSO"`

	ClientID string

	ClientSecret string
}

func Example() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	os.Setenv("LOG_LEVEL", "info")
	os.Setenv("SENDGRID_APIKEY", "****")
	os.Setenv("SENDGRID_RETRY_COUNT", "30")
	os.Setenv("OAUTH_ENABLE_SSO", "true")
	os.Setenv("OAUTH_CLIENT_ID", "example@centsys.jp")
	os.Setenv("OAUTH_CLIENT_SECRET", "####")

	var cfg Config
	if err := configbase.Load(ctx, &cfg); err != nil {
		panic(err)
	}

	b, err := json.MarshalIndent(&cfg, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	// Output:
	// {
	//   "LogLevel": "info",
	//   "LogFile": "-",
	//   "Sendgrid": {
	//     "APIKey": "****",
	//     "RetryCount": 30
	//   },
	//   "OAuth": {
	//     "EnableSingleSignOn": true,
	//     "ClientID": "example@centsys.jp",
	//     "ClientSecret": "####"
	//   }
	// }
}

func writeTempFile(s string) string {
	wc, err := os.CreateTemp(os.TempDir(), "tmp.****")
	if err != nil {
		panic(err)
	}
	defer wc.Close()

	if _, err := io.WriteString(wc, s); err != nil {
		panic(err)
	}

	return wc.Name()
}

func Example_fromFile() {
	apiKeyFile := writeTempFile("****")
	clientSecretFile := writeTempFile("####")

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	os.Setenv("LOG_LEVEL", "info")
	os.Setenv("SENDGRID_APIKEY_FROM_FILE", apiKeyFile)
	os.Setenv("SENDGRID_RETRY_COUNT", "30")
	os.Setenv("OAUTH_ENABLE_SSO", "true")
	os.Setenv("OAUTH_CLIENT_ID", "example@centsys.jp")
	os.Setenv("OAUTH_CLIENT_SECRET_FROM_FILE", clientSecretFile)

	var cfg Config
	if err := configbase.Load(ctx, &cfg); err != nil {
		panic(err)
	}

	b, err := json.MarshalIndent(&cfg, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	// Output:
	// {
	//   "LogLevel": "info",
	//   "LogFile": "-",
	//   "Sendgrid": {
	//     "APIKey": "****",
	//     "RetryCount": 30
	//   },
	//   "OAuth": {
	//     "EnableSingleSignOn": true,
	//     "ClientID": "example@centsys.jp",
	//     "ClientSecret": "####"
	//   }
	// }
}
