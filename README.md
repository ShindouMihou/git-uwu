<div align="center">
    it's not all about practicality, but more of entertainment.
</div>

##

git-uwu is a fun, expewimentaw pwoject that expewiments awound uwuifying git c-commits by using a custom uwuify impwementation
and go-git to enabwe a fun, a-and surewy funny expewience. it is nyot intended to be used in pwoduction, and doesn't impwement
aww the features of git s-such as signed commits, etc.

#### commands

- `git-uwu add`: adds a given path (`--path`) to the twacked fiwes in git.
- `git-uwu commit`: commits the twacked fiwes t-to git with a m-message (`--message`) that wiww be uwuified.

#### impwementation

git-uwu's uwuifier is simiwar to nani [daniew-wiu-c0deb0t](https://github.com/daniel-liu-c0deb0t/uwu) c-cweated but with a-a few additionaw wuwes
such as:
- aww symbows awe ignyowed.
- you can ignyowe a specific wegion of text by adding ` ur text here `
- `ur` wiww nyot be uwuified because that fewt weiwd.

#### configuring

you can configure git-uwu by adding a `config.yaml` fiwe (exact extension, must be `.yaml`) and editing the fowwowing pwopewties:
- `stutter_chance`: t-the wate of s-stutter to be done, for exampwe, if the vawue is 10 then it w-wouwd b-be 1 out of 10.
- `emoji_chance`: the chance of a emoji to be appended.
- `emojis`: the emojis t-to append at a chance evewy after a sentence.
- `replacer`: wepwaces specific wowds to anyother s-such as `cúte` t-to `kawaii`

defauwt configuration: 
```yaml
stutter_chance: 10
emoji_chance: 2
emojis:
  - rawr x3
  - owo
  - uwu
  - o.o
  - '-.-'
  - '>w<'
  - (⑅˘o˘)
  - (oᴗo)
  - (˘ω˘)
  - (u ᵕ u❁)
  - σωσ
  - òωó
  - (u ﹏ u)
  - ( ͡o ω ͡o )
  - ʘwʘ
  - ':3'
  - xd
  - nyaa~~
  - '>_<'
  - '>.<'
  - '>w<'
  - qwq
  - rawr
  - ^^
  - ^^;;
  - ' (ˆ ﻌ ˆ)♡'
  - ' ^•ﻌ•^'
  - /(^•ω•^)
  - (✿oωo)
replacer:
  - smol
  - smol
  - kawaii~
  - kawaii~
  - floof
  - floof
  - luv
  - luv
  - baka
  - baka
  - nani
  - nani
  - nya~
  - nya~
  - u
  - u
```
