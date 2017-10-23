package stablemarriage

func StableMatching(males, females []Person) {
	// by having a nil EngagedTo, a person is free... inital state all
	// people are free

	for SinglesExist(males) {
		for mi := range males {
			if males[mi].EngagedTo != nil {
				continue
			}

			fi := findIndex(females, males[mi].Preferences[0])
			if len(males[mi].Preferences) > 1 {
				males[mi].Preferences = males[mi].Preferences[1:]
			}
			if females[fi].EngagedTo == nil {
				propose(&males[mi], &females[fi])
				continue
			}

			if females[fi].Perfers(&males[mi], females[fi].EngagedTo) > 0 {
				males[findIndex(males, females[fi].EngagedTo.Name)].EngagedTo = nil

				propose(&males[mi], &females[fi])
				continue
			}
		}
	}
}
