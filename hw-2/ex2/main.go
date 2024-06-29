package main

import "sort"

type Candidate struct {
	Name  string
	Votes int
}

func calcVotes(allCandidates []string) []Candidate {
	if len(allCandidates) == 0 {
		return []Candidate{}
	}

	count := make(map[string]int)
	for _, candidate := range allCandidates {
		count[candidate]++
	}

	candidates := make([]Candidate, 0, len(count))
	for name, votes := range count {
		candidates = append(candidates, Candidate{Name: name, Votes: votes})
	}

	sort.Slice(candidates, func(i, j int) bool {
		if candidates[i].Votes == candidates[j].Votes {
			return candidates[i].Name < candidates[j].Name
		}

		return candidates[i].Votes > candidates[j].Votes
	})

	return candidates
}

func main() {
	candidates := calcVotes([]string{"A", "B", "C"})
	for _, candidate := range candidates {
		println(candidate.Name, candidate.Votes)
	}
}
