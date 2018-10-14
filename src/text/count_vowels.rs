pub fn count_vowels(input: String) -> u32 {
    let mut chars_iter = input.chars();
    let mut count: u32 = 0;

    loop {
        count += match chars_iter.next() {
            Some(_char) => match _char {
                'a' => 1,
                'e' => 1,
                'i' => 1,
                'o' => 1,
                'u' => 1,
                _ => 0
            },
            None => break
        }
    }
    count
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_count_vowels() {
        assert_eq!(count_vowels("".to_string()), 0);
        assert_eq!(count_vowels("bcdbcd".to_string()), 0);
        assert_eq!(count_vowels("one".to_string()), 2);
        assert_eq!(count_vowels("one two three".to_string()), 5);
    }
}
