# Go-Misc

模仿[lua-misc](https://github.com/trumanzhao/misc)中`log_tree`函数
go中的`log_tree`函数
```
[Display] tempData (log.TestStruct):
└─  [Struct: TestStruct]
   ├─ Id = 12 (int)
   ├─ Name = "tempData" (string)
   ├─ FloatMap  [Map <string, ?>]
   │  ├─ "test1" [Array: float64]
   │  │  ├─ 0 = 12.3 (float64)
   │  │  └─ 1 = 33 (float64)
   │  └─ "test2" [Array: float64]
   │     ├─ 0 = 88.456 (float64)
   │     └─ 1 = 78.12 (float64)
   ├─ StringMap  [Map <int, StringStruct>]
   │  ├─ 1 [Struct: StringStruct]
   │  │  └─ Str = "hello" (string)
   │  └─ 2 [Struct: StringStruct]
   │     └─ Str = "world" (string)
   ├─ EmptyStruct = {} [Struct: EmptyStruct]
   ├─ [Ptr *string] (*Address) = "tempString" (string)
   └─ [Interface] Reserver  [Map <?, ?>]
      ├─ [MapKey]interface {} value = "Reserver" (string)
      │  └─ (*[MapValue]interface {} value) = "tempString" (string)
      ├─ [MapKey]interface {} value = "Func" (string)
      ├─ [MapValue]interface {} value = func() 0x1d5680 (func)
      ├─ [MapKey]interface {} value = "Nil" (string)
      └─ [MapValue]interface {} value = nil
```