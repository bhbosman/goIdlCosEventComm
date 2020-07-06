// Code generated by me. DO NOT EDIT.

package goIdlCosEventComm

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CosEventComm::PullSupplier_pull_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CosEventCommPullSupplierPullIn struct {
	__goidl__.IdlObject
}

//noinspection GoSnakeCaseUsage
const CosEventCommPullSupplierPullInId_Const = "IDL:CosEventComm/PullSupplier_pull_In:1.0"

func (self *CosEventCommPullSupplierPullIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CosEventCommPullSupplierPullIn) GoString() string {
	return self.String()
}

func (self *CosEventCommPullSupplierPullIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CosEventCommPullSupplierPullIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CosEventCommPullSupplierPullIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CosEventCommPullSupplierPullIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CosEventCommPullSupplierPullInHelper = CosEventCommPullSupplierPullIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CosEventCommPullSupplierPullInId_Const,
			__goidl__.StructType,
			"CosEventComm.idl",
			"xdl_CosEventCommPullSupplierPullIn.go",
			func() __goidl__.IIdlObject {
				return &CosEventCommPullSupplierPullIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CosEventCommPullSupplierPullIn{}
			},
			__reflect__.TypeOf((*CosEventCommPullSupplierPullIn)(nil))))
}