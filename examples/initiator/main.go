package main

import (
	"bytes"
	"context"
	"fmt"
	simplefixgo "github.com/b2broker/simplefix-go"
	"github.com/b2broker/simplefix-go/fix"
	flow "github.com/b2broker/simplefix-go/session"
	"github.com/b2broker/simplefix-go/session/messages"
	fixgen "github.com/b2broker/simplefix-go/tests/fix44"
	"net"
	"strconv"
	"time"
)

func mustConvToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}

var PseudoGeneratedOpts = flow.SessionOpts{
	LogonBuilder:         fixgen.Logon{}.New(),
	LogoutBuilder:        fixgen.Logout{}.New(),
	RejectBuilder:        fixgen.Reject{}.New(),
	HeartbeatBuilder:     fixgen.Heartbeat{}.New(),
	TestRequestBuilder:   fixgen.TestRequest{}.New(),
	ResendRequestBuilder: fixgen.ResendRequest{}.New(),
	Tags: messages.Tags{
		MsgType:         mustConvToInt(fixgen.FieldMsgType),
		MsgSeqNum:       mustConvToInt(fixgen.FieldMsgSeqNum),
		HeartBtInt:      mustConvToInt(fixgen.FieldHeartBtInt),
		EncryptedMethod: mustConvToInt(fixgen.FieldEncryptMethod),
	},
	AllowedEncryptedMethods: map[string]struct{}{
		fixgen.EnumEncryptMethodNoneother: {},
	},
	SessionErrorCodes: messages.SessionErrorCodes{
		RequiredTagMissing: 1,
		IncorrectValue:     5,
		Other:              99,
	},
}

func main() {
	conn, err := net.Dial("tcp", fmt.Sprintf(":%d", 9091))
	if err != nil {
		panic(fmt.Errorf("could not dial: %s", err))
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	handler := simplefixgo.NewInitiatorHandler(ctx, fixgen.FieldMsgType, 10)
	client := simplefixgo.NewInitiator(conn, handler, 10)

	handler.OnConnect(func() bool {
		return true
	})

	session := flow.NewInitiatorSession(
		context.Background(),
		handler,
		PseudoGeneratedOpts,
		flow.LogonSettings{
			TargetCompID:  "Server",
			SenderCompID:  "Client",
			HeartBtInt:    5,
			EncryptMethod: fixgen.EnumEncryptMethodNoneother,
			Password:      "password",
			Username:      "login",
		},
	)

	handler.HandleIncoming(fixgen.MsgTypeLogon, func(msg []byte) {
		incomingLogon, err := fixgen.ParseLogon(msg)
		_, _ = incomingLogon, err
	})

	handler.HandleIncoming(simplefixgo.AllMsgTypes, func(msg []byte) {
		fmt.Println("incoming", string(bytes.Replace(msg, fix.Delimiter, []byte("|"), -1)))
	})
	handler.HandleOutgoing(simplefixgo.AllMsgTypes, func(msg []byte) {
		fmt.Println("outgoing", string(bytes.Replace(msg, fix.Delimiter, []byte("|"), -1)))
	})

	go func() {
		time.Sleep(time.Second * 10)
		fmt.Println("resend request after 10 seconds")
		session.Send(fixgen.ResendRequest{}.New().SetFieldBeginSeqNo(2).SetFieldEndSeqNo(3))
	}()

	_ = session.Run()

	panic(client.Serve())
}