package main

func predictPartyVictory(senate string) string {
	var radiant, dire int
	for _, s := range senate {
		if s == 'R' {
			radiant++
			continue
		}
		dire++
	}

	var banRadiant, banDire int
	for len(senate) > 0 && radiant > 0 && dire > 0 {
		temp := make([]rune, 0, len(senate))
		for _, s := range senate {
			if s == 'R' {
				if banRadiant > 0 {
					banRadiant--
					continue
				}
				banDire++
				dire--
				temp = append(temp, s)
			}
			if s == 'D' {
				if banDire > 0 {
					banDire--
					continue
				}
				banRadiant++
				radiant--
				temp = append(temp, s)
			}
		}
		senate = string(temp)
	}
	if dire > radiant {
		return "Dire"
	}
	return "Radiant"
}
