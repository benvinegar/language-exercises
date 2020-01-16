/**
 * Generate a fibonacci sequence up to n numbers, where
 * n must be 2 or larger
 */
pub fn fibonacci(n: usize) -> Vec<u32> {
    assert!(n > 1);

    let mut out: Vec<u32> = Vec::new();

    let mut last: u32 = 0;
    let mut current: u32 = 1;
    let mut t: u32;

    out.push(last);
    out.push(current);

    while out.len() < n {
        t = current;
        current += last;
        last = t;
        out.push(current);
    }
    out
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_fibonacci() {
        assert_eq!(fibonacci(2), [0, 1]);
        assert_eq!(fibonacci(3), [0, 1, 1]);
        assert_eq!(fibonacci(10), [0, 1, 1, 2, 3, 5, 8, 13, 21, 34])
    }
}
