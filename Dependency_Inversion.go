package main

import "fmt"

type VehicleCDPlayer interface {
	PlayCD()
}
type VehicleCdPlay struct {
}
func (v VehicleCdPlay) PlayCD() {
	fmt.Println("Playnig Guns n Roses")
}
type VehicleAccelerator interface {
	Accelerate()
}
type VehicleAccelerate struct {

}
func (v VehicleAccelerate) Accelerate() {
	fmt.Println("Accelerating ....")
}
type Vehicle struct {
	VehicleAccelerate
	VehicleCdPlay
}
func main(){
	Vehicle{}.VehicleAccelerate.Accelerate()
	Vehicle{}.VehicleCdPlay.PlayCD()

}