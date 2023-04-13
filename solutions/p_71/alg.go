package main

func simplifyPath(path string) string {
	stack := make([]string, 0, 100)
	cur := make([]byte, 0, 100)
	for i := 0; i < len(path); i++ {
		if len(cur) == 0 {
			periodCount := 0
			for i < len(path) && path[i] == '.' {
				periodCount++
				cur = append(cur, '.')
				i++
			}
			if ((i < len(path) && path[i] == '/') || i >= len(path)) && periodCount == 1 {
				cur = cur[:0]
				continue
			}
			if ((i < len(path) && path[i] == '/') || i >= len(path)) && periodCount == 2 {
				if len(stack) > 0 {
					stack = stack[:len(stack)-1]
				}
				cur = cur[:0]
				continue
			}
		}
		if i < len(path) && path[i] == '/' {
			if len(cur) > 0 {
				stack = append(stack, string(cur))
				cur = cur[:0]
			}
			continue
		}
		if i < len(path) && path[i] != '/' {
			cur = append(cur, path[i])
		}
	}
	if len(cur) > 0 {
		stack = append(stack, string(cur))
	}
	newPath := "/"
	if len(stack) > 0 {
		newPath += stack[0]
	}
	for i := 1; i < len(stack); i++ {
		newPath += "/" + stack[i]
	}
	return newPath
}
