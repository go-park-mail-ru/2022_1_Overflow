// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson521a5691DecodeOverflowBackendInternalModels(in *jlexer.Lexer, out *ProfileInfo) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.Id = int32(in.Int32())
		case "first_name":
			out.Firstname = string(in.String())
		case "last_name":
			out.Lastname = string(in.String())
		case "username":
			out.Username = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson521a5691EncodeOverflowBackendInternalModels(out *jwriter.Writer, in ProfileInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int32(int32(in.Id))
	}
	{
		const prefix string = ",\"first_name\":"
		out.RawString(prefix)
		out.String(string(in.Firstname))
	}
	{
		const prefix string = ",\"last_name\":"
		out.RawString(prefix)
		out.String(string(in.Lastname))
	}
	{
		const prefix string = ",\"username\":"
		out.RawString(prefix)
		out.String(string(in.Username))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ProfileInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson521a5691EncodeOverflowBackendInternalModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ProfileInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson521a5691EncodeOverflowBackendInternalModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ProfileInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson521a5691DecodeOverflowBackendInternalModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ProfileInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson521a5691DecodeOverflowBackendInternalModels(l, v)
}
func easyjson521a5691DecodeOverflowBackendInternalModels1(in *jlexer.Lexer, out *Avatar) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "name":
			out.Name = string(in.String())
		case "username":
			out.Username = string(in.String())
		case "file":
			if in.IsNull() {
				in.Skip()
				out.File = nil
			} else {
				out.File = in.Bytes()
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson521a5691EncodeOverflowBackendInternalModels1(out *jwriter.Writer, in Avatar) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"username\":"
		out.RawString(prefix)
		out.String(string(in.Username))
	}
	{
		const prefix string = ",\"file\":"
		out.RawString(prefix)
		out.Base64Bytes(in.File)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Avatar) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson521a5691EncodeOverflowBackendInternalModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Avatar) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson521a5691EncodeOverflowBackendInternalModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Avatar) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson521a5691DecodeOverflowBackendInternalModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Avatar) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson521a5691DecodeOverflowBackendInternalModels1(l, v)
}
