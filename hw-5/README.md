### Решение
1. Добавил временную мапу `temp` для передачи ее в канал. При изменении значений и отправке одной и той же мапы `currentPrice` происходит ее мутация. Если я правильно понял, то значения должны изменяться как у меня.
2. Сделал указатель на `&sync.WaitGroup{}`, чтобы передавать именно эту `wg`
3. Добавил мьютекс чтобы избавиться от гонки, которая возникает при мутации слайса мап

```
map[inst1:2.1 inst2:3.1 inst3:4.1 inst4:5.1]
map[inst1:2.1 inst2:3.1 inst3:4.1 inst4:5.1]
map[inst1:2.1 inst2:3.1 inst3:4.1 inst4:5.1]
map[inst1:2.1 inst2:3.1 inst3:4.1 inst4:5.1]
map[inst1:3.1 inst2:4.1 inst3:5.1 inst4:6.1]
map[inst1:3.1 inst2:4.1 inst3:5.1 inst4:6.1]
map[inst1:3.1 inst2:4.1 inst3:5.1 inst4:6.1]
map[inst1:3.1 inst2:4.1 inst3:5.1 inst4:6.1]
map[inst1:4.1 inst2:5.1 inst3:6.1 inst4:7.1]
map[inst1:4.1 inst2:5.1 inst3:6.1 inst4:7.1]
map[inst1:4.1 inst2:5.1 inst3:6.1 inst4:7.1]
map[inst1:4.1 inst2:5.1 inst3:6.1 inst4:7.1]
map[inst1:5.1 inst2:6.1 inst3:7.1 inst4:8.1]
map[inst1:5.1 inst2:6.1 inst3:7.1 inst4:8.1]
map[inst1:5.1 inst2:6.1 inst3:7.1 inst4:8.1]
map[inst1:5.1 inst2:6.1 inst3:7.1 inst4:8.1]
```