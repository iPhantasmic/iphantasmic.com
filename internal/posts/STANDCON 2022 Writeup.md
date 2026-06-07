---
title: "STANDCON 2022"
slug: "standcon-2022"
description: "Writeup for STANDCON 2022"
cover: "/static/content/standcon-2022/cover.png"
featured: false
published: "2022-06-24"
tags: ["CTF", "writeup"]
---

# STANDCON 2022 Writeup

![](/static/content/standcon-2022/results.png)

---

# Forensics

## Memedump

Points: 100

I too would like to be a professional meme maker. Sadly, I lack the skill or talent to make such amazing memes. BUT... and hear me out, I got hold of someone's laptop full of dank memes. Well not the laptop, just the memory of the laptop. With this, I should be able to copy all his memes right????

**Challenge File: memedump.raw**

```
volatility -f ./memedump.raw imageinfo
```

![Untitled](/static/content/standcon-2022/Untitled.png)

```
./volatility/volatility -f ./memedump.raw --profile=Win7SP1x86_23418 pslist                              130 ⨯
Volatility Foundation Volatility Framework 2.6
Offset(V)  Name                    PID   PPID   Thds     Hnds   Sess  Wow64 Start                          Exit                          
---------- -------------------- ------ ------ ------ -------- ------ ------ ------------------------------ ------------------------------
0x84e4ac78 System                    4      0     54      367 ------      0 2022-05-15 04:13:25 UTC+0000                                 
0x85cffd40 smss.exe                204      4      2       29 ------      0 2022-05-15 04:13:25 UTC+0000                                 
0x864e1d40 csrss.exe               288    272      8      209      0      0 2022-05-15 04:13:25 UTC+0000                                 
0x86510290 wininit.exe             336    272      3       75      0      0 2022-05-15 04:13:25 UTC+0000                                 
0x86512d40 csrss.exe               344    328      7      202      1      0 2022-05-15 04:13:25 UTC+0000                                 
0x865194e8 winlogon.exe            384    328      3      108      1      0 2022-05-15 04:13:25 UTC+0000                                 
0x8654e290 services.exe            428    336      8      161      0      0 2022-05-15 04:13:25 UTC+0000                                 
0x86556030 lsass.exe               436    336      6      467      0      0 2022-05-15 04:13:25 UTC+0000                                 
0x86558030 lsm.exe                 444    336      9      134      0      0 2022-05-15 04:13:25 UTC+0000                                 
0x865a2990 svchost.exe             556    428     10      333      0      0 2022-05-15 04:13:26 UTC+0000                                 
0x85da7030 svchost.exe             620    428      7      220      0      0 2022-05-15 04:13:26 UTC+0000                                 
0x865cb630 svchost.exe             672    428     16      347      0      0 2022-05-15 04:13:26 UTC+0000                                 
0x874a6588 svchost.exe             752    428     11      282      0      0 2022-05-15 04:13:26 UTC+0000                                 
0x8658d188 svchost.exe             828    428     30      704      0      0 2022-05-15 04:13:26 UTC+0000                                 
0x86626860 svchost.exe             948    428      8      154      0      0 2022-05-15 04:13:26 UTC+0000                                 
0x865fb8e8 svchost.exe            1076    428      5       93      0      0 2022-05-15 04:13:26 UTC+0000                                 
0x86691880 svchost.exe            1120    428      9      293      0      0 2022-05-15 04:13:26 UTC+0000                                 
0x86728030 dwm.exe                1404    752      3       70      1      0 2022-05-15 04:13:27 UTC+0000                                 
0x8672b530 explorer.exe           1416   1380     32      762      1      0 2022-05-15 04:13:27 UTC+0000                                 
0x8674d6e8 taskhost.exe           1480    428      8      145      1      0 2022-05-15 04:13:27 UTC+0000                                 
0x8678b530 Everything.exe         1600   1416      2       59      1      0 2022-05-15 04:13:28 UTC+0000                                 
0x867bd030 dllhost.exe            1936    556     28      537      1      0 2022-05-15 04:14:07 UTC+0000                                 
0x86800a58 WmiPrvSE.exe           1104    556      6      113      0      0 2022-05-15 04:14:27 UTC+0000                                 
0x86787d40 mspaint.exe            1464   1416     12      297      1      0 2022-05-15 04:14:31 UTC+0000                                 
0x84f66d40 svchost.exe             544    428      9      118      0      0 2022-05-15 04:15:26 UTC+0000                                 
0x867ef030 DumpIt.exe             1784   1416      2       39      1      0 2022-05-15 04:18:37 UTC+0000                                 
0x8660b030 conhost.exe            1816    344      2       52      1      0 2022-05-15 04:18:37 UTC+0000
```

```
./volatility/volatility -f ./memedump.raw --profile=Win7SP1x86_23418 memdump -p 1464 -D ./
Volatility Foundation Volatility Framework 2.6
************************************************************************
Writing mspaint.exe [  1464] to 1464.dmp

mv 1464.dmp 1464.raw
```

![Untitled](/static/content/standcon-2022/Untitled%201.png)

![Untitled](/static/content/standcon-2022/Untitled%202.png)

```
STANDCON22{meme_mem_dump}
```

---

# OSINT

## Locate Me

Points: 100

[compute_flag.py](/static/content/standcon-2022/compute_flag.py)

![locate_me.jpg](/static/content/standcon-2022/locate_me.jpg)

![Untitled](/static/content/standcon-2022/Untitled%203.png)

```
Latitude: 14° 34’ 55.2” N
Longitude: 120° 58’ 22.8” E

14° 34’ 55.2” N 120° 58’ 22.8” E
```

[14.5819634,120.9727315](https://www.google.com/maps/place/14%C2%B034'55.2%22N+120%C2%B058'22.8%22E/@14.5819634,120.9727315,20.22z/data=!4m5!3m4!1s0x0:0x844e426ce5a2374a!8m2!3d14.582!4d120.973)

![Untitled](/static/content/standcon-2022/Untitled%204.png)

```
HXJF+Q6R Manila, Metro Manila, Philippines
```

![Untitled](/static/content/standcon-2022/Untitled%205.png)

```
STANDCON22{69fb13fc17536fa7d6a5d0043a0b9f12}
```

## Trolley Trolling!

Points: 300

The thief needed to make a quick getaway, but he left his loot where he could come back for it later, though he never did.

If you manage to find it, the flag is STANDCON22{text_on_box}. Replace any spaces or punctuation with an underscore "_".

[Trolley_Trolling_Loot.pdf](/static/content/standcon-2022/Trolley_Trolling_Loot.pdf)

Free Hint:

[Trolley Trolling Text Version.pdf](/static/content/standcon-2022/Trolley_Trolling_Text_Version.pdf)

[Google Maps](https://www.google.com/maps/@1.2654759,103.8211477,3a,75y,153.14h,90.18t/data=!3m7!1e1!3m5!1sRVw9NfOCBTm6wJR3i55StA!2e0!6shttps:%2F%2Fstreetviewpixels-pa.googleapis.com%2Fv1%2Fthumbnail%3Fpanoid%3DRVw9NfOCBTm6wJR3i55StA%26cb_client%3Dmaps_sv.tactile.gps%26w%3D203%26h%3D100%26yaw%3D222.11035%26pitch%3D0%26thumbfov%3D100!7i16384!8i8192)

![Untitled](/static/content/standcon-2022/Untitled%206.png)

[Google Maps](https://www.google.com/maps/@1.2655199,103.8208475,3a,75y,221.47h,72.22t/data=!3m6!1e1!3m4!1sgP4lfKkjYB45F5yYbzC8_A!2e0!7i16384!8i8192)

![Untitled](/static/content/standcon-2022/Untitled%207.png)

[Google Maps](https://www.google.com/maps/@1.2658998,103.8208671,3a,75y,321.86h,89.42t/data=!3m6!1e1!3m4!1s7aRmFHmc8vRZOYjnq-ujQA!2e0!7i16384!8i8192)

![Untitled](/static/content/standcon-2022/Untitled%208.png)

[Google Maps](https://www.google.com/maps/@1.2668472,103.8205459,3a,18.3y,294.83h,82.52t/data=!3m7!1e1!3m5!1s-sGOixKFz1y47QgJFLnLvQ!2e0!5s20180401T000000!7i16384!8i8192)

![Untitled](/static/content/standcon-2022/Untitled%209.png)

```
STANDCON22{NIPPON_PAINT}
```

## Incoherent Symbols

Points: 300

You receive the following image. 

![Incoherent_Message.jpg](/static/content/standcon-2022/Incoherent_Message.jpg)

[incoherent_flag.zip](/static/content/standcon-2022/incoherent_flag.zip)

![Untitled](/static/content/standcon-2022/Untitled%2010.png)

[incoherent_flag.txt](/static/content/standcon-2022/incoherent_flag.txt)

After CTF:

[OOArmadeus (u/OOArmadeus) - Reddit](https://www.reddit.com/user/OOArmadeus/)

[InkoherianWIP.ttf](https://drive.google.com/file/d/12_tLne6kWo7sB-zjLQ-qGQMbh1xTn6Hc/view)

```
STANDCON22{Ink0herian_I5_n0_Match_4_U}
```

## I Sea You (Part 2)

Points: 100

This is the second challenge of the "I SEA You Series" (out of a total of three challenges)

Good Job Recruit! We investigated Agent Orca and found that he has unfortunately left Singapore! However, all hope is not lost! Although he has wiped out most of his Computer's Hard Drive, we have managed to recover a file that he stored online. Investigate the file and report back!

Access the file here: [https://docs.google.com/presentation/d/10uAbsSbh9LVSDK4Qwa4QfiFGBBl1zKi5OdSCVoPJQ5I/edit#slide=id.p](https://docs.google.com/presentation/d/10uAbsSbh9LVSDK4Qwa4QfiFGBBl1zKi5OdSCVoPJQ5I/edit#slide=id.p)

You will need to find:

1. The original title of the file (lowercase, no spaces, if any)
2. Date of File Creation (DDMMYYYY) (GMT+8, SGT)
3. Orca's Instagram Handle (including the @, lowercase)
4. The Secret Text on the second last Slide (copy it down as you see it - no spaces)

And submit it in the format STANDCON22{<original title of file>_<date of file creation in the format DDMMYYYY>_<agent orca's instagram handle>_<the secret text>}

For instance, a valid flag submission would be: STANDCON22{filenamewithnospaces_02022022_@thisisnotarealaccount_somemagicsecret}

![Untitled](/static/content/standcon-2022/Untitled%2011.png)

```
STANDCON22{orcaifydiary_17022022_<insta>_<secret>}
```

![Untitled](/static/content/standcon-2022/Untitled%2012.png)

![Untitled](/static/content/standcon-2022/Untitled%2013.png)

[atlantis forever](https://web.archive.org/web/20220217070842/https://docs.google.com/presentation/d/10uAbsSbh9LVSDK4Qwa4QfiFGBBl1zKi5OdSCVoPJQ5I/edit)

![Untitled](/static/content/standcon-2022/Untitled%2014.png)

![Untitled](/static/content/standcon-2022/Untitled%2015.png)

![Untitled](/static/content/standcon-2022/Untitled%2016.png)

![Untitled](/static/content/standcon-2022/Untitled%2017.png)

```
STANDCON22{orcaifydiary_17022022_@whaleywonka_iloveplushies}
```

---

# Misc

## Shark in the Ocean

Points: 100

Shark??? in the ocean??? Ayo?????

[SHARK_IN_THE_OCEAN.pcapng](/static/content/standcon-2022/SHARK_IN_THE_OCEAN.pcapng)

![Untitled](/static/content/standcon-2022/Untitled%2018.png)

![Untitled](/static/content/standcon-2022/Untitled%2019.png)

```
STANDCON22{W1R3SH4RK_EXP3RT?}
```

## Walks like a cat, barks like a dog

Points: 100

What makes Zebra, a Zebra? What makes Cow, a Cow? What makes Dog, a Dog? These are some tough questions. What's more tough is what makes a PDF file, a PDF? And what makes a PNG file, a PNG? I suppose only you can answer that.

[trickery.pdf](/static/content/standcon-2022/trickery.pdf)

![Untitled](/static/content/standcon-2022/Untitled%2020.png)

[251B.zip](/static/content/standcon-2022/251B.zip)

[note.txt](/static/content/standcon-2022/note.txt)

[Trickery](/static/content/standcon-2022/Trickery.txt)

![Untitled](/static/content/standcon-2022/Untitled%2021.png)

![Untitled](/static/content/standcon-2022/Untitled%2022.png)

```
STANDCON22{f1l3_f0rm4ts_4r3_t00_d4mn_c0mpl!c4t3d}
```

---

# RE

## Wait for the day

Points: 500

Momma bought a Christmas gift for you. She says it's something you love! But Christmas is like more than 6 months away :/ Try to peek in and find out what the gift is, before the Christmas eve ;)

Provided binary:

[gift_for_you](/static/content/standcon-2022/gift_for_you.txt)

![Untitled](/static/content/standcon-2022/Untitled%2023.png)

![Untitled](/static/content/standcon-2022/Untitled%2024.png)

![Untitled](/static/content/standcon-2022/Untitled%2025.png)

![Untitled](/static/content/standcon-2022/Untitled%2026.png)

![Untitled](/static/content/standcon-2022/Untitled%2027.png)

```
0x7fffffffdf20: 0x4e415453      0x4e4f4344      0x217b3232      0x345f7374
0x7fffffffdf30: 0x7934776c      0x5f345f73      0x64303067      0x6d21745f
0x7fffffffdf40: 0x00007d65      0x00000000
```

![Untitled](/static/content/standcon-2022/Untitled%2028.png)

![Untitled](/static/content/standcon-2022/Untitled%2029.png)

```
STANDCON22{!ts_4lw4ys_4_g00d_t!me}
```

---

# Web

## A Fishy Site

Points: 300

You sail across the ocean looking for lost treasures. One night, you see lights coming out somewhere deep in the ocean. You decided to dive into the ocean towards the light....

**Please do not brute force passwords. It can be guessed.**

![Untitled](/static/content/standcon-2022/Untitled%2030.png)

![Untitled](/static/content/standcon-2022/Untitled%2031.png)

![Password: ocean](/static/content/standcon-2022/Untitled%2032.png)

Password: ocean

![Untitled](/static/content/standcon-2022/Untitled%2033.png)

[Offensive Security's Exploit Database Archive](https://www.exploit-db.com/exploits/43963)

[https://github.com/flozz/p0wny-shell](https://github.com/flozz/p0wny-shell)

![Untitled](/static/content/standcon-2022/Untitled%2034.png)

![Untitled](/static/content/standcon-2022/Untitled%2035.png)

```
STANDCON22{L0ST_C1TY_TR34SUR3_1S_M1N3!}
```

## Maze Repo

Points: 500

You are about to witness a repository that has more than it meets the eye! Some think it as a maze but I would say it's the mirror that deceives you the most :)

![Untitled](/static/content/standcon-2022/Untitled%2036.png)

![Untitled](/static/content/standcon-2022/Untitled%2037.png)

![Untitled](/static/content/standcon-2022/Untitled%2038.png)

[https://github.com/arthaud/git-dumper](https://github.com/arthaud/git-dumper)

![Untitled](/static/content/standcon-2022/Untitled%2039.png)

![Untitled](/static/content/standcon-2022/Untitled%2040.png)

![Untitled](/static/content/standcon-2022/Untitled%2041.png)

![Untitled](/static/content/standcon-2022/Untitled%2042.png)

![Untitled](/static/content/standcon-2022/Untitled%2043.png)

![Untitled](/static/content/standcon-2022/Untitled%2044.png)

![Untitled](/static/content/standcon-2022/Untitled%2045.png)

![Untitled](/static/content/standcon-2022/Untitled%2046.png)

![Untitled](/static/content/standcon-2022/Untitled%2047.png)

```
STANDCON22{bl33d!ng_g!t_r3p0_!s_4_tr34sur3_tr0v3}
```

## Vim Hurts

Points: 500

Sysadmin like me love Vim! It's one of the beauties that never gets old. Its the part of the hacker culture, if you will. I have configured a web server with secure file upload functionality. My friends say it might not be as secure as I think. I cannot think of any good reason except that maybe I left some trace? Maybe its my poor Vim habit? In case you are wondering, yeah, I still don't know how to exit out of Vim.

By going to:

[http://18.141.84.203:62790/.index.php.swp](http://18.141.84.203:62790/index.php.swap)

[index.php.swp](/static/content/standcon-2022/index.php.swp)

![Untitled](/static/content/standcon-2022/Untitled%2048.png)

```php
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Upload files securely!</title>
</head>
<body>
    <div align="center">

        <?php

            define("UPLOAD_DIR", "./uploads/");
            define("ERROR", "Error uploading file!" );

            function check_file($uploaded_file) {

                $ret_val = 0;

                // 10 MB limit
                if ($uploaded_file["size"] > 10485760) {
                    $ret_val = -1;
                }

                $extension = end((explode(".", $uploaded_file["name"])));

                if ($extension !== "png" && $extension !== "jpeg" && $extension !== "jpg") {
                    $ret_val = -1;
                }

                $allowed_formats = array("image/jpeg", "image/jpg", "image/png");

                if (!in_array($uploaded_file["type"], $allowed_formats)) {
                    // Invalid file!
                    $ret_val = -1;
                }

                return $ret_val;
            }
        ?>

        <div class="text-center card mt-5 shadow p-3 mb-4 bg-white rounded" style="width: 24rem; height: 20rem;">
            <div class="h2 my-5">Secure File Uploader</div>
            <form method="post" enctype="multipart/form-data">
                <br />
                <input type="submit" value="Upload file" class="btn btn-primary" />
                </div>
            </form>
        </div>
        <?php
            if ($_SERVER["REQUEST_METHOD"] == "POST" && !empty($_FILES["uploaded_file"])) {
<html lang="en">
Last login: Thu Jun 23 10:21:58 on ttys006
❯ vim index.php
❯ mv index.php.swp .index.php.swp
❯ vim index.php
❯ ls
index.php
❯ cat index.php
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Upload files securely!</title>
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-1BmE4kWBq78iYhFldvKuhfTAU6auU8tT94WrHftjDbrCEXSU1oBoqyl2QvZ6jIW3" crossorigin="anonymous">
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/js/bootstrap.bundle.min.js" integrity="sha384-ka7Sk0Gln4gmtz2MlQnikT1wXgYsOg+OMhuP+IlRH9sENBO0LRn5q+8nbTov4+1p" crossorigin="anonymous"></script>
</head>
<body>
    <div align="center">

        <?php

            define("UPLOAD_DIR", "./uploads/");
            define("ERROR", "Error uploading file!" );

            function check_file($uploaded_file) {

                $ret_val = 0;

                // 10 MB limit
                if ($uploaded_file["size"] > 10485760) {
                    $ret_val = -1;
                }

                $extension = end((explode(".", $uploaded_file["name"])));

                if ($extension !== "png" && $extension !== "jpeg" && $extension !== "jpg") {
                    $ret_val = -1;
                }

                $allowed_formats = array("image/jpeg", "image/jpg", "image/png");

                if (!in_array($uploaded_file["type"], $allowed_formats)) {
                    // Invalid file!
                    $ret_val = -1;
                }

		// Probably add some more checks here...

                return $ret_val;
            }
        ?>

        <div class="text-center card mt-5 shadow p-3 mb-4 bg-white rounded" style="width: 24rem; height: 20rem;">
            <div class="h2 my-5">Secure File Uploader</div>
            <form method="post" enctype="multipart/form-data">
                <div class="form-group">
                <input type="file" name="uploaded_file" class="form-control-file mb-5">
                <br />
                <input type="submit" value="Upload file" class="btn btn-primary" />
                </div>
            </form>
        </div>
        <?php
            if ($_SERVER["REQUEST_METHOD"] == "POST" && !empty($_FILES["uploaded_file"])) {
                $uploaded_file = $_FILES["uploaded_file"];
                if ($uploaded_file["error"] !== UPLOAD_ERR_OK) {
                    // echo $ERROR;
                    echo '<div class="alert alert-danger" role="alert" style="width: 24rem;">File upload failed!</div>';
                }

		else {
			// Check the filename is safe
			$name = preg_replace("/[^A-Z0-9._-]/i", "_", $uploaded_file["name"]);

			// Grab file from the temp dir
			$success = move_uploaded_file($uploaded_file["tmp_name"], UPLOAD_DIR . $name);

			// Quarantine malicious/unexpected file
			if (check_file($uploaded_file) === 0) {
			    // All good..., we can send the success message back!
			    echo '<div class="alert alert-success" role="alert" style="width: 24rem;">File uploaded successfully <a target="_blank" href="/uploads/' . $name . '">here</a>!</div>';
			}
			else {
			    // Something's not right
			    // Quarantine the file!
			    // An alert will be sent to the admin automatically.
			    rename(UPLOAD_DIR . $name, "/tmp/" . $name);
			    // Send the error message back!
			    echo '<div class="alert alert-danger" role="alert" style="width: 24rem;"><b>Error:</b> Only image files (jpg/png) are allowed!</div>';
			}
		}
            }
        ?>
    </div>
</body>
</html>
```

![Exploit the race condition here](/static/content/standcon-2022/Untitled%2049.png)

Exploit the race condition here

![ngrok tcp 8080](/static/content/standcon-2022/Untitled%2050.png)

ngrok tcp 8080

![Untitled](/static/content/standcon-2022/Untitled%2051.png)

```python
import requests

while True:
    r = requests.get("http://18.141.84.203:62790/uploads/hehexd.php")
    if r.status_code != 404:
        break
    else:
        print(r.status_code)
```

![Untitled](/static/content/standcon-2022/Untitled%2052.png)

![Untitled](/static/content/standcon-2022/Untitled%2053.png)

![Untitled](/static/content/standcon-2022/Untitled%2054.png)

```
STANDCON22{v!m_swp_f!l3s_c4n_s3rv3_4tt3ck3rs}
```