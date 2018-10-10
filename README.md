2000 times - Insert
raw: 9.21s 4605024 ns/op 552 B/op 12 allocs/op
gorm: 12.94s 6472218 ns/op 7612 B/op 159 allocs/op
orm: 14.43s 7214323 ns/op 1937 B/op 40 allocs/op
xorm: 16.26s 8132163 ns/op 2295 B/op 55 allocs/op


500 times - MultiInsert 100 row
raw: 12.29s 24573092 ns/op 110865 B/op 811 allocs/op
xorm: 12.41s 24820691 ns/op 140199 B/op 2356 allocs/op
orm: 12.86s 25713363 ns/op 147330 B/op 1534 allocs/op
gorm: Not support multi insert


2000 times - Update
raw: 0.64s 318255 ns/op 616 B/op 14 allocs/op
xorm: 0.80s 400931 ns/op 2418 B/op 75 allocs/op
orm: 0.94s 470951 ns/op 1929 B/op 40 allocs/op
gorm: 1.20s 599894 ns/op 7507 B/op 185 allocs/op


4000 times - Read
gorm: 0.02s 4511 ns/op 896 B/op 5 allocs/op
raw: 0.59s 147287 ns/op 1432 B/op 37 allocs/op
orm: 1.33s 333259 ns/op 2801 B/op 97 allocs/op
xorm: 1.54s 386046 ns/op 6120 B/op 185 allocs/op


2000 times - MultiRead limit 100
raw: 1.05s 522702 ns/op 34704 B/op 1320 allocs/op
orm: 2.29s 1145430 ns/op 85121 B/op 4287 allocs/op
xorm: 3.22s 1612216 ns/op 143741 B/op 5983 allocs/op
gorm: 0.00s 0.00 ns/op 0 B/op 0 allocs/op