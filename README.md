# 密码生成器|GeneratePassword

用Go语言编写的密码生成器，在***https://www.cnblogs.com/xwxz/p/13322927.html*** 的基础上进行了适当修改。

## 参数
  ```
  -l int
        -l 生成密码的长度 (default 8)
  -n int
        -n 生成的个数 (default 5)
  -o string
        -o 将密码保存到到指定位置,
                如果不指定位置但使用了该参数, 请输入default, 将保存在默认位置: 【桌面/out.txt】
                必须对参数值添加单或双引号！ (default "nil")
           注意：
           如果已存在该文件，则会追加到文件中
  -t string
        -t 设置密码复杂度,
                num:只使用数字[0-9], 如: 637599124317,
                char:只使用英文字母[a-zA-Z], 如: SHXEXPHjoJsx,
                mix:使用数字[0-9]和字母[a-zA-Z], 如: r12srF27nfB3,
                advance:使用数字[0-9]、字母[a-zA-Z]以及特殊字符[+=-#~.!^*$/&_], 如: F[]QnK*a-2aE
                注意：
                1. 建议对参数值添加单或双引号,
                2. mix模式和advance模式下, 首字符必是字母 (default "mix")
  ```
## 用法
### 不编译
`go run generatePassword.go [...]`
### 编译运行
#### Windows
`generatePassword.exe [...]`
#### Linux
`./generatePassword [...]`

## 示例
* 全默认<br><br>
`./generatePassword` 等价于 `./generatePassword -n 5 -l 8 -t mix`<br><br>
![image](https://user-images.githubusercontent.com/37327252/162928835-633fab2e-feb2-4858-a84c-e62d6490f141.png)

* 生成8个12位纯数字密码<br><br>
`./generatePassword -n 8 -l 12 -t num`<br><br>
![image](https://user-images.githubusercontent.com/37327252/162929557-bc6ab386-aa06-43ff-b402-9bcaade2a3e9.png)

* 生成6个10位字母+数字的密码，并将密码保存到默认位置（桌面/out.txt)<br><br>
`./generatePassword -n 6 -l 10 -t mix -o default`<br><br>
![image](https://user-images.githubusercontent.com/37327252/162930558-a36fde91-0993-4847-a211-64d0eada4342.png)

* 生成8个16位字母+数字+特殊字符的密码，并保存到指定位置<br><br>
`./generatePassword -n 8 -l 16 -t advance -o "D:/Programming/Go/src/Tools/GeneratePassword/pass.txt"`<br><br>
![image](https://user-images.githubusercontent.com/37327252/162931455-feb10a51-6e84-44c5-a84f-3ae22bff1bd0.png)
