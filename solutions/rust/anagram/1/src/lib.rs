use std::collections::HashSet;

pub fn anagrams_for<'a>(word: &str, possible_anagrams: &[&'a str]) -> HashSet<&'a str> {
    let word_lower = word.to_lowercase();
    let mut target_chars: Vec<char> = word_lower.chars().collect();
    target_chars.sort_unstable();

    possible_anagrams
        .iter()
        .filter(|&candidate| {
            let candidate_lower = candidate.to_lowercase();
            if word_lower == candidate_lower {
                return false;
            }
            let mut candidate_chars: Vec<char> = candidate_lower.chars().collect();
            candidate_chars.sort_unstable();
            return target_chars == candidate_chars;
        })
        .cloned()
        .collect()
}
