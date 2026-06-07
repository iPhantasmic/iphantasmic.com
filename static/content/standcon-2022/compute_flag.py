import hashlib

plus_code = input("Enter plus code: ").strip()

print("STANDCON22{" + hashlib.md5(plus_code.encode("utf-8")).hexdigest() + "}")
