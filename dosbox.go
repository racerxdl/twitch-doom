package main

import (
	fifo "github.com/racerxdl/go.fifo"
	"net"
	"time"
)

type DosBoxKey uint8

const (
	KBD_NONE DosBoxKey = iota
	KBD_1    DosBoxKey = iota
	KBD_2    DosBoxKey = iota
	KBD_3    DosBoxKey = iota
	KBD_4    DosBoxKey = iota
	KBD_5    DosBoxKey = iota
	KBD_6    DosBoxKey = iota
	KBD_7    DosBoxKey = iota
	KBD_8    DosBoxKey = iota
	KBD_9    DosBoxKey = iota
	KBD_0    DosBoxKey = iota
	KBD_q    DosBoxKey = iota
	KBD_w    DosBoxKey = iota
	KBD_e    DosBoxKey = iota
	KBD_r    DosBoxKey = iota
	KBD_t    DosBoxKey = iota
	KBD_y    DosBoxKey = iota
	KBD_u    DosBoxKey = iota
	KBD_i    DosBoxKey = iota
	KBD_o    DosBoxKey = iota
	KBD_p    DosBoxKey = iota
	KBD_a    DosBoxKey = iota
	KBD_s    DosBoxKey = iota
	KBD_d    DosBoxKey = iota
	KBD_f    DosBoxKey = iota
	KBD_g    DosBoxKey = iota
	KBD_h    DosBoxKey = iota
	KBD_j    DosBoxKey = iota
	KBD_k    DosBoxKey = iota
	KBD_l    DosBoxKey = iota
	KBD_z    DosBoxKey = iota
	KBD_x    DosBoxKey = iota
	KBD_c    DosBoxKey = iota
	KBD_v    DosBoxKey = iota
	KBD_b    DosBoxKey = iota
	KBD_n    DosBoxKey = iota
	KBD_m    DosBoxKey = iota
	KBD_f1   DosBoxKey = iota
	KBD_f2   DosBoxKey = iota
	KBD_f3   DosBoxKey = iota
	KBD_f4   DosBoxKey = iota
	KBD_f5   DosBoxKey = iota
	KBD_f6   DosBoxKey = iota
	KBD_f7   DosBoxKey = iota
	KBD_f8   DosBoxKey = iota
	KBD_f9   DosBoxKey = iota
	KBD_f10  DosBoxKey = iota
	KBD_f11  DosBoxKey = iota
	KBD_f12  DosBoxKey = iota

	/*Now the weirder keys */

	KBD_esc        DosBoxKey = iota
	KBD_tab        DosBoxKey = iota
	KBD_backspace  DosBoxKey = iota
	KBD_enter      DosBoxKey = iota
	KBD_space      DosBoxKey = iota
	KBD_leftalt    DosBoxKey = iota
	KBD_rightalt   DosBoxKey = iota
	KBD_leftctrl   DosBoxKey = iota
	KBD_rightctrl  DosBoxKey = iota
	KBD_leftshift  DosBoxKey = iota
	KBD_rightshift DosBoxKey = iota
	KBD_capslock   DosBoxKey = iota
	KBD_scrolllock DosBoxKey = iota
	KBD_numlock    DosBoxKey = iota

	KBD_grave        DosBoxKey = iota
	KBD_minus        DosBoxKey = iota
	KBD_equals       DosBoxKey = iota
	KBD_backslash    DosBoxKey = iota
	KBD_leftbracket  DosBoxKey = iota
	KBD_rightbracket DosBoxKey = iota
	KBD_semicolon    DosBoxKey = iota
	KBD_quote        DosBoxKey = iota
	KBD_period       DosBoxKey = iota
	KBD_comma        DosBoxKey = iota
	KBD_slash        DosBoxKey = iota
	KBD_extra_lt_gt  DosBoxKey = iota

	KBD_printscreen DosBoxKey = iota
	KBD_pause       DosBoxKey = iota
	KBD_insert      DosBoxKey = iota
	KBD_home        DosBoxKey = iota
	KBD_pageup      DosBoxKey = iota
	KBD_delete      DosBoxKey = iota
	KBD_end         DosBoxKey = iota
	KBD_pagedown    DosBoxKey = iota
	KBD_left        DosBoxKey = iota
	KBD_up          DosBoxKey = iota
	KBD_down        DosBoxKey = iota
	KBD_right       DosBoxKey = iota

	KBD_kp1        DosBoxKey = iota
	KBD_kp2        DosBoxKey = iota
	KBD_kp3        DosBoxKey = iota
	KBD_kp4        DosBoxKey = iota
	KBD_kp5        DosBoxKey = iota
	KBD_kp6        DosBoxKey = iota
	KBD_kp7        DosBoxKey = iota
	KBD_kp8        DosBoxKey = iota
	KBD_kp9        DosBoxKey = iota
	KBD_kp0        DosBoxKey = iota
	KBD_kpdivide   DosBoxKey = iota
	KBD_kpmultiply DosBoxKey = iota
	KBD_kpminus    DosBoxKey = iota
	KBD_kpplus     DosBoxKey = iota
	KBD_kpenter    DosBoxKey = iota
	KBD_kpperiod   DosBoxKey = iota

	KBD_LAST
)

var keyMap = map[uint8]DosBoxKey{
	'1': KBD_1,
	'2': KBD_2,
	'3': KBD_3,
	'4': KBD_4,
	'5': KBD_5,
	'6': KBD_6,
	'7': KBD_7,
	'8': KBD_8,
	'9': KBD_9,
	'0': KBD_0,
	'q': KBD_q,
	'w': KBD_w,
	'e': KBD_e,
	'r': KBD_r,
	't': KBD_t,
	'y': KBD_y,
	'u': KBD_u,
	'i': KBD_i,
	'o': KBD_o,
	'p': KBD_p,
	'a': KBD_a,
	's': KBD_s,
	'd': KBD_d,
	'f': KBD_f,
	'g': KBD_g,
	'h': KBD_h,
	'j': KBD_j,
	'k': KBD_k,
	'l': KBD_l,
	'z': KBD_z,
	'x': KBD_x,
	'c': KBD_c,
	'v': KBD_v,
	'b': KBD_b,
	'n': KBD_n,
	'm': KBD_m,
	' ': KBD_space,
}

type dosboxCmd uint8

const (
	KEEP_ALIVE dosboxCmd = iota
	KEY_DOWN   dosboxCmd = iota
	KEY_UP     dosboxCmd = iota
)

type dosboxQueuedCmd struct {
	Type dosboxCmd
	Arg  uint8
}

func (cmd dosboxQueuedCmd) Send(c net.Conn) error {
	_, err := c.Write([]byte{uint8(cmd.Type), cmd.Arg})
	return err
}

type DosBox struct {
	socket        net.Conn
	cmdQueue      *fifo.Queue
	lastKeepAlive time.Time
}

const keepAliveInterval = time.Second * 2

func ConnectToDosBox(addr string) (*DosBox, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	dosbox := &DosBox{
		cmdQueue: fifo.NewQueue(),
		socket:   conn,
	}

	go dosbox.loop()

	return dosbox, nil
}

func (dosbox *DosBox) Close() error {
	return dosbox.socket.Close()
}

func (dosbox *DosBox) loop() {
	var err error
	keepAliveCmd := dosboxQueuedCmd{Type: KEEP_ALIVE}
	for {
		if dosbox.cmdQueue.Len() > 0 {
			cmd := dosbox.cmdQueue.Next().(dosboxQueuedCmd)
			err = cmd.Send(dosbox.socket)
			dosbox.lastKeepAlive = time.Now()
		} else if time.Since(dosbox.lastKeepAlive) > keepAliveInterval {
			err = keepAliveCmd.Send(dosbox.socket)
			dosbox.lastKeepAlive = time.Now()
		}

		if err != nil {
			panic(err)
		}
		time.Sleep(time.Millisecond)
	}
}

func (dosbox *DosBox) SendKeyDown(key DosBoxKey) {
	log.Debug("Sending KEY_DOWN %d", key)
	dosbox.cmdQueue.Add(dosboxQueuedCmd{
		Type: KEY_DOWN,
		Arg:  uint8(key),
	})
}

func (dosbox *DosBox) SendKeyUp(key DosBoxKey) {
	log.Debug("Sending KEY_UP %d", key)
	dosbox.cmdQueue.Add(dosboxQueuedCmd{
		Type: KEY_UP,
		Arg:  uint8(key),
	})
}
