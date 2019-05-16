// Code generated by thriftrw v1.19.1. DO NOT EDIT.
// @generated

package typedefs

import (
	enums "go.uber.org/thriftrw/gen/internal/tests/enums"
	structs "go.uber.org/thriftrw/gen/internal/tests/structs"
	thriftreflect "go.uber.org/thriftrw/thriftreflect"
)

// ThriftModule represents the IDL file used to generate this package.
var ThriftModule = &thriftreflect.ThriftModule{
	Name:     "typedefs",
	Package:  "go.uber.org/thriftrw/gen/internal/tests/typedefs",
	FilePath: "typedefs.thrift",
	SHA1:     "c44a85a79e17ab1c29e271bdef1909abe37fbb97",
	Includes: []*thriftreflect.ThriftModule{
		enums.ThriftModule,
		structs.ThriftModule,
	},
	Raw: rawIDL,
}

const rawIDL = "include \"./structs.thrift\"\ninclude \"./enums.thrift\"\n\n/**\n * Number of seconds since epoch.\n *\n * Deprecated: Use ISOTime instead.\n */\ntypedef i64 Timestamp  // alias of primitive\ntypedef string State\n\ntypedef i128 UUID  // alias of struct\n\ntypedef list<Event> EventGroup  // alias fo collection\n\nstruct i128 {\n    1: required i64 high\n    2: required i64 low\n}\n\nstruct Event {\n    1: required UUID uuid  // required typedef\n    2: optional Timestamp time  // optional typedef\n}\n\nstruct DefaultPrimitiveTypedef {\n    1: optional State state = \"hello\"\n}\n\nstruct Transition {\n    1: required State fromState\n    2: required State toState\n    3: optional EventGroup events\n}\n\ntypedef binary PDF  // alias of []byte\n\ntypedef set<structs.Frame> FrameGroup\n\ntypedef map<structs.Point, structs.Point> PointMap\n\ntypedef set<binary> BinarySet\n\ntypedef map<structs.Edge, structs.Edge> EdgeMap\n\ntypedef map<State, i64> StateMap\n\ntypedef enums.EnumWithValues MyEnum\n"
