#include <stdlib.h>
#include <stdio.h>

#ifdef _WIN32
    #define OS_NAME "Windows"
#elif _WIN64
    #define OS_NAME "Windows"
#elif __APPLE__
    #define OS_NAME "AppleOS"
#elif __linux__
    #define OS_NAME "Linux"
#elif TARGET_OS_MAC
    #define OS_NAME "MacOS"
#else
    #define OS_NAME "Other"
#endif

void cls() {
    if (OS_NAME == "Windows") {
        // In cmd.exe does not support ansii sooo we need to make a open command
        // promtd and write down a clear command to it.
        system("cls");
    } else {
        // Other system support ansii so we can do just this. It takes less memory 
        // than clear in Windows
        printf("\e[1;1H\e[2J");
    }
}


