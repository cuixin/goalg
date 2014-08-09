package hashset

type HashSet struct {
	items map[interface{}]struct{}
}

func New() *HashSet {
	return &HashSet{make(map[interface{}]struct{})}
}

func (self *HashSet) Len() int {
	return len(self.items)
}

func (self *HashSet) ToSlice() []interface{} {
	var result []interface{}
	for k, _ := range self.items {
		result = append(result, k)
	}
	return result
}

func (self *HashSet) IsEmpty() bool {
	return self.Len() == 0
}

func (self *HashSet) Add(objects ...interface{}) {
	for i, _ := range objects {
		self.items[objects[i]] = struct{}{}
	}
}

func (self *HashSet) Contains(k interface{}) bool {
	_, ok := self.items[k]
	return ok
}

func (self *HashSet) Remove(k interface{}) bool {
	if _, ok := self.items[k]; ok {
		delete(self.items, k)
		return true
	}
	return false
}

func (self *HashSet) Clear() {
	self.items = nil
	self.items = make(map[interface{}]struct{})
}
