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