1. Размер входных данных алгоритма `len(n)`
2. Определите основную операцию алгоритма: сравнение `if t <= a[len(a)-1]`
3. Проверьте, зависит ли число выполняемых основных операций только от размера входных данных: `Да (M(n))`
4. Составьте рекуррентное уравнение, выражающее количество выполняемых основных операций алгоритма, и укажите соответствующие начальные условия
```
M(n) = M(n-1) + C
где M(n) — количество операций для массива длины n, а C — константа.
```
5. Найдите решение рекуррентного уравнения или, если это невозможно, определите хотя бы его порядок роста

```M(n) = M(n-1) + C```

Подставляя `M(n-1)`, получим: `M(n) = (M(n-2) + C) + C = M(n-2) + 2C`

...

```M(n) = MMM(n-3) + 3C```

Возьмем случай, когда `n === k => M(n) = M(0) + n*C => Const + n*Const`

### Итоговая сложность:

`Const` - убираем, потому что константа

В итоге получаем`M(n) = O(n)`, где `n` — размер исходного массива.

Таким образом, сложность данной функции `MinEl` составляет `O(n)`.
