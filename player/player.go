package player

import (
	cha "final_project2/channel"
	"fmt"
)

//Player - Player in the game
type Player struct {
	username             string
	number               int
	ch                   cha.Channel
	otherPlayersChannels []cha.Channel
}

//ChannelsListNotGoodErrMsg - Error message
const ChannelsListNotGoodErrMsg = "Channels propbability must be between (0,1]"

//New - Player constructor
func New(username string, number int, userChannel cha.Channel, channelsList []cha.Channel) (Player, error) {
	if channelsList == nil {
		return Player{}, fmt.Errorf("%s", ChannelsListNotGoodErrMsg)
	}
	e := Player{username, number, userChannel, channelsList}
	return e, nil
}

//-----------Public functions-----------

//UserToString - Returns a representation of the player
func (e Player) UserToString() string {
	return fmt.Sprintf("username: %s, user number: %d", e.username, e.number)
}

//GetUsername - return the username
func (e Player) GetUsername() string {
	return e.username
}

//GetRandomNumber - return the random number of the user
func (e Player) GetRandomNumber() int {
	return e.number
}

//GetChannel - return the channel of the user
func (e Player) GetChannel() chan int {
	return e.ch.GetChannel()
}

//GetotherPlayersChannels - return other players channels
func (e Player) GetotherPlayersChannels() []cha.Channel {
	return e.otherPlayersChannels
}

//SendMessagesToAllPlayers - Sends the random number of the user to all the others players channels
func (e Player) SendMessagesToAllPlayers() {
	for _, element := range e.otherPlayersChannels {
		e.sendNumber(element)
	}
}

//GetSum - get all the numbers from the channel, summerizes and prints it
func (e Player) GetSum() (string, int) {
	sum := e.ch.GetSum()
	sum = sum + e.number
	return e.username, sum
}

//-----------Private functions-----------

//sendNumber - Sends the random number of the user to the channel (argument)
func (e Player) sendNumber(channel cha.Channel) {
	//e.number is always the same!!!
	channel.InsertNumber(e.number)
}
