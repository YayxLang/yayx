import os
import sys

def main():
    argv = sys.argv

    # Build and run
    if argv[1] == "bar":
        os.system("gcc src/main.c")
        if sys.platform.startswith("win32") or sys.platform.startswith("cygwin"):
            os.system(".\\a.exe")
        else:
            os.system("./a.out")
    else:
        print("Help of Builder:\nTo Build and Run write into the argv 'bar'")

main() 