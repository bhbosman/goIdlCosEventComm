// Code generated by me. DO NOT EDIT.

package goIdlCosEventComm

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CosEventComm::PushConsumer_disconnect_push_consumer_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CosEventCommPushConsumerDisconnectPushConsumerOut struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CosEventCommPushConsumerDisconnectPushConsumerOutId_Const = "IDL:CosEventComm/PushConsumer_disconnect_push_consumer_Out:1.0"

func (self *CosEventCommPushConsumerDisconnectPushConsumerOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CosEventCommPushConsumerDisconnectPushConsumerOut) GoString() string {
	return self.String()
}

func (self *CosEventCommPushConsumerDisconnectPushConsumerOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CosEventCommPushConsumerDisconnectPushConsumerOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CosEventCommPushConsumerDisconnectPushConsumerOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CosEventCommPushConsumerDisconnectPushConsumerOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CosEventCommPushConsumerDisconnectPushConsumerOutHelper = CosEventCommPushConsumerDisconnectPushConsumerOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CosEventCommPushConsumerDisconnectPushConsumerOutId_Const,
			__goidl__.StructType,
			"CosEventComm.idl",
			"xdl_CosEventCommPushConsumerDisconnectPushConsumerOut.go",
			func() __goidl__.IIdlObject {
				return &CosEventCommPushConsumerDisconnectPushConsumerOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CosEventCommPushConsumerDisconnectPushConsumerOut{}
			},
			__reflect__.TypeOf((*CosEventCommPushConsumerDisconnectPushConsumerOut)(nil))))
}
