import java.util.Deque;
import java.util.LinkedList;

public class SlidingWindowMaximum {
    
    public static int[] maxSlidingWindow(int[] nums, int k) {
        if (nums == null || nums.length == 0 || k <= 0) {
            return new int[0];
        }
        
        int n = nums.length;
        int[] result = new int[n - k + 1];
        int ri = 0; // 结果数组的索引
        
        // 双端队列，存储的是数组元素的索引
        Deque<Integer> deque = new LinkedList<>();
        
        for (int i = 0; i < nums.length; i++) {
            // 移除队列中不在当前窗口的元素
            while (!deque.isEmpty() && deque.peek() < i - k + 1) {
                deque.poll();
            }
            
            // 移除队列中所有小于当前元素的元素，因为它们不可能是最大值
            while (!deque.isEmpty() && nums[deque.peekLast()] < nums[i]) {
                deque.pollLast();
            }
            
            // 添加当前元素到队列
            deque.offer(i);
            
            // 当窗口形成后，开始记录结果
            if (i >= k - 1) {
                result[ri++] = nums[deque.peek()];
            }
        }
        
        return result;
    }
    
    public static void main(String[] args) {
        int[] nums = {1, 3, -1, -3, 5, 3, 6, 7};
        int k = 3;
        int[] result = maxSlidingWindow(nums, k);
        
        System.out.println("窗口最大值数组:");
        for (int num : result) {
            System.out.print(num + " ");
        }
        // 输出: 3 3 5 5 6 7
    }
}