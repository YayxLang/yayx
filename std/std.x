// Official Std library
// Writed by SolindekDev

name std;

made type string:letthers;
made type int:numbers;

vd println(string text) {
    if (clib("stdio.h") == True) {
        caddlib("stdio.h");
        inpc("printf(\"%(text)\\n\")");
    } else {
        inpc("printf(\"%(text)\\n\")");
    }
}

vd print(string text) {
    if (clib("stdio.h") == True) {
        caddlib("stdio.h");
        inpc("printf(\"%(text)\")");
    } else {
        inpc("printf(\"%(text)\")");
    }
}

vd exit(int exit_code) {
    if (clib("stdlib.h") == True) {
        caddlib("stdlib.h");
        inpc("exit(%(exit_code))");
    } else {
        inpc("exit(%(exit_code))");
    }
}