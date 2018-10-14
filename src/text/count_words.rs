pub fn count_words(input: String) -> usize {
    let words_iter = input.split_whitespace();
    words_iter.count()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_count_words() {
        assert_eq!(count_words("".to_string()), 0);
        assert_eq!(count_words("one".to_string()), 1);
        assert_eq!(count_words("one two three".to_string()), 3);
        assert_eq!(count_words("one    two      three".to_string()), 3);
    }
}
