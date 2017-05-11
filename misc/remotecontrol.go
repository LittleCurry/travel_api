package misc

import (
	"git.iguiyu.com/park/api/db"
	"git.iguiyu.com/park/misc/sms_message"

	"git.iguiyu.com/park/api/rpc_client"
	"git.iguiyu.com/park/struct/model"
	pb "git.iguiyu.com/park/struct/protobuf"
	"golang.org/x/net/context"
	"log"
)

func sendSMS(mobile string, content string) {
	request := &pb.SendRequest{
		Mobile:  mobile,
		Content: content,
	}
	_, err := rpc_client.GetSmsGatewayClient().Send(context.Background(), request)
	if err != nil {
		log.Println(err)
	}
}

func DownLocker(lockid int) error {
	sendContent, err := sms_message.NewSendContent(sms_message.SEND_ARM_DOWN, true)
	locker := model.Locker{}
	has, err := db.MySQL().Id(lockid).Get(&locker)
	if err == nil && has {
		sendSMS(locker.PhoneNumber, sendContent.String())
		return nil
	}
	return err
}

func UpLocker(lockid int) error {
	sendContent, err := sms_message.NewSendContent(sms_message.SEND_ARM_UP, true)
	locker := model.Locker{}
	has, err := db.MySQL().Id(lockid).Get(&locker)
	if err == nil && has {
		sendSMS(locker.PhoneNumber, sendContent.String())
		return nil
	}
	return err
}
