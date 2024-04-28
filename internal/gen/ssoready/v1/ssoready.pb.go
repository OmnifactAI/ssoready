// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: ssoready/v1/ssoready.proto

package ssoreadyv1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type AppUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Email       string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *AppUser) Reset() {
	*x = AppUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssoready_v1_ssoready_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppUser) ProtoMessage() {}

func (x *AppUser) ProtoReflect() protoreflect.Message {
	mi := &file_ssoready_v1_ssoready_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppUser.ProtoReflect.Descriptor instead.
func (*AppUser) Descriptor() ([]byte, []int) {
	return file_ssoready_v1_ssoready_proto_rawDescGZIP(), []int{0}
}

func (x *AppUser) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AppUser) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *AppUser) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type Environment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RedirectUrl string `protobuf:"bytes,2,opt,name=redirect_url,json=redirectUrl,proto3" json:"redirect_url,omitempty"`
}

func (x *Environment) Reset() {
	*x = Environment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssoready_v1_ssoready_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Environment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Environment) ProtoMessage() {}

func (x *Environment) ProtoReflect() protoreflect.Message {
	mi := &file_ssoready_v1_ssoready_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Environment.ProtoReflect.Descriptor instead.
func (*Environment) Descriptor() ([]byte, []int) {
	return file_ssoready_v1_ssoready_proto_rawDescGZIP(), []int{1}
}

func (x *Environment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Environment) GetRedirectUrl() string {
	if x != nil {
		return x.RedirectUrl
	}
	return ""
}

type Organization struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	EnvironmentId string `protobuf:"bytes,2,opt,name=environment_id,json=environmentId,proto3" json:"environment_id,omitempty"`
}

func (x *Organization) Reset() {
	*x = Organization{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssoready_v1_ssoready_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Organization) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Organization) ProtoMessage() {}

func (x *Organization) ProtoReflect() protoreflect.Message {
	mi := &file_ssoready_v1_ssoready_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Organization.ProtoReflect.Descriptor instead.
func (*Organization) Descriptor() ([]byte, []int) {
	return file_ssoready_v1_ssoready_proto_rawDescGZIP(), []int{2}
}

func (x *Organization) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Organization) GetEnvironmentId() string {
	if x != nil {
		return x.EnvironmentId
	}
	return ""
}

type SAMLConnection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OrganizationId     string `protobuf:"bytes,2,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	IdpRedirectUrl     string `protobuf:"bytes,3,opt,name=idp_redirect_url,json=idpRedirectUrl,proto3" json:"idp_redirect_url,omitempty"`
	IdpX509Certificate []byte `protobuf:"bytes,4,opt,name=idp_x509_certificate,json=idpX509Certificate,proto3" json:"idp_x509_certificate,omitempty"`
	IdpEntityId        string `protobuf:"bytes,5,opt,name=idp_entity_id,json=idpEntityId,proto3" json:"idp_entity_id,omitempty"`
}

func (x *SAMLConnection) Reset() {
	*x = SAMLConnection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssoready_v1_ssoready_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SAMLConnection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SAMLConnection) ProtoMessage() {}

func (x *SAMLConnection) ProtoReflect() protoreflect.Message {
	mi := &file_ssoready_v1_ssoready_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SAMLConnection.ProtoReflect.Descriptor instead.
func (*SAMLConnection) Descriptor() ([]byte, []int) {
	return file_ssoready_v1_ssoready_proto_rawDescGZIP(), []int{3}
}

func (x *SAMLConnection) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SAMLConnection) GetOrganizationId() string {
	if x != nil {
		return x.OrganizationId
	}
	return ""
}

func (x *SAMLConnection) GetIdpRedirectUrl() string {
	if x != nil {
		return x.IdpRedirectUrl
	}
	return ""
}

func (x *SAMLConnection) GetIdpX509Certificate() []byte {
	if x != nil {
		return x.IdpX509Certificate
	}
	return nil
}

func (x *SAMLConnection) GetIdpEntityId() string {
	if x != nil {
		return x.IdpEntityId
	}
	return ""
}

type RedeemSAMLAccessTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (x *RedeemSAMLAccessTokenRequest) Reset() {
	*x = RedeemSAMLAccessTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssoready_v1_ssoready_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedeemSAMLAccessTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedeemSAMLAccessTokenRequest) ProtoMessage() {}

func (x *RedeemSAMLAccessTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ssoready_v1_ssoready_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedeemSAMLAccessTokenRequest.ProtoReflect.Descriptor instead.
func (*RedeemSAMLAccessTokenRequest) Descriptor() ([]byte, []int) {
	return file_ssoready_v1_ssoready_proto_rawDescGZIP(), []int{4}
}

func (x *RedeemSAMLAccessTokenRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type RedeemSAMLAccessTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubjectIdpId         string            `protobuf:"bytes,1,opt,name=subject_idp_id,json=subjectIdpId,proto3" json:"subject_idp_id,omitempty"`
	SubjectIdpAttributes map[string]string `protobuf:"bytes,2,rep,name=subject_idp_attributes,json=subjectIdpAttributes,proto3" json:"subject_idp_attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	OrganizationId       string            `protobuf:"bytes,3,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	EnvironmentId        string            `protobuf:"bytes,4,opt,name=environment_id,json=environmentId,proto3" json:"environment_id,omitempty"`
}

func (x *RedeemSAMLAccessTokenResponse) Reset() {
	*x = RedeemSAMLAccessTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssoready_v1_ssoready_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedeemSAMLAccessTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedeemSAMLAccessTokenResponse) ProtoMessage() {}

func (x *RedeemSAMLAccessTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ssoready_v1_ssoready_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedeemSAMLAccessTokenResponse.ProtoReflect.Descriptor instead.
func (*RedeemSAMLAccessTokenResponse) Descriptor() ([]byte, []int) {
	return file_ssoready_v1_ssoready_proto_rawDescGZIP(), []int{5}
}

func (x *RedeemSAMLAccessTokenResponse) GetSubjectIdpId() string {
	if x != nil {
		return x.SubjectIdpId
	}
	return ""
}

func (x *RedeemSAMLAccessTokenResponse) GetSubjectIdpAttributes() map[string]string {
	if x != nil {
		return x.SubjectIdpAttributes
	}
	return nil
}

func (x *RedeemSAMLAccessTokenResponse) GetOrganizationId() string {
	if x != nil {
		return x.OrganizationId
	}
	return ""
}

func (x *RedeemSAMLAccessTokenResponse) GetEnvironmentId() string {
	if x != nil {
		return x.EnvironmentId
	}
	return ""
}

type SignInRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoogleCredential string `protobuf:"bytes,1,opt,name=google_credential,json=googleCredential,proto3" json:"google_credential,omitempty"`
}

func (x *SignInRequest) Reset() {
	*x = SignInRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssoready_v1_ssoready_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInRequest) ProtoMessage() {}

func (x *SignInRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ssoready_v1_ssoready_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInRequest.ProtoReflect.Descriptor instead.
func (*SignInRequest) Descriptor() ([]byte, []int) {
	return file_ssoready_v1_ssoready_proto_rawDescGZIP(), []int{6}
}

func (x *SignInRequest) GetGoogleCredential() string {
	if x != nil {
		return x.GoogleCredential
	}
	return ""
}

type SignInResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionToken string `protobuf:"bytes,1,opt,name=session_token,json=sessionToken,proto3" json:"session_token,omitempty"`
}

func (x *SignInResponse) Reset() {
	*x = SignInResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssoready_v1_ssoready_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInResponse) ProtoMessage() {}

func (x *SignInResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ssoready_v1_ssoready_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInResponse.ProtoReflect.Descriptor instead.
func (*SignInResponse) Descriptor() ([]byte, []int) {
	return file_ssoready_v1_ssoready_proto_rawDescGZIP(), []int{7}
}

func (x *SignInResponse) GetSessionToken() string {
	if x != nil {
		return x.SessionToken
	}
	return ""
}

var File_ssoready_v1_ssoready_proto protoreflect.FileDescriptor

var file_ssoready_v1_ssoready_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x73, 0x6f, 0x72, 0x65, 0x61, 0x64, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x73,
	0x6f, 0x72, 0x65, 0x61, 0x64, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x73,
	0x6f, 0x72, 0x65, 0x61, 0x64, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a, 0x07, 0x41, 0x70, 0x70, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x40, 0x0a, 0x0b, 0x45,
	0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x6c, 0x22, 0x45, 0x0a,
	0x0c, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a,
	0x0e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x22, 0xc9, 0x01, 0x0a, 0x0e, 0x53, 0x41, 0x4d, 0x4c, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x28, 0x0a, 0x10, 0x69, 0x64, 0x70, 0x5f, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x64, 0x70, 0x52,
	0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x30, 0x0a, 0x14, 0x69, 0x64,
	0x70, 0x5f, 0x78, 0x35, 0x30, 0x39, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x12, 0x69, 0x64, 0x70, 0x58, 0x35, 0x30,
	0x39, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0d,
	0x69, 0x64, 0x70, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x64, 0x70, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64,
	0x22, 0x41, 0x0a, 0x1c, 0x52, 0x65, 0x64, 0x65, 0x65, 0x6d, 0x53, 0x41, 0x4d, 0x4c, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0xda, 0x02, 0x0a, 0x1d, 0x52, 0x65, 0x64, 0x65, 0x65, 0x6d, 0x53, 0x41,
	0x4d, 0x4c, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x70, 0x49, 0x64, 0x12, 0x7a, 0x0a, 0x16, 0x73,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x70, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x44, 0x2e, 0x73, 0x73,
	0x6f, 0x72, 0x65, 0x61, 0x64, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x64, 0x65, 0x65, 0x6d,
	0x53, 0x41, 0x4d, 0x4c, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x64, 0x70, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x14, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x70, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x1a, 0x47, 0x0a, 0x19, 0x53, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x64, 0x70, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x3c, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2b, 0x0a, 0x11, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x22, 0x35,
	0x0a, 0x0e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xe1, 0x01, 0x0a, 0x0f, 0x53, 0x53, 0x4f, 0x52, 0x65, 0x61,
	0x64, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8a, 0x01, 0x0a, 0x15, 0x52, 0x65,
	0x64, 0x65, 0x65, 0x6d, 0x53, 0x41, 0x4d, 0x4c, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x29, 0x2e, 0x73, 0x73, 0x6f, 0x72, 0x65, 0x61, 0x64, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x64, 0x65, 0x65, 0x6d, 0x53, 0x41, 0x4d, 0x4c, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a,
	0x2e, 0x73, 0x73, 0x6f, 0x72, 0x65, 0x61, 0x64, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x64,
	0x65, 0x65, 0x6d, 0x53, 0x41, 0x4d, 0x4c, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x61, 0x6d, 0x6c, 0x2f,
	0x72, 0x65, 0x64, 0x65, 0x65, 0x6d, 0x12, 0x41, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e,
	0x12, 0x1a, 0x2e, 0x73, 0x73, 0x6f, 0x72, 0x65, 0x61, 0x64, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73,
	0x73, 0x6f, 0x72, 0x65, 0x61, 0x64, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xaf, 0x01, 0x0a, 0x0f, 0x63, 0x6f,
	0x6d, 0x2e, 0x73, 0x73, 0x6f, 0x72, 0x65, 0x61, 0x64, 0x79, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x53,
	0x73, 0x6f, 0x72, 0x65, 0x61, 0x64, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x40,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x73, 0x6f, 0x72, 0x65,
	0x61, 0x64, 0x79, 0x2f, 0x73, 0x73, 0x6f, 0x72, 0x65, 0x61, 0x64, 0x79, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x73, 0x73, 0x6f, 0x72, 0x65, 0x61,
	0x64, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x73, 0x6f, 0x72, 0x65, 0x61, 0x64, 0x79, 0x76, 0x31,
	0xa2, 0x02, 0x03, 0x53, 0x58, 0x58, 0xaa, 0x02, 0x0b, 0x53, 0x73, 0x6f, 0x72, 0x65, 0x61, 0x64,
	0x79, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0b, 0x53, 0x73, 0x6f, 0x72, 0x65, 0x61, 0x64, 0x79, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x17, 0x53, 0x73, 0x6f, 0x72, 0x65, 0x61, 0x64, 0x79, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x53,
	0x73, 0x6f, 0x72, 0x65, 0x61, 0x64, 0x79, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_ssoready_v1_ssoready_proto_rawDescOnce sync.Once
	file_ssoready_v1_ssoready_proto_rawDescData = file_ssoready_v1_ssoready_proto_rawDesc
)

func file_ssoready_v1_ssoready_proto_rawDescGZIP() []byte {
	file_ssoready_v1_ssoready_proto_rawDescOnce.Do(func() {
		file_ssoready_v1_ssoready_proto_rawDescData = protoimpl.X.CompressGZIP(file_ssoready_v1_ssoready_proto_rawDescData)
	})
	return file_ssoready_v1_ssoready_proto_rawDescData
}

var file_ssoready_v1_ssoready_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_ssoready_v1_ssoready_proto_goTypes = []interface{}{
	(*AppUser)(nil),                       // 0: ssoready.v1.AppUser
	(*Environment)(nil),                   // 1: ssoready.v1.Environment
	(*Organization)(nil),                  // 2: ssoready.v1.Organization
	(*SAMLConnection)(nil),                // 3: ssoready.v1.SAMLConnection
	(*RedeemSAMLAccessTokenRequest)(nil),  // 4: ssoready.v1.RedeemSAMLAccessTokenRequest
	(*RedeemSAMLAccessTokenResponse)(nil), // 5: ssoready.v1.RedeemSAMLAccessTokenResponse
	(*SignInRequest)(nil),                 // 6: ssoready.v1.SignInRequest
	(*SignInResponse)(nil),                // 7: ssoready.v1.SignInResponse
	nil,                                   // 8: ssoready.v1.RedeemSAMLAccessTokenResponse.SubjectIdpAttributesEntry
}
var file_ssoready_v1_ssoready_proto_depIdxs = []int32{
	8, // 0: ssoready.v1.RedeemSAMLAccessTokenResponse.subject_idp_attributes:type_name -> ssoready.v1.RedeemSAMLAccessTokenResponse.SubjectIdpAttributesEntry
	4, // 1: ssoready.v1.SSOReadyService.RedeemSAMLAccessToken:input_type -> ssoready.v1.RedeemSAMLAccessTokenRequest
	6, // 2: ssoready.v1.SSOReadyService.SignIn:input_type -> ssoready.v1.SignInRequest
	5, // 3: ssoready.v1.SSOReadyService.RedeemSAMLAccessToken:output_type -> ssoready.v1.RedeemSAMLAccessTokenResponse
	7, // 4: ssoready.v1.SSOReadyService.SignIn:output_type -> ssoready.v1.SignInResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ssoready_v1_ssoready_proto_init() }
func file_ssoready_v1_ssoready_proto_init() {
	if File_ssoready_v1_ssoready_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ssoready_v1_ssoready_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppUser); i {
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
		file_ssoready_v1_ssoready_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Environment); i {
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
		file_ssoready_v1_ssoready_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Organization); i {
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
		file_ssoready_v1_ssoready_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SAMLConnection); i {
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
		file_ssoready_v1_ssoready_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedeemSAMLAccessTokenRequest); i {
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
		file_ssoready_v1_ssoready_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedeemSAMLAccessTokenResponse); i {
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
		file_ssoready_v1_ssoready_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInRequest); i {
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
		file_ssoready_v1_ssoready_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInResponse); i {
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
			RawDescriptor: file_ssoready_v1_ssoready_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ssoready_v1_ssoready_proto_goTypes,
		DependencyIndexes: file_ssoready_v1_ssoready_proto_depIdxs,
		MessageInfos:      file_ssoready_v1_ssoready_proto_msgTypes,
	}.Build()
	File_ssoready_v1_ssoready_proto = out.File
	file_ssoready_v1_ssoready_proto_rawDesc = nil
	file_ssoready_v1_ssoready_proto_goTypes = nil
	file_ssoready_v1_ssoready_proto_depIdxs = nil
}
