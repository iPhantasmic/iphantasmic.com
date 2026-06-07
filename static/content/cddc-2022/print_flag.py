import pyminizip 
import sys 

passwd = raw_input("What is user's password?") 
passwd += raw_input("What is service's password?") 
passwd += raw_input("What is sys's password?") 
passwd += raw_input("What is root's password?") 
passwd += raw_input("What is bin's password?") 
passwd += raw_input("What is backup's password?") 
passwd += raw_input("What is daemon's password?") 

#d print "password : ", passwd
pyminizip.uncompress("flag.zip", passwd, "out", 0)
print open("flag").read()
