Command:
```
go run main.go
```

Output:
```
2016/07/13 19:52:49 File Info Mode of os.Stdin :  Dcrw--w----
2016/07/13 19:52:49 os.ModeCharDevice          :   c---------
2016/07/13 19:52:49 Binary AND Result          :   c---------
Input from Terminal
```

---

Command:
```
echo "foobar" | go run main.go
```

Output:
```
2016/07/13 19:48:24 File Info Mode of os.Stdin :  prw-rw----
2016/07/13 19:48:24 os.ModeNamedPipe           :  p---------
2016/07/13 19:48:24 Binary AND Result          :  p---------
Input from Pipe
```