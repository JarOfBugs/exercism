use unicode_segmentation::UnicodeSegmentation;

pub fn reverse(input: &str) -> String {
    let mut result = String::new();
    for i in input.graphemes(true).rev() {
        result.push_str(i);
    }
    result
}
