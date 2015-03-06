- http://stackoverflow.com.80bola.com/questions/70661/what-is-gnu-screen
- http://blog.siyelo.com/remote-pair-programming-with-screen/
- http://dustwell.com/screen_command_to_pair_program.html

The UNIX screen command is typically used to run multiple terminal programs inside a single ssh session, and be able to disconnect/re-connect to the session without the programs noticing. It's an awesome utility, but you can also use it to let multiple people interact with the same terminal screen, and hence, allow multiple people to use the same editor at the same time.

First, I'm assuming both programmers have a user account on the same machine, and are already logged in (or ssh'd in) to the machine. (If the second programmer doesn't have an account, he can use the first programmer's account, and the steps below are the same.)

Enabling multi-user with screen

There are two ways to do this. One way is to do
chmod u+s /usr/bin/screen
first, and then make sure everyone's ~/.screenrc file contains:
multiuser on
acladd second_programmer_username
The other way is to just have
multiuser on
acladd root
But then the second programmer will need to do sudo screen instead of just screen in the steps below. (There are also more advanced security options.)
First Programmer: run screen

The first programmer starts his day by doing:
screen
Hit ENTER to dismiss the screen startup message. Then, go about your normal activities, such as running vim, or grep, or whatever.
Second Programmer: attach to his screen

The second user does the following:
[sudo] screen -rx first_programmer_username/
This attaches to the other user's active screen.
That's it! Now you should both be seing the same window and can both use their keyboard.

Problems with the Delete Button?

When ssh'ing from a terminal in Mac OSX I noticed that my delete button no longer worked. (ssh'ing from PuTTY in Windows didn't have this problem.) To fix this, do:
`TERM=screen; screen`
whereever you would normally do screen. Or, it's probably easier to just put
`alias screen='TERM=screen; screen'`
in your `~/.bashrc` file.
Essential commands while inside screen

Command	Purpose

- Ctrl-a ?	Show the screen help menu
- Ctrl-a d	Dettach from the screen (without killing it)
- Ctrl-a c	Create another window inside the screen.
- Ctrl-a <space>  	Cycle to the next window.
- Ctrl-a a	Alternate to the previous window.
- Ctrl-a [	Enter "scroll mode" (use Up/Down arrows, then ESC to exit).

Other screen flags

Command-line	Purpose

- screen -ls	List all the active screens I have on this machine.
- screen	Start a new screen session.
- screen -x	Reattach to my pre-existing screen session.
