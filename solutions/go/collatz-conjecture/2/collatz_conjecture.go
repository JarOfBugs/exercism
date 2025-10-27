package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
    count := 0
    for n != 1 {
        switch {
            case n > 0 && n % 2 == 0:
            	n /= 2
            case n > 0 && n % 2 != 0:
            	n = n * 3 + 1
            default:
            	return 0, errors.New("What kinda number did you use?")
        }
        count++
    }
    return count, nil
}
