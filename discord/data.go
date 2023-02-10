package discord

import "strings"

type User struct {
	ID               string `json:"id,omitempty"`
	Username         string `json:"username,omitempty"`
	Discriminator    string `json:"discriminator,omitempty"`
	AvatarHash       string `json:"avatar,omitempty"`
	DisplayName      string `json:"display_name,omitempty"`
	AvatarDecoration string `json:"avatar_decoration,omitempty"`
	PublicFlags      int64  `json:"public_flags,omitempty"`
	Flags            int64  `json:"flags,omitempty"`
	BannerHash       string `json:"banner,omitempty"`
	BannerColor      string `json:"banner_color,omitempty"`
	AccentColor      int64  `json:"accent_color,omitempty"`
	Bio              string `json:"bio,omitempty"`
}


type ConnectedAccount struct {
	Type     string
	ID       string
	Name     string
	Verified bool
	Metadata map[string]string
}

type Guild struct {
	ID   string
	Nick string
}

type UserProfile struct {
	Bio         string
	AccentColor int64
	BannerHash  string
}

type Profile struct {
	User                  User               `json:"user,omitempty"`
	ConnectedAccounts     []ConnectedAccount `json:"connected_accounts,omitempty"`
	PremiumSince          string             `json:"premium_since,omitempty"`
	PremiumType           int                `json:"premium_type,omitempty"`
	PremiumGuildSince     string             `json:"premium_guild_since,omitempty"`
	ProfileThemeExpBucket int                `json:"profile_theme_experiment_bucket,omitempty"`
	MutualGuilds          []Guild            `json:"mutual_guilds,omitempty"`
	Userprofile           UserProfile        `json:"user_profile"`
}



/* Base Discord API stuff */

const BaseURL = "https://discord.com/api"

var ApiVersions = map[string]string{
		"10": "v10",
		"9":  "v9",
		"8":  "v8",
		"7":  "v7",
		"6":  "v6",
		"5":  "v5",
		"4":  "v4",
		"3":  "v3",
	}

type ApiEndpoint struct {
	ApiVersion   string
	Url          string
	Replacements map[string]string
	FullURL      string
}



/* Image Data */

func IsGif(imgHash string) bool {
	return strings.HasPrefix(imgHash, "a_")
}


/* Flags */

const (
	FUser_staff = 1 << 0
	FUser_partner = 1 << 1
	FUser_hypesquad = 1 << 2
	FUser_bughunter_1 = 1 << 3
	FUser_bughunter_2 = 1 << 14
	FUser_hypesquad_bravery = 1 << 6
	FUser_hypesquad_brilliance = 1 << 7
	FUser_hypesquad_balance = 1 << 8
	FUser_premium_eary_supporter = 1 << 9
	FUser_team_pseudo_user = 1 << 10
	FUser_verified_bot = 1 << 16
	FUser_verified_developer = 1 << 17
	FUser_certified_moderator = 1 << 18
	FUser_bot_http_interaction = 1 << 19
	FUser_active_developer = 1 << 22
)



// (Misc) Types

const (
	Premium_Type_None = iota
	Premium_Type_Nitro_Classic
	Premium_Type_Nitro
	Premium_Type_Nitro_Basic
)

const (
	Visibility_Type_None = iota
	Visibility_Type_Everyone
)
