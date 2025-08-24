#[derive(Debug, PartialEq, Eq)]
pub enum Comparison {
    Equal,
    Sublist,
    Superlist,
    Unequal,
}

pub fn sublist(first_list: &[i32], second_list: &[i32]) -> Comparison {
    if first_list == second_list {
        return Comparison::Equal;
    }

    if first_list.is_empty() {
        return Comparison::Sublist;
    }

    if second_list.is_empty() {
        return Comparison::Superlist;
    }

    let mut greater: &[i32] = first_list;
    let mut smaller: &[i32] = second_list;
    let mut positive_comparison = Comparison::Superlist;

    if second_list.len() > first_list.len() {
        greater = second_list;
        smaller = first_list;
        positive_comparison = Comparison::Sublist;
    }

    let result = greater
        .windows(smaller.len())
        .any(|sublist| sublist.eq(smaller));

    if result {
        return positive_comparison;
    }
    Comparison::Unequal
}
