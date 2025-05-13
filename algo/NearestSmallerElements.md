### 问题描述
给定一个不含有重复值的数组 `arr`，对于每一个位置 `i`，找到其左边和右边离 `i` 最近且值比 `arr[i]` 小的位置。返回所有位置相应的信息。

### 示例
输入：
```java
int[] arr = {3, 4, 1, 5, 6, 2, 7};
```
输出：
```java
[
  {-1, 2},  // 对于arr[0]=3，左边没有比它小的（-1），右边最近比它小的是arr[2]=1（索引2）
  {0, 2},   // 对于arr[1]=4，左边最近比它小的是arr[0]=3（索引0），右边最近比它小的是arr[2]=1（索引2）
  {-1, -1}, // 对于arr[2]=1，左边和右边都没有比它小的（-1）
  {2, 5},   // 对于arr[3]=5，左边最近比它小的是arr[2]=1（索引2），右边最近比它小的是arr[5]=2（索引5）
  {3, 5},   // 对于arr[4]=6，左边最近比它小的是arr[3]=5（索引3），右边最近比它小的是arr[5]=2（索引5）
  {2, -1},  // 对于arr[5]=2，左边最近比它小的是arr[2]=1（索引2），右边没有比它小的（-1）
  {5, -1}   // 对于arr[6]=7，左边最近比它小的是arr[5]=2（索引5），右边没有比它小的（-1）
]
```

### 解决方案
我们可以使用**单调栈**来解决这个问题。单调栈可以帮助我们在 O(n) 的时间复杂度内找到每个元素左边和右边第一个比它小的元素的位置。

#### 算法步骤
1. **初始化**：
   - 创建一个栈 `stack`，用于存储数组元素的索引。
   - 创建两个数组 `left` 和 `right`，分别存储每个元素左边和右边第一个比它小的元素的位置。初始时，`left` 填充 `-1`，`right` 填充 `-1`。

2. **遍历数组**：
   - 对于当前元素 `arr[i]`，如果栈不为空且栈顶元素对应的值大于等于 `arr[i]`，则弹出栈顶元素，并记录当前 `i` 为栈顶元素的右边第一个比它小的位置。
   - 如果栈不为空，则栈顶元素是 `arr[i]` 的左边第一个比它小的位置。
   - 将当前索引 `i` 压入栈中。

3. **处理剩余栈中的元素**：
   - 遍历完成后，栈中剩余的元素右边没有比它小的元素，因此 `right` 保持 `-1`。

#### Java 实现
```java
import java.util.Stack;

public class NearestSmallerElements {
    public static int[][] findNearestSmaller(int[] arr) {
        int n = arr.length;
        int[][] result = new int[n][2];
        Stack<Integer> stack = new Stack<>();

        // 初始化 left 和 right 数组
        int[] left = new int[n];
        int[] right = new int[n];
        for (int i = 0; i < n; i++) {
            left[i] = -1;
            right[i] = -1;
        }

        // 遍历数组，找到左边最近比它小的元素
        for (int i = 0; i < n; i++) {
            while (!stack.isEmpty() && arr[stack.peek()] >= arr[i]) {
                stack.pop();
            }
            if (!stack.isEmpty()) {
                left[i] = stack.peek();
            }
            stack.push(i);
        }

        stack.clear();

        // 遍历数组，找到右边最近比它小的元素
        for (int i = n - 1; i >= 0; i--) {
            while (!stack.isEmpty() && arr[stack.peek()] >= arr[i]) {
                stack.pop();
            }
            if (!stack.isEmpty()) {
                right[i] = stack.peek();
            }
            stack.push(i);
        }

        // 填充结果
        for (int i = 0; i < n; i++) {
            result[i][0] = left[i];
            result[i][1] = right[i];
        }

        return result;
    }

    public static void main(String[] args) {
        int[] arr = {3, 4, 1, 5, 6, 2, 7};
        int[][] result = findNearestSmaller(arr);

        for (int i = 0; i < result.length; i++) {
            System.out.println("{" + result[i][0] + ", " + result[i][1] + "}");
        }
    }
}
```

#### 输出
```java
{-1, 2}
{0, 2}
{-1, -1}
{2, 5}
{3, 5}
{2, -1}
{5, -1}
```

### 复杂度分析
- **时间复杂度**：O(n)，每个元素最多被压入和弹出栈一次。
- **空间复杂度**：O(n)，使用了额外的栈和数组。

### 关键点
- **单调栈**：维护一个单调递增的栈，可以高效地找到左边和右边第一个比当前元素小的位置。
- **两次遍历**：第一次遍历找到左边最近比它小的元素，第二次遍历找到右边最近比它小的元素。

### 扩展
如果需要处理数组中存在重复值的情况，可以在栈中存储元素的索引和值的组合，并在比较时考虑索引的顺序。

## solo
单调栈如何保证栈中存储为左边最小
大值pop，至少存放两个值 -x，-1