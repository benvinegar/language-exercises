/**
 * Validate a credit card number up to 16 digits long 
 * using Luhn's algorithm (https://en.wikipedia.org/wiki/Luhn_algorithm)
 */
pub fn validate_credit_card(num: String) -> bool {
    let reversed = num.chars().rev().collect::<String>();
    let mut doubled = [0; 16];

    for (idx, ch) in reversed.char_indices() {
        let mut value = ch.to_digit(10).unwrap();
        // multiply odd-indexed digits by 2
        if idx % 2 != 0 {
            value *= 2;
        }
        // if value is > 10, sum its digits
        // e.g. 8 doubled becomes 16, whose digits summed is 7
        if value >= 10 {
            value = sum_digits(value);
        }
        doubled[idx] = value;
    }

    // sum the transformed digits
    let mut sum = 0;
    for value in &doubled {
        sum += value;
    }
    sum % 10 == 0
}

/**
 * Given an unsigned integer, return the sum of that 
 * integers digits.
 * e.g. sum_digits(32) => 5
 */
fn sum_digits(num: u32) -> u32 {
    let mut sum = 0;
    for ch in num.to_string().chars() {
        sum += ch.to_digit(10).unwrap();
    }
    sum
}


#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_validate_credit_card() {
        assert!(validate_credit_card("4111111111111111".to_string())); // visa
        assert!(validate_credit_card("3566002020360505".to_string())); // jcb
        assert!(validate_credit_card("5555555555554444".to_string())); // mastercard
        assert!(validate_credit_card("30569309025904".to_string())); // diner's club

        assert!(!validate_credit_card("4111111111111112".to_string())); // bad
    }
}
