
func abs(a int, b int) int {
 if (a < b) {
 return b-a
 } 
 return a-b
}

func countGoodTriplets(arr []int, a int, b int, c int) int {
 ans := 0; n := len(arr)
 for i := 0; i < n-2; i++ {
 for j := i + 1; j < n-1; j++ {
 for k := j + 1; k < n; k++ {
 if abs(arr[i], arr[j]) <= a && abs(arr[j], arr[k]) <= b && abs(arr[i], arr[k]) <= c {
 ans += 1
 }
 return ans
}
