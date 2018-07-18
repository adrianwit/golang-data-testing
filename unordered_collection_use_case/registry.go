package viant



type Registry map[int]*Message

func (r *Registry) Register(message *Message) {
	(*r)[message.ID] = message
}

func (r *Registry) AsSlice() []*Message {
    var result = make([]*Message, len(*r))
    var i = 0
    for _, v := range *r {
		result[i] = v
		i++
	}
	return result
}


