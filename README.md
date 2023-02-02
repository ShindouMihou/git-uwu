<div align="center">
    it's not all about practicality, but more of entertainment.
</div>

##

git-uwu is a f-fun, expewimentaw pwoject that expewiments awound uwuifying git commits by using a custom uwuify impwementation
and go-git to enabwe a fun, and surewy funny expewience. it is nyot intended to be used in pwoduction, a-and doesn't impwement
aww the features of g-git such as s-signed commits, e-etc.

#### commands

- `git-uwu add`: a-adds a given path (`--path`) to the t-twacked fiwes in git.
- `git-uwu commit`: commits the twacked fiwes t-to git with a message (`--message`) that wiww be uwuified.

##### exampwes
```shell
git-uwu add --path .
git-uwu commit -m "some commit message"
git push
```

#### impwementation

git-uwu's uwuifier is simiwar t-to nani [daniew-wiu-c0deb0t](https://github.com/daniel-liu-c0deb0t/uwu) cweated but with a few additionaw wuwes
such a-as:
- aww symbows awe ignyowed.
- you can i-ignyowe a specific wegion of t-text by adding `<: ur text here :>`
- `ur` wiww nyot be uwuified because that fewt weiwd.

#### configuring

you can c-configure git-uwu by adding a `config.yaml` fiwe (exact extension, must be `.yaml`) and editing the fowwowing p-pwopewties:
- `stutter_chance`: the wate of stutter to be done, for exampwe, if the vawue is 10 then it wouwd be 1 out of 10.
- `emoji_chance`: t-the chance of a-a emoji to be appended.
- `emojis`: the emojis to append at a chance evewy after a sentence.
- `replacer`: wepwaces specific wowds to a-anyother such as `cúte` to `kawaii`

##### defauwt configuration
due to some issues with the uwuifier changing up stuff, u can wead the defauwt configuration at [default_config.yaml](default_config.yaml).