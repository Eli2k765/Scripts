package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  labnames := []string{
       "legacy", "blue", "devel", "optimum", "bastard", "granny", "arctic", "grandpa", "silo", "bounty", "jerry", "conceal", "chatterbox", "forest", "bankrobber", "secnotes", "bastion", "buff", "servmon", "active", "remote", "fuse", "omni", "worker", "jeeves", "bart", "tally", "netmon", "sizzle", "sniper", "control", "nest", "sauna", "cascade", "querier", "blackfield", "lame", "brainfuck", "shocker", "bashed", "nibbles", "cronos", "nineveh", "sense", "solidstate", "node", "valentine", "poison", "sunday", "tartarsauce", "irked", "friendzone", "swagshop", "networked", "jarvis", "mirai", "popcorn", "haircut", "blocky", "frolic", "postman", "mango", "traverxec", "openadmin", "magic", "admirer", "kotarak", "falafel", "devops", "hawk", "lightweight", "la casa de papel", "jail", "safe", "bitlab", "october", "mango", "book", "quick",
	}

  rand.Seed(time.Now().UnixNano())
  var length int = len(labnames)
  var indexnum int = rand.Intn(length-1)
  word := labnames[indexnum]

  fmt.Println("Your lab of the day is:", word)
}
