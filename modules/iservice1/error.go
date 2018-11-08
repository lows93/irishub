package iservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"fmt"
)

const (
	DefaultCodespace sdk.CodespaceType = 6

	CodeInvalidIDL               sdk.CodeType = 100
	CodeSvcDefExists             sdk.CodeType = 101
	CodeSvcDefNotExists          sdk.CodeType = 102
	CodeInvalidOutputPrivacyEnum sdk.CodeType = 103
	CodeInvalidOutputCachedEnum  sdk.CodeType = 104
	CodeInvalidServiceName       sdk.CodeType = 105
	CodeInvalidChainId           sdk.CodeType = 106
	CodeInvalidAuthor            sdk.CodeType = 107
	CodeInvalidMethodName        sdk.CodeType = 108
	CodeInvalidMessagingType     sdk.CodeType = 109
	CodeMoreTags                 sdk.CodeType = 110
	CodeDuplicateTags            sdk.CodeType = 111

	CodeSvcBindingExists    sdk.CodeType = 112
	CodeSvcBindingNotExists sdk.CodeType = 113
	CodeInvalidDefChainId   sdk.CodeType = 114
	CodeInvalidBindingType  sdk.CodeType = 115
	CodeInvalidLevel        sdk.CodeType = 116
	CodeInvalidPriceCount   sdk.CodeType = 117
	CodeInvalidUpdate       sdk.CodeType = 118
	CodeRefundDeposit       sdk.CodeType = 119
	CodeInvalidExpiration   sdk.CodeType = 120
)

func codeToDefaultMsg(code sdk.CodeType) string {
	switch code {
	case CodeInvalidIDL:
		return "The IDL file cannot be parsed"
	default:
		return sdk.CodeToDefaultMsg(code)
	}
}

func NewError(codespace sdk.CodespaceType, code sdk.CodeType, msg string) sdk.Error {
	msg = msgOrDefaultMsg(msg, code)
	return sdk.NewError(codespace, code, msg)
}

func msgOrDefaultMsg(msg string, code sdk.CodeType) string {
	if msg != "" {
		return msg
	}
	return codeToDefaultMsg(code)
}

func ErrSvcDefExists(codespace sdk.CodespaceType, defChainId, svcDefName string) sdk.Error {
	return sdk.NewError(codespace, CodeSvcDefExists, fmt.Sprintf("service definition name %s already exists in %s", svcDefName, defChainId))
}

func ErrSvcDefNotExists(codespace sdk.CodespaceType, defChainId, svcDefName string) sdk.Error {
	return sdk.NewError(codespace, CodeSvcDefNotExists, fmt.Sprintf("service definition name %s is not existed in %s", svcDefName, defChainId))
}

func ErrInvalidIDL(codespace sdk.CodespaceType, msg string) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidIDL, fmt.Sprintf("The IDL content cannot be parsed, %s", msg))
}

func ErrInvalidOutputPrivacyEnum(codespace sdk.CodespaceType, value string) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidOutputPrivacyEnum, fmt.Sprintf("invalid OutputPrivacyEnum %s", value))
}

func ErrInvalidOutputCachedEnum(codespace sdk.CodespaceType, value string) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidOutputCachedEnum, fmt.Sprintf("invalid OutputCachedEnum %s", value))
}

func ErrInvalidServiceName(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidServiceName, fmt.Sprintf("service name is empty"))
}

func ErrInvalidChainId(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidChainId, fmt.Sprintf("chainId is empty"))
}

func ErrInvalidAuthor(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidAuthor, fmt.Sprintf("author is empty"))
}

func ErrInvalidMethodName(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidMethodName, fmt.Sprintf("method name is empty"))
}

func ErrInvalidMessagingType(codespace sdk.CodespaceType, value MessagingType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidMessagingType, fmt.Sprintf("invalid messaging type %s", value))
}

func ErrMoreTags(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeMoreTags, fmt.Sprintf("tags are limited to %d", iserviceParams.MaxTagsNum))
}

func ErrDuplicateTags(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeDuplicateTags, "tags contains duplicate tag")
}

func ErrInvalidDefChainId(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidDefChainId, fmt.Sprintf("def-chain-id is empty"))
}

func ErrSvcBindingExists(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeSvcBindingExists, fmt.Sprintf("service binding already exists"))
}

func ErrSvcBindingNotExists(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeSvcBindingNotExists, fmt.Sprintf("service binding is not existed"))
}

func ErrInvalidBindingType(codespace sdk.CodespaceType, bindingType BindingType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidBindingType, fmt.Sprintf("invalid binding type %s", bindingType))
}

func ErrInvalidLevel(codespace sdk.CodespaceType, level Level) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidLevel, fmt.Sprintf("invalid level %v, must avg_rsp_time>0 and 0<usable_time<=10000", level))
}

func ErrInvalidPriceCount(codespace sdk.CodespaceType, priceCount int, methodCount int) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidPriceCount, fmt.Sprintf("invalid prices count %d, but methods count is %d", priceCount, methodCount))
}

func ErrInvalidUpdate(codespace sdk.CodespaceType, msg string) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidUpdate, fmt.Sprintf("invalid service binding update, %s", msg))
}

func ErrRefundDeposit(codespace sdk.CodespaceType, msg string) sdk.Error {
	return sdk.NewError(codespace, CodeRefundDeposit, fmt.Sprintf("can't refund deposit, %s", msg))
}

func ErrInvalidExpiration(codespace sdk.CodespaceType, expiration int64) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidExpiration, fmt.Sprintf("invalid expiration %v, can't be equal to zero", expiration))
}

func ErrLtMinProviderDeposit(codespace sdk.CodespaceType, coins sdk.Coins) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidExpiration, fmt.Sprintf("deposit amount must >= %s", coins.String()))
}
