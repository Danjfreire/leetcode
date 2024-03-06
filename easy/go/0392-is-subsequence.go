func isSubsequence(s string, t string) bool {
   subIdx := 0 

   if len(s) == 0 {
	return true
   }

   if len(s) > len(t) {
	return false
   }

   for i := 0; i < len(t); i++ {
		if t[i] == s[subIdx] {
			subIdx++
			if subIdx == len(s) {
				return true
			}
		}
   }

   return false
}