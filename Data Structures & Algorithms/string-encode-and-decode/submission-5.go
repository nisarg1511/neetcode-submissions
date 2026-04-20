type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	
	encoded:=""
	if len(strs) <1{
		return encoded
	}
	for i,str:=range strs{
		for _,c:=range str{
			if string(c) == "\\"{
				encoded+="\\"+"\\"
			}else if string(c) == "#"{
				encoded+="\\"+"#"
			}else{
				encoded+=string(c)
			}
		}
		if i<len(strs)-1 {
			encoded+="#"
		}
	}
	return encoded + "#"
}

func (s *Solution) Decode(encoded string) []string {
	
	decoded:=[]string{}
	
	if encoded == ""{
		return decoded
	}
	encoded = encoded[:len(encoded)-1]
	word:=""

	for i:=0;i<len(encoded);i++{
		if string(encoded[i])=="\\" &&  i<len(encoded)-1{
			if string(encoded[i+1])== "\\" || string(encoded[i+1])=="#"{
				word+=string(encoded[i+1])
				i++
			}
		}else{
			if string(encoded[i])=="#"{
				decoded = append(decoded,word)
				word = ""
			}else{
				word += string(encoded[i])
			}
		}
	}
	decoded = append(decoded,word)
	return decoded
}
