package bearychat

// Team information
type Team struct {
	Id          string `json:"id"`
	Subdomain   string `json:"subdomain"`
	Name        string `json:"name"`
	UserId      string `json:"uid"`
	Description string `json:"description"`
	EmailDomain string `json:"email_domain"`
	Inactive    bool   `json:"inactive"`
	CreatedAt   string `json:"created"` // TODO parse date
	UpdatedAt   string `json:"updated"` // TODO parse date
}

const (
	UserRoleOwner   = "owner"
	UserRoleAdmin   = "admin"
	UserRoleNormal  = "normal"
	UserRoleVisitor = "visitor"
)

const (
	UserTypeNormal    = "normal"
	UserTypeAssistant = "assistant"
	UserTypeHubot     = "hubot"
)

// User information
type User struct {
	Id         string `json:"id"`
	TeamId     string `json:"team_id"`
	VChannelId string `json:"vchannel_id"`
	Name       string `json:"name"`
	FullName   string `json:"full_name"`
	Email      string `json:"email"`
	AvatarUrl  string `json:"avatar_url"`
	Role       string `json:"role"`
	Type       string `json:"type"`
	Conn       string `json:"conn"`
	CreatedAt  string `json:"created"` // TODO parse date
	UpdatedAt  string `json:"updated"` // TODO parse date
}

// IsOnline tells user connection status.
func (u User) IsOnline() bool {
	return u.Conn == "connected"
}

// IsNormal tells if this user a normal user (owner, admin or normal)
func (u User) IsNormal() bool {
	return u.Type == UserTypeNormal && u.Role != UserRoleVisitor
}

// Channel information.
type Channel struct {
	Id         string `json:"id"`
	TeamId     string `json:"team_id"`
	UserId     string `json:"uid"`
	VChannelId string `json:"vchannel_id"`
	Name       string `json:"name"`
	IsPrivate  bool   `json:"private"`
	IsGeneral  bool   `json:"general"`
	Topic      string `json:"topic"`
	CreatedAt  string `json:"created"` // TODO parse date
	UpdatedAt  string `json:"updated"` // TODO parse date
}

// AttachedFile RTM struct in Attachments
type AttachedFile struct {
	Category    string `json:"category"`
	Created     string `json:"created"`
	Deleted     bool   `json:"deleted"`
	Description string `json:"description"`
	Height      int    `json:"height"`
	ID          string `json:"id"`
	// image URL
	ImageURL string `json:"image_url"`
	Inactive bool   `json:"inactive"`
	IsPublic bool   `json:"is_public"`
	Key      string `json:"key"`
	// File MIME
	MIME        string `json:"mime"`
	Name        string `json:"name"`
	Orientation int    `json:"orientation"`
	Original    bool   `json:"original"`
	// preview URL
	PreviewURL string `json:"preview_url"`
	// File Size in byte
	Size       int    `json:"size"`
	Source     string `json:"source"`
	TeamID     string `json:"team_id"`
	Title      string `json:"title"`
	Type       string `json:"type"`
	UID        string `json:"uid"`
	Updated    string `json:"updated"`
	UploadZone string `json:"upload_zone"`
	URL        string `json:"url"`
	Width      int    `json:"width"`
}

// UpdateAttachments RTM msg UpdateAttachments
type UpdateAttachments struct {
	// very odd data structure...
	Data struct {
		Attachments []struct {
			File          *AttachedFile `json:"file"`
			ReferText     string        `json:"refer_text"`
			ReferTextI18N struct {
				En string `json:"en"`
				Zh string `json:"zh-CN"`
			} `json:"refer_text_i18n"`
			Subtype string `json:"subtype"`
			Text    string `json:"text"`
			Type    string `json:"type"`
			UID     string `json:"uid"`
		} `json:"attachments"`
		Created         string `json:"created"`
		CreatedTs       int64  `json:"created_ts"`
		DisableMarkdown bool   `json:"disable_markdown"`
		Edited          bool   `json:"edited"`
		ID              string `json:"id"`
		IsChannel       bool   `json:"is_channel"`
		Key             string `json:"key"`
		ReferKey        string `json:"refer_key"`
		Subtype         string `json:"subtype"`
		TeamID          string `json:"team_id"`
		Text            string `json:"text"`
		UID             string `json:"uid"`
		Updated         string `json:"updated"`
		VChannelID      string `json:"vchannel_id"`
	} `json:"data"`
	Ts   int64  `json:"ts"`
	Type string `json:"type"`
}
