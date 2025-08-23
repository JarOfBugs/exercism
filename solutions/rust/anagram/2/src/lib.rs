use std::collections::HashSet;

pub fn anagrams_for<'a>(word: &str, possible_anagrams: &[&'a str]) -> HashSet<&'a str> {
    let word_lower = word.to_lowercase();
    let target_chars = get_sorted_vector_from_string(&word_lower);

    possible_anagrams
        .iter()
        .filter(|&candidate| {
            let candidate_lower = candidate.to_lowercase();
            if word_lower == candidate_lower {
                return false;
            }
            let candidate_chars = get_sorted_vector_from_string(&candidate_lower);
            return target_chars == candidate_chars;
        })
        .cloned()
        .collect()
}

fn get_sorted_vector_from_string(word: &String) -> Vec<char> {
    let mut vector: Vec<char> = word.chars().collect();
    vector.sort_unstable();
    vector
}
