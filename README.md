# GoAnyBrute
A much needed update of my old Python script to use my new favorite language!

A script for generating wordlists of feasibly guessable IDs, (intruder is BAD at this...)
Output to a wordlist, allows piping into pbcopy for VERY easy importing into BurpSuite's wordlist via the "paste" button

## Some Examples:::
`go run anyBrute.go -h`
 
 _Shows help information._

---
`go run anyBrute.go -c 4 --charset hexLower --form XXXXX-XXXXXX`

_Outputs:_
```
  26650-c9e387
  bffd8-dc905a
  179ee-966040
  e9002-2289a6
```
---
`go run anyBrute.go -c 1 --form "X-Client-ID: ZZZZ-ZZZZZZZZ-ZZZZZZ" -cc Z`

_Outputs:_
```
  X-Client-ID: 4633-fee1e3bd-237912
```
---
`go run anyBrute.go -c 1 --form XXXXXXXXXXXX-XX-XXXXXXXX`

_Outputs:_
```
  12d0c35cbdb8-e8-d827b5ef
```
---
`go run anyBrute.go -l 16 -c 50 | pbcopy`

_Copies fifty 16-length randomized hex strings to your clipboard for pasting directly into Intruder_

---
<br>

## All Charset Options
`nums`:                        `1234567890`
`hexUpper`:                    `ABCDEF1234567890`
`hexLower`:                    `abcdef1234567890`
`asciiUpper`:                  `ABCDEFGHIJKLMNOPQRSTUVWXYZ`
`asciiLower`:                  `abcdefghijklmnopqrstuvwxyz`
`asciiUpperLower`:             `ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz`
`asciiUpperAndNums`:       `ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890`
`asciiLowerAndNums`:       `abcdefghijklmnopqrstuvwxyz1234567890`
`asciiUpperLowerAndNums`: `ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890`
`asciiAllPrintables`: \\/ \\/ \\/       

```ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890<>?,./:;'{}|[]\"\\~`!@#$%^&*()-_=+ ```

---
__Additionally, any custom character set can be defined by assigning a charset value which does not align with the above named conventions__, for instance...

`ABC123!@#`:                  `ABC123!@#`
