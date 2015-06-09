# Benchmark set for Structure To Map converters

Benchmarks simulate common case of serialization: medium sized object that has to be stored into database or returned by some REST API.

The structure (with some modification for particular tools) is following:
```
type SomeItem struct {
    ID              int             `db:"id" custom_tag:"id"`
    Name            string          `db:"name"`
    somePrivate     string          `db:"some_private" custom_tag:"some_private"`
    Number          int             `db:"number" custom_tag:"num"`
    Checksum        int32           `custom_tag:"sum"`
    Created         time.Time       `custom_tag:"created_time" db:"created"`
    Updated         mysql.NullTime  `db:"updated" custom_tag:"updated_time"`
    Price           float64         `db:"price"`
    Discount        *float64        `db:"discount"`
    IsReserved      sql.NullBool    `db:"reserved" custom_tag:"is_reserved"`
    Points          sql.NullInt64   `db:"points"`
    Rating          sql.NullFloat64 `db:"rating"`
    IsVisible       bool            `db:"visible" custom_tag:"visible"`
    SomeIgnoreField int             `db:"-" custom_tag:"i_ignore_nothing"`
    Notes           string
}
```

All the benchmarks were run on 2.6 GHz i5 Macbook Pro with following command: ```go test -bench=. -benchmem | column -t ```

Tools participating in benchmark:
- Manually written (or generated) method `ToMap()` that manually converts structure's fields into map's fields
- stom: https://github.com/elgris/stom
  -  general mode of stom
  -  "individual mode" of stom where structure's metadata is cached before doing multiple conversions
- structs: https://github.com/fatih/structs

# Feature list

| feature                                      | stom | structs |
|----------------------------------------------|------|---------|
| Get structure metadata (fields, names, etc)  |      | +       |
| Cache structure metadata for fast conversion | +    |         |
| Replace nil values with given default        | +    |         |
| Support for custom conversion logic          | +    |         |

# Participants

## Method ToMap written by hand or generated
Check [ToMap implementation](https://github.com/elgris/struct-to-map-conversion-benchmark/blob/master/byhand_test.go#L29) for details.

## SToM
Good thing about SToM - it's built to be fast. That means "use as little reflection as possible". Also it allows implement custom converters for each type: you just need to implement interface `ToMappable` for it. There are benchmarks for 2 modes of working with stom.

### SToM general
There are 2 benchmarks:
1. `BenchmarkStomGeneral` - when `stom.ConvertToMap` is given a copy of the object
2. `BenchmarkStomGeneralPtr` - when `stom.ConvertToMap` is given a reference (pointer) to the object

### SToM individual
There are 2 benchmarks:
1. `BenchmarkStomIndividual` - new stom instance is initialized with empty instance of the structure and is given copies of the object during benchmark
2. `BenchmarkStomGeneralPtr` - new stom instance is initialized with empty instance of the structure and is given the same reference to the object during benchmark

## github.com/fatih/structs
Benchmark use a slightly different struct type
```
type ItemStructs struct {
    ID          int     `db:"id" custom_tag:"id"`
    Name        string  `db:"name"`
    somePrivate string  `db:"some_private" custom_tag:"some_private"`
    Number      int     `db:"number" custom_tag:"num"`
    Created     string  `custom_tag:"created_time" db:"created"`
    Updated     string  `db:"updated" custom_tag:"updated_time"`
    Price       float64 `db:"price"`
    Discount    float64 `db:"discount"`
    IsReserved  bool    `db:"reserved" custom_tag:"is_reserved"`
    Points      int64   `db:"points"`
    Rating      float64 `db:"rating"`
    IsVisible   bool    `db:"visible" custom_tag:"visible"`
}
```

I had to remove pointers because **structs** leaves them as is while I'm expecting for actual number (point dereference). But on the other hands complex types (like `sql.NullFloat64`) are converted into another `map[string]interface{}` while I'm expecting leaving it as is.

This tool is more about getting struct's metadata like Fields, Names, etc. Think about it like a high-level wrapper around `reflect`. But still it's able to convert a structure to a map, so it participates in the benchmark.
There are 2 benchmarks:
1. `BenchmarkStructs` - when `structs.Map()` is given a copy of the object
2. `BenchmarkStructsPtr` - when `structs.Map()` is given a reference (pointer) to the object

# Benchmark results

```
BenchmarkByHand             1000000    2187    ns/op  1122  B/op  13  allocs/op

BenchmarkStomGeneral        200000     7016    ns/op  2144  B/op  26  allocs/op
BenchmarkStomGeneralPtr     200000     7620    ns/op  2128  B/op  35  allocs/op

BenchmarkStomIndividual     500000     3464    ns/op  1218  B/op  8   allocs/op
BenchmarkStomIndividualPtr  500000     3843    ns/op  1202  B/op  17  allocs/op

BenchmarkStructs            200000     9943    ns/op  4530  B/op  43  allocs/op
BenchmarkStructsPtr         200000     11555   ns/op  4754  B/op  64  allocs/op
```

# Conclusion

**Manual or generated conversion logic** is the fastest, so if you're fine with tones of generated code (which can bloat in case of large struct types), go for it.

If you don't trust generated code (or feel that code generation can be unsafe), but still you need to convert structs of the same time many times during application lifecycle, you can speed things up with **(stom)[https://github.com/elgris/stom]** ("individual" mode).

If you need a handy tool to work with struct type metadata (in case you don't like core package `reflect` and looking for abstraction of higher level) which is also capable for converting a struct to map - use **(structs)[https://github.com/fatih/structs]**