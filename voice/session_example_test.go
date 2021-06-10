package voice_test

import (
	"context"
	"io"
	"log"
	"testing"

	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/internal/testenv"
	"github.com/diamondburned/arikawa/v3/state"
	"github.com/diamondburned/arikawa/v3/voice"
)

var (
	token     string
	channelID discord.ChannelID
)

func init() {
	e, err := testenv.GetEnv()
	if err == nil {
		token = e.BotToken
		channelID = e.VoiceChID
	}
}

// pseudo function for example
func writeOpusInto(w io.Writer) {}

// make godoc not show the full file
func TestNoop(t *testing.T) {
	t.Skip("noop")
}

func ExampleSession() {
	s, err := state.New("Bot " + token)
	if err != nil {
		log.Fatalln("failed to make state:", err)
	}

	// This is required for bots.
	voice.AddIntents(s.Gateway)

	if err := s.Open(context.TODO()); err != nil {
		log.Fatalln("failed to open gateway:", err)
	}
	defer s.Close()

	c, err := s.Channel(channelID)
	if err != nil {
		log.Fatalln("failed to get channel:", err)
	}

	v, err := voice.NewSession(s)
	if err != nil {
		log.Fatalln("failed to create voice session:", err)
	}

	if err := v.JoinChannel(c.GuildID, c.ID, false, false); err != nil {
		log.Fatalln("failed to join voice channel:", err)
	}
	defer v.Leave()

	// Start writing Opus frames.
	for {
		writeOpusInto(v)
	}
}
