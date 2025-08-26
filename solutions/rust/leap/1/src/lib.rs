pub fn is_leap_year(year: u64) -> bool {
    (year % 4_u64 == 0 && year % 100_u64 != 0) || (year % 100_u64 == 0 && year % 400_u64 == 0)
}
