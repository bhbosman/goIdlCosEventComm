// Code generated by me. DO NOT EDIT.

package goIdlCosEventComm

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CosEventComm::PullSupplier_try_pull_Out", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CosEventCommPullSupplierTryPullOut struct {
	__goidl__.IdlObject
	HasEvent bool `json:"HasEvent"`
}

//noinspection GoSnakeCaseUsage
const CosEventCommPullSupplierTryPullOutId_Const = "IDL:CosEventComm/PullSupplier_try_pull_Out:1.0"

func (self *CosEventCommPullSupplierTryPullOut) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CosEventCommPullSupplierTryPullOut) GoString() string {
	return self.String()
}

func (self *CosEventCommPullSupplierTryPullOut) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(BooleanType)
	self.HasEvent, err = __goidl__.IdlBooleanHelper.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CosEventCommPullSupplierTryPullOut) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CosEventCommPullSupplierTryPullOut) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(BooleanType)
	err = __goidl__.IdlBooleanHelper.Write(stream, self.HasEvent)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CosEventCommPullSupplierTryPullOut_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CosEventCommPullSupplierTryPullOutHelper = CosEventCommPullSupplierTryPullOut_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CosEventCommPullSupplierTryPullOutId_Const,
			__goidl__.StructType,
			"CosEventComm.idl",
			"xdl_CosEventCommPullSupplierTryPullOut.go",
			func() __goidl__.IIdlObject {
				return &CosEventCommPullSupplierTryPullOut{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CosEventCommPullSupplierTryPullOut{}
			},
			__reflect__.TypeOf((*CosEventCommPullSupplierTryPullOut)(nil))))
}