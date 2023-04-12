# function
* return the ip of the requestor http://xxx/ip
* return the ip info http://xx/ip/1.1.1.1
* activation http://xxx/activate?token=aaa
* redirect http://fldjsalfjds/other


# requirement 
new a config.yaml file
```
port: ":56567"
token: ["wangke","aaa"]
```


# update
23.04.12
add func return the ip info
move config into the init()
1.0.1
* allow cors
* show version
* have the function of showing ip 

1.0.3
* replace nginx with this software( have the same redirect function and show the real ip) 

1.0.4
* show the real ip in the X-Forwarded-For
  
1.0.5
* fixed the config file path logic.