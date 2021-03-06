
<!-- <div style="text-align:center"><img src="https://raw.githubusercontent.com/SolindekDev/yayx/main/website/img/yayx.png" style="width: 128px"></div> -->
<img align="center" style="width: 128px" src="https://avatars.githubusercontent.com/u/95468331?s=200&v=4" alt="Yayx Logo">

![](https://shields.io/github/languages/top/YayxLang/yayx)
![](https://img.shields.io/github/directory-file-count/yayxlang/yayx)

# Yayx
Yayx programming language is begginer friendly programming language. Yayx have:
- Easy syntax
- Dynamic types
- Compiler to outhers programming language (GoLang and C)
- Clear documentation

# Directory Tree
- **src** -> Source code of our two interpreters, read more in [README](https://github.com/SolindekDev/yayx/tree/main/src)
- **.vscode** -> Files for visual studio code 
- **builder** -> Builder of interpreter (update soon)
- **std** -> Source code of std library
- **examples** -> Examples for Yayx Programming Language

# ToDo
- [x] Lexer
- [ ] End parser
- [ ] Create translator
- [ ] Add args in function
- [ ] Shell (like python shell)
- [ ] Create website
- [ ] Create documentation
- [ ] Create std library
- [ ] Create file library

# Hello World in Yayx
```kotlin
fn main() {
    println("Hello, World!")
}
```

# Special thanks to
- [cpython repo](https://github.com/python/cpython) and [comet repo](https://github.com/chermehdi/comet)

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
