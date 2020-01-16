pub fn is_palindrome(input: String) -> bool {
    let reversed_input = input.chars().rev().collect::<String>();
    reversed_input == input
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_is_palindrome() {
        assert!(is_palindrome("racecar".to_string()));
        assert!(is_palindrome("".to_string()));
        assert!(is_palindrome("r".to_string()));
        assert!(!is_palindrome("hello".to_string()));
    }
}