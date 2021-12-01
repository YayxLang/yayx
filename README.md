
<!-- <div style="text-align:center"><img src="https://raw.githubusercontent.com/SolindekDev/yayx/main/website/img/yayx.png" style="width: 128px"></div> -->
<img align="center" style="width: 128px" src="https://raw.githubusercontent.com/SolindekDev/yayx/main/website/img/yayx.png" alt="Yayx Logo">

# Yayx
Yayx programming language i will try to write a interpreter in GoLang

# Directory Tree
- **src** -> Source code of our two interpreters more in [README](https://github.com/SolindekDev/yayx/tree/main/src)
- **.vscode** -> Files for visual studio code 
- **builder** -> A little builder of interpreter (soon update)
- **std** -> A source code of std library
- **examples** -> Example with Yayx Programming Language

# Hello World in Yayx
```kotlin
fn main() {
    println("Hello, World!")
}
```

# Helped me
- With writing the interpreter so help me [cpython repo](https://github.com/python/cpython) and [comet repo](https://github.com/chermehdi/comet)

# Keywords
```
==============================================
Create function keyword = fn

Example =========================
fn main() { }
==============================================
Return keyword = return

Example =========================
return <VARIABLE OR A DATA>
==============================================
Variable assigment keyword = var

Example =========================
var <NAME OF VARIABLE> = <DATA>
==============================================
Constant variable assigment keyword = cst

Example =========================
cst <NAME OF VARIABLE> = <DATA>
==============================================
True data keyword = true

Example =========================
true ( 1 )
==============================================
False data keyword = false

Example =========================
false ( 0 )
==============================================
If keyword = if

Example =========================
if (<CONDITION>) { }
==============================================
Else keyword = else

Example =========================
if (<CONDITION>) { } else { }
==============================================
Comment: #

Example =========================
 # Comment #
/\        /\
Comment   Comment
start     end
==============================================
```
