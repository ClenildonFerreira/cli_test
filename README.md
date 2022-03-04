# Command Line
Small Command Line Application

application adds two command lines, the first one uses to get the public IP and the other to get the name of the server where the website is hosting.

External library: 
```golang
go get https://github.com/urfave/cli
```

Code: 
```golang
.\linha-de-comando ip --host url
.\linha-de-comando server --host url
```

examples:
```golang
.\linha-de-comando ip --host google.com
.\linha-de-comando server --host google.com
```
