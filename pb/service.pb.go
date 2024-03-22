// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.1
// source: proto/service.proto

package pb

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

type BankRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BankName      string          `protobuf:"bytes,1,opt,name=bank_name,json=bankName,proto3" json:"bank_name,omitempty"`
	AccountNumber int64           `protobuf:"varint,2,opt,name=account_number,json=accountNumber,proto3" json:"account_number,omitempty"`
	AccountType   BankAccountType `protobuf:"varint,3,opt,name=account_type,json=accountType,proto3,enum=bank.BankAccountType" json:"account_type,omitempty"`
}

func (x *BankRequest) Reset() {
	*x = BankRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankRequest) ProtoMessage() {}

func (x *BankRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BankRequest.ProtoReflect.Descriptor instead.
func (*BankRequest) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{0}
}

func (x *BankRequest) GetBankName() string {
	if x != nil {
		return x.BankName
	}
	return ""
}

func (x *BankRequest) GetAccountNumber() int64 {
	if x != nil {
		return x.AccountNumber
	}
	return 0
}

func (x *BankRequest) GetAccountType() BankAccountType {
	if x != nil {
		return x.AccountType
	}
	return BankAccountType_SAVINGS
}

type BankResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BankName       string            `protobuf:"bytes,1,opt,name=bank_name,json=bankName,proto3" json:"bank_name,omitempty"`
	AccountHolder  string            `protobuf:"bytes,2,opt,name=account_holder,json=accountHolder,proto3" json:"account_holder,omitempty"`
	AccountBalance float64           `protobuf:"fixed64,3,opt,name=account_balance,json=accountBalance,proto3" json:"account_balance,omitempty"`
	AccountStatus  BankAccountStatus `protobuf:"varint,4,opt,name=account_status,json=accountStatus,proto3,enum=bank.BankAccountStatus" json:"account_status,omitempty"`
}

func (x *BankResponse) Reset() {
	*x = BankResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankResponse) ProtoMessage() {}

func (x *BankResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BankResponse.ProtoReflect.Descriptor instead.
func (*BankResponse) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{1}
}

func (x *BankResponse) GetBankName() string {
	if x != nil {
		return x.BankName
	}
	return ""
}

func (x *BankResponse) GetAccountHolder() string {
	if x != nil {
		return x.AccountHolder
	}
	return ""
}

func (x *BankResponse) GetAccountBalance() float64 {
	if x != nil {
		return x.AccountBalance
	}
	return 0
}

func (x *BankResponse) GetAccountStatus() BankAccountStatus {
	if x != nil {
		return x.AccountStatus
	}
	return BankAccountStatus_INACTIVE
}

var File_proto_service_proto protoreflect.FileDescriptor

var file_proto_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x6e,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x01, 0x0a, 0x0b, 0x42, 0x61, 0x6e, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x61, 0x6e, 0x6b, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x61, 0x6e, 0x6b,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x0c, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x15, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0xbb, 0x01, 0x0a, 0x0c, 0x42, 0x61, 0x6e, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x61, 0x6e, 0x6b, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x68,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x48, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x62, 0x61,
	0x6e, 0x6b, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x32, 0x42, 0x0a, 0x0b, 0x42, 0x61, 0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x33, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6e, 0x6b, 0x49, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_service_proto_rawDescOnce sync.Once
	file_proto_service_proto_rawDescData = file_proto_service_proto_rawDesc
)

func file_proto_service_proto_rawDescGZIP() []byte {
	file_proto_service_proto_rawDescOnce.Do(func() {
		file_proto_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_service_proto_rawDescData)
	})
	return file_proto_service_proto_rawDescData
}

var file_proto_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_service_proto_goTypes = []interface{}{
	(*BankRequest)(nil),    // 0: BankRequest
	(*BankResponse)(nil),   // 1: BankResponse
	(BankAccountType)(0),   // 2: bank.BankAccountType
	(BankAccountStatus)(0), // 3: bank.BankAccountStatus
}
var file_proto_service_proto_depIdxs = []int32{
	2, // 0: BankRequest.account_type:type_name -> bank.BankAccountType
	3, // 1: BankResponse.account_status:type_name -> bank.BankAccountStatus
	0, // 2: BankService.GetBankInformation:input_type -> BankRequest
	1, // 3: BankService.GetBankInformation:output_type -> BankResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_service_proto_init() }
func file_proto_service_proto_init() {
	if File_proto_service_proto != nil {
		return
	}
	file_proto_bank_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BankRequest); i {
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
		file_proto_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BankResponse); i {
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
			RawDescriptor: file_proto_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_service_proto_goTypes,
		DependencyIndexes: file_proto_service_proto_depIdxs,
		MessageInfos:      file_proto_service_proto_msgTypes,
	}.Build()
	File_proto_service_proto = out.File
	file_proto_service_proto_rawDesc = nil
	file_proto_service_proto_goTypes = nil
	file_proto_service_proto_depIdxs = nil
}