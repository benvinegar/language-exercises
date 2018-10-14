pub fn fizz_buzz(start: u32, finish: u32) -> Vec<String> {
    let mut out: Vec<String> = Vec::new();

    for i in start..finish + 1 {
        if i % 3 == 0 && i % 5 == 0 {
            out.push("FizzBuzz".to_string());
        } else if i % 3 == 0 {
            out.push("Fizz".to_string());
        } else if i % 5 == 0 {
            out.push("Buzz".to_string());
        } else {
            out.push(i.to_string());
        }
    }
    out
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_fizz_buzz() {
        assert_eq!(
            fizz_buzz(1, 15),
            [
                "1",
                "2",
                "Fizz",
                "4",
                "Buzz",
                "Fizz",
                "7",
                "8",
                "Fizz",
                "Buzz",
                "11",
                "Fizz",
                "13",
                "14",
                "FizzBuzz"
            ]
        );
    }
}
