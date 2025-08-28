pub fn square(s: u32) -> u64 {
    if !(1..=64).contains(&s) {
        panic!("Only values between 0 and 64 are valid")
    }
    2_u64.pow(s - 1)
}

pub fn total() -> u64 {
    (1..=64).fold(0, |acc, elem| acc + square(elem))
}
