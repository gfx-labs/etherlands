package types

import "strings"

func (W *World) sendChat(args ...string) {
	go func() {
		W.sendChan <- [2]string{
			"CHAT",
			strings.Join(args, ":"),
		}
	}()
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
