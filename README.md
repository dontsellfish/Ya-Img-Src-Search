# Ya**** Reverse Image Search integrated with Telegram

## Unnecessary Description 

Original _Ya**** Reverse Image Search_ is the best in the world in my opinion, but their mobile UX suck. 
And I don't think it's their fault, that's just because mobile browser UX is kind of garbage.

So I created this bot originally on python3 in one evening, 
but it was kind of buggy and written in spaghetti code style.
This revision contains some legacy of python, 
because I couldn't get Ya**** to not give me `403 FORBIDDEN` with my Golang code.
(It would be great if someone could figure that out for me. Maybe I'll fix it one day, 
but I think I'm ok with calling python time to time)

## Setting Up 

The bot requires config with API tokens for [Telegram Bot](https://api.imgbb.com/) 
and [Imgbb](https://core.telegram.org/bots) set up in `cfg.json`. 
Probably you'd like to provide a whitelist of users to reduce bandwidth and don't
get banned by Ya**** (also technically they'll use your IP address, 
so FBI shall go for you the first)

Example of `cfg.json` is included.

### Setup and Run
Don't forget to fill ```cfg.json``` before running. 
```shell
git clone https://github.com/dontsellfish/Ya-Img-Src-Search
python3 -m pip install bs4
go build
./main
```

### And that's it. 
Any code review is appreciated, because I'm junior at Golang dev. 
And, yeah, rewrite yaimg.py in Go would be lovely.

(btw, Ya**** don't ban me pls, I love you)
