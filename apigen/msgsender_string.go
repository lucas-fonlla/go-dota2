// Code generated by "stringer -type=MsgSender"; DO NOT EDIT.

package main

import "strconv"

const _MsgSender_name = "MsgSenderNoneMsgSenderGCMsgSenderClient"

var _MsgSender_index = [...]uint8{0, 13, 24, 39}

func (i MsgSender) String() string {
	if i >= MsgSender(len(_MsgSender_index)-1) {
		return "MsgSender(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MsgSender_name[_MsgSender_index[i]:_MsgSender_index[i+1]]
}