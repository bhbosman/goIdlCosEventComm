// Code generated by me. DO NOT EDIT.

package goIdlCosEventComm

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CosEventComm::PushConsumer_disconnect_push_consumer_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CosEventCommPushConsumerDisconnectPushConsumerIn struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CosEventCommPushConsumerDisconnectPushConsumerInId_Const = "IDL:CosEventComm/PushConsumer_disconnect_push_consumer_In:1.0"

func (self *CosEventCommPushConsumerDisconnectPushConsumerIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CosEventCommPushConsumerDisconnectPushConsumerIn) GoString() string {
	return self.String()
}

func (self *CosEventCommPushConsumerDisconnectPushConsumerIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CosEventCommPushConsumerDisconnectPushConsumerIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CosEventCommPushConsumerDisconnectPushConsumerIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CosEventCommPushConsumerDisconnectPushConsumerIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CosEventCommPushConsumerDisconnectPushConsumerInHelper = CosEventCommPushConsumerDisconnectPushConsumerIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CosEventCommPushConsumerDisconnectPushConsumerInId_Const,
			__goidl__.StructType,
			"CosEventComm.idl",
			"xdl_CosEventCommPushConsumerDisconnectPushConsumerIn.go",
			func() __goidl__.IIdlObject {
				return &CosEventCommPushConsumerDisconnectPushConsumerIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CosEventCommPushConsumerDisconnectPushConsumerIn{}
			},
			__reflect__.TypeOf((*CosEventCommPushConsumerDisconnectPushConsumerIn)(nil))))
}