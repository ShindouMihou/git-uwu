<div align="center">
    it's not all about practicality, but more of entertainment.
</div>

##

git-uwu is a-a f-fun, expewimentaw pwoject that expewiments awound uwuifying git commits by using a custom uwuify impwementation
and g-go-git to enabwe a fun, and s-surewy funny expewience. it is nyot intended to be u-used in p-pwoduction, and d-doesn't i-impwement
aww the features of git such as signed commits, e-etc.

#### c-commands

- `git-uwu a-add`: a-adds a given path (`--path`) to the t-twacked fiwes in git.
- `git-uwu commit`: commits the twacked fiwes to git with a message (`--message`) that w-wiww b-be uwuified.

##### exampwes
```shell
git-uwu add --path .
git-uwu commit -m "some commit message"
git push
```

#### impwementation

git-uwu's u-uwuifier is simiwar to nani [daniew-wiu-c0deb0t](https://github.com/daniel-liu-c0deb0t/uwu) cweated but with a few additionaw wuwes
such as:
- aww symbows a-awe ignyowed.
- y-you can ignyowe a specific wegion of t-text by adding `<: ur text here :>`
- `ur` wiww n-nyot be uwuified because that fewt w-weiwd.

#### configuring

you can configure git-uwu by adding a `config.yaml` f-fiwe (exact extension, must be `.yaml`) and editing the fowwowing pwopewties:
- `stutter_chance`: the wate of stutter t-to be done, for exampwe, i-if t-the vawue is 10 t-then it wouwd be 1 out of 10.
- `emoji_chance`: the chance of a emoji to b-be a-appended.
- `emojis`: t-the emojis to append at a-a chance evewy after a sentence.
- `replacer`: wepwaces specific wowds to anyother such as `cúte` to `kawaii`

##### defauwt configuration 
due to some issues with the uwuifier changing up stuff, u can read the default configuration at [default_config.yaml](default_config.yaml).
:>