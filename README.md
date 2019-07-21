# ifsun
**ifsun** is a nice little linux command line tool for people like me who are obsessed over checking sunrise/sunset time for the day. This tool is written in go-lang. Contributions welcome!

<p align="center">
  <img src="https://github.com/sudharssun/ifsun/blob/master/icons/sunset-2.jpg" width="100" height="100">
</p>

#### Background
If you had lived in Seattle, USA you would know why this tool would be appreciated: In the months November to January the sun sets before 5.00PM and Seattlites typically are eagerly waiting for that day to arrive when its 5:01PM and sun has still not set. Trust me, it is INSANELY depressing if it is dark both when you get to work and when you get home.

#### Installing the release binary
Download the v0.1 beta release from [here](https://github.com/sudharssun/ifsun/releases/download/v0.1-beta/release-0.1-beta.tar.gz)

Unzip it in a directory, copy it to a directory in your PATH variable
```
tar -xvzf release-0.1-beta.tar.gz -C <your directory name>
cd <directory name>
cp ifsun /usr/local/bin
```
Check the optional [preparation](#Preparation) step and [Running the executable](#Running-the-executable) section below for instructions on how to run it.

#### Build and Execution

##### Preparation (optional)
Note: This step is only needed if you want to get sunset times of cities other than your current location.

**ifsun** uses [opencagedata API](https://opencagedata.com/api) to get timezone and geographical coordinates of places. In order to use their API, you need to register with a free key [here](https://opencagedata.com/users/sign_up). Once you get your key in the email, do the following step to add this to the environment variable

<code> export APIKEY=[YOUR KEY]</code>

##### Building the code
Install go lang compiler from [here](https://golang.org/doc/install). This will soon be integrated in the makefile so that it pulls it automatically when you build the project.
```
cd src/ifsun
make prod
```

##### Add to your path (optional)
```
cp src/ifsun/build/ifsun /usr/local/bin
(you may have to execute this as sudo)
```

##### Running the executable
When executed without arguments, it picks the current location

<code>ifsun</code>

```
Today: 21:04
Tomorrow: 21:03
Day after: 21:02
```

You can provide a city or country name or even an address! Note: If the address contains spaces, pass it in double quotes.

<code>ifsun Dubai</code>

```
Today: 19:10
Tomorrow: 19:10
Day after: 19:10
```

You can provide more details about a city like its state or country to get more accurate results. 

<code>ifsun Redmond,OR</code>

```
Today: 20:43
Tomorrow: 20:42
Day after: 20:42
```

Million Thanks to [arradon](https://github.com/araddon/dateparse) for the date parse go-lang library and sunrise-sunset.org for their REST API
