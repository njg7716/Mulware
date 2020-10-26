package main

import (
	"fmt"
	"log"
	"os/exec"

	"golang.org/x/net/icmp"
)

func main() {
	conn, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		log.Fatal(err)
	}

	for {
		msg := make([]byte, 4)
		_, _, err := conn.ReadFrom(msg)
		if err != nil {
			fmt.Println(err)
			continue
		}
		//Do Things!
		var spoopy string
		spoopy = "               _.---._\n"
		spoopy = spoopy + "             .'       `.\n"
		spoopy = spoopy + "             :)       (:\n"
		spoopy = spoopy + "             \\ (@) (@) /\n"
		spoopy = spoopy + "              \\   A   / \n"
		spoopy = spoopy + "               )     ( \n"
		spoopy = spoopy + "               \\\"\"\"\"\"/ \n"
		spoopy = spoopy + "                `._.'\n"
		spoopy = spoopy + "                 .=.\n"
		spoopy = spoopy + "         .---._.-.=.-._.---.\n"
		spoopy = spoopy + "        / ':-(_.-: :-._)-:` \\ \n"
		spoopy = spoopy + "       / /' (__.-: :-.__) `\\ \\ \n"
		spoopy = spoopy + "      / /  (___.-` '-.___)  \\ \\ \n"
		spoopy = spoopy + "     / /   (___.-'^`-.___)   \\ \\ \n"
		spoopy = spoopy + "    / /    (___.-'=`-.___)    \\ \\ \n"
		spoopy = spoopy + "   / /     (____.'=`.____)     \\ \\ \n"
		spoopy = spoopy + "  / /       (___.'=`.___)       \\ \\ \n"
		spoopy = spoopy + " (_.;       `---'.=.`---'       ;._)\n"
		spoopy = spoopy + " ;||        __  _.=._  __        ||;\n"
		spoopy = spoopy + " ;||       (  `.-.=.-.'  )       ||;\n"
		spoopy = spoopy + " ;||       \\    `.=.'    /       ||;\n"
		spoopy = spoopy + " ;||        \\    .=.    /        ||;\n"
		spoopy = spoopy + " ;||       .-`.`-._.-'.'-.       ||;\n"
		spoopy = spoopy + ".:::\\      ( ,): O O :(, )      /:::.\n"
		spoopy = spoopy + "|||| `     / /'`--'--'`\\ \\     ' ||||\n"
		spoopy = spoopy + "''''      / /           \\ \\      ''''\n"
		spoopy = spoopy + "         / /             \\ \\ \n"
		spoopy = spoopy + "        / /               \\ \\ \n"
		spoopy = spoopy + "       / /                 \\ \\ \n"
		spoopy = spoopy + "      / /                   \\ \\ \n"
		spoopy = spoopy + "     / /                     \\ \\ \n"
		spoopy = spoopy + "    /.'                       `.\\ \n"
		spoopy = spoopy + "   (_)'                       `(_)\n"
		spoopy = spoopy + "    \\\\.                       .//\n"
		spoopy = spoopy + "     \\\\.                     .//\n"
		spoopy = spoopy + "      \\\\.                   .//\n"
		spoopy = spoopy + "       \\\\.                 .//\n"
		spoopy = spoopy + "        \\\\.               .//\n"
		spoopy = spoopy + "         \\\\.             .//\n"
		spoopy = spoopy + "          \\\\.           .//\n"
		spoopy = spoopy + "          ///)         (\\\\\\ \n"
		spoopy = spoopy + "        ,///'           `\\\\\\, \n"
		spoopy = spoopy + "       ///'               `\\\\\\ \n"
		spoopy = spoopy + "      \"\"'                   '\"\"\n"
		spoopy = spoopy + "Spooky Scary Skeletons Send Shivers Down Your Spine!\n"
		fmt.Println("Ping!")
		cmd := exec.Command("wall", "hello")
		err = cmd.Run()
	}
}
