	███╗   ███╗██╗ ██████╗██████╗  ██████╗ ████████╗██╗███████╗
	████╗ ████║██║██╔════╝██╔══██╗██╔═══██╗╚══██╔══╝██║██╔════╝
	██╔████╔██║██║██║     ██████╔╝██║   ██║   ██║   ██║███████╗
	██║╚██╔╝██║██║██║     ██╔══██╗██║   ██║   ██║   ██║╚════██║
	██║ ╚═╝ ██║██║╚██████╗██║  ██║╚██████╔╝   ██║   ██║███████║
	╚═╝     ╚═╝╚═╝ ╚═════╝╚═╝  ╚═╝ ╚═════╝    ╚═╝   ╚═╝╚══════╝


# installation #


You will need a `.bash_profile` named file in your home folder.
get the given binary from https://github.com/fehersanyi/bin, run the appropriate one with the update command,
it will put it to it's place, the run alias with the argument you want to use the app with
after this you want to run `source ~/.bash_profile` in your terminal, and from this it will work.
to update it later you can run the update command, and it will get the latest version.

### or simply copy to terminal ###

#### for mac #### 🍎

```
touch ~/.bash_profile
git clone https://github.com/fehersanyi/bin
./bin/mac/microtis update
./bin/mac/microtis alias mc
rm -rf ./bin
source ~/.bash_profile
```

#### for linux #### 🐧

```
touch ~/.bash_profile
git clone https://github.com/fehersanyi/bin
./bin/unix/microtis update
./bin/unix/microtis alias mc
rm -rf ./bin
source ~/.bash_profile
```

#### for windows #### 💩

https://www.ubuntu.com/download

(just kidding, I will try to come up with something soon!)


this way you will be able to run the app with `mc` as a default alias, you will always be able to change it, or create another one

...note to self: create a proper installer

## feature requests/issues ##

if you have a nice idea to add to this app, please create an issue ticket 🙃
also if something is not working as expected please do create an issue report!


## usage ##

microtis [command] [args...]


`microtis holvagyok` will print the current repo and branch

`microtis kiscica` [args] will add . and commit

`microtis cica` [args] will kiscica and push

`microtis alias` [arg] will set up an alias for the command the next terminal session will be able to use your alias 😮