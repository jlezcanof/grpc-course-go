# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: robot.proto

from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='robot.proto',
  package='robot',
  syntax='proto3',
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x0brobot.proto\x12\x05robot\"\x83\x01\n\x05Robot\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x0e\n\x06powers\x18\x03 \x03(\t\x12\x17\n\x0bspeed_level\x18\n \x01(\x05\x42\x02\x18\x01\x12\x1d\n\x06helper\x18\x0b \x01(\x0b\x32\r.robot.HelperJ\x04\x08\x05\x10\x06J\x04\x08\x06\x10\x07R\x05laserR\x05\x61gile\"\"\n\x06Helper\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t*)\n\x08SCENARIO\x12\x08\n\x04\x43ITY\x10\x00\x12\t\n\x05SPACE\x10\x01\x12\x08\n\x04MALL\x10\x02\x62\x06proto3'
)

_SCENARIO = _descriptor.EnumDescriptor(
  name='SCENARIO',
  full_name='robot.SCENARIO',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='CITY', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='SPACE', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='MALL', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=192,
  serialized_end=233,
)
_sym_db.RegisterEnumDescriptor(_SCENARIO)

SCENARIO = enum_type_wrapper.EnumTypeWrapper(_SCENARIO)
CITY = 0
SPACE = 1
MALL = 2



_ROBOT = _descriptor.Descriptor(
  name='Robot',
  full_name='robot.Robot',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='robot.Robot.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='name', full_name='robot.Robot.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='powers', full_name='robot.Robot.powers', index=2,
      number=3, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='speed_level', full_name='robot.Robot.speed_level', index=3,
      number=10, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\030\001', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='helper', full_name='robot.Robot.helper', index=4,
      number=11, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=23,
  serialized_end=154,
)


_HELPER = _descriptor.Descriptor(
  name='Helper',
  full_name='robot.Helper',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='robot.Helper.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='name', full_name='robot.Helper.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=156,
  serialized_end=190,
)

_ROBOT.fields_by_name['helper'].message_type = _HELPER
DESCRIPTOR.message_types_by_name['Robot'] = _ROBOT
DESCRIPTOR.message_types_by_name['Helper'] = _HELPER
DESCRIPTOR.enum_types_by_name['SCENARIO'] = _SCENARIO
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Robot = _reflection.GeneratedProtocolMessageType('Robot', (_message.Message,), {
  'DESCRIPTOR' : _ROBOT,
  '__module__' : 'robot_pb2'
  # @@protoc_insertion_point(class_scope:robot.Robot)
  })
_sym_db.RegisterMessage(Robot)

Helper = _reflection.GeneratedProtocolMessageType('Helper', (_message.Message,), {
  'DESCRIPTOR' : _HELPER,
  '__module__' : 'robot_pb2'
  # @@protoc_insertion_point(class_scope:robot.Helper)
  })
_sym_db.RegisterMessage(Helper)


_ROBOT.fields_by_name['speed_level']._options = None
# @@protoc_insertion_point(module_scope)
