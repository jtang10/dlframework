// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package  dlframework

import (
  easyjson "github.com/mailru/easyjson"
  jlexer "github.com/mailru/easyjson/jlexer"
  json "encoding/json"
  jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
   _ *json.RawMessage
   _ *jlexer.Lexer
   _ *jwriter.Writer
   _ easyjson.Marshaler
)

func easyjson5f1d405aDecodeGithubComRaiProjectDlframework(in *jlexer.Lexer, out *Features) {
 isTopLevel := in.IsStart()
  if in.IsNull() {
    in.Skip()
    *out = nil
  } else {
    in.Delim('[')
    if *out == nil {
      if !in.IsDelim(']') {
        *out = make(Features, 0, 8)
      } else {
        *out = Features{}
      }
    } else { 
      *out = (*out)[:0]
    }
    for !in.IsDelim(']') {
      var v1 *Feature
      if in.IsNull() {
        in.Skip()
        v1 = nil
      } else {
        if v1 == nil {
          v1 = new(Feature)
        }
        easyjson5f1d405aDecodeGithubComRaiProjectDlframework1(in, &*v1)
      }
      *out = append(*out, v1)
      in.WantComma()
    }
    in.Delim(']')
  }
  if isTopLevel {
    in.Consumed()
  }
}
func easyjson5f1d405aEncodeGithubComRaiProjectDlframework(out *jwriter.Writer, in Features) {
  if in == nil && (out.Flags & jwriter.NilSliceAsEmpty) == 0 {
    out.RawString("null")
  } else {
    out.RawByte('[')
    for v2, v3 := range in {
      if v2 > 0 {
        out.RawByte(',')
      }
      if v3 == nil {
        out.RawString("null")
      } else {
        easyjson5f1d405aEncodeGithubComRaiProjectDlframework1(out, *v3)
      }
    }
    out.RawByte(']')
  }
}
// MarshalJSON supports json.Marshaler interface
func (v Features) MarshalJSON() ([]byte, error) {
  w := jwriter.Writer{}
  easyjson5f1d405aEncodeGithubComRaiProjectDlframework(&w, v)
  return w.Buffer.BuildBytes(), w.Error
}
// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Features) MarshalEasyJSON(w *jwriter.Writer) {
  easyjson5f1d405aEncodeGithubComRaiProjectDlframework(w, v)
}
// UnmarshalJSON supports json.Unmarshaler interface
func (v *Features) UnmarshalJSON(data []byte) error {
  r := jlexer.Lexer{Data: data}
  easyjson5f1d405aDecodeGithubComRaiProjectDlframework(&r, v)
  return r.Error()
}
// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Features) UnmarshalEasyJSON(l *jlexer.Lexer) {
  easyjson5f1d405aDecodeGithubComRaiProjectDlframework(l, v)
}
func easyjson5f1d405aDecodeGithubComRaiProjectDlframework1(in *jlexer.Lexer, out *Feature) {
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
    case "index":
      out.Index = int32(in.Int32())
    case "name":
      out.Name = string(in.String())
    case "probability":
      out.Probability = float32(in.Float32())
    case "metadata":
      if in.IsNull() {
        in.Skip()
      } else {
        in.Delim('{')
        if !in.IsDelim('}') {
        out.Metadata = make(map[string]string)
        } else {
        out.Metadata = nil
        }
        for !in.IsDelim('}') {
          key := string(in.String())
          in.WantColon()
          var v4 string
          v4 = string(in.String())
          (out.Metadata)[key] = v4
          in.WantComma()
        }
        in.Delim('}')
      }
    default:
      in.AddError(&jlexer.LexerError{
          Offset: in.GetPos(),
          Reason: "unknown field",
          Data: key,
      })
    }
    in.WantComma()
  }
  in.Delim('}')
  if isTopLevel {
    in.Consumed()
  }
}
func easyjson5f1d405aEncodeGithubComRaiProjectDlframework1(out *jwriter.Writer, in Feature) {
  out.RawByte('{')
  first := true
  _ = first
  if in.Index != 0 {
    const prefix string = ",\"index\":"
    if first {
      first = false
      out.RawString(prefix[1:])
    } else {
      out.RawString(prefix)
    }
    out.Int32(int32(in.Index))
  }
  if in.Name != "" {
    const prefix string = ",\"name\":"
    if first {
      first = false
      out.RawString(prefix[1:])
    } else {
      out.RawString(prefix)
    }
    out.String(string(in.Name))
  }
  if in.Probability != 0 {
    const prefix string = ",\"probability\":"
    if first {
      first = false
      out.RawString(prefix[1:])
    } else {
      out.RawString(prefix)
    }
    out.Float32(float32(in.Probability))
  }
  if len(in.Metadata) != 0 {
    const prefix string = ",\"metadata\":"
    if first {
      first = false
      out.RawString(prefix[1:])
    } else {
      out.RawString(prefix)
    }
    {
      out.RawByte('{')
      v5First := true
      for v5Name, v5Value := range in.Metadata {
        if v5First { v5First = false } else { out.RawByte(',') }
        out.String(string(v5Name))
        out.RawByte(':')
        out.String(string(v5Value))
      }
      out.RawByte('}')
    }
  }
  out.RawByte('}')
}
