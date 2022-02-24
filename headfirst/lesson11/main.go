package main

import (
	"fmt"
	"headfirst/headfirst/lesson11/gadget"
)

type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("Tweet")
}

type Horn string

func (h Horn) MakeSound() {
	fmt.Println("Honk")
}

type Robot string

func (r Robot) MakeSound() {
	fmt.Println("Beep Boop")
}

func (r Robot) Walk() {
	fmt.Println("Powering legs")
}

type NoiseMaker interface {
	MakeSound()
}

func paly(n NoiseMaker) {
	n.MakeSound()

}

func TryOut(player gadget.Player) {
	player.Play("Test Track")
	player.Stop()
	record, ok := player.(gadget.TapePlayer)
	if ok {
		record.Record()
	}
}

func main() {

	fmt.Println("Hello lesson11")

	mixtape := []string{"Jessi's Girl", "Whip It", "9 to 5"}
	var player gadget.Player = gadget.TapePlayer{}
	playList(player, mixtape)

	player = gadget.TapeRecorder{}
	playList(player, mixtape)

	/*var value mypkg.MyInterface
	value = mypkg.MyType(5)
	value.MethodWithoutParameters()
	value.MethodWithParameter(127.3)
	fmt.Println(value.MethodWithReturnValue())*/
	/*	var toy NoiseMaker

		toy = Whistle("Toyco Canary")
		toy.MakeSound()

		toy = Horn("Toyco Blaster")
		toy.MakeSound()

		toy = Robot("Botco Amber")

		var robot Robot = toy.(Robot)
		robot.Walk()*/
}

func playList(device gadget.Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}
