package utils

// Count 传入数据里面包含在另一个slice中的个数
func Count(slice1 PairList, slice2 []string) int {
  count := 0
  for i := 0; i < len(slice1); i++ {
    for j := 0; j < len(slice2); j++ {
      if slice1[i].Key == slice2[j] {
        count += slice1[i].Value
      }
    }
  }
  return count
}
