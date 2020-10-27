# Mulware
Bad stuff I wrote
## ET
### This is a Windows Persistance tool that will open a TCP random port and create a firewall rule so that someone can telnet to it and get a powershell shell

## Ping of Memes
### Linux
#### Listens for ICMP requests and then will wall a message
#### Instructions
1. Run unixPingOfMemes as sudo on the target machine and let it run in the background
2. Whenever you want to send memes, just ping the IP of the target machine
3. Sit back and enjoy the memes

### Windows
#### Listens for ICMP requests and then will open a URL
#### Instructions
1. Run winPingOfMemes.exe with elevated privledges and leave it running in the background
2. In order to set of the memes on windows machines use the ping binary written in go. There is a bug where any ICMP requests from either Unix or other Windows boxes will not set off the listener.
3. Usage: ./path/to/ping.exe \<IP of target machine\>
4. Sit back and enjoy the memes
