# console
script executor's console functions recreated
# docs

### example
```lua
console:initialize('localhost', 8080, false)
console:clear()
console:print(console.colors.white..'hello, ' .. console.colors.bright_green .. 'user' .. console.colors.white .. '!')
console:warn('warn example')
console:error('error example')
```
### print
```lua
<number> console:print(<string> text)
```
### warn
```lua
<number> console:warn(<string> text)
```
### error
```lua
<number> console:error(<string> text)
```
### clear
```lua
<number> console:clear(<void>)
```
### input
```lua
<string> console:input(<string> prompt)
```
### title
```lua
<string> console:title(<string> title)
```
### colors
- default
- black
- white
- red
- dark_green
- dark_yellow
- dark_blue
- dark_magenta
- dark_cyan
- dark_white
- bright_black
- bright_red
- bright_green
- bright_yellow
- bright_blue
- bright_magenta
- bright_cyan
