# function
* return the ip of the requestor http://localhost:56567/ip
* activation http://localhost:56567/activate?token=aaa

# requirement 
new a config.yaml file
```
port: ":56567"
token: ["wangke","aaa"]
```


# update
1.0.1
* allow cors
* show version
* have the function of showing ip 

1.0.2
* replace nginx with this software( have the same redirect function and show the real ip) 
* 