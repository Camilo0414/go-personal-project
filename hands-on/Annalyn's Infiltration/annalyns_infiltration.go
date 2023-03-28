package annalyn

// CanFastAttack can be executed only when the knight is sleeping.
func CanFastAttack(knightIsAwake bool) bool {
	att := false
	if !knightIsAwake {
		att = true
	}
	return att
}

// CanSpy can be executed if at least one of the characters is awake.
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	spy := false
	if knightIsAwake || archerIsAwake || prisonerIsAwake {
		spy = true
	}
	return spy
}

// CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping.
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	sig := false
	if !archerIsAwake && prisonerIsAwake {
		sig = true
	}
	return sig
}

// CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping.
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	free := false
	if (prisonerIsAwake && !knightIsAwake && !archerIsAwake) || (petDogIsPresent && !archerIsAwake) {
		free = true
	}
	return free
}
