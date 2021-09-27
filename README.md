php serialize to json (Go)
```
Array
(
    [0] => Array
        (
            [id] => 1
            [price] => 80
            [name] => 一个名字
        )

)
```

phpserialize string
```
a:1:{i:0;a:3:{s:2:"id";i:1;s:5:"price";i:80;s:4:"name";s:12:"一个名字";}}
```

to json
```
[{"Id":1,"Price":80,"Name":"一个名字"}]
```

run go test