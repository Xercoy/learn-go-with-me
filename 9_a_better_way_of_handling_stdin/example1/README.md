Command:
```
go run main.go
```

Output:
```
2016/07/13 20:07:57 os.Stdin File Mode  : Dcrw--w----
```

---

Command:
```
echo "foobar" | go run main.go
```

Output:
```
2016/07/13 20:08:24 os.Stdin File Mode  : prw-rw----
```