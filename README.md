# HGOTOOL

HGOTOOL is CLI tools build Golang for monitoring server condition
It monitor server Processes,files

### HGOTOOL = Tools for using Hardening Project build by Golang 



## How to Use

**darwin,linux(Ubuntu,CentOS) test is success**


### process

- ps
 
```
$ hgotool ps show
```

- monitor 

```
$ hgotool ps show | grep NAME
$ hgotool ps monitor PID


```


### File

- file Detail

```
$ hgotool file detail ~/PATH/FILE
```

- monitor 

```
$hgotool file monitor /PATH/NAME
```


### PortScan

```
$ hgotool port scan
```

