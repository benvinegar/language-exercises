/**
 * Generate a list of prime numbers up to n numbers, where
 * n must be 1 or larger
 */
pub fn primes(n: usize) -> Vec<u32> {
    assert!(n > 0);

    let mut _primes: Vec<u32> = Vec::new();
    
    _primes.push(2);

    let mut current: u32 = 3;
    while _primes.len() < n {
        let divisors = _primes.clone();
        let mut divisor_found = false;
        for divisor in divisors {
            if current % divisor == 0 {
                divisor_found = true;
                break;
            }
        }
        if !divisor_found {
            _primes.push(current);
        }
        current += 1;
    }
    _primes
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_primes() {
        assert_eq!(primes(1), [2]);
        assert_eq!(primes(2), [2, 3]);
        assert_eq!(primes(10), [2, 3, 5, 7, 11, 13, 17, 19, 23, 29]);
    }
}