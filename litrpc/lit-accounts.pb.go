// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.6.1
// source: lit-accounts.proto

package litrpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The initial account balance in satoshis representing the maximum amount that
	// can be spent.
	AccountBalance uint64 `protobuf:"varint,1,opt,name=account_balance,json=accountBalance,proto3" json:"account_balance,omitempty"`
	// The expiration date of the account as a timestamp. Set to 0 to never expire.
	ExpirationDate int64 `protobuf:"varint,2,opt,name=expiration_date,json=expirationDate,proto3" json:"expiration_date,omitempty"`
	// An optional label to identify the account. If the label is not empty, then
	// it must be unique, otherwise it couldn't be used to query a single account.
	Label string `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *CreateAccountRequest) Reset() {
	*x = CreateAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lit_accounts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountRequest) ProtoMessage() {}

func (x *CreateAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lit_accounts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountRequest.ProtoReflect.Descriptor instead.
func (*CreateAccountRequest) Descriptor() ([]byte, []int) {
	return file_lit_accounts_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAccountRequest) GetAccountBalance() uint64 {
	if x != nil {
		return x.AccountBalance
	}
	return 0
}

func (x *CreateAccountRequest) GetExpirationDate() int64 {
	if x != nil {
		return x.ExpirationDate
	}
	return 0
}

func (x *CreateAccountRequest) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type CreateAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The new account that was created.
	Account *Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	// The macaroon with all permissions required to access the account.
	Macaroon []byte `protobuf:"bytes,2,opt,name=macaroon,proto3" json:"macaroon,omitempty"`
}

func (x *CreateAccountResponse) Reset() {
	*x = CreateAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lit_accounts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountResponse) ProtoMessage() {}

func (x *CreateAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lit_accounts_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountResponse.ProtoReflect.Descriptor instead.
func (*CreateAccountResponse) Descriptor() ([]byte, []int) {
	return file_lit_accounts_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAccountResponse) GetAccount() *Account {
	if x != nil {
		return x.Account
	}
	return nil
}

func (x *CreateAccountResponse) GetMacaroon() []byte {
	if x != nil {
		return x.Macaroon
	}
	return nil
}

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the account.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The initial balance in satoshis that was set when the account was created.
	InitialBalance uint64 `protobuf:"varint,2,opt,name=initial_balance,json=initialBalance,proto3" json:"initial_balance,omitempty"`
	// The current balance in satoshis.
	CurrentBalance int64 `protobuf:"varint,3,opt,name=current_balance,json=currentBalance,proto3" json:"current_balance,omitempty"`
	// Timestamp of the last time the account was updated.
	LastUpdate int64 `protobuf:"varint,4,opt,name=last_update,json=lastUpdate,proto3" json:"last_update,omitempty"`
	// Timestamp of the account's expiration date. Zero means it does not expire.
	ExpirationDate int64 `protobuf:"varint,5,opt,name=expiration_date,json=expirationDate,proto3" json:"expiration_date,omitempty"`
	// The list of invoices created by the account. An invoice created by an
	// account will credit the account balance if it is settled.
	Invoices []*AccountInvoice `protobuf:"bytes,6,rep,name=invoices,proto3" json:"invoices,omitempty"`
	// The list of payments made by the account. A payment made by an account will
	// debit the account balance if it is settled.
	Payments []*AccountPayment `protobuf:"bytes,7,rep,name=payments,proto3" json:"payments,omitempty"`
	// An optional label to identify the account. If this is not empty, then it is
	// guaranteed to be unique.
	Label string `protobuf:"bytes,8,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lit_accounts_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_lit_accounts_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_lit_accounts_proto_rawDescGZIP(), []int{2}
}

func (x *Account) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Account) GetInitialBalance() uint64 {
	if x != nil {
		return x.InitialBalance
	}
	return 0
}

func (x *Account) GetCurrentBalance() int64 {
	if x != nil {
		return x.CurrentBalance
	}
	return 0
}

func (x *Account) GetLastUpdate() int64 {
	if x != nil {
		return x.LastUpdate
	}
	return 0
}

func (x *Account) GetExpirationDate() int64 {
	if x != nil {
		return x.ExpirationDate
	}
	return 0
}

func (x *Account) GetInvoices() []*AccountInvoice {
	if x != nil {
		return x.Invoices
	}
	return nil
}

func (x *Account) GetPayments() []*AccountPayment {
	if x != nil {
		return x.Payments
	}
	return nil
}

func (x *Account) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type AccountInvoice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The payment hash of the invoice.
	Hash []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *AccountInvoice) Reset() {
	*x = AccountInvoice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lit_accounts_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountInvoice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountInvoice) ProtoMessage() {}

func (x *AccountInvoice) ProtoReflect() protoreflect.Message {
	mi := &file_lit_accounts_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountInvoice.ProtoReflect.Descriptor instead.
func (*AccountInvoice) Descriptor() ([]byte, []int) {
	return file_lit_accounts_proto_rawDescGZIP(), []int{3}
}

func (x *AccountInvoice) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

type AccountPayment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The payment hash.
	Hash []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	// The state of the payment as reported by lnd.
	State string `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	// The full amount in satoshis reserved for this payment. This includes the
	// routing fee estimated by the fee limit of the payment request. The actual
	// debited amount will likely be lower if the fee is below the limit.
	FullAmount int64 `protobuf:"varint,3,opt,name=full_amount,json=fullAmount,proto3" json:"full_amount,omitempty"`
}

func (x *AccountPayment) Reset() {
	*x = AccountPayment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lit_accounts_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountPayment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountPayment) ProtoMessage() {}

func (x *AccountPayment) ProtoReflect() protoreflect.Message {
	mi := &file_lit_accounts_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountPayment.ProtoReflect.Descriptor instead.
func (*AccountPayment) Descriptor() ([]byte, []int) {
	return file_lit_accounts_proto_rawDescGZIP(), []int{4}
}

func (x *AccountPayment) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *AccountPayment) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *AccountPayment) GetFullAmount() int64 {
	if x != nil {
		return x.FullAmount
	}
	return 0
}

type UpdateAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the account to update. Either the ID or the label must be set.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The new account balance to set. Set to -1 to not update the balance.
	AccountBalance int64 `protobuf:"varint,2,opt,name=account_balance,json=accountBalance,proto3" json:"account_balance,omitempty"`
	// The new account expiry to set. Set to -1 to not update the expiry. Set to 0
	// to never expire.
	ExpirationDate int64 `protobuf:"varint,3,opt,name=expiration_date,json=expirationDate,proto3" json:"expiration_date,omitempty"`
	// The label of the account to update. If an account has no label, then the ID
	// must be used instead.
	Label string `protobuf:"bytes,4,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *UpdateAccountRequest) Reset() {
	*x = UpdateAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lit_accounts_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAccountRequest) ProtoMessage() {}

func (x *UpdateAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lit_accounts_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAccountRequest.ProtoReflect.Descriptor instead.
func (*UpdateAccountRequest) Descriptor() ([]byte, []int) {
	return file_lit_accounts_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateAccountRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateAccountRequest) GetAccountBalance() int64 {
	if x != nil {
		return x.AccountBalance
	}
	return 0
}

func (x *UpdateAccountRequest) GetExpirationDate() int64 {
	if x != nil {
		return x.ExpirationDate
	}
	return 0
}

func (x *UpdateAccountRequest) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type ListAccountsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListAccountsRequest) Reset() {
	*x = ListAccountsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lit_accounts_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAccountsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAccountsRequest) ProtoMessage() {}

func (x *ListAccountsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lit_accounts_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAccountsRequest.ProtoReflect.Descriptor instead.
func (*ListAccountsRequest) Descriptor() ([]byte, []int) {
	return file_lit_accounts_proto_rawDescGZIP(), []int{6}
}

type ListAccountsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// All accounts in the account database.
	Accounts []*Account `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts,omitempty"`
}

func (x *ListAccountsResponse) Reset() {
	*x = ListAccountsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lit_accounts_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAccountsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAccountsResponse) ProtoMessage() {}

func (x *ListAccountsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lit_accounts_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAccountsResponse.ProtoReflect.Descriptor instead.
func (*ListAccountsResponse) Descriptor() ([]byte, []int) {
	return file_lit_accounts_proto_rawDescGZIP(), []int{7}
}

func (x *ListAccountsResponse) GetAccounts() []*Account {
	if x != nil {
		return x.Accounts
	}
	return nil
}

type AccountInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The hexadecimal ID of the account to remove. Either the ID or the label must
	// be set.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The label of the account to remove. If an account has no label, then the ID
	// must be used instead.
	Label string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *AccountInfoRequest) Reset() {
	*x = AccountInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lit_accounts_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountInfoRequest) ProtoMessage() {}

func (x *AccountInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lit_accounts_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountInfoRequest.ProtoReflect.Descriptor instead.
func (*AccountInfoRequest) Descriptor() ([]byte, []int) {
	return file_lit_accounts_proto_rawDescGZIP(), []int{8}
}

func (x *AccountInfoRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AccountInfoRequest) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type RemoveAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The hexadecimal ID of the account to remove. Either the ID or the label must
	// be set.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The label of the account to remove. If an account has no label, then the ID
	// must be used instead.
	Label string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *RemoveAccountRequest) Reset() {
	*x = RemoveAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lit_accounts_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveAccountRequest) ProtoMessage() {}

func (x *RemoveAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lit_accounts_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveAccountRequest.ProtoReflect.Descriptor instead.
func (*RemoveAccountRequest) Descriptor() ([]byte, []int) {
	return file_lit_accounts_proto_rawDescGZIP(), []int{9}
}

func (x *RemoveAccountRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RemoveAccountRequest) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type RemoveAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveAccountResponse) Reset() {
	*x = RemoveAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lit_accounts_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveAccountResponse) ProtoMessage() {}

func (x *RemoveAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lit_accounts_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveAccountResponse.ProtoReflect.Descriptor instead.
func (*RemoveAccountResponse) Descriptor() ([]byte, []int) {
	return file_lit_accounts_proto_rawDescGZIP(), []int{10}
}

var File_lit_accounts_proto protoreflect.FileDescriptor

var file_lit_accounts_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6c, 0x69, 0x74, 0x2d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6c, 0x69, 0x74, 0x72, 0x70, 0x63, 0x22, 0x7e, 0x0a, 0x14,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x27, 0x0a,
	0x0f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x5e, 0x0a, 0x15,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6c, 0x69, 0x74, 0x72, 0x70, 0x63, 0x2e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x61, 0x63, 0x61, 0x72, 0x6f, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x08, 0x6d, 0x61, 0x63, 0x61, 0x72, 0x6f, 0x6f, 0x6e, 0x22, 0xb3, 0x02, 0x0a,
	0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x61, 0x6c, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0e, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x32, 0x0a, 0x08, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6c, 0x69, 0x74, 0x72, 0x70, 0x63, 0x2e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x08,
	0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x08, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6c, 0x69, 0x74,
	0x72, 0x70, 0x63, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x08, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x22, 0x24, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x5b, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61,
	0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x66, 0x75, 0x6c, 0x6c, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x8e, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x27,
	0x0a, 0x0f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x15, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x43, 0x0a,
	0x14, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6c, 0x69, 0x74, 0x72, 0x70, 0x63,
	0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x22, 0x3a, 0x0a, 0x12, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x3c,
	0x0a, 0x14, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x17, 0x0a, 0x15,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xed, 0x02, 0x0a, 0x08, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x12, 0x4c, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1c, 0x2e, 0x6c, 0x69, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x6c, 0x69, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3e, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1c, 0x2e, 0x6c, 0x69, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0f, 0x2e, 0x6c, 0x69, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x49, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x12, 0x1b, 0x2e, 0x6c, 0x69, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x6c, 0x69, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0b, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x2e, 0x6c, 0x69, 0x74,
	0x72, 0x70, 0x63, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x6c, 0x69, 0x74, 0x72, 0x70, 0x63, 0x2e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x4c, 0x0a, 0x0d, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x2e, 0x6c, 0x69, 0x74, 0x72, 0x70,
	0x63, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6c, 0x69, 0x74, 0x72, 0x70, 0x63, 0x2e,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x6e, 0x69, 0x6e, 0x67, 0x6c, 0x61, 0x62,
	0x73, 0x2f, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x6e, 0x69, 0x6e, 0x67, 0x2d, 0x74, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x61, 0x6c, 0x2f, 0x6c, 0x69, 0x74, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_lit_accounts_proto_rawDescOnce sync.Once
	file_lit_accounts_proto_rawDescData = file_lit_accounts_proto_rawDesc
)

func file_lit_accounts_proto_rawDescGZIP() []byte {
	file_lit_accounts_proto_rawDescOnce.Do(func() {
		file_lit_accounts_proto_rawDescData = protoimpl.X.CompressGZIP(file_lit_accounts_proto_rawDescData)
	})
	return file_lit_accounts_proto_rawDescData
}

var file_lit_accounts_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_lit_accounts_proto_goTypes = []interface{}{
	(*CreateAccountRequest)(nil),  // 0: litrpc.CreateAccountRequest
	(*CreateAccountResponse)(nil), // 1: litrpc.CreateAccountResponse
	(*Account)(nil),               // 2: litrpc.Account
	(*AccountInvoice)(nil),        // 3: litrpc.AccountInvoice
	(*AccountPayment)(nil),        // 4: litrpc.AccountPayment
	(*UpdateAccountRequest)(nil),  // 5: litrpc.UpdateAccountRequest
	(*ListAccountsRequest)(nil),   // 6: litrpc.ListAccountsRequest
	(*ListAccountsResponse)(nil),  // 7: litrpc.ListAccountsResponse
	(*AccountInfoRequest)(nil),    // 8: litrpc.AccountInfoRequest
	(*RemoveAccountRequest)(nil),  // 9: litrpc.RemoveAccountRequest
	(*RemoveAccountResponse)(nil), // 10: litrpc.RemoveAccountResponse
}
var file_lit_accounts_proto_depIdxs = []int32{
	2,  // 0: litrpc.CreateAccountResponse.account:type_name -> litrpc.Account
	3,  // 1: litrpc.Account.invoices:type_name -> litrpc.AccountInvoice
	4,  // 2: litrpc.Account.payments:type_name -> litrpc.AccountPayment
	2,  // 3: litrpc.ListAccountsResponse.accounts:type_name -> litrpc.Account
	0,  // 4: litrpc.Accounts.CreateAccount:input_type -> litrpc.CreateAccountRequest
	5,  // 5: litrpc.Accounts.UpdateAccount:input_type -> litrpc.UpdateAccountRequest
	6,  // 6: litrpc.Accounts.ListAccounts:input_type -> litrpc.ListAccountsRequest
	8,  // 7: litrpc.Accounts.AccountInfo:input_type -> litrpc.AccountInfoRequest
	9,  // 8: litrpc.Accounts.RemoveAccount:input_type -> litrpc.RemoveAccountRequest
	1,  // 9: litrpc.Accounts.CreateAccount:output_type -> litrpc.CreateAccountResponse
	2,  // 10: litrpc.Accounts.UpdateAccount:output_type -> litrpc.Account
	7,  // 11: litrpc.Accounts.ListAccounts:output_type -> litrpc.ListAccountsResponse
	2,  // 12: litrpc.Accounts.AccountInfo:output_type -> litrpc.Account
	10, // 13: litrpc.Accounts.RemoveAccount:output_type -> litrpc.RemoveAccountResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_lit_accounts_proto_init() }
func file_lit_accounts_proto_init() {
	if File_lit_accounts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lit_accounts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccountRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_lit_accounts_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccountResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_lit_accounts_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_lit_accounts_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountInvoice); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_lit_accounts_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountPayment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_lit_accounts_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAccountRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_lit_accounts_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAccountsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_lit_accounts_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAccountsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_lit_accounts_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountInfoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_lit_accounts_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveAccountRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_lit_accounts_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveAccountResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_lit_accounts_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_lit_accounts_proto_goTypes,
		DependencyIndexes: file_lit_accounts_proto_depIdxs,
		MessageInfos:      file_lit_accounts_proto_msgTypes,
	}.Build()
	File_lit_accounts_proto = out.File
	file_lit_accounts_proto_rawDesc = nil
	file_lit_accounts_proto_goTypes = nil
	file_lit_accounts_proto_depIdxs = nil
}
