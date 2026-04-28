type MinStack struct {
	Stack []int
	Min *MinTrace
}
type MinTrace struct{
	Val int
	Prev *MinTrace
}
func Constructor() MinStack {
	return MinStack{
		Stack:make([]int,0),
		Min:&MinTrace{
			Val:math.MaxInt,
			Prev:nil,
		},
	}
}

func (this *MinStack) Push(val int) {
	this.Stack = append(this.Stack,val)
	if this.Min.Val >= val {
		this.Min = &MinTrace{
			Prev:this.Min,
			Val:val,
		}
	}
}

func (this *MinStack) Pop() {
	if this.Top() == this.Min.Val{
		this.Min = this.Min.Prev
	}
	if len(this.Stack)>0{
		this.Stack = this.Stack[:len(this.Stack)-1]
	}
	
}

func (this *MinStack) Top() int {
	if len(this.Stack) > 0{
		return this.Stack[len(this.Stack)-1]
	}
	return 0 
}

func (this *MinStack) GetMin() int {
	if len(this.Stack)>0{
		return this.Min.Val
	}
	return 0
}
