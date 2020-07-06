// Code generated by me. DO NOT EDIT.

package goIdlCosEventComm

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CosEventComm::PullSupplier_try_pull_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CosEventCommPullSupplierTryPullIn struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CosEventCommPullSupplierTryPullInId_Const = "IDL:CosEventComm/PullSupplier_try_pull_In:1.0"

func (self *CosEventCommPullSupplierTryPullIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CosEventCommPullSupplierTryPullIn) GoString() string {
	return self.String()
}

func (self *CosEventCommPullSupplierTryPullIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CosEventCommPullSupplierTryPullIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CosEventCommPullSupplierTryPullIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CosEventCommPullSupplierTryPullIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CosEventCommPullSupplierTryPullInHelper = CosEventCommPullSupplierTryPullIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CosEventCommPullSupplierTryPullInId_Const,
			__goidl__.StructType,
			"CosEventComm.idl",
			"xdl_CosEventCommPullSupplierTryPullIn.go",
			func() __goidl__.IIdlObject {
				return &CosEventCommPullSupplierTryPullIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CosEventCommPullSupplierTryPullIn{}
			},
			__reflect__.TypeOf((*CosEventCommPullSupplierTryPullIn)(nil))))
}