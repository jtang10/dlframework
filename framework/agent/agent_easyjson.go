// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package  agent

import (
  dlframework "github.com/rai-project/dlframework"
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

func easyjsonCb9d4455DecodeGithubComRaiProjectDlframeworkFrameworkAgent(in *jlexer.Lexer, out *base) {
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
    case "Framework":
      easyjsonCb9d4455DecodeGithubComRaiProjectDlframework(in, &out.Framework)
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
func easyjsonCb9d4455EncodeGithubComRaiProjectDlframeworkFrameworkAgent(out *jwriter.Writer, in base) {
  out.RawByte('{')
  first := true
  _ = first
  {
    const prefix string = ",\"Framework\":"
    if first {
      first = false
      out.RawString(prefix[1:])
    } else {
      out.RawString(prefix)
    }
    easyjsonCb9d4455EncodeGithubComRaiProjectDlframework(out, in.Framework)
  }
  out.RawByte('}')
}
// MarshalJSON supports json.Marshaler interface
func (v base) MarshalJSON() ([]byte, error) {
  w := jwriter.Writer{}
  easyjsonCb9d4455EncodeGithubComRaiProjectDlframeworkFrameworkAgent(&w, v)
  return w.Buffer.BuildBytes(), w.Error
}
// MarshalEasyJSON supports easyjson.Marshaler interface
func (v base) MarshalEasyJSON(w *jwriter.Writer) {
  easyjsonCb9d4455EncodeGithubComRaiProjectDlframeworkFrameworkAgent(w, v)
}
// UnmarshalJSON supports json.Unmarshaler interface
func (v *base) UnmarshalJSON(data []byte) error {
  r := jlexer.Lexer{Data: data}
  easyjsonCb9d4455DecodeGithubComRaiProjectDlframeworkFrameworkAgent(&r, v)
  return r.Error()
}
// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *base) UnmarshalEasyJSON(l *jlexer.Lexer) {
  easyjsonCb9d4455DecodeGithubComRaiProjectDlframeworkFrameworkAgent(l, v)
}
func easyjsonCb9d4455DecodeGithubComRaiProjectDlframework(in *jlexer.Lexer, out *dlframework.FrameworkManifest) {
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
    case "name":
      out.Name = string(in.String())
    case "version":
      out.Version = string(in.String())
    case "container":
      if in.IsNull() {
        in.Skip()
      } else {
        in.Delim('{')
        if !in.IsDelim('}') {
        out.Container = make(map[string]*dlframework.ContainerHardware)
        } else {
        out.Container = nil
        }
        for !in.IsDelim('}') {
          key := string(in.String())
          in.WantColon()
          var v1 *dlframework.ContainerHardware
          if in.IsNull() {
            in.Skip()
            v1 = nil
          } else {
            if v1 == nil {
              v1 = new(dlframework.ContainerHardware)
            }
            easyjsonCb9d4455DecodeGithubComRaiProjectDlframework1(in, &*v1)
          }
          (out.Container)[key] = v1
          in.WantComma()
        }
        in.Delim('}')
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
func easyjsonCb9d4455EncodeGithubComRaiProjectDlframework(out *jwriter.Writer, in dlframework.FrameworkManifest) {
  out.RawByte('{')
  first := true
  _ = first
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
  if in.Version != "" {
    const prefix string = ",\"version\":"
    if first {
      first = false
      out.RawString(prefix[1:])
    } else {
      out.RawString(prefix)
    }
    out.String(string(in.Version))
  }
  if len(in.Container) != 0 {
    const prefix string = ",\"container\":"
    if first {
      first = false
      out.RawString(prefix[1:])
    } else {
      out.RawString(prefix)
    }
    {
      out.RawByte('{')
      v2First := true
      for v2Name, v2Value := range in.Container {
        if v2First { v2First = false } else { out.RawByte(',') }
        out.String(string(v2Name))
        out.RawByte(':')
        if v2Value == nil {
          out.RawString("null")
        } else {
          easyjsonCb9d4455EncodeGithubComRaiProjectDlframework1(out, *v2Value)
        }
      }
      out.RawByte('}')
    }
  }
  out.RawByte('}')
}
func easyjsonCb9d4455DecodeGithubComRaiProjectDlframework1(in *jlexer.Lexer, out *dlframework.ContainerHardware) {
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
    case "gpu":
      out.Gpu = string(in.String())
    case "cpu":
      out.Cpu = string(in.String())
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
func easyjsonCb9d4455EncodeGithubComRaiProjectDlframework1(out *jwriter.Writer, in dlframework.ContainerHardware) {
  out.RawByte('{')
  first := true
  _ = first
  if in.Gpu != "" {
    const prefix string = ",\"gpu\":"
    if first {
      first = false
      out.RawString(prefix[1:])
    } else {
      out.RawString(prefix)
    }
    out.String(string(in.Gpu))
  }
  if in.Cpu != "" {
    const prefix string = ",\"cpu\":"
    if first {
      first = false
      out.RawString(prefix[1:])
    } else {
      out.RawString(prefix)
    }
    out.String(string(in.Cpu))
  }
  out.RawByte('}')
}
func easyjsonCb9d4455DecodeGithubComRaiProjectDlframeworkFrameworkAgent1(in *jlexer.Lexer, out *Registry) {
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
    case "Framework":
      easyjsonCb9d4455DecodeGithubComRaiProjectDlframework(in, &out.Framework)
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
func easyjsonCb9d4455EncodeGithubComRaiProjectDlframeworkFrameworkAgent1(out *jwriter.Writer, in Registry) {
  out.RawByte('{')
  first := true
  _ = first
  {
    const prefix string = ",\"Framework\":"
    if first {
      first = false
      out.RawString(prefix[1:])
    } else {
      out.RawString(prefix)
    }
    easyjsonCb9d4455EncodeGithubComRaiProjectDlframework(out, in.Framework)
  }
  out.RawByte('}')
}
// MarshalJSON supports json.Marshaler interface
func (v Registry) MarshalJSON() ([]byte, error) {
  w := jwriter.Writer{}
  easyjsonCb9d4455EncodeGithubComRaiProjectDlframeworkFrameworkAgent1(&w, v)
  return w.Buffer.BuildBytes(), w.Error
}
// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Registry) MarshalEasyJSON(w *jwriter.Writer) {
  easyjsonCb9d4455EncodeGithubComRaiProjectDlframeworkFrameworkAgent1(w, v)
}
// UnmarshalJSON supports json.Unmarshaler interface
func (v *Registry) UnmarshalJSON(data []byte) error {
  r := jlexer.Lexer{Data: data}
  easyjsonCb9d4455DecodeGithubComRaiProjectDlframeworkFrameworkAgent1(&r, v)
  return r.Error()
}
// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Registry) UnmarshalEasyJSON(l *jlexer.Lexer) {
  easyjsonCb9d4455DecodeGithubComRaiProjectDlframeworkFrameworkAgent1(l, v)
}
func easyjsonCb9d4455DecodeGithubComRaiProjectDlframeworkFrameworkAgent2(in *jlexer.Lexer, out *Options) {
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
func easyjsonCb9d4455EncodeGithubComRaiProjectDlframeworkFrameworkAgent2(out *jwriter.Writer, in Options) {
  out.RawByte('{')
  first := true
  _ = first
  out.RawByte('}')
}
// MarshalJSON supports json.Marshaler interface
func (v Options) MarshalJSON() ([]byte, error) {
  w := jwriter.Writer{}
  easyjsonCb9d4455EncodeGithubComRaiProjectDlframeworkFrameworkAgent2(&w, v)
  return w.Buffer.BuildBytes(), w.Error
}
// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Options) MarshalEasyJSON(w *jwriter.Writer) {
  easyjsonCb9d4455EncodeGithubComRaiProjectDlframeworkFrameworkAgent2(w, v)
}
// UnmarshalJSON supports json.Unmarshaler interface
func (v *Options) UnmarshalJSON(data []byte) error {
  r := jlexer.Lexer{Data: data}
  easyjsonCb9d4455DecodeGithubComRaiProjectDlframeworkFrameworkAgent2(&r, v)
  return r.Error()
}
// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Options) UnmarshalEasyJSON(l *jlexer.Lexer) {
  easyjsonCb9d4455DecodeGithubComRaiProjectDlframeworkFrameworkAgent2(l, v)
}
func easyjsonCb9d4455DecodeGithubComRaiProjectDlframeworkFrameworkAgent3(in *jlexer.Lexer, out *Agent) {
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
    case "Framework":
      easyjsonCb9d4455DecodeGithubComRaiProjectDlframework(in, &out.Framework)
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
func easyjsonCb9d4455EncodeGithubComRaiProjectDlframeworkFrameworkAgent3(out *jwriter.Writer, in Agent) {
  out.RawByte('{')
  first := true
  _ = first
  {
    const prefix string = ",\"Framework\":"
    if first {
      first = false
      out.RawString(prefix[1:])
    } else {
      out.RawString(prefix)
    }
    easyjsonCb9d4455EncodeGithubComRaiProjectDlframework(out, in.Framework)
  }
  out.RawByte('}')
}
// MarshalJSON supports json.Marshaler interface
func (v Agent) MarshalJSON() ([]byte, error) {
  w := jwriter.Writer{}
  easyjsonCb9d4455EncodeGithubComRaiProjectDlframeworkFrameworkAgent3(&w, v)
  return w.Buffer.BuildBytes(), w.Error
}
// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Agent) MarshalEasyJSON(w *jwriter.Writer) {
  easyjsonCb9d4455EncodeGithubComRaiProjectDlframeworkFrameworkAgent3(w, v)
}
// UnmarshalJSON supports json.Unmarshaler interface
func (v *Agent) UnmarshalJSON(data []byte) error {
  r := jlexer.Lexer{Data: data}
  easyjsonCb9d4455DecodeGithubComRaiProjectDlframeworkFrameworkAgent3(&r, v)
  return r.Error()
}
// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Agent) UnmarshalEasyJSON(l *jlexer.Lexer) {
  easyjsonCb9d4455DecodeGithubComRaiProjectDlframeworkFrameworkAgent3(l, v)
}
