# go-dot

I don't know if there is already a tool available for this or not, but, I thought that I would try my hand at doing this.

This tool is aimed at managing your dot files, so what does that mean?

If you are like me then you have various dot files, weather that's your `.zsh` profile or your `init.vim` file.

also if you are like me then when you pull your repo of your dot files from Github then go into the file, then run the various commands to copy the files to their
respective directory path.

That's where this CLI comes in

What this application works to do is make it easy, for example some ideas:
- `go-dot init` will start and ask you to fill in the requred information of your files you wish to have tracked, if you set up already then will be prompted if you wish to overwrite
- `go-dot stage` will fetch all your dot files path locations that you have listed when starting
- `go-dot commit` will push them to your website where it is saved (Github/Bitbucket/Gitlab)
- `go-dot update` will fetch the information and save it accordingly
