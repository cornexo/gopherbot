package bot

import (
	"encoding/json"
	"fmt"
)

/* Handle incoming messages and other callbacks from the connector. */

// Handler is the interface that defines the callback API for Connectors
type Handler interface {
	// ChannelMessage is called by the connector for all messages the bot
	// can hear. The channelName and userName should be human-readable,
	// not internal representations.
	ChannelMessage(channelName, userName, message string)
	DirectMessage(userName, message string)
	GetProtocolConfig() json.RawMessage
	SetName(n string)
	BotLogger
}

// ChannelMessage accepts an incoming channel message from the connector.
func (b *Bot) ChannelMessage(channelName, userName, messageFull string) {
	// When command == true, the message was directed at the bot
	isCommand := false
	var message string

	b.RLock()
	for _, user := range b.ignoreUsers {
		if userName == user {
			b.Log(Debug, "Ignoring user", userName)
			b.RUnlock()
			return
		}
	}
	b.RUnlock()
	if b.preRegex != nil {
		matches := b.preRegex.FindAllStringSubmatch(messageFull, 2)
		if matches != nil && len(matches[0]) == 3 {
			isCommand = true
			message = matches[0][2]
		}
	}
	if !isCommand && b.postRegex != nil {
		matches := b.postRegex.FindAllStringSubmatch(messageFull, 2)
		if matches != nil && len(matches[0]) == 4 {
			isCommand = true
			message = matches[0][1] + matches[0][3]
		}
	}
	if !isCommand {
		message = messageFull
	}
	b.Log(Trace, fmt.Sprintf("Command \"%s\" in channel \"%s\"", message, channelName))
	b.handleMessage(isCommand, channelName, userName, message)
}

// DirectMessage accepts an incoming direct message from the connector.
func (b *Bot) DirectMessage(userName, message string) {
	b.Log(Trace, "Direct message", message, "from user", userName)
	b.RLock()
	for _, user := range b.ignoreUsers {
		if userName == user {
			b.Log(Trace, "Ignoring user", userName)
			b.RUnlock()
			return
		}
	}
	b.RUnlock()
	b.handleMessage(true, "", userName, message)
}

// GetProtocolConfig returns the connector protocol's json.RawMessage to the connector
func (b *Bot) GetProtocolConfig() json.RawMessage {
	var pc []byte
	b.RLock()
	// Make of copy of the protocol config for the plugin
	pc = append(pc, []byte(b.protocolConfig)...)
	b.RUnlock()
	return pc
}

// Connectors that support it can call SetName; otherwise it should
// be configured in gobot.conf.
func (b *Bot) SetName(n string) {
	b.Lock()
	b.Log(Debug, "Setting name to: "+n)
	b.name = n
	b.Unlock()
	b.updateRegexes()
}