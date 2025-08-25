const DIGIT_RADIX: u32 = 10;

pub fn is_armstrong_number(num: u32) -> bool {
    let num_str = num.to_string();
    let number_of_digits = num_str.len() as u32;
    let sum = num_str.chars().fold(0, |acc, digit_str| {
       let digit = digit_str.to_digit(DIGIT_RADIX).unwrap();
        acc + digit.pow(number_of_digits)
    });
    sum == num
}
