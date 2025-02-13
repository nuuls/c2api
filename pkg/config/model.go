package config

type APIConfig struct {
	// Core

	BaseURL                  string `mapstructure:"base-url" json:"base-url"`
	BindAddress              string `mapstructure:"bind-address" json:"bind-address"`
	MaxContentLength         uint64 `mapstructure:"max-content-length" json:"max-content-length"`
	EnableAnimatedThumbnails bool   `mapstructure:"enable-animated-thumbnails" json:"enable-animated-thumbnails"`
	MaxThumbnailSize         uint   `mapstructure:"max-thumbnail-size" json:"max-thumbnail-size"`

	LogLevel       string `mapstructure:"log-level" json:"log-level"`
	LogDevelopment bool   `mapstructure:"log-development" json:"log-development"`

	DSN string `mapstructure:"dsn" json:"dsn"`

	EnablePrometheus      bool   `mapstructure:"enable-prometheus" json:"enable-prometheus"`
	PrometheusBindAddress string `mapstructure:"prometheus-bind-address" json:"prometheus-bind-address"`

	// Secrets

	DiscordToken            string `mapstructure:"discord-token" json:"discord-token"`
	TwitchClientID          string `mapstructure:"twitch-client-id" json:"twitch-client-id"`
	TwitchClientSecret      string `mapstructure:"twitch-client-secret" json:"twitch-client-secret"`
	YoutubeApiKey           string `mapstructure:"youtube-api-key" json:"youtube-api-key"`
	TwitterBearerToken      string `mapstructure:"twitter-bearer-token" json:"twitter-bearer-token"`
	ImgurClientID           string `mapstructure:"imgur-client-id" json:"imgur-client-id"`
	OembedFacebookAppID     string `mapstructure:"oembed-facebook-app-id" json:"oembed-facebook-app-id"`
	OembedFacebookAppSecret string `mapstructure:"oembed-facebook-app-secret" json:"oembed-facebook-app-secret"`
	OembedProvidersPath     string `mapstructure:"oembed-providers-path" json:"oembed-providers-path"`
}
