# golang_generate_parentheses

Given `n` pairs of parentheses, write a function to *generate all combinations of well-formed parentheses*.

## Examples

**Example 1:**

```
Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]

```

**Example 2:**

```
Input: n = 1
Output: ["()"]
```

**Constraints:**

- `1 <= n <= 8`

## 解析

給定一個正整數 n

要求實作一個演算法 算出所所有由 n 組字元配對 '(', ')' 所組成的合法配對字元

假設有一個字串 s 是合法的字元配對

1 代表 s 具有 n 組字元配對 '(', ')' 所組成, 所以字串 s 具有 2*n 個字元 

2 因為每個 '(' 必須找到對應在其右方的 ')', 所以要能配對完成, 必須符合以下條件

   2.1 從字串左方往右逐步計算 '(' 與 ')' 個數，會發現每個位置 '(' 個數 ≥ ')' 個數 

當 ')' 個數大於 '(' 如下圖

![](https://i.imgur.com/nO0IRpP.png)

所以我們可以使用 backTracking 的方式逐步找出所有可能的字串

所謂回溯法也就是 BackTracking 就是先試所有可能當遇到不符合條件時，再往前一步做回溯找尋下一個可能

每次增加一個字元，並且紀錄 '(' 與 ')' 的個數

使用下面的原則來找可能的字串

1 當 蒐集的字串長度 == 2n 時代表已經找到把該字串加入 結果

2 當 '(' 個數 < n , 則首先把下一個字元使用 '(' 帶入 繼續往下找

3 當 '(' 個數 > ')' 個數 , 把下一個字元使用 ')' 帶入繼續往下找

以n = 3 為例作法如下圖


![](https://i.imgur.com/fB7AGrJ.png)


## 程式碼
```go
package sol

func generateParenthesis(n int) []string {
	if n == 1 {
		return []string{"()"}
	}
	ans := []string{}
	var backTrack func(cur string, openCount int, closeCount int)
	backTrack = func(cur string, openCount int, closeCount int) {
		if len(cur) == 2*n {
			ans = append(ans, cur)
			return
		}
		if openCount < n {
			cur += "("
			backTrack(cur, openCount+1, closeCount)
			cur = cur[:len(cur)-1]
		}
		if closeCount < openCount {
			cur += ")"
			backTrack(cur, openCount, closeCount+1)
			cur = cur[:len(cur)-1]
		}
	}
	backTrack("", 0, 0)
	return ans
}

```
## 困難點

1. 需要想出合法字元配對的規則
2. 需要知道 BackTracking 的概念

## Solve Point

- [x]  初始化 openCount =0, closeCount =0 分別用來紀錄 '(', ')' 的字元個數, 初始化 ans = []string{}, 初始化 currentString = “”
- [x]  當 len(currentString) == 2*n  則把 currentString 放入 ans 之中
- [x]  當 openCount  < n 時 , 先把 '(' append 到 currentString 之中並且把 openCount+1 來繼續往下走
- [x]  當 closeCount < openCount 時，把 ')' append 到 currentString 之中 並且把 closeCount + 1 來繼續往下尋找
- [x]  當所有可能都找完了 ans 即為所求