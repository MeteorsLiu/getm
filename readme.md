# This is a hack tool for Go runtime

DON'T USE IT IN THE PRODUCTION ENVIRONMENT!

this will cause your programe unstable.

# Example(如何使用)

**请确保你的类型与结构体名是存在的，否则会panic**
**Please make sure the type of value is correct, or it will panic**

## Get a value in G struct.

```
CustomInG[uint64]("goid")
```

## Set/Modify a value in G struct.

```
SetCustomInG[uint64]("goid", 123456)
```

Check wthether it's affected or not


```
fmt.Println(CustomInG[uint64]("goid"))
```


## Get a value in M struct.
```
CustomInM[int64]("id")
```


## Set/Modify a value in M struct.
```
SetCustomInG[int64]("id", 123456)
```
Check wthether it's affected or not


```
CustomInM[int64]("id")
```

