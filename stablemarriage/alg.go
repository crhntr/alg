package stablemarriage

func StableMatching(males, females []Person) {
	// by having a nil EngagedTo, a person is free... inital state all
	// people are free

	// while (∃ free man m who still has a woman w to propose to)
	for SinglesExist(males) {
		for mi := range males {
			if !males[mi].IsSingle() {
				continue
			}

			fi := males[mi].GetPreferedMatch(females)
			if females[fi].IsSingle() {
				match(&males[mi], &females[fi])
				continue
			}

			if females[fi].CompareToExistingMatch(&males[mi]) {
				females[fi].BreakUp()
				match(&males[mi], &females[fi])
				continue
			}
		}
	}
}
