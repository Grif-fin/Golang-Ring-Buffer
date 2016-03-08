package ringbuffer

type RingBuffer struct{
	first int
	last int
	wrap bool
	buffer []string
}

func CreateNewRingBuffer(size int) *RingBuffer {
	msgs := make([]string, size, size)
	return &RingBuffer{-1,-1,false,msgs}
}

func (rb *RingBuffer) AddToRingBuffer(msg string){
	// first insertion
	if(rb.first == -1 && rb.last == -1){
		rb.buffer[0] = msg
		rb.first = 0
		rb.last = 0
		return
	}

	// second insertion
	if (rb.first == 0 && rb.last == 0){
		rb.buffer[1] = msg
		rb.first = 0
		rb.last = 1
		return
	}

	// insertion after wrap
	if(rb.last > rb.first){
		if(rb.last != len(rb.buffer)-1){
			if(rb.wrap){
				rb.first += 1
				rb.last += 1
			}else{
				rb.last += 1
			}
			rb.buffer[rb.last] = msg
			return
		}else{
			if(rb.first != rb.last-1){
				rb.buffer[rb.first] = msg
				rb.first = 1
				rb.last = 0
				rb.wrap = true
				return
			}else{
				rb.buffer[rb.first] = msg
				rb.first = rb.last
				rb.last = 0
				return
			}
		}
	}

	// insertions after wrap
	if(rb.last < rb.first){
		if(rb.first == len(rb.buffer) -1){
			rb.buffer[rb.first] = msg
			rb.last = rb.first
			rb.first = 0
			return
		}else{
			rb.first += 1
			rb.last += 1
			rb.buffer[rb.last] = msg
			return
		}
	}
}

func (rb *RingBuffer) GetBuffer()[]string{
	retBuffer := make([]string, len(rb.buffer), len(rb.buffer))
	
	if(rb.wrap){		
		reachedEnd := false
		counter := 0
		for i:=rb.first;i<len(rb.buffer);i++{
			retBuffer[counter] = rb.buffer[i]
			counter++
			if(i == rb.last && reachedEnd){
				break
			}

			if(i==len(rb.buffer)-1 && reachedEnd == false && counter != len(rb.buffer)){
				i = -1
				reachedEnd = true
			}
		}
	}else{
		for i:=0;i<len(rb.buffer);i++{
			retBuffer[i] = rb.buffer[i]
			if(i==rb.last){
				break
			}
		}
	}

	return retBuffer
}

