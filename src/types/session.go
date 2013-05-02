package types

import "time"

type Session struct {
	MQ      chan interface{}		// ASYNC
	CALL	chan interface{}		// SYNC
	OUT		chan []byte			// server internal sending queue, like heartbeat.

	User User

	SESSID    [128]byte // UNIQUE session ID
	HeartBeat time.Time
}