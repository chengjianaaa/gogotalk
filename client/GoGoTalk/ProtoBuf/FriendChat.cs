// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: FriendChat.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace GoGoTalk.ProtoBuf {

  /// <summary>Holder for reflection information generated from FriendChat.proto</summary>
  public static partial class FriendChatReflection {

    #region Descriptor
    /// <summary>File descriptor for FriendChat.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static FriendChatReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "ChBGcmllbmRDaGF0LnByb3RvEhFHb0dvVGFsay5Qcm90b0J1ZiJKChFGcmll",
            "bmRDaGF0UmVxdWVzdBISCgpGcm9tVXNlcklEGAEgASgJEhAKCFRvVXNlcklE",
            "GAIgASgJEg8KB0NvbnRlbnQYAyABKAwiSQoQRnJpZW5kQ2hhdE5vdGljZRIS",
            "CgpGcm9tVXNlcklEGAEgASgJEhAKCFRvVXNlcklEGAIgASgJEg8KB0NvbnRl",
            "bnQYAyABKAxiBnByb3RvMw=="));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { },
          new pbr::GeneratedClrTypeInfo(null, null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::GoGoTalk.ProtoBuf.FriendChatRequest), global::GoGoTalk.ProtoBuf.FriendChatRequest.Parser, new[]{ "FromUserID", "ToUserID", "Content" }, null, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::GoGoTalk.ProtoBuf.FriendChatNotice), global::GoGoTalk.ProtoBuf.FriendChatNotice.Parser, new[]{ "FromUserID", "ToUserID", "Content" }, null, null, null, null)
          }));
    }
    #endregion

  }
  #region Messages
  public sealed partial class FriendChatRequest : pb::IMessage<FriendChatRequest>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<FriendChatRequest> _parser = new pb::MessageParser<FriendChatRequest>(() => new FriendChatRequest());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<FriendChatRequest> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::GoGoTalk.ProtoBuf.FriendChatReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public FriendChatRequest() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public FriendChatRequest(FriendChatRequest other) : this() {
      fromUserID_ = other.fromUserID_;
      toUserID_ = other.toUserID_;
      content_ = other.content_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public FriendChatRequest Clone() {
      return new FriendChatRequest(this);
    }

    /// <summary>Field number for the "FromUserID" field.</summary>
    public const int FromUserIDFieldNumber = 1;
    private string fromUserID_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string FromUserID {
      get { return fromUserID_; }
      set {
        fromUserID_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "ToUserID" field.</summary>
    public const int ToUserIDFieldNumber = 2;
    private string toUserID_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string ToUserID {
      get { return toUserID_; }
      set {
        toUserID_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "Content" field.</summary>
    public const int ContentFieldNumber = 3;
    private pb::ByteString content_ = pb::ByteString.Empty;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public pb::ByteString Content {
      get { return content_; }
      set {
        content_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as FriendChatRequest);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(FriendChatRequest other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (FromUserID != other.FromUserID) return false;
      if (ToUserID != other.ToUserID) return false;
      if (Content != other.Content) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (FromUserID.Length != 0) hash ^= FromUserID.GetHashCode();
      if (ToUserID.Length != 0) hash ^= ToUserID.GetHashCode();
      if (Content.Length != 0) hash ^= Content.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      output.WriteRawMessage(this);
    #else
      if (FromUserID.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(FromUserID);
      }
      if (ToUserID.Length != 0) {
        output.WriteRawTag(18);
        output.WriteString(ToUserID);
      }
      if (Content.Length != 0) {
        output.WriteRawTag(26);
        output.WriteBytes(Content);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    void pb::IBufferMessage.InternalWriteTo(ref pb::WriteContext output) {
      if (FromUserID.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(FromUserID);
      }
      if (ToUserID.Length != 0) {
        output.WriteRawTag(18);
        output.WriteString(ToUserID);
      }
      if (Content.Length != 0) {
        output.WriteRawTag(26);
        output.WriteBytes(Content);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(ref output);
      }
    }
    #endif

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (FromUserID.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(FromUserID);
      }
      if (ToUserID.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(ToUserID);
      }
      if (Content.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeBytesSize(Content);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(FriendChatRequest other) {
      if (other == null) {
        return;
      }
      if (other.FromUserID.Length != 0) {
        FromUserID = other.FromUserID;
      }
      if (other.ToUserID.Length != 0) {
        ToUserID = other.ToUserID;
      }
      if (other.Content.Length != 0) {
        Content = other.Content;
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      input.ReadRawMessage(this);
    #else
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 10: {
            FromUserID = input.ReadString();
            break;
          }
          case 18: {
            ToUserID = input.ReadString();
            break;
          }
          case 26: {
            Content = input.ReadBytes();
            break;
          }
        }
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    void pb::IBufferMessage.InternalMergeFrom(ref pb::ParseContext input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, ref input);
            break;
          case 10: {
            FromUserID = input.ReadString();
            break;
          }
          case 18: {
            ToUserID = input.ReadString();
            break;
          }
          case 26: {
            Content = input.ReadBytes();
            break;
          }
        }
      }
    }
    #endif

  }

  public sealed partial class FriendChatNotice : pb::IMessage<FriendChatNotice>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<FriendChatNotice> _parser = new pb::MessageParser<FriendChatNotice>(() => new FriendChatNotice());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<FriendChatNotice> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::GoGoTalk.ProtoBuf.FriendChatReflection.Descriptor.MessageTypes[1]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public FriendChatNotice() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public FriendChatNotice(FriendChatNotice other) : this() {
      fromUserID_ = other.fromUserID_;
      toUserID_ = other.toUserID_;
      content_ = other.content_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public FriendChatNotice Clone() {
      return new FriendChatNotice(this);
    }

    /// <summary>Field number for the "FromUserID" field.</summary>
    public const int FromUserIDFieldNumber = 1;
    private string fromUserID_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string FromUserID {
      get { return fromUserID_; }
      set {
        fromUserID_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "ToUserID" field.</summary>
    public const int ToUserIDFieldNumber = 2;
    private string toUserID_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string ToUserID {
      get { return toUserID_; }
      set {
        toUserID_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "Content" field.</summary>
    public const int ContentFieldNumber = 3;
    private pb::ByteString content_ = pb::ByteString.Empty;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public pb::ByteString Content {
      get { return content_; }
      set {
        content_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as FriendChatNotice);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(FriendChatNotice other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (FromUserID != other.FromUserID) return false;
      if (ToUserID != other.ToUserID) return false;
      if (Content != other.Content) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (FromUserID.Length != 0) hash ^= FromUserID.GetHashCode();
      if (ToUserID.Length != 0) hash ^= ToUserID.GetHashCode();
      if (Content.Length != 0) hash ^= Content.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      output.WriteRawMessage(this);
    #else
      if (FromUserID.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(FromUserID);
      }
      if (ToUserID.Length != 0) {
        output.WriteRawTag(18);
        output.WriteString(ToUserID);
      }
      if (Content.Length != 0) {
        output.WriteRawTag(26);
        output.WriteBytes(Content);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    void pb::IBufferMessage.InternalWriteTo(ref pb::WriteContext output) {
      if (FromUserID.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(FromUserID);
      }
      if (ToUserID.Length != 0) {
        output.WriteRawTag(18);
        output.WriteString(ToUserID);
      }
      if (Content.Length != 0) {
        output.WriteRawTag(26);
        output.WriteBytes(Content);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(ref output);
      }
    }
    #endif

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (FromUserID.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(FromUserID);
      }
      if (ToUserID.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(ToUserID);
      }
      if (Content.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeBytesSize(Content);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(FriendChatNotice other) {
      if (other == null) {
        return;
      }
      if (other.FromUserID.Length != 0) {
        FromUserID = other.FromUserID;
      }
      if (other.ToUserID.Length != 0) {
        ToUserID = other.ToUserID;
      }
      if (other.Content.Length != 0) {
        Content = other.Content;
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      input.ReadRawMessage(this);
    #else
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 10: {
            FromUserID = input.ReadString();
            break;
          }
          case 18: {
            ToUserID = input.ReadString();
            break;
          }
          case 26: {
            Content = input.ReadBytes();
            break;
          }
        }
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    void pb::IBufferMessage.InternalMergeFrom(ref pb::ParseContext input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, ref input);
            break;
          case 10: {
            FromUserID = input.ReadString();
            break;
          }
          case 18: {
            ToUserID = input.ReadString();
            break;
          }
          case 26: {
            Content = input.ReadBytes();
            break;
          }
        }
      }
    }
    #endif

  }

  #endregion

}

#endregion Designer generated code
