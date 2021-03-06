// Copyright (c) 2013, Vastech SA (PTY) LTD. All rights reserved.
// http://code.google.com/p/gogoprotobuf
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package gogoproto

import google_protobuf "github.com/dropbox/goprotoc/protoc-gen-dgo/descriptor"
import proto "github.com/dropbox/goprotoc/proto"

func IsEmbed(field *google_protobuf.FieldDescriptorProto) bool {
	return proto.GetBoolExtension(field.Options, E_Embed, false)
}

func IsCustomType(field *google_protobuf.FieldDescriptorProto) bool {
	typ := GetCustomType(field)
	if len(typ) > 0 {
		return true
	}
	return false
}

func GetCustomType(field *google_protobuf.FieldDescriptorProto) string {
	if field.Options != nil {
		v, err := proto.GetExtension(field.Options, E_Customtype)
		if err == nil && v.(*string) != nil {
			return *(v.(*string))
		}
	}
	return ""
}

func IsCustomName(field *google_protobuf.FieldDescriptorProto) bool {
	name := GetCustomName(field)
	if len(name) > 0 {
		return true
	}
	return false
}

func GetCustomName(field *google_protobuf.FieldDescriptorProto) string {
	if field.Options != nil {
		v, err := proto.GetExtension(field.Options, E_Customname)
		if err == nil && v.(*string) != nil {
			return *(v.(*string))
		}
	}
	return ""
}

func GetJsonTag(field *google_protobuf.FieldDescriptorProto) *string {
	if field.Options != nil {
		v, err := proto.GetExtension(field.Options, E_Jsontag)
		if err == nil && v.(*string) != nil {
			return (v.(*string))
		}
	}
	return nil
}

func GetMoreTags(field *google_protobuf.FieldDescriptorProto) *string {
	if field.Options != nil {
		v, err := proto.GetExtension(field.Options, E_Moretags)
		if err == nil && v.(*string) != nil {
			return (v.(*string))
		}
	}
	return nil
}

type EnableFunc func(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool

func EnabledGoEnumPrefix(file *google_protobuf.FileDescriptorProto, enum *google_protobuf.EnumDescriptorProto) bool {
	return proto.GetBoolExtension(enum.Options, E_GoprotoEnumPrefix, proto.GetBoolExtension(file.Options, E_GoprotoEnumPrefixAll, true))
}

func IsUnion(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
	return proto.GetBoolExtension(message.Options, E_Onlyone, proto.GetBoolExtension(file.Options, E_OnlyoneAll, false))
}

func HasEqual(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
	return proto.GetBoolExtension(message.Options, E_Equal, proto.GetBoolExtension(file.Options, E_EqualAll, false))
}

func HasVerboseEqual(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
	return proto.GetBoolExtension(message.Options, E_VerboseEqual, proto.GetBoolExtension(file.Options, E_VerboseEqualAll, false))
}

func IsStringer(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
	return proto.GetBoolExtension(message.Options, E_Stringer, proto.GetBoolExtension(file.Options, E_StringerAll, true))
}

func IsFace(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
	return proto.GetBoolExtension(message.Options, E_Face, proto.GetBoolExtension(file.Options, E_FaceAll, false))
}

func HasDescription(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
	return proto.GetBoolExtension(message.Options, E_Description, proto.GetBoolExtension(file.Options, E_DescriptionAll, false))
}

func HasPopulate(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
	return proto.GetBoolExtension(message.Options, E_Populate, proto.GetBoolExtension(file.Options, E_PopulateAll, false))
}

func HasTestGen(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
	return proto.GetBoolExtension(message.Options, E_Testgen, proto.GetBoolExtension(file.Options, E_TestgenAll, false))
}

func HasBenchGen(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
	return proto.GetBoolExtension(message.Options, E_Benchgen, proto.GetBoolExtension(file.Options, E_BenchgenAll, false))
}

func IsMarshaler(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
	return proto.GetBoolExtension(message.Options, E_Marshaler, proto.GetBoolExtension(file.Options, E_MarshalerAll, true))
}

func IsUnmarshaler(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
	return proto.GetBoolExtension(message.Options, E_Unmarshaler, proto.GetBoolExtension(file.Options, E_UnmarshalerAll, true))
}

func HasBufferTo(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
	return proto.GetBoolExtension(message.Options, E_Bufferto, proto.GetBoolExtension(file.Options, E_BuffertoAll, false))
}

func IsSizer(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
	return proto.GetBoolExtension(message.Options, E_Sizer, proto.GetBoolExtension(file.Options, E_SizerAll, true))
}

func IsGoEnumStringer(file *google_protobuf.FileDescriptorProto, enum *google_protobuf.EnumDescriptorProto) bool {
	return proto.GetBoolExtension(enum.Options, E_GoprotoEnumStringer, proto.GetBoolExtension(file.Options, E_GoprotoEnumStringerAll, true))
}

func IsEnumStringer(file *google_protobuf.FileDescriptorProto, enum *google_protobuf.EnumDescriptorProto) bool {
	return proto.GetBoolExtension(enum.Options, E_EnumStringer, proto.GetBoolExtension(file.Options, E_EnumStringerAll, false))
}

func HasExtensionsMap(file *google_protobuf.FileDescriptorProto, message *google_protobuf.DescriptorProto) bool {
	return proto.GetBoolExtension(message.Options, E_GoprotoExtensionsMap, proto.GetBoolExtension(file.Options, E_GoprotoExtensionsMapAll, true))
}
