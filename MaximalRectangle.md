### 问题描述
给定一个仅包含 `0` 和 `1` 的二维整型矩阵 `map`，找到其中所有由 `1` 组成的矩形区域，并返回这些矩形中 `1` 的数量最大的那个矩形的 `1` 的数量。

### 示例
输入：
```java
int[][] map = {
    {1, 0, 1, 1, 1},
    {1, 1, 1, 1, 1},
    {1, 0, 0, 1, 0}
};
```
输出：
```java
6
```
解释：最大的全 `1` 矩形区域是右下角的 `2x3` 矩形，包含 `6` 个 `1`。

### 解决方案
这个问题可以通过**单调栈**和**动态规划**的思想来解决。具体步骤如下：

1. **预处理**：
   - 对于每一行，计算以该行为底部的“直方图”高度。即，对于每一列，从当前行向上连续 `1` 的数量。
   - 例如，对于第一行 `{1, 0, 1, 1, 1}`，直方图高度为 `[1, 0, 1, 1, 1]`。
   - 对于第二行 `{1, 1, 1, 1, 1}`，直方图高度为 `[2, 1, 2, 2, 2]`（因为第一列有两个连续的 `1`，第二列有一个 `1`，依此类推）。

2. **计算最大矩形面积**：
   - 对于每一行的直方图，使用单调栈的方法计算最大矩形面积。
   - 单调栈的方法可以在 O(n) 的时间内找到直方图中的最大矩形。

3. **更新全局最大值**：
   - 在每一行的直方图处理完成后，更新全局的最大 `1` 的数量。

### Java 实现
```java
import java.util.Stack;

public class MaximalRectangle {
    public static int maximalRectangle(int[][] map) {
        if (map == null || map.length == 0 || map[0].length == 0) {
            return 0;
        }
        int rows = map.length;
        int cols = map[0].length;
        int[] heights = new int[cols];
        int maxArea = 0;

        for (int i = 0; i < rows; i++) {
            // 更新直方图高度
            for (int j = 0; j < cols; j++) {
                if (map[i][j] == 1) {
                    heights[j] += 1;
                } else {
                    heights[j] = 0;
                }
            }
            // 计算当前直方图的最大矩形面积
            maxArea = Math.max(maxArea, largestRectangleArea(heights));
        }
        return maxArea;
    }

    private static int largestRectangleArea(int[] heights) {
        Stack<Integer> stack = new Stack<>();
        int maxArea = 0;
        int n = heights.length;

        for (int i = 0; i <= n; i++) {
            int h = (i == n) ? 0 : heights[i];
            while (!stack.isEmpty() && h < heights[stack.peek()]) {
                int height = heights[stack.pop()];
                int width = stack.isEmpty() ? i : i - stack.peek() - 1;
                maxArea = Math.max(maxArea, height * width);
            }
            stack.push(i);
        }
        return maxArea;
    }

    public static void main(String[] args) {
        int[][] map = {
            {1, 0, 1, 1, 1},
            {1, 1, 1, 1, 1},
            {1, 0, 0, 1, 0}
        };
        System.out.println(maximalRectangle(map)); // 输出：6
    }
}
```

### 复杂度分析
- **时间复杂度**：O(m * n)，其中 `m` 是矩阵的行数，`n` 是矩阵的列数。每一行处理直方图的时间是 O(n)。
- **空间复杂度**：O(n)，用于存储直方图的高度和单调栈。

### 关键点
- **直方图预处理**：将二维问题转化为一系列的一维直方图问题。
- **单调栈**：高效计算直方图中的最大矩形面积。
- **动态更新**：逐行更新直方图高度，并实时更新最大面积。

### 扩展
如果矩阵非常大（例如行数和列数都超过 1000），可以考虑进一步优化空间复杂度，例如复用直方图数组。此外，如果矩阵是稀疏的（即大部分是 `0`），可以尝试基于稀疏矩阵的优化方法。