// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package http2Protocol

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

func easyjsonE4425964DecodeNwayplayServerTestHttp2Protocol(in *jlexer.Lexer, out *TestResponse) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "msg":
			out.Msg = string(in.String())
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
func easyjsonE4425964EncodeNwayplayServerTestHttp2Protocol(out *jwriter.Writer, in TestResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"msg\":"
		out.RawString(prefix[1:])
		out.String(string(in.Msg))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v TestResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE4425964EncodeNwayplayServerTestHttp2Protocol(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v TestResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE4425964EncodeNwayplayServerTestHttp2Protocol(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *TestResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE4425964DecodeNwayplayServerTestHttp2Protocol(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *TestResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE4425964DecodeNwayplayServerTestHttp2Protocol(l, v)
}
func easyjsonE4425964DecodeNwayplayServerTestHttp2Protocol1(in *jlexer.Lexer, out *TestRequest) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "msg":
			out.Msg = string(in.String())
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
func easyjsonE4425964EncodeNwayplayServerTestHttp2Protocol1(out *jwriter.Writer, in TestRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"msg\":"
		out.RawString(prefix[1:])
		out.String(string(in.Msg))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v TestRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE4425964EncodeNwayplayServerTestHttp2Protocol1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v TestRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE4425964EncodeNwayplayServerTestHttp2Protocol1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *TestRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE4425964DecodeNwayplayServerTestHttp2Protocol1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *TestRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE4425964DecodeNwayplayServerTestHttp2Protocol1(l, v)
}
