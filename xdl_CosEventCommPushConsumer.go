// Code generated by me. DO NOT EDIT.

package goIdlCosEventComm

import __goidl__ "github.com/bhbosman/goidl"

// Interface declaration: "CosEventComm::PushConsumer", generatedBy by: "WriteInterface"
type CosEventCommPushConsumer interface {
	// Original name: "disconnect_push_consumer"
	DisconnectPushConsumer(params CosEventCommPushConsumerDisconnectPushConsumerIn) (CosEventCommPushConsumerDisconnectPushConsumerOut, error)
	//Exceptions for : Push
	//	CosEventCommDisconnected
	// Original name: "push"
	Push(params CosEventCommPushConsumerPushIn) (CosEventCommPushConsumerPushOut, error)
}

const CosEventCommPushConsumerDisconnectPushConsumerOperation = "disconnect_push_consumer"
const CosEventCommPushConsumerPushOperation = "push"
//noinspection GoSnakeCaseUsage
type CosEventCommPushConsumer_Helper struct {
}

//noinspection GoSnakeCaseUsage
const CosEventCommPushConsumerId_Const = "IDL:omg.org/CosEventComm/PushConsumer:1.0"

func (self CosEventCommPushConsumer_Helper) Id() string {
	return CosEventCommPushConsumerId_Const
}

func (self CosEventCommPushConsumer_Helper) Read(stream __goidl__.IReadAny) (CosEventCommPushConsumer, error) {
	return nil, nil
}

func (self CosEventCommPushConsumer_Helper) Write(stream __goidl__.IWriteAny, v CosEventCommPushConsumer) error {
	return nil
}


//noinspection GoUnusedGlobalVariable
var CosEventCommPushConsumerHelper = CosEventCommPushConsumer_Helper{}

func init() {
}
