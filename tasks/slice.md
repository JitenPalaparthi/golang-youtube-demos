Create a function to delete element from the slice

func DeleteElem(slice []int, index int)([]int, error)

- The error should be if the index is beyond the len of the slice, the return an error with appropriate message
- If the slice is nil , return an error with message saying slice is nil
