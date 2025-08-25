const DIGIT_RADIX: u32 = 10;

/// Check a Luhn checksum.
pub fn is_valid(code: &str) -> bool {
    let parsed: Vec<char> = code.chars().filter(|e| !e.is_whitespace()).collect();

    if parsed.iter().any(|e| !e.is_digit(DIGIT_RADIX)) {
        return false;
    }

    if parsed.len() <= 1 {
        return false;
    }

    let sum = parsed
        .iter()
        .rev()
        .enumerate()
        .map(|(index, elem)| {
            let digit = elem.to_digit(DIGIT_RADIX).unwrap();
            if index % 2 != 0 {
                let doubled = digit * 2;
                if doubled > 9 {
                    return doubled - 9;
                }
                return doubled;
            }
            digit
        })
        .sum::<u32>();

    sum % 10 == 0
}
