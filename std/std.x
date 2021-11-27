// Official Std library
// Writed by SolindekDev

name std;

made type string:letthers;
made type int:numbers;

vd println(string text) {
    if (clib("stdio.h") == True) {
        caddlib("stdio.h");
        inpc("printf(\"%(text)\")");
    } else {
        inpc("printf(\"%(text)\")");
    }
}