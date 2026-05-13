import (
	"slices"
	"cmp"
)

type TimeMap struct {
	Store map[string][]Item
}

type Item struct{
	TimeStamp int
	Value string
}

func Constructor() TimeMap {
	return TimeMap{
		Store: make(map[string][]Item),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {

	if this.Store[key] ==  nil{
		this.Store[key] = make([]Item,0)
	}

	item:=Item{
		TimeStamp:timestamp,
		Value:value,
	}

	this.Store[key] = append(this.Store[key],item)
}

func (this *TimeMap) Get(key string, timestamp int) string {

	slices.SortFunc(this.Store[key],func (a,b Item) int{
		return cmp.Compare(a.TimeStamp,b.TimeStamp)
	})

	left:=0
	right:=len(this.Store[key])-1
	
	val:=""
	// biggest:=this.Store[key][left].TimeStamp
	for left<=right{
		mid:=(left+right)/2
		if this.Store[key][mid].TimeStamp == timestamp{
			return this.Store[key][mid].Value
		}else if this.Store[key][mid].TimeStamp<timestamp{
			// biggest = this.Store[key][mid].TimeStamp
			val = this.Store[key][mid].Value
		}
		if this.Store[key][mid].TimeStamp > timestamp{
			right = mid -1
		}else {
			left = mid +1
		}
	}
	return val
}
