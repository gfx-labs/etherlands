package types

import (
	"fmt"
	"strings"
)

func (W *World) sendChat(args ...string) {
	go func() {
		W.sendChan <- [2]string{
			"CHAT",
			strings.Join(args, ":"),
		}
	}()
}
func (W *World) checkUUIDError(gamer *Gamer, err error) bool {
	if err != nil {
		W.sendChan <- [2]string{
			"CHAT",
			fmt.Sprintf("gamer:%s:[Error] %s", gamer.MinecraftId().String(), err.Error()),
		}
		return true
	}
	return false
}
func (W *World) sendGamerChat(gamer *Gamer, args ...string) {
	go func() {
		W.sendChan <- [2]string{
			"CHAT",
			"gamer:" + gamer.MinecraftId().String() + ":" + strings.Join(args, ":"),
		}
	}()
}

func (W *World) sendTeamChat(teamname string, args ...string) {
	go func() {
		W.sendChan <- [2]string{
			"CHAT",
			"team:" + teamname + ":" + strings.Join(args, ":"),
		}
	}()

}
