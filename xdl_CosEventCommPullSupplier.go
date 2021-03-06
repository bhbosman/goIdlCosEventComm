// Code generated by me. DO NOT EDIT.

package goIdlCosEventComm

import __goidl__ "github.com/bhbosman/goidl"

// Interface declaration: "CosEventComm::PullSupplier", generatedBy by: "WriteInterface"
type CosEventCommPullSupplier interface {
	// Original name: "disconnect_pull_supplier"
	DisconnectPullSupplier(params CosEventCommPullSupplierDisconnectPullSupplierIn) (CosEventCommPullSupplierDisconnectPullSupplierOut, error)
	//Exceptions for : Pull
	//	CosEventCommDisconnected
	// Original name: "pull"
	Pull(params CosEventCommPullSupplierPullIn) (CosEventCommPullSupplierPullOut, error)
	//Exceptions for : TryPull
	//	CosEventCommDisconnected
	// Original name: "try_pull"
	TryPull(params CosEventCommPullSupplierTryPullIn) (CosEventCommPullSupplierTryPullOut, error)
}

const CosEventCommPullSupplierDisconnectPullSupplierOperation = "disconnect_pull_supplier"
const CosEventCommPullSupplierPullOperation = "pull"
const CosEventCommPullSupplierTryPullOperation = "try_pull"
//noinspection GoSnakeCaseUsage
type CosEventCommPullSupplier_Helper struct {
}

//noinspection GoSnakeCaseUsage
const CosEventCommPullSupplierId_Const = "IDL:omg.org/CosEventComm/PullSupplier:1.0"

func (self CosEventCommPullSupplier_Helper) Id() string {
	return CosEventCommPullSupplierId_Const
}

func (self CosEventCommPullSupplier_Helper) Read(stream __goidl__.IReadAny) (CosEventCommPullSupplier, error) {
	return nil, nil
}

func (self CosEventCommPullSupplier_Helper) Write(stream __goidl__.IWriteAny, v CosEventCommPullSupplier) error {
	return nil
}


//noinspection GoUnusedGlobalVariable
var CosEventCommPullSupplierHelper = CosEventCommPullSupplier_Helper{}

func init() {
}
