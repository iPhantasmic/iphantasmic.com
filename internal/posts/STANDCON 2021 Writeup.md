---
title: "STANDCON 2021"
slug: "standcon-2021"
description: "Writeup for STANDCON 2021"
cover: "/static/content/standcon-2021/cover.png"
featured: false
published: "2021-07-24"
tags: ["CTF", "writeup"]
---

# STANDCON 2021 Writeup

![](/static/content/standcon-2021/results.png)

---

# Forensics

## Ancient Computing

Points: 1000

Challenge File:

[BOOK1.WK1](/static/content/standcon-2021/BOOK1.wk1)

[Read/Convert old Lotus WK1 and WK4 files](https://answers.microsoft.com/en-us/msoffice/forum/all/readconvert-old-lotus-wk1-and-wk4-files/e61de246-533b-4747-b17a-b1a3ba66b120?auth=1)

[Gnumeric](http://www.gnumeric.org/download.html)

```bash
kali@kali:~/Desktop/STANDCON$ sudo apt-get install gnumeric
```

![](/static/content/standcon-2021/Untitled.png)

Flag:

```bash
# Round 0.2524130133933795 up to 14th decimal place

STC{0.25241301339338}
```

## Astronaut Document

Points: 1000

Challenge File:

[astronaut_document.txt](/static/content/standcon-2021/astronaut_document.txt)

![](/static/content/standcon-2021/Untitled%201.png)

- base64 String
    
    ```bash
    $e = '';
    $e .= 'R';
    $e .= 'X';
    $e .= 'J';
    $e .= 'y';
    $e .= 'b';
    $e .= '3';
    $e .= 'I';
    $e .= 'g';
    $e .= 'Y';
    $e .= 'X';
    $e .= 'Q';
    $e .= 'g';
    $e .= 'b';
    $e .= '2';
    $e .= 'J';
    $e .= 'q';
    $e .= 'Z';
    $e .= 'W';
    $e .= 'N';
    $e .= '0';
    $e .= 'c';
    $e .= 'y';
    $e .= 'A';
    $e .= '3';
    $e .= 'L';
    $e .= 'D';
    $e .= 'k';
    $e .= 's';
    $e .= 'M';
    $e .= 'T';
    $e .= 'A';
    $e .= 'g';
    $e .= 'Y';
    $e .= 'W';
    $e .= '5';
    $e .= 'k';
    $e .= 'I';
    $e .= 'D';
    $e .= 'E';
    $e .= 'z';
    $e .= 'I';
    $e .= 'G';
    $e .= 'l';
    $e .= 'u';
    $e .= 'I';
    $e .= 'H';
    $e .= 'R';
    $e .= 'o';
    $e .= 'Z';
    $e .= 'S';
    $e .= 'B';
    $e .= 'k';
    $e .= 'b';
    $e .= '2';
    $e .= 'N';
    $e .= '1';
    $e .= 'b';
    $e .= 'W';
    $e .= 'V';
    $e .= 'u';
    $e .= 'd';
    $e .= 'A';
    $e .= 'p';
    $e .= 'Q';
    $e .= 'b';
    $e .= 'G';
    $e .= 'V';
    $e .= 'h';
    $e .= 'c';
    $e .= '2';
    $e .= 'U';
    $e .= 'g';
    $e .= 'Z';
    $e .= 'm';
    $e .= 'l';
    $e .= '4';
    $e .= 'I';
    $e .= 'H';
    $e .= 'R';
    $e .= 'o';
    $e .= 'Z';
    $e .= 'S';
    $e .= 'B';
    $e .= 'l';
    $e .= 'c';
    $e .= 'n';
    $e .= 'J';
    $e .= 'v';
    $e .= 'c';
    $e .= 'n';
    $e .= 'M';
    $e .= 's';
    $e .= 'I';
    $e .= 'H';
    $e .= 'R';
    $e .= 'o';
    $e .= 'Z';
    $e .= 'S';
    $e .= 'B';
    $e .= 'm';
    $e .= 'b';
    $e .= '2';
    $e .= 'x';
    $e .= 's';
    $e .= 'b';
    $e .= '3';
    $e .= 'd';
    $e .= 'p';
    $e .= 'b';
    $e .= 'm';
    $e .= 'c';
    $e .= 'g';
    $e .= 'b';
    $e .= 'G';
    $e .= 'l';
    $e .= 'u';
    $e .= 'a';
    $e .= 'y';
    $e .= 'B';
    $e .= 't';
    $e .= 'Y';
    $e .= 'X';
    $e .= 'k';
    $e .= 'g';
    $e .= 'a';
    $e .= 'G';
    $e .= 'V';
    $e .= 's';
    $e .= 'c';
    $e .= 'D';
    $e .= 'o';
    $e .= 'K';
    $e .= 'a';
    $e .= 'H';
    $e .= 'R';
    $e .= '0';
    $e .= 'c';
    $e .= 'H';
    $e .= 'M';
    $e .= '6';
    $e .= 'L';
    $e .= 'y';
    $e .= '9';
    $e .= 'n';
    $e .= 'Z';
    $e .= 'W';
    $e .= '5';
    $e .= 'k';
    $e .= 'a';
    $e .= 'W';
    $e .= 'd';
    $e .= 'u';
    $e .= 'b';
    $e .= '3';
    $e .= 'V';
    $e .= '4';
    $e .= 'L';
    $e .= 'm';
    $e .= 'N';
    $e .= 'v';
    $e .= 'b';
    $e .= 'S';
    $e .= '9';
    $e .= 'i';
    $e .= 'b';
    $e .= 'G';
    $e .= '9';
    $e .= 'n';
    $e .= 'L';
    $e .= 'z';
    $e .= 'I';
    $e .= 'w';
    $e .= 'M';
    $e .= 'T';
    $e .= 'Y';
    $e .= 'v';
    $e .= 'M';
    $e .= 'T';
    $e .= 'A';
    $e .= 'v';
    $e .= 'M';
    $e .= 'D';
    $e .= 'Q';
    $e .= 'v';
    $e .= 'c';
    $e .= 'G';
    $e .= 'R';
    $e .= 'm';
    $e .= 'L';
    $e .= 'W';
    $e .= 'J';
    $e .= 'h';
    $e .= 'c';
    $e .= '2';
    $e .= 'l';
    $e .= 'j';
    $e .= 'c';
    $e .= 'y';
    $e .= '5';
    $e .= 'o';
    $e .= 'd';
    $e .= 'G';
    $e .= '1';
    $e .= 's';
    print(decode_base64($e));
    ```
    

```bash
RXJyb3IgYXQgb2JqZWN0cyA3LDksMTAgYW5kIDEzIGluIHRoZSBkb2N1bWVudApQbGVhc2UgZml4IHRoZSBlcnJvcnMsIHRoZSBmb2xsb3dpbmcgbGluayBtYXkgaGVscDoKaHR0cHM6Ly9nZW5kaWdub3V4LmNvbS9ibG9nLzIwMTYvMTAvMDQvcGRmLWJhc2ljcy5odG1s
```

![](/static/content/standcon-2021/Untitled%202.png)

[Introduction to PDF syntax](https://gendignoux.com/blog/2016/10/04/pdf-basics.html)

![](/static/content/standcon-2021/Untitled%203.png)

![](/static/content/standcon-2021/Untitled%204.png)

![](/static/content/standcon-2021/Untitled%205.png)

![](/static/content/standcon-2021/Untitled%206.png)

![](/static/content/standcon-2021/Untitled%207.png)

![](/static/content/standcon-2021/Untitled%208.png)

![](/static/content/standcon-2021/Untitled%209.png)

![](/static/content/standcon-2021/Untitled%2010.png)

![](/static/content/standcon-2021/Untitled%2011.png)

![](/static/content/standcon-2021/Untitled%2012.png)

![](/static/content/standcon-2021/Untitled%2013.png)

[astronaut_document.pdf](/static/content/standcon-2021/astronaut_document.pdf)

---

# Crypto

## Space Noise

Points: 1000

Challenge File:

[space_noise.pcap](/static/content/standcon-2021/space_noise.pcap)

![Notice how SYN and PSH act as delimiters while we have URG and RST acting as binary 1 and 0](/static/content/standcon-2021/Untitled%2014.png)

Notice how SYN and PSH act as delimiters while we have URG and RST acting as binary 1 and 0

![](/static/content/standcon-2021/Untitled%2015.png)

```python
import json

MORSE_CODE_DICT = { 'A':'.-', 'B':'-...',
                    'C':'-.-.', 'D':'-..', 'E':'.',
                    'F':'..-.', 'G':'--.', 'H':'....',
                    'I':'..', 'J':'.---', 'K':'-.-',
                    'L':'.-..', 'M':'--', 'N':'-.',
                    'O':'---', 'P':'.--.', 'Q':'--.-',
                    'R':'.-.', 'S':'...', 'T':'-',
                    'U':'..-', 'V':'...-', 'W':'.--',
                    'X':'-..-', 'Y':'-.--', 'Z':'--..',
                    '1':'.----', '2':'..---', '3':'...--',
                    '4':'....-', '5':'.....', '6':'-....',
                    '7':'--...', '8':'---..', '9':'----.',
                    '0':'-----', ', ':'--..--', '.':'.-.-.-',
                    '?':'..--..', '/':'-..-.', '-':'-....-',
                    '(':'-.--.', ')':'-.--.-'}

data = json.loads(open('space_noise.json').read())
result = ''

curr = ''
for packet in data:
    flags = packet['_source']['layers']['tcp']['tcp.flags_tree']

    if flags['tcp.flags.reset'] == '1':
        curr += 'R'

    elif flags['tcp.flags.push'] == '1':

        if not curr:
            continue

        morse_code = ''
        for char in curr:
            if char == 'R':
                morse_code += '.'
            else:
                morse_code += '-'

        print(curr, morse_code)

        for key in MORSE_CODE_DICT:
            if MORSE_CODE_DICT[key] == morse_code:
                result += key

        curr = ''

    elif flags['tcp.flags.urg'] == '1':
        curr += 'U'

print(result)
```

```bash
python3 solve.py
...
5354437B492062656C6965766520746861742074686973204E6174696F6E2073686F756C6420636F6D6D697420697473656C6620746F20616368696576696E672074686520676F616C2C206265666F7265207468697320646563616465206973206F75742C206F66206C616E64696E672061206D616E206F6E20746865204D6F6F6E20616E642072657475726E696E672068696D20736166656C7920746F2045617274682E7D
```

![](/static/content/standcon-2021/Untitled%2016.png)

```python
STC{I believe that this Nation should commit itself to achieving the goal, before this decade is out, of landing a man on the Moon and returning him safely to Earth.}
```

## Substitution Extreme

Points: 1000

Challenge File:

[transmission.txt](/static/content/standcon-2021/transmission.txt)

[](https://www.quipqiup.com/)

```
--- BEGIN INTERCEPTED TRANSMISSION ---

EZWUFROD HOPVGWUEG GIGRGB HZFUGB PZVGEG WURGU AGPV MZERZMGD IO HZRGMGP HZJZPGPQUPV JZRGAU IO GHOG MZPVVGEG. HOPVGWUEG MZEWOHGB IGEO JGRGAHOG IZPVGP HZRGM QXBXE IO UMGEG, IGP DZWURGUGP EOGU OPIXPZHOG IZPVGP HZRGM HOPVGWUEG IO HZRGMGP. HOPVGWUEG GIGRGB HGRGB HGMU WUHGM DZKGPVGP MZEUPVVUR, IGP IODZPGRO HZFGVGO HZFUGB FGPIGEGAG VRXFGR JZMEXWXROMGP AGPV JZJGOPDGP WZEGPGP WZPMOPV IGRGJ WZEIGVGPVGP IGP DZKGPVGP GPMGEGFGPVHG. WZRGFUBGP HOPVGWUEG JZEUWGDGP HGRGB HZFUGB IGEOWGIG ROJG WZRGFUBGP MZEHOFUD IUPOG.

HUJFZE: KODOWZIOG

--- END INTERCEPTED TRANSMISSION ---
```

![](/static/content/standcon-2021/Untitled%2017.png)

```
REPUBLIK SINGAPURA ADALAH SEBUAH NEGARA PULAU YANG TERLETAK DI SELATAN SEMENANJUNG MELAYU DI ASIA TENGGARA. SINGAPURA TERPISAH DARI MALAYSIA DENGAN SELAT JOHOR DI UTARA, DAN KEPULAUAN RIAU INDONESIA DENGAN SELAT SINGAPURA DI SELATAN. SINGAPURA ADALAH SALAH SATU PUSAT KEWANGAN TERUNGGUL, DAN DIKENALI SEBAGAI SEBUAH BANDARAYA GLOBAL METROPOLITAN YANG MEMAINKAN PERANAN PENTING DALAM PERDAGANGAN DAN KEWANGAN ANTARABANGSA. PELABUHAN SINGAPURA MERUPAKAN SALAH SEBUAH DARIPADA LIMA PELABUHAN TERSIBUK DUNIA. SUMBER: WIKIPEDIA

STC{REPUBLIK SINGAPURA ADALAH SEBUAH NEGARA PULAU}
```

---

# Pwn

## Rocket Science

Points: 1000

Challenge Files:

[requirements.txt](/static/content/standcon-2021/requirements.txt)

[rocket_science.py](/static/content/standcon-2021/rocket_science.py)

[](https://raw.githubusercontent.com/pyupio/safety-db/master/data/insecure_full.json)

![](/static/content/standcon-2021/Untitled%2018.png)

![](/static/content/standcon-2021/Untitled%2019.png)

![](/static/content/standcon-2021/Untitled%2020.png)

![](/static/content/standcon-2021/Untitled%2021.png)

---

# Misc

## Random Encode

Points: 1000

Challenge File:

[flag.txt](/static/content/standcon-2021/flag.txt)

![](/static/content/standcon-2021/Untitled%2022.png)

```python
nc 20.198.209.142 55051
Welcome to my random encryption server! please enter your plaintext..
STC{testing}
2 0 1 S B \ @ ^ R T { % C 8 T ^ { t   Q e i y s K e a f 4 t o R \ w D i B ^ G " b n Y   N g S Z m ( }

STC{hello}
2 0 1 S % } B \ @ T ^ R C { % 8 { h T ^ e   Q l i y K e a l f 4 o R \ o w D B ^ G } 
```

```
8 9 2 % S B - T j # c > e C _ w ^ ` o { " r % c x 4 R a k } ) n 7 D v d k 0 z m ) . " ? a 1 2 y ? _ ~ i p w # l z r r / w b [ 3 b + _ : Z d 2 E Z _ i $ k E > c o + R j 5 C a X H b V z ] [ l e - M R x }

STC{ " r % c x 4 R a k } ) n 7 D v d k 0 z m ) . " ? a 1 2 y ? _ ~ i p w # l z r r / w b [ 3 b + _ : Z d 2 E Z _ i $ k E > c o + R j 5 C a X H b V z ] [ l e - M R x }
```

## Mend the lift to Space

Points: 1000

Challenge File:

[liftoff.pass](/static/content/standcon-2021/liftoff.pass)

Mendeleeve created the Periodic Table

[Periodic Table Cipher - Chemical Element Number Online Translator](https://www.dcode.fr/atomic-number-substitution)

![](/static/content/standcon-2021/Untitled%2023.png)

```
52 53 36 92 36 73 20 78 18 84 53 83 90 64 77 21 79 40 86 34 72 20 64 17 71 64 53 73 38 64 86 79 18 87 20 83 84 38 94
```

![](/static/content/standcon-2021/Untitled%2024.png)

![](/static/content/standcon-2021/Untitled%2025.png)

[Atomic.csv](/static/content/standcon-2021/Atomic.csv)

```
STC{Ch3m1sTry_l4nGuAg3_0f_ThE_un1v3rsE}
```

## Transmission

Points: 1000

Challenge File:

![](/static/content/standcon-2021/Space.jpg)

```bash
# Extract MPEG from Space.jpg using dd command
```

[/static/content/standcon-2021/extract.mp4](/static/content/standcon-2021/extract.mp4)

```
ffmpeg -i extract.mp4 frames/out-%01d.jpg
```

[jpg.zip](/static/content/standcon-2021/jpg.zip)

[solve.py](/static/content/standcon-2021/solve.py)

![](/static/content/standcon-2021/1.png)

![](/static/content/standcon-2021/2.png)

![](/static/content/standcon-2021/3.png)

[solve_v2.py](/static/content/standcon-2021/solve_v2.py)

![](/static/content/standcon-2021/2%201.png)

---