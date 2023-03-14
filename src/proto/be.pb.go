// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: src/proto/be.proto

package src

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_be_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_be_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_src_proto_be_proto_rawDescGZIP(), []int{0}
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Des   string `protobuf:"bytes,1,opt,name=Des,proto3" json:"Des,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_be_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_be_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_src_proto_be_proto_rawDescGZIP(), []int{1}
}

func (x *Status) GetDes() string {
	if x != nil {
		return x.Des
	}
	return ""
}

func (x *Status) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type CourseDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        int32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	FacultyID int32  `protobuf:"varint,3,opt,name=FacultyID,proto3" json:"FacultyID,omitempty"`
	Faculty   *User  `protobuf:"bytes,4,opt,name=Faculty,proto3" json:"Faculty,omitempty"`
}

func (x *CourseDetails) Reset() {
	*x = CourseDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_be_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourseDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourseDetails) ProtoMessage() {}

func (x *CourseDetails) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_be_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourseDetails.ProtoReflect.Descriptor instead.
func (*CourseDetails) Descriptor() ([]byte, []int) {
	return file_src_proto_be_proto_rawDescGZIP(), []int{2}
}

func (x *CourseDetails) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *CourseDetails) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CourseDetails) GetFacultyID() int32 {
	if x != nil {
		return x.FacultyID
	}
	return 0
}

func (x *CourseDetails) GetFaculty() *User {
	if x != nil {
		return x.Faculty
	}
	return nil
}

type CourseList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Courses []*CourseDetails `protobuf:"bytes,1,rep,name=Courses,proto3" json:"Courses,omitempty"`
}

func (x *CourseList) Reset() {
	*x = CourseList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_be_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourseList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourseList) ProtoMessage() {}

func (x *CourseList) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_be_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourseList.ProtoReflect.Descriptor instead.
func (*CourseList) Descriptor() ([]byte, []int) {
	return file_src_proto_be_proto_rawDescGZIP(), []int{3}
}

func (x *CourseList) GetCourses() []*CourseDetails {
	if x != nil {
		return x.Courses
	}
	return nil
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        int32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	FirstName string `protobuf:"bytes,2,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName  string `protobuf:"bytes,3,opt,name=LastName,proto3" json:"LastName,omitempty"`
	Email     string `protobuf:"bytes,4,opt,name=Email,proto3" json:"Email,omitempty"`
	Password  string `protobuf:"bytes,5,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_be_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_be_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_src_proto_be_proto_rawDescGZIP(), []int{4}
}

func (x *User) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UserList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *UserList) Reset() {
	*x = UserList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_be_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserList) ProtoMessage() {}

func (x *UserList) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_be_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserList.ProtoReflect.Descriptor instead.
func (*UserList) Descriptor() ([]byte, []int) {
	return file_src_proto_be_proto_rawDescGZIP(), []int{5}
}

func (x *UserList) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type EnrollRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentID uint32 `protobuf:"varint,1,opt,name=StudentID,proto3" json:"StudentID,omitempty"`
	CourseID  uint32 `protobuf:"varint,2,opt,name=CourseID,proto3" json:"CourseID,omitempty"`
}

func (x *EnrollRequest) Reset() {
	*x = EnrollRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_be_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnrollRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnrollRequest) ProtoMessage() {}

func (x *EnrollRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_be_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnrollRequest.ProtoReflect.Descriptor instead.
func (*EnrollRequest) Descriptor() ([]byte, []int) {
	return file_src_proto_be_proto_rawDescGZIP(), []int{6}
}

func (x *EnrollRequest) GetStudentID() uint32 {
	if x != nil {
		return x.StudentID
	}
	return 0
}

func (x *EnrollRequest) GetCourseID() uint32 {
	if x != nil {
		return x.CourseID
	}
	return 0
}

type Assignment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          int32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	Deadline    string `protobuf:"bytes,4,opt,name=Deadline,proto3" json:"Deadline,omitempty"`
	CourseId    int32  `protobuf:"varint,5,opt,name=CourseId,proto3" json:"CourseId,omitempty"`
}

func (x *Assignment) Reset() {
	*x = Assignment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_be_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Assignment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Assignment) ProtoMessage() {}

func (x *Assignment) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_be_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Assignment.ProtoReflect.Descriptor instead.
func (*Assignment) Descriptor() ([]byte, []int) {
	return file_src_proto_be_proto_rawDescGZIP(), []int{7}
}

func (x *Assignment) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Assignment) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Assignment) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Assignment) GetDeadline() string {
	if x != nil {
		return x.Deadline
	}
	return ""
}

func (x *Assignment) GetCourseId() int32 {
	if x != nil {
		return x.CourseId
	}
	return 0
}

type AssignmentList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Assignments []*Assignment `protobuf:"bytes,1,rep,name=Assignments,proto3" json:"Assignments,omitempty"`
}

func (x *AssignmentList) Reset() {
	*x = AssignmentList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_be_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssignmentList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignmentList) ProtoMessage() {}

func (x *AssignmentList) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_be_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignmentList.ProtoReflect.Descriptor instead.
func (*AssignmentList) Descriptor() ([]byte, []int) {
	return file_src_proto_be_proto_rawDescGZIP(), []int{8}
}

func (x *AssignmentList) GetAssignments() []*Assignment {
	if x != nil {
		return x.Assignments
	}
	return nil
}

var File_src_proto_be_proto protoreflect.FileDescriptor

var file_src_proto_be_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x30, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x44, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x44, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x7f, 0x0a,
	0x0d, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12,
	0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x46, 0x61, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x49, 0x44, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x46, 0x61, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x49, 0x44,
	0x12, 0x2c, 0x0a, 0x07, 0x46, 0x61, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x07, 0x46, 0x61, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x22, 0x43,
	0x0a, 0x0a, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x07,
	0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x07, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x73, 0x22, 0x82, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09,
	0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x61,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x34, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x49,
	0x0a, 0x0d, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1a, 0x0a,
	0x08, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x44, 0x22, 0x8c, 0x01, 0x0a, 0x0a, 0x41, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x22, 0x4c, 0x0a, 0x0e, 0x41, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x0b, 0x41, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41,
	0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x41, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x32, 0x9b, 0x09, 0x0a, 0x0e, 0x55, 0x6e, 0x69, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x74, 0x79, 0x46, 0x75, 0x6e, 0x63, 0x12, 0x47, 0x0a, 0x09, 0x41, 0x64, 0x64,
	0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x1b, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x1a, 0x1b, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x22, 0x00, 0x12, 0x36, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x12, 0x12, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x1a, 0x12, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0a, 0x41, 0x64,
	0x64, 0x46, 0x61, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x12, 0x12, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x12, 0x2e, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x22, 0x00, 0x12, 0x44, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x1b, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x41,
	0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x2e, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x1a, 0x18, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x12,
	0x51, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x42, 0x79, 0x43, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x1a, 0x1b, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x22, 0x00, 0x12, 0x3d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x12, 0x2e, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22,
	0x00, 0x12, 0x3d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x46, 0x61, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x42,
	0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x12, 0x2e, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00,
	0x12, 0x40, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x73, 0x12, 0x13, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x22, 0x00, 0x12, 0x52, 0x0a, 0x20, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73,
	0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x64, 0x42, 0x79, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x18, 0x2e, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x73, 0x42, 0x79, 0x46, 0x61, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x12, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x18, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x22, 0x00, 0x12, 0x49, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x46, 0x61, 0x63, 0x75, 0x6c, 0x74, 0x79,
	0x42, 0x79, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x2e, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x1a, 0x12, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x12, 0x56, 0x0a,
	0x1d, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x64, 0x42, 0x79, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x12, 0x1b,
	0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x1a, 0x16, 0x2e, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49,
	0x64, 0x12, 0x1b, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x1a, 0x1c,
	0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x4a,
	0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x1b,
	0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x1a, 0x1b, 0x2e, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x10, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18,
	0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x18, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x62, 0x65, 0x70, 0x72, 0x6a, 0x2f, 0x73, 0x72,
	0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_proto_be_proto_rawDescOnce sync.Once
	file_src_proto_be_proto_rawDescData = file_src_proto_be_proto_rawDesc
)

func file_src_proto_be_proto_rawDescGZIP() []byte {
	file_src_proto_be_proto_rawDescOnce.Do(func() {
		file_src_proto_be_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_proto_be_proto_rawDescData)
	})
	return file_src_proto_be_proto_rawDescData
}

var file_src_proto_be_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_src_proto_be_proto_goTypes = []interface{}{
	(*Empty)(nil),          // 0: backendProto.Empty
	(*Status)(nil),         // 1: backendProto.Status
	(*CourseDetails)(nil),  // 2: backendProto.CourseDetails
	(*CourseList)(nil),     // 3: backendProto.CourseList
	(*User)(nil),           // 4: backendProto.User
	(*UserList)(nil),       // 5: backendProto.UserList
	(*EnrollRequest)(nil),  // 6: backendProto.EnrollRequest
	(*Assignment)(nil),     // 7: backendProto.Assignment
	(*AssignmentList)(nil), // 8: backendProto.AssignmentList
}
var file_src_proto_be_proto_depIdxs = []int32{
	4,  // 0: backendProto.CourseDetails.Faculty:type_name -> backendProto.User
	2,  // 1: backendProto.CourseList.Courses:type_name -> backendProto.CourseDetails
	4,  // 2: backendProto.UserList.users:type_name -> backendProto.User
	7,  // 3: backendProto.AssignmentList.Assignments:type_name -> backendProto.Assignment
	2,  // 4: backendProto.UniversityFunc.AddCourse:input_type -> backendProto.CourseDetails
	4,  // 5: backendProto.UniversityFunc.AddStudent:input_type -> backendProto.User
	4,  // 6: backendProto.UniversityFunc.AddFaculty:input_type -> backendProto.User
	6,  // 7: backendProto.UniversityFunc.AddEnrollment:input_type -> backendProto.EnrollRequest
	7,  // 8: backendProto.UniversityFunc.AddAssignment:input_type -> backendProto.Assignment
	2,  // 9: backendProto.UniversityFunc.GetCourseByCourseId:input_type -> backendProto.CourseDetails
	4,  // 10: backendProto.UniversityFunc.GetStudentByEmail:input_type -> backendProto.User
	4,  // 11: backendProto.UniversityFunc.GetFacultyByEmail:input_type -> backendProto.User
	0,  // 12: backendProto.UniversityFunc.GetAllCourses:input_type -> backendProto.Empty
	4,  // 13: backendProto.UniversityFunc.GetCoursesEnrolledByStudentEmail:input_type -> backendProto.User
	4,  // 14: backendProto.UniversityFunc.GetCoursesByFacultyEmail:input_type -> backendProto.User
	2,  // 15: backendProto.UniversityFunc.GetFacultyByCourseId:input_type -> backendProto.CourseDetails
	2,  // 16: backendProto.UniversityFunc.GetStudentsEnrolledByCourseId:input_type -> backendProto.CourseDetails
	2,  // 17: backendProto.UniversityFunc.GetAssignmentsByCourseId:input_type -> backendProto.CourseDetails
	2,  // 18: backendProto.UniversityFunc.UpdateCourse:input_type -> backendProto.CourseDetails
	7,  // 19: backendProto.UniversityFunc.UpdateAssignment:input_type -> backendProto.Assignment
	2,  // 20: backendProto.UniversityFunc.AddCourse:output_type -> backendProto.CourseDetails
	4,  // 21: backendProto.UniversityFunc.AddStudent:output_type -> backendProto.User
	4,  // 22: backendProto.UniversityFunc.AddFaculty:output_type -> backendProto.User
	1,  // 23: backendProto.UniversityFunc.AddEnrollment:output_type -> backendProto.Status
	7,  // 24: backendProto.UniversityFunc.AddAssignment:output_type -> backendProto.Assignment
	2,  // 25: backendProto.UniversityFunc.GetCourseByCourseId:output_type -> backendProto.CourseDetails
	4,  // 26: backendProto.UniversityFunc.GetStudentByEmail:output_type -> backendProto.User
	4,  // 27: backendProto.UniversityFunc.GetFacultyByEmail:output_type -> backendProto.User
	3,  // 28: backendProto.UniversityFunc.GetAllCourses:output_type -> backendProto.CourseList
	3,  // 29: backendProto.UniversityFunc.GetCoursesEnrolledByStudentEmail:output_type -> backendProto.CourseList
	3,  // 30: backendProto.UniversityFunc.GetCoursesByFacultyEmail:output_type -> backendProto.CourseList
	4,  // 31: backendProto.UniversityFunc.GetFacultyByCourseId:output_type -> backendProto.User
	5,  // 32: backendProto.UniversityFunc.GetStudentsEnrolledByCourseId:output_type -> backendProto.UserList
	8,  // 33: backendProto.UniversityFunc.GetAssignmentsByCourseId:output_type -> backendProto.AssignmentList
	2,  // 34: backendProto.UniversityFunc.UpdateCourse:output_type -> backendProto.CourseDetails
	7,  // 35: backendProto.UniversityFunc.UpdateAssignment:output_type -> backendProto.Assignment
	20, // [20:36] is the sub-list for method output_type
	4,  // [4:20] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_src_proto_be_proto_init() }
func file_src_proto_be_proto_init() {
	if File_src_proto_be_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_proto_be_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_src_proto_be_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_src_proto_be_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CourseDetails); i {
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
		file_src_proto_be_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CourseList); i {
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
		file_src_proto_be_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_src_proto_be_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserList); i {
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
		file_src_proto_be_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnrollRequest); i {
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
		file_src_proto_be_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Assignment); i {
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
		file_src_proto_be_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssignmentList); i {
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
			RawDescriptor: file_src_proto_be_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_src_proto_be_proto_goTypes,
		DependencyIndexes: file_src_proto_be_proto_depIdxs,
		MessageInfos:      file_src_proto_be_proto_msgTypes,
	}.Build()
	File_src_proto_be_proto = out.File
	file_src_proto_be_proto_rawDesc = nil
	file_src_proto_be_proto_goTypes = nil
	file_src_proto_be_proto_depIdxs = nil
}
