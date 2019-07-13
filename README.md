# ifsun
**ifsun** is a nice little linux command line tool for people like me who are obsessed over checking sunrise/sunset time for the day. This tool is written in go-lang.

<p align="center">
  <img src="https://github.com/sudharssun/ifsun/blob/master/icons/sunset-2.jpg" width="100" height="100">
</p>

#### Background
If you had lived in Seattle, USA you would know why this tool would be appreciated: In the months November to January the sun sets before 5.00PM and Seattlites typically are eagerly waiting for that day to arrive when its 5:01PM and sun has still not set. Trust me, it is WAY depressing if it is dark both when you get to work and when you get home.

#### Build and Execution [Work in progress]
<code>go build src/fetch/fetch.go</code>

<code>./fetch</code>

Output:
```
Today: 21:04
Tomorrow: 21:03
Day after: 21:02
(Credit: https://sunrise-sunset.org/api)
```

Million Thanks to [arradon](https://github.com/araddon/dateparse) for the date parse go-lang library and sunrise-sunset.org for their REST API
