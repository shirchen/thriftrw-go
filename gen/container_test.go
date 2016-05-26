// Copyright (c) 2015 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package gen

import (
	"testing"

	tc "github.com/thriftrw/thriftrw-go/gen/testdata/containers"
	te "github.com/thriftrw/thriftrw-go/gen/testdata/enums"
	ts "github.com/thriftrw/thriftrw-go/gen/testdata/structs"
	"github.com/thriftrw/thriftrw-go/wire"
)

func TestCollectionsOfPrimitives(t *testing.T) {
	tests := []struct {
		desc string
		p    tc.PrimitiveContainers
		v    wire.Value
	}{
		// Lists /////////////////////////////////////////////////////////////
		{
			"empty list",
			tc.PrimitiveContainers{ListOfInts: []int64{}},
			wire.NewValueStruct(wire.Struct{Fields: []wire.Field{{
				ID: 2,
				Value: wire.NewValueList(wire.List{
					ValueType: wire.TI64,
					Size:      0,
					Items:     wire.ValueListFromSlice([]wire.Value{}),
				}),
			}}}),
		},
		{
			"list of ints",
			tc.PrimitiveContainers{ListOfInts: []int64{1, 2, 3}},
			wire.NewValueStruct(wire.Struct{Fields: []wire.Field{{
				ID: 2,
				Value: wire.NewValueList(wire.List{
					ValueType: wire.TI64,
					Size:      3,
					Items: wire.ValueListFromSlice([]wire.Value{
						wire.NewValueI64(1),
						wire.NewValueI64(2),
						wire.NewValueI64(3),
					}),
				}),
			}}}),
		},
		{
			"list of binary",
			tc.PrimitiveContainers{
				ListOfBinary: [][]byte{
					[]byte("foo"), []byte("bar"), []byte("baz"),
				},
			},
			wire.NewValueStruct(wire.Struct{Fields: []wire.Field{{
				ID: 1,
				Value: wire.NewValueList(wire.List{
					ValueType: wire.TBinary,
					Size:      3,
					Items: wire.ValueListFromSlice([]wire.Value{
						wire.NewValueBinary([]byte("foo")),
						wire.NewValueBinary([]byte("bar")),
						wire.NewValueBinary([]byte("baz")),
					}),
				}),
			}}}),
		},
		// Sets //////////////////////////////////////////////////////////////
		{
			"empty set",
			tc.PrimitiveContainers{SetOfStrings: map[string]struct{}{}},
			wire.NewValueStruct(wire.Struct{Fields: []wire.Field{{
				ID: 3,
				Value: wire.NewValueSet(wire.Set{
					ValueType: wire.TBinary,
					Size:      0,
					Items:     wire.ValueListFromSlice([]wire.Value{}),
				}),
			}}}),
		},
		{
			"set of strings",
			tc.PrimitiveContainers{SetOfStrings: map[string]struct{}{
				"foo": struct{}{},
				"bar": struct{}{},
				"baz": struct{}{},
			}},
			wire.NewValueStruct(wire.Struct{Fields: []wire.Field{{
				ID: 3,
				Value: wire.NewValueSet(wire.Set{
					ValueType: wire.TBinary,
					Size:      3,
					Items: wire.ValueListFromSlice([]wire.Value{
						wire.NewValueString("foo"),
						wire.NewValueString("bar"),
						wire.NewValueString("baz"),
					}),
				}),
			}}}),
		},
		{
			"set of bytes",
			tc.PrimitiveContainers{SetOfBytes: map[int8]struct{}{
				-1:  struct{}{},
				1:   struct{}{},
				125: struct{}{},
			}},
			wire.NewValueStruct(wire.Struct{Fields: []wire.Field{{
				ID: 4,
				Value: wire.NewValueSet(wire.Set{
					ValueType: wire.TI8,
					Size:      3,
					Items: wire.ValueListFromSlice([]wire.Value{
						wire.NewValueI8(-1),
						wire.NewValueI8(1),
						wire.NewValueI8(125),
					}),
				}),
			}}}),
		},
		// Maps //////////////////////////////////////////////////////////////
		{
			"empty map",
			tc.PrimitiveContainers{MapOfStringToBool: map[string]bool{}},
			wire.NewValueStruct(wire.Struct{Fields: []wire.Field{{
				ID: 6,
				Value: wire.NewValueMap(wire.Map{
					KeyType:   wire.TBinary,
					ValueType: wire.TBool,
					Size:      0,
					Items:     wire.MapItemListFromSlice([]wire.MapItem{}),
				}),
			}}}),
		},
		{
			"map of int to string",
			tc.PrimitiveContainers{MapOfIntToString: map[int32]string{
				-1:    "foo",
				1234:  "bar",
				-9876: "baz",
			}},
			wire.NewValueStruct(wire.Struct{Fields: []wire.Field{{
				ID: 5,
				Value: wire.NewValueMap(wire.Map{
					KeyType:   wire.TI32,
					ValueType: wire.TBinary,
					Size:      3,
					Items: wire.MapItemListFromSlice([]wire.MapItem{
						{Key: wire.NewValueI32(-1), Value: wire.NewValueString("foo")},
						{Key: wire.NewValueI32(1234), Value: wire.NewValueString("bar")},
						{Key: wire.NewValueI32(-9876), Value: wire.NewValueString("baz")},
					}),
				}),
			}}}),
		},
		{
			"map of string to bool",
			tc.PrimitiveContainers{MapOfStringToBool: map[string]bool{
				"foo": true,
				"bar": false,
				"baz": true,
			}},
			wire.NewValueStruct(wire.Struct{Fields: []wire.Field{{
				ID: 6,
				Value: wire.NewValueMap(wire.Map{
					KeyType:   wire.TBinary,
					ValueType: wire.TBool,
					Size:      3,
					Items: wire.MapItemListFromSlice([]wire.MapItem{
						{Key: wire.NewValueString("foo"), Value: wire.NewValueBool(true)},
						{Key: wire.NewValueString("bar"), Value: wire.NewValueBool(false)},
						{Key: wire.NewValueString("baz"), Value: wire.NewValueBool(true)},
					}),
				}),
			}}}),
		},
	}

	for _, tt := range tests {
		assertRoundTrip(t, &tt.p, tt.v, tt.desc)
	}
}

func TestEnumContainers(t *testing.T) {
	tests := []struct {
		s tc.EnumContainers
		v wire.Value
	}{
		{
			tc.EnumContainers{
				ListOfEnums: []te.EnumDefault{
					te.EnumDefaultFoo,
					te.EnumDefaultBar,
				},
			},
			singleFieldStruct(1, wire.NewValueList(wire.List{
				ValueType: wire.TI32,
				Size:      2,
				Items: wire.ValueListFromSlice([]wire.Value{
					wire.NewValueI32(0),
					wire.NewValueI32(1),
				}),
			})),
		},
		{
			tc.EnumContainers{
				SetOfEnums: map[te.EnumWithValues]struct{}{
					te.EnumWithValuesX: struct{}{},
					te.EnumWithValuesZ: struct{}{},
				},
			},
			singleFieldStruct(2, wire.NewValueSet(wire.Set{
				ValueType: wire.TI32,
				Size:      2,
				Items: wire.ValueListFromSlice([]wire.Value{
					wire.NewValueI32(123),
					wire.NewValueI32(789),
				}),
			})),
		},
		{
			tc.EnumContainers{
				MapOfEnums: map[te.EnumWithDuplicateValues]int32{
					te.EnumWithDuplicateValuesP: 123,
					te.EnumWithDuplicateValuesQ: 456,
				},
			},
			singleFieldStruct(3, wire.NewValueMap(wire.Map{
				KeyType:   wire.TI32,
				ValueType: wire.TI32,
				Size:      2,
				Items: wire.MapItemListFromSlice([]wire.MapItem{
					wire.MapItem{Key: wire.NewValueI32(0), Value: wire.NewValueI32(123)},
					wire.MapItem{Key: wire.NewValueI32(-1), Value: wire.NewValueI32(456)},
				}),
			})),
		},
		{
			// this is the same as the one above except we're using "R" intsead
			// of "P" (they both have the same value)
			tc.EnumContainers{
				MapOfEnums: map[te.EnumWithDuplicateValues]int32{
					te.EnumWithDuplicateValuesR: 123,
					te.EnumWithDuplicateValuesQ: 456,
				},
			},
			singleFieldStruct(3, wire.NewValueMap(wire.Map{
				KeyType:   wire.TI32,
				ValueType: wire.TI32,
				Size:      2,
				Items: wire.MapItemListFromSlice([]wire.MapItem{
					wire.MapItem{Key: wire.NewValueI32(0), Value: wire.NewValueI32(123)},
					wire.MapItem{Key: wire.NewValueI32(-1), Value: wire.NewValueI32(456)},
				}),
			})),
		},
	}

	for _, tt := range tests {
		assertRoundTrip(t, &tt.s, tt.v, "EnumContainers")
	}
}

func TestListOfStructs(t *testing.T) {
	tests := []struct {
		s ts.Graph
		v wire.Value
	}{
		{
			ts.Graph{Edges: []*ts.Edge{}},
			singleFieldStruct(1, wire.NewValueList(wire.List{
				ValueType: wire.TStruct,
				Size:      0,
				Items:     wire.ValueListFromSlice(nil),
			})),
		},
		{
			ts.Graph{Edges: []*ts.Edge{
				{
					Start: &ts.Point{X: 1.0, Y: 2.0},
					End:   &ts.Point{X: 3.0, Y: 4.0},
				},
				{
					Start: &ts.Point{X: 5.0, Y: 6.0},
					End:   &ts.Point{X: 7.0, Y: 8.0},
				},
				{
					Start: &ts.Point{X: 9.0, Y: 10.0},
					End:   &ts.Point{X: 11.0, Y: 12.0},
				},
			}},
			singleFieldStruct(1, wire.NewValueList(wire.List{
				ValueType: wire.TStruct,
				Size:      3,
				Items: wire.ValueListFromSlice([]wire.Value{
					wire.NewValueStruct(wire.Struct{Fields: []wire.Field{
						{
							ID: 1,
							Value: wire.NewValueStruct(wire.Struct{Fields: []wire.Field{
								{ID: 1, Value: wire.NewValueDouble(1.0)},
								{ID: 2, Value: wire.NewValueDouble(2.0)},
							}}),
						},
						{
							ID: 2,
							Value: wire.NewValueStruct(wire.Struct{Fields: []wire.Field{
								{ID: 1, Value: wire.NewValueDouble(3.0)},
								{ID: 2, Value: wire.NewValueDouble(4.0)},
							}}),
						},
					}}),
					wire.NewValueStruct(wire.Struct{Fields: []wire.Field{
						{
							ID: 1,
							Value: wire.NewValueStruct(wire.Struct{Fields: []wire.Field{
								{ID: 1, Value: wire.NewValueDouble(5.0)},
								{ID: 2, Value: wire.NewValueDouble(6.0)},
							}}),
						},
						{
							ID: 2,
							Value: wire.NewValueStruct(wire.Struct{Fields: []wire.Field{
								{ID: 1, Value: wire.NewValueDouble(7.0)},
								{ID: 2, Value: wire.NewValueDouble(8.0)},
							}}),
						},
					}}),
					wire.NewValueStruct(wire.Struct{Fields: []wire.Field{
						{
							ID: 1,
							Value: wire.NewValueStruct(wire.Struct{Fields: []wire.Field{
								{ID: 1, Value: wire.NewValueDouble(9.0)},
								{ID: 2, Value: wire.NewValueDouble(10.0)},
							}}),
						},
						{
							ID: 2,
							Value: wire.NewValueStruct(wire.Struct{Fields: []wire.Field{
								{ID: 1, Value: wire.NewValueDouble(11.0)},
								{ID: 2, Value: wire.NewValueDouble(12.0)},
							}}),
						},
					}}),
				}),
			})),
		},
	}

	for _, tt := range tests {
		assertRoundTrip(t, &tt.s, tt.v, "Graph")
	}
}

func TestCrazyTown(t *testing.T) {
	tests := []struct {
		desc string
		x    tc.ContainersOfContainers
		v    wire.Value
	}{
		{
			"ListOfLists",
			tc.ContainersOfContainers{
				ListOfLists: [][]int32{
					{1, 2, 3},
					{4, 5, 6},
				},
			},
			wire.NewValueStruct(wire.Struct{Fields: []wire.Field{
				{ID: 1, Value: wire.NewValueList(wire.List{
					ValueType: wire.TList,
					Size:      2,
					Items: wire.ValueListFromSlice([]wire.Value{
						wire.NewValueList(wire.List{
							ValueType: wire.TI32,
							Size:      3,
							Items: wire.ValueListFromSlice([]wire.Value{
								wire.NewValueI32(1),
								wire.NewValueI32(2),
								wire.NewValueI32(3),
							}),
						}),
						wire.NewValueList(wire.List{
							ValueType: wire.TI32,
							Size:      3,
							Items: wire.ValueListFromSlice([]wire.Value{
								wire.NewValueI32(4),
								wire.NewValueI32(5),
								wire.NewValueI32(6),
							}),
						}),
					}),
				})},
			}}),
		},
		{
			"ListOfSets",
			tc.ContainersOfContainers{
				ListOfSets: []map[int32]struct{}{
					{
						1: struct{}{},
						2: struct{}{},
						3: struct{}{},
					},
					{
						4: struct{}{},
						5: struct{}{},
						6: struct{}{},
					},
				},
			},
			wire.NewValueStruct(wire.Struct{Fields: []wire.Field{
				{ID: 2, Value: wire.NewValueList(wire.List{
					ValueType: wire.TSet,
					Size:      2,
					Items: wire.ValueListFromSlice([]wire.Value{
						wire.NewValueSet(wire.Set{
							ValueType: wire.TI32,
							Size:      3,
							Items: wire.ValueListFromSlice([]wire.Value{
								wire.NewValueI32(1),
								wire.NewValueI32(2),
								wire.NewValueI32(3),
							}),
						}),
						wire.NewValueSet(wire.Set{
							ValueType: wire.TI32,
							Size:      3,
							Items: wire.ValueListFromSlice([]wire.Value{
								wire.NewValueI32(4),
								wire.NewValueI32(5),
								wire.NewValueI32(6),
							}),
						}),
					}),
				})},
			}}),
		},
		{
			"ListOfMaps",
			tc.ContainersOfContainers{
				ListOfMaps: []map[int32]int32{
					{
						1: 100,
						2: 200,
						3: 300,
					},
					{
						4: 400,
						5: 500,
						6: 600,
					},
				},
			},
			wire.NewValueStruct(wire.Struct{Fields: []wire.Field{
				{ID: 3, Value: wire.NewValueList(wire.List{
					ValueType: wire.TMap,
					Size:      2,
					Items: wire.ValueListFromSlice([]wire.Value{
						wire.NewValueMap(wire.Map{
							KeyType:   wire.TI32,
							ValueType: wire.TI32,
							Size:      3,
							Items: wire.MapItemListFromSlice([]wire.MapItem{
								{Key: wire.NewValueI32(1), Value: wire.NewValueI32(100)},
								{Key: wire.NewValueI32(2), Value: wire.NewValueI32(200)},
								{Key: wire.NewValueI32(3), Value: wire.NewValueI32(300)},
							}),
						}),
						wire.NewValueMap(wire.Map{
							KeyType:   wire.TI32,
							ValueType: wire.TI32,
							Size:      3,
							Items: wire.MapItemListFromSlice([]wire.MapItem{
								{Key: wire.NewValueI32(4), Value: wire.NewValueI32(400)},
								{Key: wire.NewValueI32(5), Value: wire.NewValueI32(500)},
								{Key: wire.NewValueI32(6), Value: wire.NewValueI32(600)},
							}),
						}),
					}),
				})},
			}}),
		},
		{
			"SetOfSets",
			tc.ContainersOfContainers{
				SetOfSets: []map[string]struct{}{
					{
						"1": struct{}{},
						"2": struct{}{},
						"3": struct{}{},
					},
					{
						"4": struct{}{},
						"5": struct{}{},
						"6": struct{}{},
					},
				},
			},
			wire.NewValueStruct(wire.Struct{Fields: []wire.Field{
				{ID: 4, Value: wire.NewValueSet(wire.Set{
					ValueType: wire.TSet,
					Size:      2,
					Items: wire.ValueListFromSlice([]wire.Value{
						wire.NewValueSet(wire.Set{
							ValueType: wire.TBinary,
							Size:      3,
							Items: wire.ValueListFromSlice([]wire.Value{
								wire.NewValueString("1"),
								wire.NewValueString("2"),
								wire.NewValueString("3"),
							}),
						}),
						wire.NewValueSet(wire.Set{
							ValueType: wire.TBinary,
							Size:      3,
							Items: wire.ValueListFromSlice([]wire.Value{
								wire.NewValueString("4"),
								wire.NewValueString("5"),
								wire.NewValueString("6"),
							}),
						}),
					}),
				})},
			}}),
		},
		{
			"SetOfLists",
			tc.ContainersOfContainers{
				SetOfLists: [][]string{
					{"1", "2", "3"},
					{"4", "5", "6"},
				},
			},
			wire.NewValueStruct(wire.Struct{Fields: []wire.Field{
				{ID: 5, Value: wire.NewValueSet(wire.Set{
					ValueType: wire.TList,
					Size:      2,
					Items: wire.ValueListFromSlice([]wire.Value{
						wire.NewValueList(wire.List{
							ValueType: wire.TBinary,
							Size:      3,
							Items: wire.ValueListFromSlice([]wire.Value{
								wire.NewValueString("1"),
								wire.NewValueString("2"),
								wire.NewValueString("3"),
							}),
						}),
						wire.NewValueList(wire.List{
							ValueType: wire.TBinary,
							Size:      3,
							Items: wire.ValueListFromSlice([]wire.Value{
								wire.NewValueString("4"),
								wire.NewValueString("5"),
								wire.NewValueString("6"),
							}),
						}),
					}),
				})},
			}}),
		},
		{
			"SetOfMaps",
			tc.ContainersOfContainers{
				SetOfMaps: []map[string]string{
					{
						"1": "one",
						"2": "two",
						"3": "three",
					},
					{
						"4": "four",
						"5": "five",
						"6": "six",
					},
				},
			},
			wire.NewValueStruct(wire.Struct{Fields: []wire.Field{
				{ID: 6, Value: wire.NewValueSet(wire.Set{
					ValueType: wire.TMap,
					Size:      2,
					Items: wire.ValueListFromSlice([]wire.Value{
						wire.NewValueMap(wire.Map{
							KeyType:   wire.TBinary,
							ValueType: wire.TBinary,
							Size:      3,
							Items: wire.MapItemListFromSlice([]wire.MapItem{
								{Key: wire.NewValueString("1"), Value: wire.NewValueString("one")},
								{Key: wire.NewValueString("2"), Value: wire.NewValueString("two")},
								{Key: wire.NewValueString("3"), Value: wire.NewValueString("three")},
							}),
						}),
						wire.NewValueMap(wire.Map{
							KeyType:   wire.TBinary,
							ValueType: wire.TBinary,
							Size:      3,
							Items: wire.MapItemListFromSlice([]wire.MapItem{
								{Key: wire.NewValueString("4"), Value: wire.NewValueString("four")},
								{Key: wire.NewValueString("5"), Value: wire.NewValueString("five")},
								{Key: wire.NewValueString("6"), Value: wire.NewValueString("six")},
							}),
						}),
					}),
				})},
			}}),
		},
		{
			"MapOfMapToInt",
			tc.ContainersOfContainers{
				MapOfMapToInt: []struct {
					Key   map[string]int32
					Value int64
				}{
					{
						Key:   map[string]int32{"1": 1, "2": 2, "3": 3},
						Value: 123,
					},
					{
						Key:   map[string]int32{"4": 4, "5": 5, "6": 6},
						Value: 456,
					},
				},
			},
			wire.NewValueStruct(wire.Struct{Fields: []wire.Field{
				{ID: 7, Value: wire.NewValueMap(wire.Map{
					KeyType:   wire.TMap,
					ValueType: wire.TI64,
					Size:      2,
					Items: wire.MapItemListFromSlice([]wire.MapItem{
						{
							Key: wire.NewValueMap(wire.Map{
								KeyType:   wire.TBinary,
								ValueType: wire.TI32,
								Size:      3,
								Items: wire.MapItemListFromSlice([]wire.MapItem{
									{Key: wire.NewValueString("1"), Value: wire.NewValueI32(1)},
									{Key: wire.NewValueString("2"), Value: wire.NewValueI32(2)},
									{Key: wire.NewValueString("3"), Value: wire.NewValueI32(3)},
								}),
							}),
							Value: wire.NewValueI64(123),
						},
						{
							Key: wire.NewValueMap(wire.Map{
								KeyType:   wire.TBinary,
								ValueType: wire.TI32,
								Size:      3,
								Items: wire.MapItemListFromSlice([]wire.MapItem{
									{Key: wire.NewValueString("4"), Value: wire.NewValueI32(4)},
									{Key: wire.NewValueString("5"), Value: wire.NewValueI32(5)},
									{Key: wire.NewValueString("6"), Value: wire.NewValueI32(6)},
								}),
							}),
							Value: wire.NewValueI64(456),
						},
					}),
				})},
			}}),
		},
		{
			"MapOfListToSet",
			tc.ContainersOfContainers{
				MapOfListToSet: []struct {
					Key   []int32
					Value map[int64]struct{}
				}{
					{
						Key: []int32{1, 2, 3},
						Value: map[int64]struct{}{
							1: struct{}{},
							2: struct{}{},
							3: struct{}{},
						},
					},
					{
						Key: []int32{4, 5, 6},
						Value: map[int64]struct{}{
							4: struct{}{},
							5: struct{}{},
							6: struct{}{},
						},
					},
				},
			},
			wire.NewValueStruct(wire.Struct{Fields: []wire.Field{
				{ID: 8, Value: wire.NewValueMap(wire.Map{
					KeyType:   wire.TList,
					ValueType: wire.TSet,
					Size:      2,
					Items: wire.MapItemListFromSlice([]wire.MapItem{
						{
							Key: wire.NewValueList(wire.List{
								ValueType: wire.TI32,
								Size:      3,
								Items: wire.ValueListFromSlice([]wire.Value{
									wire.NewValueI32(1),
									wire.NewValueI32(2),
									wire.NewValueI32(3),
								}),
							}),
							Value: wire.NewValueSet(wire.Set{
								ValueType: wire.TI64,
								Size:      3,
								Items: wire.ValueListFromSlice([]wire.Value{
									wire.NewValueI64(1),
									wire.NewValueI64(2),
									wire.NewValueI64(3),
								}),
							}),
						},
						{
							Key: wire.NewValueList(wire.List{
								ValueType: wire.TI32,
								Size:      3,
								Items: wire.ValueListFromSlice([]wire.Value{
									wire.NewValueI32(4),
									wire.NewValueI32(5),
									wire.NewValueI32(6),
								}),
							}),
							Value: wire.NewValueSet(wire.Set{
								ValueType: wire.TI64,
								Size:      3,
								Items: wire.ValueListFromSlice([]wire.Value{
									wire.NewValueI64(4),
									wire.NewValueI64(5),
									wire.NewValueI64(6),
								}),
							}),
						},
					}),
				})},
			}}),
		},
		{
			"MapOfSetToListOfDouble",
			tc.ContainersOfContainers{
				MapOfSetToListOfDouble: []struct {
					Key   map[int32]struct{}
					Value []float64
				}{
					{
						Key: map[int32]struct{}{
							1: struct{}{},
							2: struct{}{},
							3: struct{}{},
						},
						Value: []float64{1.0, 2.0, 3.0},
					},
					{
						Key: map[int32]struct{}{
							4: struct{}{},
							5: struct{}{},
							6: struct{}{},
						},
						Value: []float64{4.0, 5.0, 6.0},
					},
				},
			},
			wire.NewValueStruct(wire.Struct{Fields: []wire.Field{
				{ID: 9, Value: wire.NewValueMap(wire.Map{
					KeyType:   wire.TSet,
					ValueType: wire.TList,
					Size:      2,
					Items: wire.MapItemListFromSlice([]wire.MapItem{
						{
							Key: wire.NewValueSet(wire.Set{
								ValueType: wire.TI32,
								Size:      3,
								Items: wire.ValueListFromSlice([]wire.Value{
									wire.NewValueI32(1),
									wire.NewValueI32(2),
									wire.NewValueI32(3),
								}),
							}),
							Value: wire.NewValueList(wire.List{
								ValueType: wire.TDouble,
								Size:      3,
								Items: wire.ValueListFromSlice([]wire.Value{
									wire.NewValueDouble(1.0),
									wire.NewValueDouble(2.0),
									wire.NewValueDouble(3.0),
								}),
							}),
						},
						{
							Key: wire.NewValueSet(wire.Set{
								ValueType: wire.TI32,
								Size:      3,
								Items: wire.ValueListFromSlice([]wire.Value{
									wire.NewValueI32(4),
									wire.NewValueI32(5),
									wire.NewValueI32(6),
								}),
							}),
							Value: wire.NewValueList(wire.List{
								ValueType: wire.TDouble,
								Size:      3,
								Items: wire.ValueListFromSlice([]wire.Value{
									wire.NewValueDouble(4.0),
									wire.NewValueDouble(5.0),
									wire.NewValueDouble(6.0),
								}),
							}),
						},
					}),
				})},
			}}),
		},
	}

	for _, tt := range tests {
		assertRoundTrip(t, &tt.x, tt.v, tt.desc)
	}
}