package bearychat

import (
	"encoding/json"
	"regexp"

	"github.com/pkg/errors"
)

type RTMMessageType string

const (
	RTMMessageTypeUnknown              RTMMessageType = "unknown"
	RTMMessageTypePing                                = "ping"
	RTMMessageTypePong                                = "pong"
	RTMMessageTypeReply                               = "reply"
	RTMMessageTypeOk                                  = "ok"
	RTMMessageTypeP2PMessage                          = "message"
	RTMMessageTypeP2PTyping                           = "typing"
	RTMMessageTypeChannelMessage                      = "channel_message"
	RTMMessageTypeChannelTyping                       = "channel_typing"
	RTMMessageTypeUpdateUserConnection                = "update_user_connection"
	RTMMessageTypeUpdateAttachments                   = "update_attachments"
)

// RTMMessage represents a message entity send over RTM protocol.
type RTMMessage map[string]interface{}

func (m RTMMessage) Type() RTMMessageType {
	if t, present := m["type"]; present {
		if mtype, ok := t.(string); ok {
			return RTMMessageType(mtype)
		}
		if mtype, ok := t.(RTMMessageType); ok {
			return mtype
		}
	}

	return RTMMessageTypeUnknown
}

// Reply a message (with copying type, vchannel_id)
func (m RTMMessage) Reply(text string) RTMMessage {
	reply := RTMMessage{
		"text":        text,
		"vchannel_id": m["vchannel_id"],
	}

	if m.IsP2P() {
		reply["type"] = RTMMessageTypeP2PMessage
		reply["to_uid"] = m["uid"]
	} else {
		reply["type"] = RTMMessageTypeChannelMessage
		reply["channel_id"] = m["channel_id"]
	}

	return reply
}

// Refer a message
func (m RTMMessage) Refer(text string) RTMMessage {
	refer := m.Reply(text)
	refer["refer_key"] = m["key"]

	return refer
}

func (m RTMMessage) IsP2P() bool {
	mt := m.Type()
	if mt == RTMMessageTypeP2PMessage || mt == RTMMessageTypeP2PTyping {
		return true
	}

	return false
}

func (m RTMMessage) IsChatMessage() bool {
	mt := m.Type()
	if mt == RTMMessageTypeP2PMessage || mt == RTMMessageTypeChannelMessage {
		return true
	}

	return false
}

func (m RTMMessage) IsFromUser(u User) bool {
	return m.IsFromUID(u.Id)
}

func (m RTMMessage) IsFromUID(uid string) bool {
	return m["uid"] == uid
}

func (m RTMMessage) Text() string {
	if text, ok := m["text"].(string); ok {
		return text
	}

	return ""
}

func (m RTMMessage) ParseMentionUser(u User) (bool, string) {
	return m.ParseMentionUID(u.Id)
}

var mentionUserRegex = regexp.MustCompile("@<=(=[A-Za-z0-9]+)=> ")

func (m RTMMessage) ParseMentionUID(uid string) (bool, string) {
	text := m.Text()

	if m.IsP2P() {
		return true, text
	}

	if text == "" {
		return false, text
	}

	locs := mentionUserRegex.FindAllStringSubmatchIndex(text, -1)

	if len(locs) == 0 {
		return false, text
	}

	for _, loc := range locs {
		// "@<==1=> xxx" -> [0 8 3 5]
		// [3:5] "=1" [8:] "xxx"
		if text[loc[2]:loc[3]] == uid {
			return true, text[loc[1]:]
		}
	}

	return false, text
}

// ParseReferImageURL tries to get
func (m RTMMessage) ParseReferredFile() (file AttachedFile, err error) {
	rawMessageInterface, rawExist := m[JSONRawTag]
	if !rawExist {
		err = errors.New("rawMsg not exist")
		return
	}
	rawMessage, ok := rawMessageInterface.([]byte)
	if !ok {
		err = errors.New("rawMsg not []byte")
		return
	}

	var msg UpdateAttachments
	err = json.Unmarshal(rawMessage, &msg)
	if err != nil {
		err = errors.Wrap(err, "json.Unmarshal")
		return
	}

	if len(msg.Data.Attachments) == 0 {
		err = errors.New("Attachments len is 0")
		return
	}

	if msg.Data.Attachments[0].File == nil {
		err = errors.New("no file detected in first attachment")
		return
	}

	file = *msg.Data.Attachments[0].File

	return
}
