type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	encoded:=""
	for _,str:=range strs{
		length:=strconv.Itoa(len(str)) 
		encoded= encoded + length + "#"+str
	}
	return encoded
}

func (s *Solution) Decode(encoded string) []string {
	decoded:=[]string{}
	length:=""
	for i:=0;i<len(encoded);i++{
		if string(encoded[i])=="#"{
			l,_:=strconv.Atoi(length)
			word:=encoded[i+1:i+l+1]
			decoded = append(decoded,word)
			i=i+l
			length = ""
		}else{
			length+=string(encoded[i])
		}
	}
	return decoded
}
