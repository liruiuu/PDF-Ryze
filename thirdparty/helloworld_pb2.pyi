from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class HelloRequest(_message.Message):
    __slots__ = ("name", "myStrings")
    NAME_FIELD_NUMBER: _ClassVar[int]
    MYSTRINGS_FIELD_NUMBER: _ClassVar[int]
    name: str
    myStrings: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, name: _Optional[str] = ..., myStrings: _Optional[_Iterable[str]] = ...) -> None: ...

class HelloReply(_message.Message):
    __slots__ = ("message",)
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    message: str
    def __init__(self, message: _Optional[str] = ...) -> None: ...
