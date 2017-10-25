package stablemarriage

func StableMatching(males, females []Person) {
	// by having a nil EngagedTo, a person is free... inital state all
	// people are free

	for SinglesExist(males) { // ∃ free man m who still has a woman w to propose to
		for mi := range males {
			if !males[mi].IsSingle() {
				continue
			}

			fi := males[mi].GetPrefered(females)
			if females[fi].IsSingle() {
				propose(&males[mi], &females[fi])
				continue
			}

			if females[fi].PrefersToExistingEngagement(&males[mi]) {
				females[fi].BreakUp()
				propose(&males[mi], &females[fi])
				continue
			}
		}
	}
}
