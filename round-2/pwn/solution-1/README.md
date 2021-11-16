## Is Perfect Shit ?

Энэхүү бодлого нь анх харахад **heap / UAF** гэж харагдах бөгөөд эх кодыг сайн анзаарах юм бол **case 5:** дээр **gets()** ашигласан байна.

```sh
[~/]$ pwn checksec ./challenge
Arch:     amd64-64-little
RELRO:    Partial RELRO
Stack:    No canary found
NX:       NX enabled
PIE:      No PIE (0x400000)
```

Дээрх үр дүнгээс бид stack дотор код ажиллуулж чадахгүй ч stack -ийг удирдах боломжтой юм. Тиймээс бид ret2libc хийж shell авах болно. Энэ нь өмнө жилийн харуул зангийн финалын бодлоготой ижил тул доорх бичлэгээс үзнэ үү.

[![youtube](http://img.youtube.com/vi/LUSy4lxmFkA/0.jpg)](http://www.youtube.com/watch?v=LUSy4lxmFkA "Writeup")

Тэмцээний үеийн [скрипт](https://github.com/haruulzangi/2021/blob/main/round-2/pwn/solution-1/xpl.py)
