package bearychat

import (
	"fmt"
	"testing"
)

func TestRTMMessage_Type(t *testing.T) {
	cases := [][]RTMMessageType{
		{RTMMessageTypeUnknown, RTMMessageTypeUnknown},
		{RTMMessageTypePing, RTMMessageTypePing},
		{RTMMessageTypePong, RTMMessageTypePong},
		{RTMMessageTypeReply, RTMMessageTypeReply},
		{RTMMessageTypeOk, RTMMessageTypeOk},
		{RTMMessageTypeP2PMessage, RTMMessageTypeP2PMessage},
		{RTMMessageTypeP2PTyping, RTMMessageTypeP2PTyping},
		{RTMMessageTypeChannelMessage, RTMMessageTypeChannelMessage},
		{RTMMessageTypeChannelTyping, RTMMessageTypeChannelTyping},
		{RTMMessageTypeUpdateUserConnection, RTMMessageTypeUpdateUserConnection},
	}

	for _, c := range cases {
		m := RTMMessage{"type": c[0]}
		if m.Type() != c[1] {
			t.Errorf("expected type: %s, got: %s", c[0], m.Type())
		}
	}
}

func TestRTMMessage_IsP2P(t *testing.T) {
	cases := []struct {
		mt       RTMMessageType
		expected bool
	}{
		{RTMMessageTypeUnknown, false},
		{RTMMessageTypePing, false},
		{RTMMessageTypePong, false},
		{RTMMessageTypeReply, false},
		{RTMMessageTypeOk, false},
		{RTMMessageTypeP2PMessage, true},
		{RTMMessageTypeP2PTyping, true},
		{RTMMessageTypeChannelMessage, false},
		{RTMMessageTypeChannelTyping, false},
		{RTMMessageTypeUpdateUserConnection, false},
	}

	for _, c := range cases {
		m := RTMMessage{"type": c.mt}
		if m.IsP2P() != c.expected {
			t.Errorf("expected: %+v, got: %+v", c.expected, m.IsP2P())
		}
	}
}

func TestRTMMessage_IsChatMessage(t *testing.T) {
	cases := []struct {
		mt       RTMMessageType
		expected bool
	}{
		{RTMMessageTypeUnknown, false},
		{RTMMessageTypePing, false},
		{RTMMessageTypePong, false},
		{RTMMessageTypeReply, false},
		{RTMMessageTypeOk, false},
		{RTMMessageTypeP2PMessage, true},
		{RTMMessageTypeP2PTyping, false},
		{RTMMessageTypeChannelMessage, true},
		{RTMMessageTypeChannelTyping, false},
		{RTMMessageTypeUpdateUserConnection, false},
	}

	for _, c := range cases {
		m := RTMMessage{"type": c.mt}
		if m.IsChatMessage() != c.expected {
			t.Errorf("expected: %+v, got: %+v", c.expected, m.IsChatMessage())
		}
	}
}

func TestRTMMessage_IsFrom(t *testing.T) {
	uid := "1"
	user := User{Id: uid}
	var m RTMMessage

	m = RTMMessage{"uid": uid}
	if !m.IsFromUser(user) {
		t.Errorf("expected from user: %+v", m)
	}
	if !m.IsFromUID(uid) {
		t.Errorf("expected from uid: %+v", m)
	}

	m = RTMMessage{"uid": uid + "1"}
	if m.IsFromUser(user) {
		t.Errorf("unexpected from user: %+v", m)
	}
	if m.IsFromUID(uid) {
		t.Errorf("expected from uid: %+v", m)
	}
}

func TestRTMMessage_Refer_ChannelMessage(t *testing.T) {
	m := RTMMessage{
		"type":        RTMMessageTypeChannelMessage,
		"channel_id":  "foobar",
		"vchannel_id": "foobar",
		"key":         "foobar",
	}

	referText := "foobar"
	refer := m.Refer(referText)
	if refer["text"] != referText {
		t.Errorf("unexpected %s", refer["text"])
	}
	if refer.Type() != RTMMessageTypeChannelMessage {
		t.Errorf("unexpected %s", refer.Type())
	}
	if refer["channel_id"] != m["channel_id"] {
		t.Errorf("unexpected %s", refer["channel_id"])
	}
	if refer["vchannel_id"] != m["vchannel_id"] {
		t.Errorf("unexpected %s", refer["vchannel_id"])
	}
	if refer["refer_key"] != m["key"] {
		t.Errorf("unexpected %s", refer["refer_key"])
	}
}

func TestRTMMessage_Refer_P2PMessage(t *testing.T) {
	m := RTMMessage{
		"type":        RTMMessageTypeP2PMessage,
		"uid":         "foobar",
		"vchannel_id": "foobar",
		"key":         "foobar",
	}

	referText := "foobar"
	refer := m.Refer(referText)
	if refer["text"] != referText {
		t.Errorf("unexpected %s", refer["text"])
	}
	if refer.Type() != RTMMessageTypeP2PMessage {
		t.Errorf("unexpected %s", refer.Type())
	}
	if refer["to_uid"] != m["uid"] {
		t.Errorf("unexpected %s", refer["to_uid"])
	}
	if refer["vchannel_id"] != m["vchannel_id"] {
		t.Errorf("unexpected %s", refer["vchannel_id"])
	}
	if refer["refer_key"] != m["key"] {
		t.Errorf("unexpected %s", refer["refer_key"])
	}
}

func TestRTMMessage_Reply_ChannelMessage(t *testing.T) {
	m := RTMMessage{
		"type":        RTMMessageTypeChannelMessage,
		"channel_id":  "foobar",
		"vchannel_id": "foobar",
	}

	replyText := "foobar"
	reply := m.Reply(replyText)
	if reply["text"] != replyText {
		t.Errorf("unexpected %s", reply["text"])
	}
	if reply.Type() != RTMMessageTypeChannelMessage {
		t.Errorf("unexpected %s", reply.Type())
	}
	if reply["channel_id"] != m["channel_id"] {
		t.Errorf("unexpected %s", reply["channel_id"])
	}
	if reply["vchannel_id"] != m["vchannel_id"] {
		t.Errorf("unexpected %s", reply["vchannel_id"])
	}
}

func TestRTMMessage_Reply_P2PMessage(t *testing.T) {
	m := RTMMessage{
		"type":        RTMMessageTypeChannelMessage,
		"channel_id":  "foobar",
		"vchannel_id": "foobar",
	}

	replyText := "foobar"
	reply := m.Reply(replyText)
	if reply["text"] != replyText {
		t.Errorf("unexpected %s", reply["text"])
	}
	if reply.Type() != RTMMessageTypeChannelMessage {
		t.Errorf("unexpected %s", reply.Type())
	}
	if reply["to_uid"] != m["uid"] {
		t.Errorf("unexpected %s", reply["to_uid"])
	}
	if reply["vchannel_id"] != m["vchannel_id"] {
		t.Errorf("unexpected %s", reply["vchannel_id"])
	}
}

func TestRTMMessage_ParseMention(t *testing.T) {
	uid := "=1"
	text := "abc"
	user := User{Id: uid}

	m := RTMMessage{}
	var mentioned bool
	var content string

	expect := func(expectMentioned bool, expectContent string) {
		if mentioned != expectMentioned {
			t.Errorf("expected mentioned: '%v', got '%v', m: %+v", expectMentioned, mentioned, m)
		}
		if content != expectContent {
			t.Errorf("expected content: '%v', got '%v'", expectContent, content)
		}
	}

	m["text"] = text
	m["type"] = RTMMessageTypeP2PMessage
	mentioned, content = m.ParseMentionUser(user)
	expect(true, text)
	mentioned, content = m.ParseMentionUID(uid)
	expect(true, text)

	m["type"] = RTMMessageTypeChannelMessage
	mentioned, content = m.ParseMentionUser(user)
	expect(false, text)
	mentioned, content = m.ParseMentionUID(uid)
	expect(false, text)

	m["text"] = fmt.Sprintf("@<=%s=> %s", uid, text)
	mentioned, content = m.ParseMentionUser(user)
	expect(true, text)
	mentioned, content = m.ParseMentionUID(uid)
	expect(true, text)

	m["text"] = fmt.Sprintf("123123123 12312 123@<=%s=> %s", uid, text)
	mentioned, content = m.ParseMentionUser(user)
	expect(true, text)
	mentioned, content = m.ParseMentionUID(uid)
	expect(true, text)

	m["text"] = fmt.Sprintf("@<=%s=>", uid)
	mentioned, content = m.ParseMentionUser(user)
	expect(false, m.Text())
	mentioned, content = m.ParseMentionUID(uid)
	expect(false, m.Text())

	m["text"] = fmt.Sprintf("@<=%s=> ", uid)
	mentioned, content = m.ParseMentionUser(user)
	expect(true, "")
	mentioned, content = m.ParseMentionUID(uid)
	expect(true, "")

	m["text"] = fmt.Sprintf("@<=%s=> 你和 @<==bwOwr=> 谁聪明", uid)
	mentioned, content = m.ParseMentionUser(user)
	expect(true, "你和 @<==bwOwr=> 谁聪明")
	mentioned, content = m.ParseMentionUID(uid)
	expect(true, "你和 @<==bwOwr=> 谁聪明")

	m["text"] = fmt.Sprintf("@<==bwOwr=> @<=%s=> hello", uid)
	mentioned, content = m.ParseMentionUser(user)
	expect(true, "hello")
	mentioned, content = m.ParseMentionUID(uid)
	expect(true, "hello")
}

func TestRTMMessage_ParseReferredFile(t *testing.T) {
	var msg = RTMMessage{
		"type": "update_attachments",
		JSONRawTag: []byte(`{
   "data":{
      "attachments":[
         {
            "file":{
               "category":"image",
               "created":"2018-10-28T14:16:52.000+0000",
               "deleted":false,
               "description":"",
               "height":500,
               "id":"=b93bD",
               "image_url":"https://file.bearychat.com/45c7a7dd13a29e046e5a20a9a774b5de",
               "inactive":false,
               "is_public":true,
               "key":"45c7a7dd13a29e046e5a20a9a774b5de",
               "mime":"image/png",
               "name":"clipboard_2018-10-28_22:16.png",
               "orientation":1,
               "original":true,
               "preview_url":"https://file.bearychat.com/45c7a7dd13a29e046e5a20a9a774b5de",
               "size":5677,
               "source":"internal",
               "summary":null,
               "team_id":"=bwECm",
               "title":"clipboard_2018-10-28_22:16.png",
               "type":"png",
               "uid":"=bxcKY",
               "updated":"2018-10-28T14:16:52.000+0000",
               "upload_zone":"unknown",
               "url":"https://file.bearychat.com/45c7a7dd13a29e046e5a20a9a774b5de",
               "width":500
            },
            "refer_text":"上传了图片",
            "refer_text_i18n":{
               "en":"uploaded an image",
               "zh-CN":"上传了图片"
            },
            "subtype":"file",
            "text":"nanmu: 上传了图片",
            "type":"refer",
            "uid":"=bxcKY"
         }
      ],
      "created":"2018-10-28T14:26:26.000+0000",
      "created_ts":1540736786063,
      "disable_markdown":false,
      "edited":false,
      "id":"=hmjlz",
      "is_channel":true,
      "key":"1540736786063.0177",
      "pin_id":null,
      "reactions":[

      ],
      "refer_key":"1540736212397.0229",
      "repost":null,
      "robot_id":null,
      "subtype":"normal",
      "team_id":"=bwECm",
      "text":"@<==bxcNg=> ",
      "thread_key":null,
      "uid":"=bxcKY",
      "updated":"2018-10-28T14:26:26.000+0000",
      "vchannel_id":"=bwTfi"
   },
   "ts":1540736786090,
   "type":"update_attachments"
}`),
	}

	file, err := msg.ParseAttachedFile()
	if err != nil {
		t.Fatal(err)
	}
	if file.ImageURL == "" {
		t.Fatal("ImageURL is empty")
	}

	var msg2 = RTMMessage{
		JSONRawTag: []byte(`{
   "created":"2018-10-29T07:26:56.000+0000",
   "created_ts":1540798015600,
   "edited":false,
   "file":{
      "category":"image",
      "channel_id":null,
      "comments_count":0,
      "comments_ids":[

      ],
      "created":"2018-10-29T07:26:56.000+0000",
      "deleted":false,
      "description":"",
      "height":120,
      "id":"=b937z",
      "image_url":"https://file.bearychat.com/77ec54654e09243a8064747b248cb73d",
      "inactive":false,
      "is_public":false,
      "key":"77ec54654e09243a8064747b248cb73d",
      "mime":"image/gif",
      "name":"cat.gif",
      "orientation":1,
      "original":true,
      "preview_url":"https://file.bearychat.com/77ec54654e09243a8064747b248cb73d",
      "size":4081,
      "source":"internal",
      "summary":null,
      "team_id":"=bwECm",
      "title":"cat.gif",
      "type":"gif",
      "uid":"=bxcKY",
      "updated":"2018-10-29T07:26:56.000+0000",
      "upload_zone":"unknown",
      "url":"https://file.bearychat.com/77ec54654e09243a8064747b248cb73d",
      "vcids":[
         "=c6fkE6xfO"
      ],
      "width":120
   },
   "go_json_raw_message":"eyJjcmVhdGVkIjoiMjAxOC0xMC0yOVQwNzoyNjo1Ni4wMDArMDAwMCIsImNyZWF0ZWRfdHMiOjE1NDA3OTgwMTU2MDAsImVkaXRlZCI6ZmFsc2UsImZpbGUiOnsiY2F0ZWdvcnkiOiJpbWFnZSIsImNoYW5uZWxfaWQiOm51bGwsImNvbW1lbnRzX2NvdW50IjowLCJjb21tZW50c19pZHMiOltdLCJjcmVhdGVkIjoiMjAxOC0xMC0yOVQwNzoyNjo1Ni4wMDArMDAwMCIsImRlbGV0ZWQiOmZhbHNlLCJkZXNjcmlwdGlvbiI6IiIsImhlaWdodCI6MTIwLCJpZCI6Ij1iOTM3eiIsImltYWdlX3VybCI6Imh0dHBzOi8vZmlsZS5iZWFyeWNoYXQuY29tLzc3ZWM1NDY1NGUwOTI0M2E4MDY0NzQ3YjI0OGNiNzNkIiwiaW5hY3RpdmUiOmZhbHNlLCJpc19wdWJsaWMiOmZhbHNlLCJrZXkiOiI3N2VjNTQ2NTRlMDkyNDNhODA2NDc0N2IyNDhjYjczZCIsIm1pbWUiOiJpbWFnZS9naWYiLCJuYW1lIjoiY2F0LmdpZiIsIm9yaWVudGF0aW9uIjoxLCJvcmlnaW5hbCI6dHJ1ZSwicHJldmlld191cmwiOiJodHRwczovL2ZpbGUuYmVhcnljaGF0LmNvbS83N2VjNTQ2NTRlMDkyNDNhODA2NDc0N2IyNDhjYjczZCIsInNpemUiOjQwODEsInNvdXJjZSI6ImludGVybmFsIiwic3VtbWFyeSI6bnVsbCwidGVhbV9pZCI6Ij1id0VDbSIsInRpdGxlIjoiY2F0LmdpZiIsInR5cGUiOiJnaWYiLCJ1aWQiOiI9YnhjS1kiLCJ1cGRhdGVkIjoiMjAxOC0xMC0yOVQwNzoyNjo1Ni4wMDArMDAwMCIsInVwbG9hZF96b25lIjoidW5rbm93biIsInVybCI6Imh0dHBzOi8vZmlsZS5iZWFyeWNoYXQuY29tLzc3ZWM1NDY1NGUwOTI0M2E4MDY0NzQ3YjI0OGNiNzNkIiwidmNpZHMiOlsiPWM2ZmtFNnhmTyJdLCJ3aWR0aCI6MTIwfSwiaWQiOiI9aG1zbEwiLCJpc19jaGFubmVsIjpmYWxzZSwia2V5IjoiMTU0MDc5ODAxNTYwMC4wNDA0IiwicGluX2lkIjpudWxsLCJyZWFjdGlvbnMiOltdLCJyZWZlcl9rZXkiOm51bGwsInJlcG9zdCI6bnVsbCwicmVzb3VyY2Vfa2V5IjoiMGY2OWIwMDktMzE0ZS00Njc4LWEyZDItZWQwZWRmNzZkNzk4IiwicmlkIjoiMjEwYmZhM2ZkNGQ1ZDcxNzE4ZjlmOGNmMDQzMmRlYjEiLCJyb2JvdF9pZCI6bnVsbCwic3VidHlwZSI6ImZpbGUiLCJ0ZWFtX2lkIjoiPWJ3RUNtIiwidGV4dCI6IuS4iuS8oOS6huWbvueJhyIsInRleHRfaTE4biI6eyJlbiI6InVwbG9hZGVkIGFuIGltYWdlIiwiemgtQ04iOiLkuIrkvKDkuoblm77niYcifSwidGhyZWFkX2tleSI6bnVsbCwidHMiOjE1NDA3OTgwMTU2MzcsInR5cGUiOiJtZXNzYWdlIiwidWlkIjoiPWJ4Y0tZIiwidXBkYXRlZCI6IjIwMTgtMTAtMjlUMDc6MjY6NTYuMDAwKzAwMDAiLCJ2Y2hhbm5lbF9pZCI6Ij1jNmZrRTZ4Zk8ifQ==",
   "id":"=hmslL",
   "is_channel":false,
   "key":"1540798015600.0404",
   "pin_id":null,
   "reactions":[

   ],
   "refer_key":null,
   "repost":null,
   "resource_key":"0f69b009-314e-4678-a2d2-ed0edf76d798",
   "rid":"210bfa3fd4d5d71718f9f8cf0432deb1",
   "robot_id":null,
   "subtype":"file",
   "team_id":"=bwECm",
   "text":"上传了图片",
   "text_i18n":{
      "en":"uploaded an image",
      "zh-CN":"上传了图片"
   },
   "thread_key":null,
   "ts":1540798015637,
   "type":"message",
   "uid":"=bxcKY",
   "updated":"2018-10-29T07:26:56.000+0000",
   "vchannel_id":"=c6fkE6xfO"
}`),
		"type": "message",
	}

	file, err = msg2.ParseAttachedFile()
	if err != nil {
		t.Fatal(err)
	}
	if file.ImageURL == "" {
		t.Fatal("ImageURL is empty")
	}

}
