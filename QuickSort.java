public class QuickSort {

    public void quickSort(int arr[], int begin, int end) {
        if (begin < end) {
            int partitionIndex = partition(arr, begin, end);

            quickSort(arr, begin, partitionIndex - 1);
            quickSort(arr, partitionIndex + 1, end);
        }
    }

    private int partition(int arr[], int begin, int end) {
        int pivot = arr[end];
        int i = (begin - 1);

        for (int j = begin; j < end; j++) {
            if (arr[j] <= pivot) {
                i++;

                int swapTemp = arr[i];
                arr[i] = arr[j];
                arr[j] = swapTemp;
            }
        }

        int swapTemp = arr[i + 1];
        arr[i + 1] = arr[end];
        arr[end] = swapTemp;

        return i + 1;
    }
    public static void main(String args[]) {
        QuickSort ob = new QuickSort();
        int arr[] = {12, 11, 13, 5, 6};
        ob.quickSort(arr, 0, arr.length - 1);

        System.out.println("Sorted array");
        for (int i : arr) {
            System.out.print(i + " ");
        }
    }
}
