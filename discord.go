// Discordgo - Discord bindings for Go
// Available at https://github.com/courtier/kolizey

// Copyright 2015-2016 Bruce Marriner <bruce@sqls.net>.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// package kolizey provides Discord client bindings for developing
// "selfbots"/automated user accounts in Go
package kolizey

import (
	"errors"
	"net/http"
	"strings"
	"time"
)

var (
	errEmptyToken = errors.New("empty auth token")
	errBotToken   = errors.New("bot tokens are disallowed")
)

// New creates a new Discord session and will automate some startup
// tasks if given enough information to do so. You must pass it a user account
// auth token. Do note that this is against Discord TOS and might get your
// account banned.
func New(token string) (s *Session, err error) {

	//return empty session
	if token == "" {
		return nil, errEmptyToken
	} else if strings.HasPrefix(token, "Bot ") {
		return nil, errBotToken
	}

	// Create an empty Session interface.
	s = &Session{
		State:                  NewState(),
		Ratelimiter:            NewRatelimiter(),
		StateEnabled:           true,
		ShouldReconnectOnError: true,
		MaxRestRetries:         3,
		Client:                 &http.Client{Timeout: (20 * time.Second)},
		UserAgent:              "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36",
		sequence:               new(int64),
		LastHeartbeatAck:       time.Now().UTC(),
	}

	// Initilize the Identify packet with defaults.
	// These can be modified prior to calling Open()
	// I will be using values that I got on my browser,
	// but you can set to anything else if desired.

	s.Identify.Token = token
	s.Identify.Properties = IdentifyProperties{
		// Windows, Linux etc.
		OS: "Mac OS X",
		// Brave, Safari etc.
		Browser: "Chrome",
		// Perhaps filled if its a phone like iPhone or Samsung ? idk
		Device: "",
		// en-US, de-DE etc.
		SystemLocale: "en-AU",
		// Any user agent that matches OS and Browser
		BrowserUserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36",
		BrowserVersion:   "92.0.4515.159",
		// This is actually NOT my OS version, idk how they get this value,
		// perhaps its the OS version you signed up to your account with,
		// but then how does the webssite know to send that in the packet ? idk
		OSVersion: "10.15.7",
		// idk
		Referrer:               "",
		ReferringDomain:        "",
		ReferrerCurrent:        "",
		ReferringDomainCurrent: "",
		ReleaseChannel:         "stable",
		ClientBuildNumber:      94294,
		ClientEventSource:      nil,
	}
	s.Identify.Presence = IdentifyPresence{
		// Online etcc.
		Status: "invisible",
		// Leave at 0
		Since: 0,
		// idk how you would fill this in
		Activities: nil,
		// idk what this does
		AFK: false,
	}
	s.Identify.Compress = false
	// No clue, I suggest you do not touch this.
	s.Identify.ClientState = IdentifyClientState{
		GuildHashes:              struct{}{},
		HighestLastMessageID:     "0",
		ReadStateVersion:         0,
		UserGuildSettingsVersion: -1,
	}
	// 61 for Discord client and 125 for browser?
	// Reddit people report seeing 61, I saw 125 on browser.
	s.Identify.Capabilities = 125
	return
}
