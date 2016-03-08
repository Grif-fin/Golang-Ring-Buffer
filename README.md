# Golang-Ring-Buffer
Simple circular buffer of strings. 

#Example
```javascript
package main

import "strconv"

func main(){
	rb:=CreateNewRingBuffer(10)
	for i:=0;i<25;i++{
		rb.AddToRingBuffer(strconv.Itoa(i))
		fmt.Println(rb.GetBuffer())
	}
}
```

```javascript
[0         ]
[0 1        ]
[0 1 2       ]
[0 1 2 3      ]
[0 1 2 3 4     ]
[0 1 2 3 4 5    ]
[0 1 2 3 4 5 6   ]
[0 1 2 3 4 5 6 7  ]
[0 1 2 3 4 5 6 7 8 ]
[0 1 2 3 4 5 6 7 8 9]
[1 2 3 4 5 6 7 8 9 10]
[2 3 4 5 6 7 8 9 10 11]
[3 4 5 6 7 8 9 10 11 12]
[4 5 6 7 8 9 10 11 12 13]
[5 6 7 8 9 10 11 12 13 14]
[6 7 8 9 10 11 12 13 14 15]
[7 8 9 10 11 12 13 14 15 16]
[8 9 10 11 12 13 14 15 16 17]
[9 10 11 12 13 14 15 16 17 18]
[10 11 12 13 14 15 16 17 18 19]
[11 12 13 14 15 16 17 18 19 20]
[12 13 14 15 16 17 18 19 20 21]
[13 14 15 16 17 18 19 20 21 22]
[14 15 16 17 18 19 20 21 22 23]
[15 16 17 18 19 20 21 22 23 24]
```
