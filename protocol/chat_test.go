package protocol

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ChatTestSuite struct {
	suite.Suite
}

func TestChatSuite(t *testing.T) {
	suite.Run(t, new(ChatTestSuite))
}

func (s *ChatTestSuite) TestValidateChat() {
	testCases := []struct {
		Name  string
		Valid bool
		Chat  Chat
	}{
		{
			Name:  "valid one to one chat",
			Valid: true,
			Chat: Chat{
				ID:       "0x0424a68f89ba5fcd5e0640c1e1f591d561fa4125ca4e2a43592bc4123eca10ce064e522c254bb83079ba404327f6eafc01ec90a1444331fe769d3f3a7f90b0dde1",
				Name:     "",
				ChatType: ChatTypeOneToOne,
			},
		},
		{
			Name:  "valid public chat",
			Valid: true,
			Chat: Chat{
				ID:       "status",
				Name:     "status",
				ChatType: ChatTypePublic,
			},
		},
		{
			Name:  "empty chatID",
			Valid: false,
			Chat: Chat{
				ID:       "",
				Name:     "status",
				ChatType: ChatTypePublic,
			},
		},
		{
			Name:  "invalid one to one chat, wrong public key",
			Valid: false,
			Chat: Chat{
				ID:       "0xnotvalid",
				Name:     "",
				ChatType: ChatTypeOneToOne,
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.Name, func() {
			err := tc.Chat.Validate()
			if tc.Valid {
				s.Require().NoError(err)
			} else {
				s.Require().Error(err)
			}
		})
	}

}

func (s *ChatTestSuite) TestUpdateFromMessage() {

	// Base case, clock is higher
	message := &Message{}
	chat := &Chat{}

	message.Clock = 1
	s.Require().NoError(chat.UpdateFromMessage(message, &testTimeSource{}))
	s.Require().NotNil(chat.LastMessage)
	s.Require().Equal(uint64(1), chat.LastClockValue)

	// Clock is lower and lastMessage is not nil
	message = &Message{}
	lastMessage := []byte("test")
	chat = &Chat{LastClockValue: 2, LastMessage: lastMessage}

	message.Clock = 1
	s.Require().NoError(chat.UpdateFromMessage(message, &testTimeSource{}))
	s.Require().Equal(lastMessage, chat.LastMessage)
	s.Require().Equal(uint64(2), chat.LastClockValue)

	// Clock is lower and lastMessage is nil
	message = &Message{}
	chat = &Chat{LastClockValue: 2}

	message.Clock = 1
	s.Require().NoError(chat.UpdateFromMessage(message, &testTimeSource{}))
	s.Require().NotNil(chat.LastMessage)
	s.Require().Equal(uint64(2), chat.LastClockValue)
}
