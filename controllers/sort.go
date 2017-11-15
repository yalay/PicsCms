package controllers

type IdsFreq struct {
	ids   []int
	freqs map[int]int
}

func NewIdsFreq() *IdsFreq {
	return &IdsFreq{
		ids:make([]int, 0),
		freqs:make(map[int]int, 0),
	}
}

func (f *IdsFreq) Append(newIds []int) {
	for _, newId := range newIds {
		if num, ok := f.freqs[newId]; ok {
			f.freqs[newId] = num + 1
		} else {
			f.freqs[newId] = 1
			f.ids = append(f.ids, newId)
		}
	}
}

func (f *IdsFreq) Top(num int) []int {
	if num <= 0 {
		return nil
	}

	if len(f.ids) < num {
		return f.ids
	}
	return f.ids[:num]
}

func (f *IdsFreq) Len() int      { return len(f.ids) }
func (f *IdsFreq) Swap(i, j int) { f.ids[i], f.ids[j] = f.ids[j], f.ids[i] }
func (f *IdsFreq) Less(i, j int) bool {
	return f.freqs[f.ids[i]] < f.freqs[f.ids[j]]
}
