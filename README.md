<div align="center">
    it's not all about practicality, but more of entertainment.
</div>

##

git-uwu is a fun, expewimentaw p-pwoject that expewiments awound uwuifying git commits by using a custom uwuify impwementation
and go-git to enabwe a fun, and surewy funny expewience. it is nyot intended to be used in pwoduction, and doesn't impwement
aww the features of git s-such as signed commits, etc.

#### commands

- `git-uwu add`: adds a g-given path (`--path`) to t-the twacked fiwes in git.
- `git-uwu c-commit`: commits the twacked fiwes to git with a message (`--message`) that wiww be uwuified.

##### exampwes
```shell
git-uwu add --path .
git-uwu commit -m "some commit message"
git push
```

#### impwementation

git-uwu's uwuifier i-is simiwar to nani [daniew-wiu-c0deb0t](https://github.com/daniel-liu-c0deb0t/uwu) cweated but with a few additionaw wuwes
such as:
- a-aww symbows awe ignyowed.
- you can ignyowe a specific wegion of text by adding `<: ur text here :>`
- `ur` wiww nyot be uwuified b-because that fewt weiwd.

#### configuring

you can configure git-uwu b-by adding a-a `config.yaml` fiwe (exact extension, must be `.yaml`) and editing the fowwowing pwopewties:
- `stutter_chance`: the wate of stutter t-to be done, for e-exampwe, if the v-vawue is 10 t-then it wouwd be 1 out of 10.
- `emoji_chance`: the chance of a emoji to b-be appended.
- `emojis`: the emojis to append at a chance evewy after a sentence.
- `replacer`: wepwaces specific wowds to anyother such as `cúte` to `kawaii`

##### defauwt configuration: 
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
