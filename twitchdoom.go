package main

import (
	"github.com/quan-to/slog"
	"github.com/racerxdl/twitchdoom/twitch"
	"strconv"
	"strings"
	"time"
)

var log = slog.Scope("DOOM")
var dosbox *DosBox

func GodMode() {
	godcmd := "iddqd"
	for _, v := range godcmd {
		k := keyMap[uint8(v)]
		dosbox.SendKeyDown(k)
		time.Sleep(time.Millisecond * 10)
		dosbox.SendKeyUp(k)
	}
}

func ParseCommand(msg string) {
	msg = strings.Trim(msg, "\r\n")
	spMsg := strings.Split(msg, " ")
	if len(spMsg) > 1 {
		count := spMsg[0]
		cmd := spMsg[1]
		if count != "long" {
			n, err := strconv.ParseInt(count, 10, 32)
			if err != nil {
				return
			}
			if n > 10 {
				return
			}
			for t := int64(0); t < n; t++ {
				ParseCommand(cmd)
			}
			return
		}
	}

	key := KBD_LAST
	islong := false
	switch msg {
	case "0":
		key = KBD_0
	case "1":
		key = KBD_1
	case "2":
		key = KBD_2
	case "3":
		key = KBD_3
	case "4":
		key = KBD_4
	case "5":
		key = KBD_5
	case "6":
		key = KBD_6
	case "7":
		key = KBD_7
	case "8":
		key = KBD_8
	case "9":
		key = KBD_9
	case "up":
		key = KBD_up
	case "down":
		key = KBD_down
	case "left":
		key = KBD_left
	case "right":
		key = KBD_right
	case "ctrl":
		key = KBD_leftctrl
	case "space":
		key = KBD_space
	case "enter":
		key = KBD_enter
	case "long up":
		key = KBD_up
		islong = true
	case "long down":
		key = KBD_down
		islong = true
	case "long left":
		key = KBD_left
		islong = true
	case "long right":
		key = KBD_right
		islong = true
	case "long ctrl":
		key = KBD_leftctrl
		islong = true
	case "godmode":
		GodMode()
		return
	}

	if key != KBD_LAST {
		dosbox.SendKeyDown(key)
		if islong {
			time.Sleep(time.Second * 2)
		} else {
			time.Sleep(time.Millisecond * 100)
		}
		dosbox.SendKeyUp(key)
	}
}

func main() {
	var err error
	log.Info("Connecting to DOSBOX")
	dosbox, err = ConnectToDosBox("127.0.0.1:1337")
	if err != nil {
		panic(err)
	}
	defer dosbox.Close()

	log.Info("Connected to dosbox. Connecting to twitch")

	chat, err := twitch.MakeChat("justinfan1337", "racerxdl", "SCHMOOPIIE")

	if err != nil {
		panic(err)
	}

	log.Info("Connected to twitch")

	for {
		ev := <-chat.Events
		switch ev.GetType() {
		case twitch.EventError:
			errData := ev.GetData().(*twitch.ErrorEventData)
			log.Error("Received error: %s", errData.Error())
			break
		case twitch.EventMessage:
			msgData := ev.GetData().(*twitch.MessageEventData)
			log.Info("User %s typed %q", msgData.Username, msgData.Message)
			msgData.Message = strings.ToLower(msgData.Message)
			ParseCommand(msgData.Message)
		}
	}

	//chat.Events
	//
	//stringToSend := "HELLO WORLD"
	//
	//stringToSend = strings.ToLower(stringToSend)
	//for _, v := range stringToSend {
	//	kbdChar, ok := keyMap[uint8(v)]
	//	if !ok {
	//		fmt.Printf("Char %c not supported...\n", v)
	//		continue
	//	}
	//}
	//
	//time.Sleep(time.Second * 5)

}
