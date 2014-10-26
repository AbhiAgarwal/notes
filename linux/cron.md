You can also use intervals, they are defined like this: */20 this example means every 20th and if in the minutes column this will be equivalent to 0,20,40
So to run a command every monday at 5:30 in afternoon:

`30 17 * * 1 /path/to/command`

or every 15 minutes

`*/15 * * * * /path/to/command`
