public class InsertionSort {
    public static void main(String[] args) {
        int arr[] = {12, 11, 13, 5, 6};
        insertionSort(arr);
        System.out.println("Sorted array: ");
        for (int i : arr) {
            System.out.print(i + " ");
        }
    }
    static void insertionSort(int arr[])
    {
           int n = arr.length;
           for (int i = 1; i < n; ++i) {
               int key = arr[i];
               int j = i - 1;
    
               while (j >= 0 && arr[j] > key) {
                   arr[j + 1] = arr[j];
                   j = j - 1;
               }
               arr[j + 1] = key;
           }
    }

}
# 插入排序（在有序序列中）
    // This function sorts arr[] using insertion sort
