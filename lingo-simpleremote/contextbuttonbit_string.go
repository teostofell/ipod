// Code generated by "stringer -type=ContextButtonBit"; DO NOT EDIT.

package simpleremote

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ContextButtonVolumeUp-1]
	_ = x[ContextButtonVolumeDown-2]
	_ = x[ContextButtonNextTrack-4]
	_ = x[ContextButtonPreviousTrack-8]
	_ = x[ContextButtonNextAlbum-16]
	_ = x[ContextButtonPreviousAlbum-32]
	_ = x[ContextButtonStop-64]
	_ = x[ContextButtonPlayResume-128]
	_ = x[ContextButtonPause-256]
	_ = x[ContextButtonMuteToggle-512]
	_ = x[ContextButtonNextChapter-1024]
	_ = x[ContextButtonPreviousChapter-2048]
	_ = x[ContextButtonNextPlaylist-4096]
	_ = x[ContextButtonPreviousPlaylist-8192]
	_ = x[ContextButtonShuffleSettingAdvance-16384]
	_ = x[ContextButtonRepeatSettingAdvance-32768]
	_ = x[ContextButtonPowerOn-65536]
	_ = x[ContextButtonPowerOff-131072]
	_ = x[ContextButtonBacklightfor30Seconds-262144]
	_ = x[ContextButtonBeginFastForward-524288]
	_ = x[ContextButtonBeginRewind-1048576]
	_ = x[ContextButtonMenu-2097152]
	_ = x[ContextButtonSelect-4194304]
	_ = x[ContextButtonUpArrow-8388608]
	_ = x[ContextButtonDownArrow-16777216]
	_ = x[ContextButtonBacklightOff-33554432]
}

const _ContextButtonBit_name = "ContextButtonVolumeUpContextButtonVolumeDownContextButtonNextTrackContextButtonPreviousTrackContextButtonNextAlbumContextButtonPreviousAlbumContextButtonStopContextButtonPlayResumeContextButtonPauseContextButtonMuteToggleContextButtonNextChapterContextButtonPreviousChapterContextButtonNextPlaylistContextButtonPreviousPlaylistContextButtonShuffleSettingAdvanceContextButtonRepeatSettingAdvanceContextButtonPowerOnContextButtonPowerOffContextButtonBacklightfor30SecondsContextButtonBeginFastForwardContextButtonBeginRewindContextButtonMenuContextButtonSelectContextButtonUpArrowContextButtonDownArrowContextButtonBacklightOff"

var _ContextButtonBit_map = map[ContextButtonBit]string{
	1:        _ContextButtonBit_name[0:21],
	2:        _ContextButtonBit_name[21:44],
	4:        _ContextButtonBit_name[44:66],
	8:        _ContextButtonBit_name[66:92],
	16:       _ContextButtonBit_name[92:114],
	32:       _ContextButtonBit_name[114:140],
	64:       _ContextButtonBit_name[140:157],
	128:      _ContextButtonBit_name[157:180],
	256:      _ContextButtonBit_name[180:198],
	512:      _ContextButtonBit_name[198:221],
	1024:     _ContextButtonBit_name[221:245],
	2048:     _ContextButtonBit_name[245:273],
	4096:     _ContextButtonBit_name[273:298],
	8192:     _ContextButtonBit_name[298:327],
	16384:    _ContextButtonBit_name[327:361],
	32768:    _ContextButtonBit_name[361:394],
	65536:    _ContextButtonBit_name[394:414],
	131072:   _ContextButtonBit_name[414:435],
	262144:   _ContextButtonBit_name[435:469],
	524288:   _ContextButtonBit_name[469:498],
	1048576:  _ContextButtonBit_name[498:522],
	2097152:  _ContextButtonBit_name[522:539],
	4194304:  _ContextButtonBit_name[539:558],
	8388608:  _ContextButtonBit_name[558:578],
	16777216: _ContextButtonBit_name[578:600],
	33554432: _ContextButtonBit_name[600:625],
}

func (i ContextButtonBit) String() string {
	if str, ok := _ContextButtonBit_map[i]; ok {
		return str
	}
	return "ContextButtonBit(" + strconv.FormatInt(int64(i), 10) + ")"
}