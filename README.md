# go-discord-bot
A very simple discord bot written in go

## To install and use this Discord bot, you will need to follow these steps:

1.) Install Go programming language on your computer if you haven't already. You can download Go from the official website at https://golang.org/dl/.<br>
2.) Clone or download the code from the repository to your local machine.<br>
3.) Open a command prompt or terminal window and navigate to the directory where the code is located.<br>
4.) Replace "YOUR_BOT_TOKEN" in the code with your actual Discord bot token.<br>
5.) Save the code and run the following command to build the executable:<br>
`go build` <br>
6.) Run the executable to start the bot:<br>
`./discord-bot` (for linux/mac os)<br>
(If you are using Windows, you will need to run `discord-bot.exe` instead.)<br>
7.) Invite the bot to your Discord server using the following link: https://discordapp.com/oauth2/authorize?client_id=YOUR_CLIENT_ID&scope=bot<br>
Replace "YOUR_CLIENT_ID" with your actual bot client ID, which you can find on the Discord Developer Portal.<br>
8.) Once the bot is invited to your server, you can use the commands you've implemented in the code by typing them into any channel that the bot has access to.<br>

<br><br>

## Usage / Commands 

For example, you can use the "ping" command by typing "ping" into a channel, and the bot will respond with "Pong!" and the response time.<br>
You can also create custom commands by adding them to a text file named "commands.txt" in the same directory as the executable. <br>
Each line in the file should contain a command and its response separated by a comma, like this:<br>
`!hello, Hello, world!`<br>
`!bye, Goodbye, world!`<br>
Then, when you type "!hello" or "!bye" in a channel, the bot will respond with "Hello, world!" or "Goodbye, world!", respectively.<br>



