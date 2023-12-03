package audio

import (
	"github.com/teostofell/ipod"
)

type DeviceAudio interface {
}

// func ackSuccess(req ipod.Packet) ACK {
// 	return ACK{Status: ACKStatusSuccess, CmdID: uint8(req.ID.CmdID())}
// }

// func ackPending(req ipod.Packet, maxWait uint32) ACKPending {
// 	return ACKPending{Status: ACKStatusPending, CmdID: uint8(req.ID.CmdID()), MaxWait: maxWait}
// }

func Start(tr ipod.CommandWriter) {
	ipod.Send(tr, &GetAccSampleRateCaps{})
}

func HandleAudio(req *ipod.Command, tr ipod.CommandWriter, dev DeviceAudio) error {
	switch msg := req.Payload.(type) {
	case *AccAck:
	case *RetAccSampleRateCaps:
		ipod.Respond(req, tr, &TrackNewAudioAttributes{
			SampleRate: 44100,
		})

	default:
		_ = msg
	}
	return nil
}
